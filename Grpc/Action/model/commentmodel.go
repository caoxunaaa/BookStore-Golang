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
	commentFieldNames          = builderx.RawFieldNames(&Comment{})
	commentRows                = strings.Join(commentFieldNames, ",")
	commentRowsExpectAutoSet   = strings.Join(stringx.Remove(commentFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	commentRowsWithPlaceHolder = strings.Join(stringx.Remove(commentFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	CommentModel interface {
		Insert(data Comment) (sql.Result, error)
		FindOne(id int64) (*Comment, error)
		Update(data Comment) error
		Delete(id int64) error
		FindCommentsByBookContentId(contentId int64) ([]*Comment, error)
	}

	defaultCommentModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Comment struct {
		Id              int64  `db:"id"`
		ParentId        int64  `db:"parent_id"`
		BookContentId   int64  `db:"book_content_id"`
		Comment         string `db:"comment"`
		CommentByUserId int64  `db:"comment_by_user_id"` // 评论者
		CommentToUserId int64  `db:"comment_to_user_id"` // 被评论者
	}
)

func NewCommentModel(conn sqlx.SqlConn) CommentModel {
	return &defaultCommentModel{
		conn:  conn,
		table: "`comment`",
	}
}

func (m *defaultCommentModel) Insert(data Comment) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, commentRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.ParentId, data.BookContentId, data.Comment, data.CommentByUserId, data.CommentToUserId)
	return ret, err
}

func (m *defaultCommentModel) FindOne(id int64) (*Comment, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", commentRows, m.table)
	var resp Comment
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

func (m *defaultCommentModel) Update(data Comment) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, commentRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.ParentId, data.BookContentId, data.Comment, data.CommentByUserId, data.CommentToUserId, data.Id)
	return err
}

func (m *defaultCommentModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
