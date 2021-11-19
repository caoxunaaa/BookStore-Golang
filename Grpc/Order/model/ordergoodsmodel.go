package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	orderGoodsFieldNames          = builderx.RawFieldNames(&OrderGoods{})
	orderGoodsRows                = strings.Join(orderGoodsFieldNames, ",")
	orderGoodsRowsExpectAutoSet   = strings.Join(stringx.Remove(orderGoodsFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	orderGoodsRowsWithPlaceHolder = strings.Join(stringx.Remove(orderGoodsFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheBsOrderOrderGoodsIdPrefix = "cache:bsOrder:orderGoods:id:"
)

type (
	OrderGoodsModel interface {
		Insert(data OrderGoods) (sql.Result, error)
		FindOne(id int64) (*OrderGoods, error)
		Update(data OrderGoods) error
		Delete(id int64) error
	}

	defaultOrderGoodsModel struct {
		sqlc.CachedConn
		table string
	}

	OrderGoods struct {
		Id       int64  `db:"id"`
		OrderNum string `db:"order_num"` // 订单号
		BookId   int64  `db:"book_id"`   // 书籍id
		Count    int64  `db:"count"`     // 数量
	}
)

func NewOrderGoodsModel(conn sqlx.SqlConn, c cache.CacheConf) OrderGoodsModel {
	return &defaultOrderGoodsModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`order_goods`",
	}
}

func (m *defaultOrderGoodsModel) Insert(data OrderGoods) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, orderGoodsRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.OrderNum, data.BookId, data.Count)

	return ret, err
}

func (m *defaultOrderGoodsModel) FindOne(id int64) (*OrderGoods, error) {
	bsOrderOrderGoodsIdKey := fmt.Sprintf("%s%v", cacheBsOrderOrderGoodsIdPrefix, id)
	var resp OrderGoods
	err := m.QueryRow(&resp, bsOrderOrderGoodsIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", orderGoodsRows, m.table)
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

func (m *defaultOrderGoodsModel) Update(data OrderGoods) error {
	bsOrderOrderGoodsIdKey := fmt.Sprintf("%s%v", cacheBsOrderOrderGoodsIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, orderGoodsRowsWithPlaceHolder)
		return conn.Exec(query, data.OrderNum, data.BookId, data.Count, data.Id)
	}, bsOrderOrderGoodsIdKey)
	return err
}

func (m *defaultOrderGoodsModel) Delete(id int64) error {

	bsOrderOrderGoodsIdKey := fmt.Sprintf("%s%v", cacheBsOrderOrderGoodsIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, bsOrderOrderGoodsIdKey)
	return err
}

func (m *defaultOrderGoodsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheBsOrderOrderGoodsIdPrefix, primary)
}

func (m *defaultOrderGoodsModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", orderGoodsRows, m.table)
	return conn.QueryRow(v, query, primary)
}
