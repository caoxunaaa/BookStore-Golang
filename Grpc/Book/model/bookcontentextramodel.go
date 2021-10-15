package model

import "fmt"

func (m *defaultBookContentModel) FindAllBookContentsByBookId(bookId int64) ([]*BookContent, error) {
	query := fmt.Sprintf("select %s from %s where book_id=%d", bookBasicInfoRows, m.table, bookId)
	var v = make([]*BookContent, 0)
	err := m.QueryRowsNoCache(&v, query)
	if err != nil {
		return nil, err
	}
	return v, err
}
