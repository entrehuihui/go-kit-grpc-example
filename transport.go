package ordering

import (
	"context"
	"errors"

	"./proto"
)

// ErrorDecode .
var ErrorDecode = errors.New("decode error")

func encodeResponse(_ context.Context, rsp interface{}) (interface{}, error) {
	return rsp, nil
}

// ---------商品----------//

func decodeCommodityRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*proto.GetCommodityListsReq)
	if !ok {
		return nil, ErrorDecode
	}
	return req, nil
}

// ---------订单----------//

func decodeGetorderinfoListsRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*proto.GetorderinfoListsReq)
	if !ok {
		return nil, ErrorDecode
	}
	return req, nil
}

func decodeCreateOrderRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*proto.CreateOrderReq)
	if !ok {
		return nil, ErrorDecode
	}
	return req, nil
}

func decodeDelOrderRespRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*proto.DelOrderReq)
	if !ok {
		return nil, ErrorDecode
	}
	return req, nil
}

// ---------付款----------//

func decodeGetPayListRespRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*proto.GetPayListReq)
	if !ok {
		return nil, ErrorDecode
	}
	return req, nil
}

func decodeCreatePayInfoRespRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*proto.CreatePayInfoReq)
	if !ok {
		return nil, ErrorDecode
	}
	return req, nil
}

func decodeDelPayInfoRespRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*proto.DelPayInfoReq)
	if !ok {
		return nil, ErrorDecode
	}
	return req, nil
}

func decodePaySuccessRespRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*proto.PaySuccessReq)
	if !ok {
		return nil, ErrorDecode
	}
	return req, nil
}

// ---------退款----------//

func decodeGetRefundInfoRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*proto.GetRefundReq)
	if !ok {
		return nil, ErrorDecode
	}
	return req, nil
}

func decodeCreateReFundInfoRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*proto.CreateReFundReq)
	if !ok {
		return nil, ErrorDecode
	}
	return req, nil
}

func decodeDelRefundinfoRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*proto.DelRefundReq)
	if !ok {
		return nil, ErrorDecode
	}
	return req, nil
}

func decodeRefundSuccessRequest(_ context.Context, r interface{}) (interface{}, error) {
	req, ok := r.(*proto.RefundSuccessReq)
	if !ok {
		return nil, ErrorDecode
	}
	return req, nil
}
