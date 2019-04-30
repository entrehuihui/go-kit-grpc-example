package ordering

import (
	"context"

	"./operating"
	"./proto"
	"github.com/go-kit/kit/endpoint"
)

// makeGetBookListEndpoint ...
func makeGetCommodityListsEndpoint(s *operating.Orderoperating) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(*proto.GetCommodityListsReq)
		data, err := s.GetCommodityLists(req.ID, req.Limit, req.Offset)
		if err != nil {
			return &proto.GetCommodityListsResp{
				Code: 500,
				Err:  err.Error(),
			}, nil
		}
		return &proto.GetCommodityListsResp{
			Code:   200,
			Err:    "",
			Allnum: int64(data.AllNum),
			Data:   data.Data.([]*proto.CommodityInfo),
		}, nil
	}
}

// makeGetBookListEndpoint ...
func makeGetorderinfoListsEndpoint(s *operating.Orderoperating) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(*proto.GetorderinfoListsReq)
		data, err := s.GetOrderLists("1", req.Offset, req.Limit, req.ID, req.CID, req.Status, req.PID, req.StartTime, req.EndTime)
		if err != nil {
			return &proto.GetorderinfoListsResp{
				Code: 500,
				Err:  err.Error(),
			}, nil
		}
		return &proto.GetorderinfoListsResp{
			Code:   200,
			Err:    "",
			Allnum: int64(data.AllNum),
			Data:   data.Data.([]*proto.OrderinfoInfo),
		}, nil
	}
}

// makeGetBookListEndpoint ...
func makeCreateOrderEndpoint(s *operating.Orderoperating) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(*proto.CreateOrderReq)
		data, err := s.CreateOrder("1", req.CID, req.Newprice)
		if err != nil {
			return &proto.CreateOrderResp{
				Code: 500,
				Err:  err.Error(),
			}, nil
		}
		return &proto.CreateOrderResp{
			Code: 200,
			Err:  "",
			ID:   data,
		}, nil
	}
}

// makeDelOrderEndpoint ...
func makeDelOrderEndpoint(s *operating.Orderoperating) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(*proto.DelOrderReq)
		err := s.DelOrder("1", req.ID)
		if err != nil {
			return &proto.DelOrderResp{
				Code: 500,
				Err:  err.Error(),
			}, nil
		}
		return &proto.DelOrderResp{
			Code: 200,
			Err:  "",
		}, nil
	}
}

// makeDelOrderEndpoint ...
func makeGetPayListEndpoint(s *operating.Orderoperating) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(*proto.GetPayListReq)
		data, err := s.GetPayList("1", req.Offset, req.Limit, req.ID, req.Way, req.Status, req.StartTime, req.EndTime)
		if err != nil {
			return &proto.GetPayListResp{
				Code: 500,
				Err:  err.Error(),
			}, nil
		}
		return &proto.GetPayListResp{
			Code:   200,
			Err:    "",
			Allnum: int64(data.AllNum),
			Data:   data.Data.([]*proto.GetPayListInfo),
		}, nil
	}
}

// makeCreatePayInfoEndpoint ...
func makeCreatePayInfoEndpoint(s *operating.Orderoperating) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(*proto.CreatePayInfoReq)
		data, err := s.CreatePayInfo("1", req.CIDs, req.Price)
		if err != nil {
			return &proto.CreatePayInfoResp{
				Code: 500,
				Err:  err.Error(),
			}, nil
		}
		return &proto.CreatePayInfoResp{
			Code: 200,
			Err:  "",
			ID:   data,
		}, nil
	}
}

// makeCreatePayInfoEndpoint ...
func makeDelPayInfoEndpoint(s *operating.Orderoperating) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(*proto.DelPayInfoReq)
		err := s.DelPayInfo("1", req.ID)
		if err != nil {
			return &proto.DelPayInfoResp{
				Code: 500,
				Err:  err.Error(),
			}, nil
		}
		return &proto.DelPayInfoResp{
			Code: 200,
			Err:  "",
		}, nil
	}
}

// makeCreatePayInfoEndpoint ...
func makePaySuccessEndpoint(s *operating.Orderoperating) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(*proto.PaySuccessReq)
		err := s.PaySuccess("1", req.ID, req.Price, req.Way, req.Number, req.Serial)
		if err != nil {
			return &proto.PaySuccessResp{
				Code: 500,
				Err:  err.Error(),
			}, nil
		}
		return &proto.PaySuccessResp{
			Code: 200,
			Err:  "",
		}, nil
	}
}

// makeGetRefundInfoEndpoint ...
func makeGetRefundInfoEndpoint(s *operating.Orderoperating) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(*proto.GetRefundReq)
		data, err := s.GetRefundInfo("1", req.Offset, req.Limit, req.ID, req.PID, req.OID, req.Status, req.Way, req.StartTime, req.EndTime)
		if err != nil {
			return &proto.GetRefundResp{
				Code: 500,
				Err:  err.Error(),
			}, nil
		}
		return &proto.GetRefundResp{
			Code:   200,
			Err:    "",
			Allnum: int64(data.AllNum),
			Data:   data.Data.([]*proto.GetRefundInfo),
		}, nil
	}
}

// makeCreateReFundInfoEndpoint ...
func makeCreateReFundInfoEndpoint(s *operating.Orderoperating) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(*proto.CreateReFundReq)
		data, err := s.CreateReFundInfo("1", req.OID, req.Price, req.Way, req.Number)
		if err != nil {
			return &proto.CreateReFundResp{
				Code: 500,
				Err:  err.Error(),
			}, nil
		}
		return &proto.CreateReFundResp{
			Code: 200,
			Err:  "",
			ID:   data,
		}, nil
	}
}

// makeDelRefundinfoEndpoint ...
func makeDelRefundinfoEndpoint(s *operating.Orderoperating) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(*proto.DelRefundReq)
		err := s.DelRefundinfo("1", req.ID)
		if err != nil {
			return &proto.DelRefundResp{
				Code: 500,
				Err:  err.Error(),
			}, nil
		}
		return &proto.DelRefundResp{
			Code: 200,
			Err:  "",
		}, nil
	}
}

func makeRefundSuccessEndpoint(s *operating.Orderoperating) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(*proto.RefundSuccessReq)
		err := s.RefundSuccess("1", req.ID, req.Price, req.Way, req.Number, req.Serial)
		if err != nil {
			return &proto.RefundSuccessResp{
				Code: 500,
				Err:  err.Error(),
			}, nil
		}
		return &proto.RefundSuccessResp{
			Code: 200,
			Err:  "",
		}, nil
	}
}
