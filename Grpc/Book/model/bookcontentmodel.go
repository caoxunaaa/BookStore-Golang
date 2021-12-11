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
	bookContentFieldNames          = builderx.RawFieldNames(&BookContent{})
	bookContentRows                = strings.Join(bookContentFieldNames, ",")
	bookContentRowsExpectAutoSet   = strings.Join(stringx.Remove(bookContentFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	bookContentRowsWithPlaceHolder = strings.Join(stringx.Remove(bookContentFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheBsBooksBookContentIdPrefix               = "cache:bsBooks:bookContent:id:"
	cacheBsBooksBookContentBookIdChapterNumPrefix = "cache:bsBooks:bookContent:bookId:chapterNum:"
)

type (
	BookContentModel interface {
		Insert(data BookContent) (sql.Result, error)
		FindOne(id int64) (*BookContent, error)
		FindOneByBookIdChapterNum(bookId int64, chapterNum int64) (*BookContent, error)
		Update(data BookContent) error
		Delete(id int64) error
		FindAllBookContentsByBookId(bookId int64) ([]*BookContent, error)
	}

	defaultBookContentModel struct {
		sqlc.CachedConn
		table string
	}

	BookContent struct {
		Id                int64     `db:"id"`
		BookId            int64     `db:"book_id"`             // 书籍的ID，外键
		ChapterNum        int64     `db:"chapter_num"`         // 章节数
		ChapterName       string    `db:"chapter_name"`        // 章节名
		ChapterContent    string    `db:"chapter_content"`     // 章节内容
		CreateContentTime time.Time `db:"create_content_time"` // 创建事件
	}
)

func NewBookContentModel(conn sqlx.SqlConn, c cache.CacheConf) BookContentModel {
	return &defaultBookContentModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`book_content`",
	}
}

func (m *defaultBookContentModel) Insert(data BookContent) (sql.Result, error) {
	bsBooksBookContentBookIdChapterNumKey := fmt.Sprintf("%s%v:%v", cacheBsBooksBookContentBookIdChapterNumPrefix, data.BookId, data.ChapterNum)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, bookContentRowsExpectAutoSet)
		return conn.Exec(query, data.BookId, data.ChapterNum, data.ChapterName, data.ChapterContent, data.CreateContentTime)
	}, bsBooksBookContentBookIdChapterNumKey)
	return ret, err
}

func (m *defaultBookContentModel) FindOne(id int64) (*BookContent, error) {
	bsBooksBookContentIdKey := fmt.Sprintf("%s%v", cacheBsBooksBookContentIdPrefix, id)
	var resp BookContent
	err := m.QueryRow(&resp, bsBooksBookContentIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", bookContentRows, m.table)
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

func (m *defaultBookContentModel) FindOneByBookIdChapterNum(bookId int64, chapterNum int64) (*BookContent, error) {
	bsBooksBookContentBookIdChapterNumKey := fmt.Sprintf("%s%v:%v", cacheBsBooksBookContentBookIdChapterNumPrefix, bookId, chapterNum)
	var resp BookContent
	err := m.QueryRowIndex(&resp, bsBooksBookContentBookIdChapterNumKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `book_id` = ? and `chapter_num` = ? limit 1", bookContentRows, m.table)
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

func (m *defaultBookContentModel) Update(data BookContent) error {
	bsBooksBookContentIdKey := fmt.Sprintf("%s%v", cacheBsBooksBookContentIdPrefix, data.Id)
	bsBooksBookContentBookIdChapterNumKey := fmt.Sprintf("%s%v:%v", cacheBsBooksBookContentBookIdChapterNumPrefix, data.BookId, data.ChapterNum)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, bookContentRowsWithPlaceHolder)
		return conn.Exec(query, data.BookId, data.ChapterNum, data.ChapterName, data.ChapterContent, data.CreateContentTime, data.Id)
	}, bsBooksBookContentIdKey, bsBooksBookContentBookIdChapterNumKey)
	return err
}

func (m *defaultBookContentModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	bsBooksBookContentIdKey := fmt.Sprintf("%s%v", cacheBsBooksBookContentIdPrefix, id)
	bsBooksBookContentBookIdChapterNumKey := fmt.Sprintf("%s%v:%v", cacheBsBooksBookContentBookIdChapterNumPrefix, data.BookId, data.ChapterNum)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, bsBooksBookContentIdKey, bsBooksBookContentBookIdChapterNumKey)
	return err
}

func (m *defaultBookContentModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheBsBooksBookContentIdPrefix, primary)
}

func (m *defaultBookContentModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", bookContentRows, m.table)
	return conn.QueryRow(v, query, primary)
}
