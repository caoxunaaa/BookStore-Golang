package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	orderInfoFieldNames          = builderx.RawFieldNames(&OrderInfo{})
	orderInfoRows                = strings.Join(orderInfoFieldNames, ",")
	orderInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(orderInfoFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	orderInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(orderInfoFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheBsOrderOrderInfoIdPrefix       = "cache:bsOrder:orderInfo:id:"
	cacheBsOrderOrderInfoOrderNumPrefix = "cache:bsOrder:orderInfo:orderNum:"
)

type (
	OrderInfoModel interface {
		Insert(data OrderInfo) (sql.Result, error)
		FindOne(id int64) (*OrderInfo, error)
		FindOneByOrderNum(orderNum string) (*OrderInfo, error)
		Update(data OrderInfo) error
		Delete(id int64) error
	}

	defaultOrderInfoModel struct {
		sqlc.CachedConn
		table string
	}

	OrderInfo struct {
		Id          int64     `db:"id"`
		BuyerId     int64     `db:"buyer_id"`     // 购买者id
		OrderNum    string    `db:"order_num"`    // 订单号
		CreateTime  time.Time `db:"create_time"`  // 订单创建时间
		Cost        float64   `db:"cost"`         // 费用
		IsPaid      int64     `db:"is_paid"`      // 是否支付
		OrderStatus string    `db:"order_status"` // 订单状态：待付款，关闭
	}
)

func NewOrderInfoModel(conn sqlx.SqlConn, c cache.CacheConf) OrderInfoModel {
	return &defaultOrderInfoModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`order_info`",
	}
}

func (m *defaultOrderInfoModel) Insert(data OrderInfo) (sql.Result, error) {
	bsOrderOrderInfoOrderNumKey := fmt.Sprintf("%s%v", cacheBsOrderOrderInfoOrderNumPrefix, data.OrderNum)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, orderInfoRowsExpectAutoSet)
		return conn.Exec(query, data.BuyerId, data.OrderNum, data.Cost, data.IsPaid, data.OrderStatus)
	}, bsOrderOrderInfoOrderNumKey)
	return ret, err
}

func (m *defaultOrderInfoModel) FindOne(id int64) (*OrderInfo, error) {
	bsOrderOrderInfoIdKey := fmt.Sprintf("%s%v", cacheBsOrderOrderInfoIdPrefix, id)
	var resp OrderInfo
	err := m.QueryRow(&resp, bsOrderOrderInfoIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", orderInfoRows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultOrderInfoModel) FindOneByOrderNum(orderNum string) (*OrderInfo, error) {
	bsOrderOrderInfoOrderNumKey := fmt.Sprintf("%s%v", cacheBsOrderOrderInfoOrderNumPrefix, orderNum)
	var resp OrderInfo
	err := m.QueryRowIndex(&resp, bsOrderOrderInfoOrderNumKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `order_num` = ? limit 1", orderInfoRows, m.table)
		if err := conn.QueryRow(&resp, query, orderNum); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultOrderInfoModel) Update(data OrderInfo) error {
	bsOrderOrderInfoIdKey := fmt.Sprintf("%s%v", cacheBsOrderOrderInfoIdPrefix, data.Id)
	bsOrderOrderInfoOrderNumKey := fmt.Sprintf("%s%v", cacheBsOrderOrderInfoOrderNumPrefix, data.OrderNum)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, orderInfoRowsWithPlaceHolder)
		return conn.Exec(query, data.BuyerId, data.OrderNum, data.Cost, data.IsPaid, data.OrderStatus, data.Id)
	}, bsOrderOrderInfoIdKey, bsOrderOrderInfoOrderNumKey)
	return err
}

func (m *defaultOrderInfoModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	bsOrderOrderInfoIdKey := fmt.Sprintf("%s%v", cacheBsOrderOrderInfoIdPrefix, id)
	bsOrderOrderInfoOrderNumKey := fmt.Sprintf("%s%v", cacheBsOrderOrderInfoOrderNumPrefix, data.OrderNum)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, bsOrderOrderInfoIdKey, bsOrderOrderInfoOrderNumKey)
	return err
}

func (m *defaultOrderInfoModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheBsOrderOrderInfoIdPrefix, primary)
}

func (m *defaultOrderInfoModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", orderInfoRows, m.table)
	return conn.QueryRow(v, query, primary)
}
