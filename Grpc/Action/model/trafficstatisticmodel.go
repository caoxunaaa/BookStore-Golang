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
	trafficStatisticFieldNames          = builderx.RawFieldNames(&TrafficStatistic{})
	trafficStatisticRows                = strings.Join(trafficStatisticFieldNames, ",")
	trafficStatisticRowsExpectAutoSet   = strings.Join(stringx.Remove(trafficStatisticFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	trafficStatisticRowsWithPlaceHolder = strings.Join(stringx.Remove(trafficStatisticFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheBsActionTrafficStatisticIdPrefix               = "cache:bsAction:trafficStatistic:id:"
	cacheBsActionTrafficStatisticBookIdChapterNumPrefix = "cache:bsAction:trafficStatistic:bookId:chapterNum:"
)

type (
	TrafficStatisticModel interface {
		Insert(data TrafficStatistic) (sql.Result, error)
		FindOne(id int64) (*TrafficStatistic, error)
		FindOneByBookIdChapterNum(bookId int64, chapterNum int64) (*TrafficStatistic, error)
		Update(data TrafficStatistic) error
		Delete(id int64) error
		FindAll() ([]*TrafficStatistic, error)
	}

	defaultTrafficStatisticModel struct {
		sqlc.CachedConn
		table string
	}

	TrafficStatistic struct {
		Id            int64 `db:"id"`
		BookId        int64 `db:"book_id"`
		ChapterNum    int64 `db:"chapter_num"`
		TrafficNumber int64 `db:"traffic_number"`
	}
)

func NewTrafficStatisticModel(conn sqlx.SqlConn, c cache.CacheConf) TrafficStatisticModel {
	return &defaultTrafficStatisticModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`traffic_statistic`",
	}
}

func (m *defaultTrafficStatisticModel) Insert(data TrafficStatistic) (sql.Result, error) {
	bsActionTrafficStatisticBookIdChapterNumKey := fmt.Sprintf("%s%v:%v", cacheBsActionTrafficStatisticBookIdChapterNumPrefix, data.BookId, data.ChapterNum)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, trafficStatisticRowsExpectAutoSet)
		return conn.Exec(query, data.BookId, data.ChapterNum, data.TrafficNumber)
	}, bsActionTrafficStatisticBookIdChapterNumKey)
	return ret, err
}

func (m *defaultTrafficStatisticModel) FindOne(id int64) (*TrafficStatistic, error) {
	bsActionTrafficStatisticIdKey := fmt.Sprintf("%s%v", cacheBsActionTrafficStatisticIdPrefix, id)
	var resp TrafficStatistic
	err := m.QueryRow(&resp, bsActionTrafficStatisticIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", trafficStatisticRows, m.table)
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

func (m *defaultTrafficStatisticModel) FindOneByBookIdChapterNum(bookId int64, chapterNum int64) (*TrafficStatistic, error) {
	bsActionTrafficStatisticBookIdChapterNumKey := fmt.Sprintf("%s%v:%v", cacheBsActionTrafficStatisticBookIdChapterNumPrefix, bookId, chapterNum)
	var resp TrafficStatistic
	err := m.QueryRowIndex(&resp, bsActionTrafficStatisticBookIdChapterNumKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `book_id` = ? and `chapter_num` = ? limit 1", trafficStatisticRows, m.table)
		if err := conn.QueryRow(&resp, query, bookId, chapterNum); err != nil {
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

func (m *defaultTrafficStatisticModel) Update(data TrafficStatistic) error {
	bsActionTrafficStatisticIdKey := fmt.Sprintf("%s%v", cacheBsActionTrafficStatisticIdPrefix, data.Id)
	bsActionTrafficStatisticBookIdChapterNumKey := fmt.Sprintf("%s%v:%v", cacheBsActionTrafficStatisticBookIdChapterNumPrefix, data.BookId, data.ChapterNum)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, trafficStatisticRowsWithPlaceHolder)
		return conn.Exec(query, data.BookId, data.ChapterNum, data.TrafficNumber, data.Id)
	}, bsActionTrafficStatisticIdKey, bsActionTrafficStatisticBookIdChapterNumKey)
	return err
}

func (m *defaultTrafficStatisticModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	bsActionTrafficStatisticBookIdChapterNumKey := fmt.Sprintf("%s%v:%v", cacheBsActionTrafficStatisticBookIdChapterNumPrefix, data.BookId, data.ChapterNum)
	bsActionTrafficStatisticIdKey := fmt.Sprintf("%s%v", cacheBsActionTrafficStatisticIdPrefix, id)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, bsActionTrafficStatisticBookIdChapterNumKey, bsActionTrafficStatisticIdKey)
	return err
}

func (m *defaultTrafficStatisticModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheBsActionTrafficStatisticIdPrefix, primary)
}

func (m *defaultTrafficStatisticModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", trafficStatisticRows, m.table)
	return conn.QueryRow(v, query, primary)
}
