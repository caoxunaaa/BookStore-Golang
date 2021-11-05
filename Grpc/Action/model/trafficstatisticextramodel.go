package model

import "fmt"

func (m *defaultTrafficStatisticModel) FindAll() ([]*TrafficStatistic, error) {
	query := fmt.Sprintf("select %s from %s", trafficStatisticRows, m.table)
	var v = make([]*TrafficStatistic, 0)
	err := m.QueryRowsNoCache(&v, query)
	if err != nil {
		return nil, err
	}
	return v, err
}
