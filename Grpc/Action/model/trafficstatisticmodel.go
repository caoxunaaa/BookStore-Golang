package model

import (
	"database/sql"
	"fmt"
	"strings"

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
		conn  sqlx.SqlConn
		table string
	}

	TrafficStatistic struct {
		Id            int64 `db:"id"`
		BookId        int64 `db:"book_id"`
		ChapterNum    int64 `db:"chapter_num"`
		TrafficNumber int64 `db:"traffic_number"`
	}
)

func NewTrafficStatisticModel(conn sqlx.SqlConn) TrafficStatisticModel {
	return &defaultTrafficStatisticModel{
		conn:  conn,
		table: "`traffic_statistic`",
	}
}

func (m *defaultTrafficStatisticModel) Insert(data TrafficStatistic) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, trafficStatisticRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.BookId, data.ChapterNum, data.TrafficNumber)
	return ret, err
}

func (m *defaultTrafficStatisticModel) FindOne(id int64) (*TrafficStatistic, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", trafficStatisticRows, m.table)
	var resp TrafficStatistic
	err := m.conn.QueryRow(&resp, query, id)
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
	var resp TrafficStatistic
	query := fmt.Sprintf("select %s from %s where `book_id` = ? and `chapter_num` = ? limit 1", trafficStatisticRows, m.table)
	err := m.conn.QueryRow(&resp, query, bookId, chapterNum)
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
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, trafficStatisticRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.BookId, data.ChapterNum, data.TrafficNumber, data.Id)
	return err
}

func (m *defaultTrafficStatisticModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
