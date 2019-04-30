package operating

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"../proto"
)

// GetPayList 付款列表
func (operat *Orderoperating) GetPayList(uid string, offser, limit, id, way, status, startTime, endTime int64) (RetResponse, error) {
	payinfos := make([]*proto.GetPayListInfo, 0)
	if limit == 0 {
		limit = 10
	}
	query := fmt.Sprintf("WHERE uid = '%s' and del = 1 ", uid)
	if id != 0 {
		query = fmt.Sprintf("%s and id = %d ", query, id)
	} else {
		if way != 0 {
			query = fmt.Sprintf("%s and way = %d ", query, way)
		}
		if status != 0 {
			query = fmt.Sprintf("%s and status = %d ", query, status)
		}
		if startTime != 0 {
			query = fmt.Sprintf("%s createtime >= %d ", query, startTime)
		}
		if endTime != 0 {
			query = fmt.Sprintf("%s createtime <= %d ", query, endTime)
		}
	}
	err := operat.db.Select(&payinfos, fmt.Sprintf("select id, price, way, number, status, createtime from pay %s order by id asc offset %d limit %d ", query, offser, limit))
	if err != nil {
		return RetResponse{}, err
	}

	allNum := 0
	err = operat.db.Get(&allNum, fmt.Sprintf("select count(*) from pay %s", query))
	if err != nil {
		return RetResponse{}, err
	}

	for i := range payinfos {
		lens := len(payinfos[i].Number)
		if lens > 5 {
			payinfos[i].Number = payinfos[i].Number[0:3] + "******" + payinfos[i].Number[lens-3:]
		}
	}

	return RetResponse{
		AllNum: allNum,
		Data:   payinfos,
	}, nil
}

// CreatePayInfo 创建支付订单
func (operat *Orderoperating) CreatePayInfo(uid string, cids []int64, price int64) (int64, error) {
	if len(cids) == 0 || price < 0 {
		return 0, errors.New("pay creation failed:Parameter error")
	}
	tx, err := operat.db.Begin()
	if err != nil {
		return 0, err
	}

	//创建支付订单
	stmt, err := tx.Prepare("insert into pay (uid, price, retprice, way, number, serial, status, del, createtime, updatetime, details) values($1, $2, 0, 0, '', '', 1, 1, $3, 0, '')")
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	defer stmt.Close()
	result1, err := stmt.Exec(uid, price, time.Now().Unix())
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	numI, err := result1.RowsAffected()
	if err != nil || numI != 1 {
		tx.Rollback()
		if numI != 1 {
			return 0, errors.New("insert error")
		}
		return 0, err
	}

	//查询刚刚创建的支付订单id
	rowsq, err := tx.Query(fmt.Sprintf("select id from pay where uid = '%s' and price = %d and status = 1 order by id desc limit 1", uid, price))
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	var id int64
	if !rowsq.Next() {
		return 0, errors.New("insert error")
	}
	err = rowsq.Scan(&id)
	rowsq.Close()

	//查询订单是否正确并改变订单状态
	idsString := ""
	for _, v := range cids {
		idsString = idsString + "," + strconv.FormatInt(v, 10)
	}
	idsString = idsString[1:]
	query := fmt.Sprintf("update orderinfo set status = 2, updatetime = %d, pid = %d where uid = '%s' and status = 1 and del =1 and id in (%s) and (select sum(cprice) from orderinfo where id in (%s)) = %d", time.Now().Unix(), id, uid, idsString, idsString, price)
	result, err := tx.Exec(query)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	num, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if int(num) != len(cids) {
		tx.Rollback()
		return 0, errors.New("create order fail")
	}

	tx.Commit()
	return id, nil
}

// DelPayInfo 删除支付订单
func (operat *Orderoperating) DelPayInfo(uid string, id int64) error {
	if id < 1 {
		return errors.New("pay delete failed:Parameter error")
	}
	// 事务执行
	tx, err := operat.db.Begin()
	if err != nil {
		return err
	}
	//更新支付订单
	delresult, err := tx.Exec(fmt.Sprintf("update pay set del = 2, updatetime = %d where id = %d and uid = '%s' and status != 5 and del = 1", time.Now().Unix(), id, uid))
	if err != nil {
		tx.Rollback()
		return err
	}
	delnum, err := delresult.RowsAffected()
	if err != nil || delnum != 1 {
		tx.Rollback()
		if delnum != 1 {
			return errors.New("pay delete fail")
		}
		return err
	}
	//更新订单 -- 未付款订单更新为付款失败订单
	_, err = tx.Exec(fmt.Sprintf("update orderinfo set status = 4 where uid = '%s' and pid = %d and status = 2", uid, id))
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// PaySuccess 支付成功
func (operat *Orderoperating) PaySuccess(uid string, id, price, way int64, number, serial string) error {
	if id < 0 || price < 0 || serial == "" || number == "" {
		return errors.New("pay error")
	}
	//--------------确保付款流水已经被确认
	// 事务执行
	tx, err := operat.db.Begin()
	if err != nil {
		return err
	}
	//更新付款订单
	payresult, err := tx.Exec(fmt.Sprintf("update pay set status = 2, updatetime = %d, way = %d, number = '%s', serial = '%s' where uid = '%s' and status = 1 and del =1 and price = %d", time.Now().Unix(), way, number, serial, uid, price))
	if err != nil {
		tx.Rollback()
		return err
	}
	paynum, err := payresult.RowsAffected()
	if err != nil || paynum != 1 {
		tx.Rollback()
		if paynum != 1 {
			return errors.New("pay error")
		}
		return err
	}

	//更新订单状态
	orderresult, err := tx.Exec(fmt.Sprintf("update orderinfo set status = 3, updatetime = %d where uid = '%s' and pid = %d and status = 2", time.Now().Unix(), uid, id))
	if err != nil {
		tx.Rollback()
		return err
	}
	ordernum, err := orderresult.RowsAffected()
	if err != nil || ordernum < 1 {
		tx.Rollback()
		if ordernum < 1 {
			return errors.New("pay error")
		}
		return err
	}
	tx.Commit()
	return nil
}
