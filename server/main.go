package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"go.mongodb.org/mongo-driver/bson"

	"net"

	salespb "../proto"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var db *mongo.Client
var salesdb *mongo.Collection
var mongoCtx context.Context

type SalesServer struct{}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Starting server on port :4040...")

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	salespb.RegisterSalesServiceServer(srv, &SalesServer{})
	reflection.Register(srv)
	fmt.Println("Connecting to MongoDB...")
	mongoCtx = context.Background()

	db, err := mongo.Connect(mongoCtx, options.Client().ApplyURI("mongodb://poc-sales:poc-sales@127.0.0.1:27017/poc-sales"))
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping(mongoCtx, nil)
	if err != nil {
		log.Fatalf("Could not connect to mongoDB %v\n", err)
	} else {
		fmt.Println("Connected to MongoDB")
	}
	salesdb = db.Database("poc-sales").Collection("sales")

	go func() {
		if e := srv.Serve(listener); e != nil {
			panic(e)
		}
	}()

	fmt.Println("Server succesfully started on port :50051")

	c := make(chan os.Signal)

	signal.Notify(c, os.Interrupt)
	<-c

	fmt.Println("\nStopping server...")
	srv.Stop()
	listener.Close()
	fmt.Println("Closing mongoDb connection...")
	db.Disconnect(mongoCtx)
	fmt.Println("Bye Bye")
}

type SaleItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Product  string             `bson:"product"`
	Quantity int64              `bson:"quantity"`
}

// AddSale {}
func (s *SalesServer) AddSale(ctx context.Context, request *salespb.Request) (*salespb.Response, error) {
	sale := request.GetSale()
	data := SaleItem{
		Product:  sale.GetProduct(),
		Quantity: sale.GetQuantity(),
	}
	result, err := salesdb.InsertOne(mongoCtx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	oid := result.InsertedID.(primitive.ObjectID)
	sale.Id = oid.Hex()
	response := &salespb.Response{}
	response.Sale = sale
	return response, nil
}

// ListSales {}
func (s *SalesServer) ListSales(req *salespb.ListSalesReq, stream salespb.SalesService_ListSalesServer) error {
	data := &SaleItem{}
	cursor, err := salesdb.Find(mongoCtx, bson.M{})
	if err != nil {
		status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	defer cursor.Close(mongoCtx)

	for cursor.Next(mongoCtx) {
		err := cursor.Decode(data)
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}
		stream.Send(&salespb.ListSalesRes{
			Sale: &salespb.Sale{
				Id:       data.ID.Hex(),
				Product:  data.Product,
				Quantity: data.Quantity,
			},
		})
	}
	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %v", err))
	}
	return nil
}

func (s *SalesServer) ListById(ctx context.Context, req *salespb.RequestById) (*salespb.Sale, error) {
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		status.Errorf(
			codes.Internal,
			fmt.Sprintf("Inválid Id: %v", err),
		)
	}

	result := salesdb.FindOne(mongoCtx, bson.M{"_id": oid})
	data := &SaleItem{}

	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find sale with Object Id %s: %v", req.GetId(), err))
	}
	response := &salespb.Sale{
		Id:       oid.Hex(),
		Product:  data.Product,
		Quantity: data.Quantity,
	}
	return response, nil
}
