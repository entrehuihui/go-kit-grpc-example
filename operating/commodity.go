package operating

import "../proto"

// CommodityInfo 商品信息
type CommodityInfo struct {
	ID          int
	Name        string
	OldPrice    int
	NewPrice    int
	PriceRemark string
	UpdateTime  int
}

// GetCommodityLists 获取商品信息列表
func (operat *Orderoperating) GetCommodityLists(id, limit, offset int64) (RetResponse, error) {
	// commodityInfos := []CommodityInfo{}
	commodityInfos := make([]*proto.CommodityInfo, 0)
	allNum := 0
	if limit == 0 {
		limit = 10
	}
	if id == 0 {
		err := operat.db.Select(&commodityInfos, "select id, name, oldprice, newprice, priceremark, updatetime from commodity where status = 2 order by id offset $1 limit $2", offset, limit)
		if err != nil {
			return RetResponse{}, err
		}
		err = operat.db.Get(&allNum, "select count(*) from commodity where status = 2")
		if err != nil {
			return RetResponse{}, err
		}
	} else {
		err := operat.db.Select(&commodityInfos, "select id, name, oldprice, newprice, priceremark, updatetime from commodity where id = $1 and status = 2 order by id asc", id)
		if err != nil {
			return RetResponse{}, err
		}
		allNum = 1
	}
	return RetResponse{
		AllNum: allNum,
		Data:   commodityInfos,
	}, nil
}
