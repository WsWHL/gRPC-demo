package main

import (
	"context"
	"grpc-demo/product"
	"log"

	"google.golang.org/grpc"
)

const (
	address = "127.0.0.1:5001"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Println("did not connect.", err)
		return
	}
	defer conn.Close()

	client := product.NewProductInfoClient(conn)
	ctx := context.Background()

	productId := addProduct(ctx, client)
	getProduct(ctx, client, productId)
}

// 添加商品
func addProduct(ctx context.Context, client product.ProductInfoClient) (id string) {
	mac := &product.Product{Name: "iMac Pro 2021", Description: "Ths is iMac pro"}
	productId, err := client.AddProduct(ctx, mac)
	if err != nil {
		log.Println("add product failed.", err)
		return
	}
	id = productId.Value
	log.Println("add product success. id = ", id)
	return
}

// 获取商品
func getProduct(ctx context.Context, client product.ProductInfoClient, id string) {
	productId := &product.ProductId{Value: id}
	item, err := client.GetProduct(ctx, productId)
	if err != nil {
		log.Println("get product failed.", err)
		return
	}
	log.Printf("get product success. %v", item)
}
