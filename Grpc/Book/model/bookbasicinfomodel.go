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
	bookBasicInfoFieldNames          = builderx.RawFieldNames(&BookBasicInfo{})
	bookBasicInfoRows                = strings.Join(bookBasicInfoFieldNames, ",")
	bookBasicInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(bookBasicInfoFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	bookBasicInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(bookBasicInfoFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheBsBooksBookBasicInfoIdPrefix = "cache:bsBooks:bookBasicInfo:id:"
)

type (
	BookBasicInfoModel interface {
		Insert(data BookBasicInfo) (sql.Result, error)
		FindOne(id int64) (*BookBasicInfo, error)
		Update(data BookBasicInfo) error
		Delete(id int64) error
		FindAll() ([]*BookBasicInfo, error)
		FindBooksSortedByMonth(year, month int64) ([]*BookBasicInfo, error)
		FindBooksByLikeName(name string) ([]*BookBasicInfo, error)
	}

	defaultBookBasicInfoModel struct {
		sqlc.CachedConn
		table string
	}

	BookBasicInfo struct {
		Id          int64        `db:"id"`
		Name        string       `db:"name"`         // 书籍名称
		Author      string       `db:"author"`       // 作者
		Image       string       `db:"image"`        // 书籍图片的路径
		StorageTime sql.NullTime `db:"storage_time"` // 入库时间
	}
)

func NewBookBasicInfoModel(conn sqlx.SqlConn, c cache.CacheConf) BookBasicInfoModel {
	return &defaultBookBasicInfoModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`book_basic_info`",
	}
}

func (m *defaultBookBasicInfoModel) Insert(data BookBasicInfo) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, bookBasicInfoRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Name, data.Author, data.Image, data.StorageTime)

	return ret, err
}

func (m *defaultBookBasicInfoModel) FindOne(id int64) (*BookBasicInfo, error) {
	bsBooksBookBasicInfoIdKey := fmt.Sprintf("%s%v", cacheBsBooksBookBasicInfoIdPrefix, id)
	var resp BookBasicInfo
	err := m.QueryRow(&resp, bsBooksBookBasicInfoIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", bookBasicInfoRows, m.table)
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

func (m *defaultBookBasicInfoModel) Update(data BookBasicInfo) error {
	bsBooksBookBasicInfoIdKey := fmt.Sprintf("%s%v", cacheBsBooksBookBasicInfoIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, bookBasicInfoRowsWithPlaceHolder)
		return conn.Exec(query, data.Name, data.Author, data.Image, data.StorageTime, data.Id)
	}, bsBooksBookBasicInfoIdKey)
	return err
}

func (m *defaultBookBasicInfoModel) Delete(id int64) error {

	bsBooksBookBasicInfoIdKey := fmt.Sprintf("%s%v", cacheBsBooksBookBasicInfoIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, bsBooksBookBasicInfoIdKey)
	return err
}

func (m *defaultBookBasicInfoModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheBsBooksBookBasicInfoIdPrefix, primary)
}

func (m *defaultBookBasicInfoModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", bookBasicInfoRows, m.table)
	return conn.QueryRow(v, query, primary)
}
