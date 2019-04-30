package ordering

import (
	"context"

	"./proto"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

// GRPCService .
type GRPCService struct {
	//商品
	Commodity kitgrpc.Handler
	//订单
	OrderinfoLists kitgrpc.Handler
	OrderCreate    kitgrpc.Handler
	OrderDel       kitgrpc.Handler
	// 付款
	PayList    kitgrpc.Handler
	PayCreate  kitgrpc.Handler
	PayDel     kitgrpc.Handler
	SuccessPay kitgrpc.Handler
	// 退款
	RefundInfoLists kitgrpc.Handler
	ReFundCreate    kitgrpc.Handler
	RefundDel       kitgrpc.Handler
	SuccessRefund   kitgrpc.Handler
}

// ---------商品----------//

// GetCommodityLists 只做透传
func (s *GRPCService) GetCommodityLists(ctx context.Context, in *proto.GetCommodityListsReq) (*proto.GetCommodityListsResp, error) {
	_, rsp, err := s.Commodity.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*proto.GetCommodityListsResp), nil
}

// ---------订单----------//

// GetorderinfoLists 只做透传
func (s *GRPCService) GetorderinfoLists(ctx context.Context, in *proto.GetorderinfoListsReq) (*proto.GetorderinfoListsResp, error) {
	_, rsp, err := s.OrderinfoLists.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*proto.GetorderinfoListsResp), nil
}

// CreateOrder 只做透传
func (s *GRPCService) CreateOrder(ctx context.Context, in *proto.CreateOrderReq) (*proto.CreateOrderResp, error) {
	_, rsp, err := s.OrderCreate.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*proto.CreateOrderResp), nil
}

// DelOrder 只做透传
func (s *GRPCService) DelOrder(ctx context.Context, in *proto.DelOrderReq) (*proto.DelOrderResp, error) {
	_, rsp, err := s.OrderDel.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*proto.DelOrderResp), nil
}

// ---------付款----------//

// GetPayList 只做透传
func (s *GRPCService) GetPayList(ctx context.Context, in *proto.GetPayListReq) (*proto.GetPayListResp, error) {
	_, rsp, err := s.PayList.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*proto.GetPayListResp), nil
}

// CreatePayInfo 只做透传
func (s *GRPCService) CreatePayInfo(ctx context.Context, in *proto.CreatePayInfoReq) (*proto.CreatePayInfoResp, error) {
	_, rsp, err := s.PayCreate.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*proto.CreatePayInfoResp), nil
}

// DelPayInfo 只做透传
func (s *GRPCService) DelPayInfo(ctx context.Context, in *proto.DelPayInfoReq) (*proto.DelPayInfoResp, error) {
	_, rsp, err := s.PayDel.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*proto.DelPayInfoResp), nil
}

// PaySuccess 只做透传
func (s *GRPCService) PaySuccess(ctx context.Context, in *proto.PaySuccessReq) (*proto.PaySuccessResp, error) {
	_, rsp, err := s.SuccessPay.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*proto.PaySuccessResp), nil
}

// ---------退款----------//

// GetRefundInfo 只做透传
func (s *GRPCService) GetRefundInfo(ctx context.Context, in *proto.GetRefundReq) (*proto.GetRefundResp, error) {
	_, rsp, err := s.RefundInfoLists.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*proto.GetRefundResp), nil
}

// CreateReFundInfo 只做透传
func (s *GRPCService) CreateReFundInfo(ctx context.Context, in *proto.CreateReFundReq) (*proto.CreateReFundResp, error) {
	_, rsp, err := s.ReFundCreate.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*proto.CreateReFundResp), nil
}

// DelRefundinfo 只做透传
func (s *GRPCService) DelRefundinfo(ctx context.Context, in *proto.DelRefundReq) (*proto.DelRefundResp, error) {
	_, rsp, err := s.RefundDel.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*proto.DelRefundResp), nil
}

// RefundSuccess 只做透传
func (s *GRPCService) RefundSuccess(ctx context.Context, in *proto.RefundSuccessReq) (*proto.RefundSuccessResp, error) {
	_, rsp, err := s.SuccessRefund.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*proto.RefundSuccessResp), nil
}
