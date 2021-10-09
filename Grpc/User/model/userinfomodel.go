package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	userInfoFieldNames          = builderx.RawFieldNames(&UserInfo{})
	userInfoRows                = strings.Join(userInfoFieldNames, ",")
	userInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(userInfoFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	userInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(userInfoFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheBsUserUserInfoIdPrefix       = "cache:bsUser:userInfo:id:"
	cacheBsUserUserInfoEmailPrefix    = "cache:bsUser:userInfo:email:"
	cacheBsUserUserInfoNicknamePrefix = "cache:bsUser:userInfo:nickname:"
	cacheBsUserUserInfoPhonePrefix    = "cache:bsUser:userInfo:phone:"
	cacheBsUserUserInfoUsernamePrefix = "cache:bsUser:userInfo:username:"
)

type (
	UserInfoModel interface {
		Insert(data UserInfo) (sql.Result, error)
		FindOne(id int64) (*UserInfo, error)
		FindOneByEmail(email string) (*UserInfo, error)
		FindOneByNickname(nickname string) (*UserInfo, error)
		FindOneByPhone(phone string) (*UserInfo, error)
		FindOneByUsername(username string) (*UserInfo, error)
		Update(data UserInfo) error
		Delete(id int64) error
		FindAll() ([]*UserInfo, error)
	}

	defaultUserInfoModel struct {
		sqlc.CachedConn
		table string
	}

	UserInfo struct {
		Id       int64  `db:"id"`
		Username string `db:"username"`
		Password string `db:"password"`
		Nickname string `db:"nickname"`
		Phone    string `db:"phone"`
		Email    string `db:"email"`
	}
)

func NewUserInfoModel(conn sqlx.SqlConn, c cache.CacheConf) UserInfoModel {
	return &defaultUserInfoModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user_info`",
	}
}

func (m *defaultUserInfoModel) Insert(data UserInfo) (sql.Result, error) {
	bsUserUserInfoEmailKey := fmt.Sprintf("%s%v", cacheBsUserUserInfoEmailPrefix, data.Email)
	bsUserUserInfoNicknameKey := fmt.Sprintf("%s%v", cacheBsUserUserInfoNicknamePrefix, data.Nickname)
	bsUserUserInfoPhoneKey := fmt.Sprintf("%s%v", cacheBsUserUserInfoPhonePrefix, data.Phone)
	bsUserUserInfoUsernameKey := fmt.Sprintf("%s%v", cacheBsUserUserInfoUsernamePrefix, data.Username)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, userInfoRowsExpectAutoSet)
		return conn.Exec(query, data.Username, data.Password, data.Nickname, data.Phone, data.Email)
	}, bsUserUserInfoEmailKey, bsUserUserInfoNicknameKey, bsUserUserInfoPhoneKey, bsUserUserInfoUsernameKey)
	return ret, err
}

func (m *defaultUserInfoModel) FindOne(id int64) (*UserInfo, error) {
	bsUserUserInfoIdKey := fmt.Sprintf("%s%v", cacheBsUserUserInfoIdPrefix, id)
	var resp UserInfo
	err := m.QueryRow(&resp, bsUserUserInfoIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userInfoRows, m.table)
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

func (m *defaultUserInfoModel) FindOneByEmail(email string) (*UserInfo, error) {
	bsUserUserInfoEmailKey := fmt.Sprintf("%s%v", cacheBsUserUserInfoEmailPrefix, email)
	var resp UserInfo
	err := m.QueryRowIndex(&resp, bsUserUserInfoEmailKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `email` = ? limit 1", userInfoRows, m.table)
		if err := conn.QueryRow(&resp, query, email); err != nil {
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

func (m *defaultUserInfoModel) FindOneByNickname(nickname string) (*UserInfo, error) {
	bsUserUserInfoNicknameKey := fmt.Sprintf("%s%v", cacheBsUserUserInfoNicknamePrefix, nickname)
	var resp UserInfo
	err := m.QueryRowIndex(&resp, bsUserUserInfoNicknameKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `nickname` = ? limit 1", userInfoRows, m.table)
		if err := conn.QueryRow(&resp, query, nickname); err != nil {
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

func (m *defaultUserInfoModel) FindOneByPhone(phone string) (*UserInfo, error) {
	bsUserUserInfoPhoneKey := fmt.Sprintf("%s%v", cacheBsUserUserInfoPhonePrefix, phone)
	var resp UserInfo
	err := m.QueryRowIndex(&resp, bsUserUserInfoPhoneKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `phone` = ? limit 1", userInfoRows, m.table)
		if err := conn.QueryRow(&resp, query, phone); err != nil {
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

func (m *defaultUserInfoModel) FindOneByUsername(username string) (*UserInfo, error) {
	bsUserUserInfoUsernameKey := fmt.Sprintf("%s%v", cacheBsUserUserInfoUsernamePrefix, username)
	var resp UserInfo
	err := m.QueryRowIndex(&resp, bsUserUserInfoUsernameKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `username` = ? limit 1", userInfoRows, m.table)
		if err := conn.QueryRow(&resp, query, username); err != nil {
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

func (m *defaultUserInfoModel) Update(data UserInfo) error {
	bsUserUserInfoIdKey := fmt.Sprintf("%s%v", cacheBsUserUserInfoIdPrefix, data.Id)
	bsUserUserInfoEmailKey := fmt.Sprintf("%s%v", cacheBsUserUserInfoEmailPrefix, data.Email)
	bsUserUserInfoNicknameKey := fmt.Sprintf("%s%v", cacheBsUserUserInfoNicknamePrefix, data.Nickname)
	bsUserUserInfoPhoneKey := fmt.Sprintf("%s%v", cacheBsUserUserInfoPhonePrefix, data.Phone)
	bsUserUserInfoUsernameKey := fmt.Sprintf("%s%v", cacheBsUserUserInfoUsernamePrefix, data.Username)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userInfoRowsWithPlaceHolder)
		return conn.Exec(query, data.Username, data.Password, data.Nickname, data.Phone, data.Email, data.Id)
	}, bsUserUserInfoIdKey, bsUserUserInfoEmailKey, bsUserUserInfoNicknameKey, bsUserUserInfoPhoneKey, bsUserUserInfoUsernameKey)
	return err
}

func (m *defaultUserInfoModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	bsUserUserInfoIdKey := fmt.Sprintf("%s%v", cacheBsUserUserInfoIdPrefix, id)
	bsUserUserInfoEmailKey := fmt.Sprintf("%s%v", cacheBsUserUserInfoEmailPrefix, data.Email)
	bsUserUserInfoNicknameKey := fmt.Sprintf("%s%v", cacheBsUserUserInfoNicknamePrefix, data.Nickname)
	bsUserUserInfoPhoneKey := fmt.Sprintf("%s%v", cacheBsUserUserInfoPhonePrefix, data.Phone)
	bsUserUserInfoUsernameKey := fmt.Sprintf("%s%v", cacheBsUserUserInfoUsernamePrefix, data.Username)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, bsUserUserInfoEmailKey, bsUserUserInfoNicknameKey, bsUserUserInfoPhoneKey, bsUserUserInfoUsernameKey, bsUserUserInfoIdKey)
	return err
}

func (m *defaultUserInfoModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheBsUserUserInfoIdPrefix, primary)
}

func (m *defaultUserInfoModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userInfoRows, m.table)
	return conn.QueryRow(v, query, primary)
}
