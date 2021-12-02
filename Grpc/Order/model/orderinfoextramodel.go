package model

import (
	"fmt"
)

func (m *defaultOrderInfoModel) FindNotPaidOrdersByBuyerId(BuyerId int64) (*OrderInfo, error) {
	query := fmt.Sprintf("select %s from %s where buyer_id=%d and order_status='待支付'", orderInfoRows, m.table, BuyerId)
	var resp OrderInfo
	err := m.CachedConn.QueryRowNoCache(&resp, query)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
