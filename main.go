package main

import (
	"log"
	"net"

	pb "github.com/hoosu418/consigment_service/proto/consignment"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// IRepository aa
type IRepository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
}

// Repository -- 模拟一个数据库
type Repository struct {
	consignments []*pb.Consignment
}

// Create 实现Reposigory
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated = append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

// service 要实现在.proto定义的所有方法
type service {
	repo IRepository
}

// CreateConsignment -- 在proto中，定义了这个方法
// 这个方法接受一个context以及proto里面定义的Consignment消息
// 这个消息是由gRPC服务器处理之后提供给你的
func (s *service) CreateConsignment (ctx context.Context, req *pb.Consignment) (*pb.Response, error){

	// 保存consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	// 返回的数据也要复合.proto里面定义的数据结构
	return &pb.Response{Created: true, Consigment: consignment}, nil
}

func main() {
	repo := &Repository{}
	lis, err = net.Listen("tpc",port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	// 在服务器上注册微服务，让自己的代码与*.pb.go中的各种interface对应起来
	pb.RegisterShippingServiceServer(s, &service{repo})

	// 在gRPC服务器上注册reflection
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
