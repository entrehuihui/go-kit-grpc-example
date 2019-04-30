package main

import (
	"context"
	"fmt"
	"log"

	pb "../../proto"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:8181"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// getCommodityLists(conn)

	// id := createOrder(conn)
	// getorderinfoLists(conn)
	// delOrder(conn, id)

	// getPayList(conn)
	// createPayInfo(conn)
	// delPayInfo(conn)
	// paySuccess(conn)

	// createReFundInfo(conn)
	// getRefundInfo(conn)
	// delRefundinfo(conn)
	refundSuccess(conn)
}

func getCommodityLists(conn *grpc.ClientConn) {
	c := pb.NewCommodityClient(conn)
	a := pb.GetCommodityListsReq{
		Limit:  3,
		Offset: 0,
		ID:     0,
	}
	fmt.Println()
	r, err := c.GetCommodityLists(context.Background(), &a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
}

func getorderinfoLists(conn *grpc.ClientConn) {
	c := pb.NewOrderinfoClient(conn)
	a := pb.GetorderinfoListsReq{
		Limit:  3,
		Offset: 0,
		ID:     0,
	}
	fmt.Println()
	r, err := c.GetorderinfoLists(context.Background(), &a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
}

func createOrder(conn *grpc.ClientConn) int64 {
	c := pb.NewOrderinfoClient(conn)
	a := pb.CreateOrderReq{
		Newprice: 400000,
		CID:      1,
	}
	fmt.Println()
	r, err := c.CreateOrder(context.Background(), &a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
	return (r.ID)
}

func delOrder(conn *grpc.ClientConn, id int64) {
	c := pb.NewOrderinfoClient(conn)
	a := pb.DelOrderReq{
		ID: id,
	}
	fmt.Println()
	r, err := c.DelOrder(context.Background(), &a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
}

func getPayList(conn *grpc.ClientConn) {
	c := pb.NewPayClient(conn)
	a := pb.GetPayListReq{
		Limit:  3,
		Offset: 0,
		ID:     0,
	}
	fmt.Println()
	r, err := c.GetPayList(context.Background(), &a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
}

func createPayInfo(conn *grpc.ClientConn) {
	c := pb.NewPayClient(conn)
	a := pb.CreatePayInfoReq{
		CIDs:  []int64{5},
		Price: 400000,
	}
	fmt.Println()
	r, err := c.CreatePayInfo(context.Background(), &a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
}

func delPayInfo(conn *grpc.ClientConn) {
	c := pb.NewPayClient(conn)
	a := pb.DelPayInfoReq{
		ID: 16,
	}
	fmt.Println()
	r, err := c.DelPayInfo(context.Background(), &a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
}

func paySuccess(conn *grpc.ClientConn) {
	c := pb.NewPayClient(conn)
	a := pb.PaySuccessReq{
		ID:     21,
		Price:  400000,
		Way:    1,
		Number: "1",
		Serial: "1",
	}
	fmt.Println()
	r, err := c.PaySuccess(context.Background(), &a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
}

//
func getRefundInfo(conn *grpc.ClientConn) {
	c := pb.NewRefundClient(conn)
	a := pb.GetRefundReq{
		Limit:  3,
		Offset: 0,
		ID:     0,
	}
	fmt.Println()
	r, err := c.GetRefundInfo(context.Background(), &a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
}

func createReFundInfo(conn *grpc.ClientConn) {
	c := pb.NewRefundClient(conn)
	a := pb.CreateReFundReq{
		OID:   5,
		Price: 400000,
	}
	fmt.Println()
	r, err := c.CreateReFundInfo(context.Background(), &a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
}

func delRefundinfo(conn *grpc.ClientConn) {
	c := pb.NewRefundClient(conn)
	a := pb.DelRefundReq{
		ID: 1,
	}
	r, err := c.DelRefundinfo(context.Background(), &a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
}

func refundSuccess(conn *grpc.ClientConn) {
	c := pb.NewRefundClient(conn)
	a := pb.RefundSuccessReq{
		ID:     2,
		Price:  400000,
		Way:    1,
		Number: "1",
		Serial: "1",
	}
	fmt.Println()
	r, err := c.RefundSuccess(context.Background(), &a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
}
