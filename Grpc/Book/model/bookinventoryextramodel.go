package model

import "fmt"

func (m *defaultBookInventoryModel) FindAll() ([]*BookInventory, error) {
	query := fmt.Sprintf("select %s from %s", bookInventoryRows, m.table)
	var v = make([]*BookInventory, 0)
	err := m.QueryRowsNoCache(&v, query)
	if err != nil {
		return nil, err
	}
	return v, err
}
