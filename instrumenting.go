package ordering

import (
	"./operating"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

// NewInstrumentingService ..
func NewInstrumentingService(svc *operating.Orderoperating) *GRPCService {
	server := GRPCService{
		// ------ 商品 ------//
		Commodity: kitgrpc.NewServer(
			makeGetCommodityListsEndpoint(svc),
			decodeCommodityRequest,
			encodeResponse,
		),
		// ----- 订单------//
		OrderinfoLists: kitgrpc.NewServer(
			makeGetorderinfoListsEndpoint(svc),
			decodeGetorderinfoListsRequest,
			encodeResponse,
		),
		OrderCreate: kitgrpc.NewServer(
			makeCreateOrderEndpoint(svc),
			decodeCreateOrderRequest,
			encodeResponse,
		),
		OrderDel: kitgrpc.NewServer(
			makeDelOrderEndpoint(svc),
			decodeDelOrderRespRequest,
			encodeResponse,
		),
		//-----付款-----//
		PayList: kitgrpc.NewServer(
			makeGetPayListEndpoint(svc),
			decodeGetPayListRespRequest,
			encodeResponse,
		),
		PayCreate: kitgrpc.NewServer(
			makeCreatePayInfoEndpoint(svc),
			decodeCreatePayInfoRespRequest,
			encodeResponse,
		),
		PayDel: kitgrpc.NewServer(
			makeDelPayInfoEndpoint(svc),
			decodeDelPayInfoRespRequest,
			encodeResponse,
		),
		SuccessPay: kitgrpc.NewServer(
			makePaySuccessEndpoint(svc),
			decodePaySuccessRespRequest,
			encodeResponse,
		),
		//----退款----//
		RefundInfoLists: kitgrpc.NewServer(
			makeGetRefundInfoEndpoint(svc),
			decodeGetRefundInfoRequest,
			encodeResponse,
		),
		ReFundCreate: kitgrpc.NewServer(
			makeCreateReFundInfoEndpoint(svc),
			decodeCreateReFundInfoRequest,
			encodeResponse,
		),
		RefundDel: kitgrpc.NewServer(
			makeDelRefundinfoEndpoint(svc),
			decodeDelRefundinfoRequest,
			encodeResponse,
		),
		SuccessRefund: kitgrpc.NewServer(
			makeRefundSuccessEndpoint(svc),
			decodeRefundSuccessRequest,
			encodeResponse,
		),
	}
	return &server
}
