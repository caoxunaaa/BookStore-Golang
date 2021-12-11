package model

import "fmt"

func (m *defaultUserInfoModel) FindAll() ([]*UserInfo, error) {
	query := fmt.Sprintf("select %s from %s", userInfoRows, m.table)
	var v = make([]*UserInfo, 0)
	err := m.QueryRowsNoCache(&v, query)
	if err != nil {
		return nil, err
	}
	return v, err
}
