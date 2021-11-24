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
	bookInventoryFieldNames          = builderx.RawFieldNames(&BookInventory{})
	bookInventoryRows                = strings.Join(bookInventoryFieldNames, ",")
	bookInventoryRowsExpectAutoSet   = strings.Join(stringx.Remove(bookInventoryFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	bookInventoryRowsWithPlaceHolder = strings.Join(stringx.Remove(bookInventoryFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheBsBooksBookInventoryIdPrefix     = "cache:bsBooks:bookInventory:id:"
	cacheBsBooksBookInventoryBookIdPrefix = "cache:bsBooks:bookInventory:bookId:"
)

type (
	BookInventoryModel interface {
		Insert(data BookInventory) (sql.Result, error)
		FindOne(id int64) (*BookInventory, error)
		FindOneByBookId(bookId int64) (*BookInventory, error)
		Update(data BookInventory) error
		Delete(id int64) error
		FindAll() ([]*BookInventory, error)
	}

	defaultBookInventoryModel struct {
		sqlc.CachedConn
		table string
	}

	BookInventory struct {
		Id        int64 `db:"id"`
		BookId    int64 `db:"book_id"`   // 书籍id
		Inventory int64 `db:"inventory"` // 库存量
	}
)

func NewBookInventoryModel(conn sqlx.SqlConn, c cache.CacheConf) BookInventoryModel {
	return &defaultBookInventoryModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`book_inventory`",
	}
}

func (m *defaultBookInventoryModel) Insert(data BookInventory) (sql.Result, error) {
	bsBooksBookInventoryBookIdKey := fmt.Sprintf("%s%v", cacheBsBooksBookInventoryBookIdPrefix, data.BookId)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, bookInventoryRowsExpectAutoSet)
		return conn.Exec(query, data.BookId, data.Inventory)
	}, bsBooksBookInventoryBookIdKey)
	return ret, err
}

func (m *defaultBookInventoryModel) FindOne(id int64) (*BookInventory, error) {
	bsBooksBookInventoryIdKey := fmt.Sprintf("%s%v", cacheBsBooksBookInventoryIdPrefix, id)
	var resp BookInventory
	err := m.QueryRow(&resp, bsBooksBookInventoryIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", bookInventoryRows, m.table)
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

func (m *defaultBookInventoryModel) FindOneByBookId(bookId int64) (*BookInventory, error) {
	bsBooksBookInventoryBookIdKey := fmt.Sprintf("%s%v", cacheBsBooksBookInventoryBookIdPrefix, bookId)
	var resp BookInventory
	err := m.QueryRowIndex(&resp, bsBooksBookInventoryBookIdKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `book_id` = ? limit 1", bookInventoryRows, m.table)
		if err := conn.QueryRow(&resp, query, bookId); err != nil {
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

func (m *defaultBookInventoryModel) Update(data BookInventory) error {
	bsBooksBookInventoryIdKey := fmt.Sprintf("%s%v", cacheBsBooksBookInventoryIdPrefix, data.Id)
	bsBooksBookInventoryBookIdKey := fmt.Sprintf("%s%v", cacheBsBooksBookInventoryBookIdPrefix, data.BookId)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, bookInventoryRowsWithPlaceHolder)
		return conn.Exec(query, data.BookId, data.Inventory, data.Id)
	}, bsBooksBookInventoryIdKey, bsBooksBookInventoryBookIdKey)
	return err
}

func (m *defaultBookInventoryModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	bsBooksBookInventoryBookIdKey := fmt.Sprintf("%s%v", cacheBsBooksBookInventoryBookIdPrefix, data.BookId)
	bsBooksBookInventoryIdKey := fmt.Sprintf("%s%v", cacheBsBooksBookInventoryIdPrefix, id)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, bsBooksBookInventoryBookIdKey, bsBooksBookInventoryIdKey)
	return err
}

func (m *defaultBookInventoryModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheBsBooksBookInventoryIdPrefix, primary)
}

func (m *defaultBookInventoryModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", bookInventoryRows, m.table)
	return conn.QueryRow(v, query, primary)
}
