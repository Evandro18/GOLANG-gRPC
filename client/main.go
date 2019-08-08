package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	salespb "../proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type Sale struct {
	// ID       string `bson:"_id,omitempty"`
	Product  string `json:"product"`
	Quantity int64  `json:"quantity"`
}

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := salespb.NewSalesServiceClient(conn)
	g := gin.Default()
	g.GET("/sales", func(ctx *gin.Context) {
		req := &salespb.ListSalesReq{}
		stream, err := client.ListSales(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprint(err),
			})
		}
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": fmt.Sprint(err),
				})
			}
			ctx.JSON(http.StatusOK, res)
		}
	})

	g.POST("/sale", func(ctx *gin.Context) {
		var json Sale
		if err := ctx.ShouldBindJSON(&json); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		sale := &salespb.Sale{}
		sale.Product = json.Product
		sale.Quantity = json.Quantity
		req := &salespb.Request{Sale: sale}

		if response, err := client.AddSale(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, response.Sale)
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
