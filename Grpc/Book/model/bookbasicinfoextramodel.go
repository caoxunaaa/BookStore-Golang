package model

import "fmt"

func (m *defaultBookBasicInfoModel) FindAll() ([]*BookBasicInfo, error) {
	query := fmt.Sprintf("select %s from %s", bookBasicInfoRows, m.table)
	var v = make([]*BookBasicInfo, 0)
	err := m.QueryRowsNoCache(&v, query)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *defaultBookBasicInfoModel) FindBooksSortedByMonth(year, month int64) ([]*BookBasicInfo, error) {
	query := fmt.Sprintf("select %s from %s where date_format(storage_time, '%%Y-%%m')='%d-%d'", bookBasicInfoRows, m.table, year, month)
	var v = make([]*BookBasicInfo, 0)
	err := m.QueryRowsNoCache(&v, query)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *defaultBookBasicInfoModel) FindBooksByLikeName(name string) ([]*BookBasicInfo, error) {
	query := fmt.Sprintf("select %s from %s where name like '%%%s%%'", bookBasicInfoRows, m.table, name)
	var v = make([]*BookBasicInfo, 0)
	err := m.QueryRowsNoCache(&v, query)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *defaultBookBasicInfoModel) FindBooksByStorageUserId(storageUserId int64) ([]*BookBasicInfo, error) {
	query := fmt.Sprintf("select %s from %s where storage_user_id=%d", bookBasicInfoRows, m.table, storageUserId)
	var v = make([]*BookBasicInfo, 0)
	err := m.QueryRowsNoCache(&v, query)
	if err != nil {
		return nil, err
	}
	return v, err
}
