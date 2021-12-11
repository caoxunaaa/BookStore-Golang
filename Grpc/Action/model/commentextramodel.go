package model

import "fmt"

func (m *defaultCommentModel) FindCommentsByBookContentId(contentId int64) ([]*Comment, error) {
	query := fmt.Sprintf("select %s from %s where book_content_id=?", commentRows, m.table)
	var v = make([]*Comment, 0)
	err := m.conn.QueryRows(&v, query, contentId)
	if err != nil {
		return nil, err
	}
	return v, err
}
