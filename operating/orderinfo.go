package operating

import (
	"errors"
	"fmt"
	"time"

	"../proto"
)

// GetOrderLists 订单列表
func (operat *Orderoperating) GetOrderLists(uid string, offser, limit, id, cid, status, pid, startTime, endTime int64) (RetResponse, error) {
	orderinfos := make([]*proto.OrderinfoInfo, 0)
	if limit == 0 {
		limit = 10
	}
	query := fmt.Sprintf("where uid = '%s' and del = 1 ", uid)
	if id != 0 {
		query = fmt.Sprintf(" and id = %d ", id)
	} else if pid != 0 {
		query = fmt.Sprintf(" and id = %d ", pid)
	} else {
		if cid != 0 {
			query = fmt.Sprintf(" %s and cid = %d ", query, cid)
		}
		if status != 0 {
			query = fmt.Sprintf(" %s and status = %d ", query, status)
		}
		if startTime != 0 {
			query = fmt.Sprintf(" %s createtime >= %d ", query, startTime)
		}
		if endTime != 0 {
			query = fmt.Sprintf(" %s createtime <= %d ", query, endTime)
		}
	}
	err := operat.db.Select(&orderinfos, fmt.Sprintf("select id, cid, pid, cprice, priceremark, createtime, status from orderinfo %s order by id asc  offset %d limit %d", query, offser, limit))
	if err != nil {
		return RetResponse{}, err
	}

	allNum := 0
	err = operat.db.Get(&allNum, fmt.Sprintf("select count(*) from orderinfo %s", query))
	if err != nil {
		return RetResponse{}, err
	}

	return RetResponse{
		AllNum: allNum,
		Data:   orderinfos,
	}, nil
}

// CreateOrder 创建订单
func (operat *Orderoperating) CreateOrder(uid string, cid, newprice int64) (int64, error) {
	if cid < 1 {
		return 0, errors.New("order creation failed:Parameter error")
	}
	// 事务执行
	tx, err := operat.db.Begin()
	if err != nil {
		return 0, err
	}

	// 查询 确认价格无误
	query := fmt.Sprintf("select priceremark from commodity where id = %d and newprice = %d and status = 2", cid, newprice)
	rows, err := tx.Query(query)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	var priceremark string
	if !rows.Next() {
		tx.Rollback()
		return 0, errors.New("do not find the commodity")
	}
	err = rows.Scan(&priceremark)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	rows.Close() // ....

	// 插入订单
	stmt, err := tx.Prepare("INSERT INTO orderinfo (uid, cid, pid, cprice, priceremark, createtime, updatetime,status, del, details) VALUES($1, $2, 0, $3, $4, $5, 0, $6, $7, '')")
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(uid, cid, newprice, priceremark, time.Now().Unix(), 1, 1)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	innum, err := result.RowsAffected()
	if err != nil || innum != 1 {
		tx.Rollback()
		if innum != 1 {
			return 0, errors.New("insert error")
		}
		return 0, err
	}
	query = fmt.Sprintf("select id from orderinfo where uid = '%s' order by id desc limit 1", uid)
	rowsq, err := tx.Query(query)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	var id int64
	if !rowsq.Next() {
		return 0, errors.New("insert error")
	}
	err = rowsq.Scan(&id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	rowsq.Close() // ....

	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	// 返回订单id
	return id, nil
}

// DelOrder 删除订单
func (operat *Orderoperating) DelOrder(uid string, id int64) error {
	_, err := operat.db.Exec("update orderinfo set del =2, updatetime =$1 where id =$2 and uid =$3", time.Now().Unix(), id, uid)
	if err != nil {
		return err
	}
	return nil
}
