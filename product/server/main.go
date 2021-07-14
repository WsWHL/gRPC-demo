package main

import (
	"context"
	"grpc-demo/product"
	"log"
	"net"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	address = "127.0.0.1:5001"
)

type server struct {
	products map[string]*product.Product
}

// 添加商品
func (s *server) AddProduct(c context.Context, req *product.Product) (resp *product.ProductId, err error) {
	resp = &product.ProductId{}
	out, err := uuid.NewV4()
	if err != nil {
		return resp, status.Errorf(codes.Internal, "err while generate the uuid ", err)
	}
	req.Id = out.String()
	if s.products == nil {
		s.products = make(map[string]*product.Product)
	}
	s.products[req.Id] = req
	resp.Value = req.Id
	return
}

// 获取商品
func (s *server) GetProduct(c context.Context, req *product.ProductId) (resp *product.Product, err error) {
	if req.Value == "" {
		return resp, status.Error(codes.InvalidArgument, "Product id cannot be empty")
	}
	if s.products == nil {
		s.products = make(map[string]*product.Product)
	}
	if product, ok := s.products[req.Value]; ok {
		resp = product
		return
	}
	return resp, status.Error(codes.NotFound, "Can't find this product.")
}

func main() {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Println("net listen error", err)
	}

	s := grpc.NewServer()
	product.RegisterProductInfoServer(s, &server{})
	log.Println("start gRPC listen on ", address)
	if err := s.Serve(listener); err != nil {
		log.Println("filed error", err)
		return
	}
}
