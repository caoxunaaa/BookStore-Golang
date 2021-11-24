package model

import "fmt"

func (m *defaultOrderGoodsModel) FindGoodsByOrderNum(orderNum string) ([]*OrderGoods, error) {
	query := fmt.Sprintf("select %s from %s where `order_num` = %s", orderGoodsRows, m.table, orderNum)
	var v = make([]*OrderGoods, 0)
	err := m.QueryRowsNoCache(&v, query)
	if err != nil {
		return nil, err
	}
	return v, err
}
