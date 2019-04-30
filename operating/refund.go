package operating

import (
	"errors"
	"fmt"
	"time"

	"../proto"
)

// GetRefundInfo 退款订单列表
func (operat *Orderoperating) GetRefundInfo(uid string, offset, limit, id, pid, oid, status, way, startTime, endTime int64) (RetResponse, error) {
	reFuncdInfos := make([]*proto.GetRefundInfo, 0)
	if limit == 0 {
		limit = 10
	}
	query := fmt.Sprintf("where uid = '%s' and del = 1 ", uid)
	if id != 0 {
		query = fmt.Sprintf("%s and id = %d", query, id)
	} else if oid != 0 {
		query = fmt.Sprintf("%s and oid = %d", query, oid)
	} else {
		if pid != 0 {
			query = fmt.Sprintf("%s and pid = %d ", query, pid)
		}
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
	err := operat.db.Select(&reFuncdInfos, fmt.Sprintf("select id, oid, oprice, price, way, number, status, createtime from refund %s offset %d limit %d", query, offset, limit))
	if err != nil {
		return RetResponse{}, err
	}
	for i := range reFuncdInfos {
		lens := len(reFuncdInfos[i].Number)
		if lens > 5 {
			reFuncdInfos[i].Number = reFuncdInfos[i].Number[0:3] + "******" + reFuncdInfos[i].Number[lens-3:]
		}
	}

	allNum := 0
	err = operat.db.Get(&allNum, fmt.Sprintf("select count(*) from refund %s", query))
	if err != nil {
		return RetResponse{}, err
	}
	return RetResponse{
		AllNum: allNum,
		Data:   reFuncdInfos,
	}, nil
}

// CreateReFundInfo 创建退款订单
func (operat *Orderoperating) CreateReFundInfo(uid string, oid, price, way int64, number string) (int64, error) {
	if oid < 1 {
		return 0, errors.New("refund creation failed:Parameter error")
	}
	// 事务执行
	tx, err := operat.db.Begin()
	if err != nil {
		return 0, err
	}
	//查询支付订单是否可以退款
	rows, err := tx.Query(fmt.Sprintf("select cprice, pid from orderinfo where id = %d and status = 3 and cprice >= %d and  uid = '%s'", oid, price, uid))
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	pid := 0
	selectprice := 0
	for rows.Next() {
		err = rows.Scan(&selectprice, &pid)
		if err != nil || pid == 0 || selectprice == 0 {
			tx.Rollback()
			if pid == 0 || selectprice == 0 {
				return 0, errors.New("order error")
			}
			return 0, err
		}
	}
	//查询订单是否可以退款并改变状态
	result, err := tx.Exec(fmt.Sprintf("update orderinfo set status = 5, updatetime = %d where pid = %d and status = 3 and cprice >= %d and uid = '%s'", time.Now().Unix(), pid, price, uid))
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	num, err := result.RowsAffected()
	if err != nil || num != 1 {
		tx.Rollback()
		if num != 1 {
			return 0, errors.New("fail")
		}
		return 0, err
	}
	//创建退款订单
	stmt, err := tx.Prepare("insert into refund (uid, oid, pid, oprice, price, way, number, serial, status, del, createtime, updatetime, details) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)")
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	defer stmt.Close()
	insertresult, err := stmt.Exec(uid, oid, pid, selectprice, price, way, number, "", 1, 1, time.Now().Unix(), 0, "")
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	refundnum, err := insertresult.RowsAffected()
	if err != nil || refundnum != 1 {
		tx.Rollback()
		if refundnum != 1 {
			return 0, errors.New("insert error")
		}
		return 0, err
	}

	//查询刚刚创建的支付订单id
	rowsq, err := tx.Query(fmt.Sprintf("select id from refund where uid = '%s' and price = %d and status = 1 order by id desc limit 1", uid, price))
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

	tx.Commit()
	return id, nil
}

// DelRefundinfo 删除退款订单
func (operat *Orderoperating) DelRefundinfo(uid string, id int64) error {
	if id < 1 {
		return errors.New("refund delete failed:Parameter error")
	}
	// 事务执行
	tx, err := operat.db.Begin()
	if err != nil {
		return err
	}
	// 更新退款订单状态
	delresult, err := tx.Exec(fmt.Sprintf("update refund set del = 2, updatetime = %d where id = %d and uid = '%s' and status != 4", time.Now().Unix(), id, uid))
	if err != nil {
		tx.Rollback()
		return err
	}
	delnum, err := delresult.RowsAffected()
	if err != nil || delnum != 1 {
		tx.Rollback()
		if delnum != 1 {
			return errors.New("refund delete fail")
		}
		return err
	}

	//检索退款订单状态
	rowsq, err := tx.Query(fmt.Sprintf("select oid from refund where id = %d and status = 1", id))
	if err != nil {
		tx.Rollback()
		return err
	}
	var oid int64
	if !rowsq.Next() {
		return errors.New("del error")
	}
	err = rowsq.Scan(&oid)
	rowsq.Close()
	// 如果退款订单处于初始状态, 则更改订单状态
	if oid != 0 {
		//更新订单状态
		_, err = tx.Exec(fmt.Sprintf("update orderinfo set status = 3 where uid = '%s' and status = 5", uid))
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

// RefundSuccess 退款成功
func (operat *Orderoperating) RefundSuccess(uid string, id, price, way int64, number, serial string) error {
	if id < 1 || price < 0 || way < 1 || number == "" || serial == "" {
		return errors.New("refund error")
	}
	// //检索订单信息
	type Info struct {
		PID int
		OID int
	}
	info := Info{}
	err := operat.db.Get(&info, "select pid, oid from refund where id = $1", id)
	if err != nil {
		return err
	}

	// 事务执行
	tx, err := operat.db.Begin()
	if err != nil {
		return err
	}
	//更新退款订单
	refundresult, err := tx.Exec(fmt.Sprintf("update refund set status = 2, updatetime = %d, way = %d, number = '%s', serial = '%s' where id = %d and uid = '%s' and price = %d and status = 1", time.Now().Unix(), way, number, serial, id, uid, price))
	if err != nil {
		tx.Rollback()
		return err
	}
	refundnum, err := refundresult.RowsAffected()
	if err != nil || refundnum != 1 {
		tx.Rollback()
		if refundnum != 1 {
			return errors.New("refund fail")
		}
		return err
	}
	//更新付款订单状态
	payresult, err := tx.Exec(fmt.Sprintf("update pay set status = 4, updatetime = %d, retprice= %d where id = %d and status = 2 and price >= %d + retprice", time.Now().Unix(), price, info.PID, price))
	if err != nil {
		tx.Rollback()
		return err
	}
	payresultnum, err := payresult.RowsAffected()
	if err != nil || payresultnum != 1 {
		tx.Rollback()
		if payresultnum != 1 {
			return errors.New("refund fail")
		}
		return err
	}

	//更新订单状态
	orderresult, err := tx.Exec(fmt.Sprintf("update orderinfo set status = 6, updatetime = %d where id = %d and status = 5 and cprice >= %d", time.Now().Unix(), info.OID, price))
	if err != nil {
		tx.Rollback()
		return err
	}
	ordernum, err := orderresult.RowsAffected()
	if err != nil || ordernum != 1 {
		tx.Rollback()
		if ordernum != 1 {
			return errors.New("refund fail")
		}
		return err
	}
	//提交事务
	tx.Commit()
	return nil
}
