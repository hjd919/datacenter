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
	appTaizhangFieldNames          = builderx.RawFieldNames(&AppTaizhang{})
	appTaizhangRows                = strings.Join(appTaizhangFieldNames, ",")
	appTaizhangRowsExpectAutoSet   = strings.Join(stringx.Remove(appTaizhangFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	appTaizhangRowsWithPlaceHolder = strings.Join(stringx.Remove(appTaizhangFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAppTaizhangIdPrefix = "cache#appTaizhang#id#"
)

type (
	AppTaizhangModel interface {
		Insert(data AppTaizhang) (sql.Result, error)
		FindOne(id int64) (*AppTaizhang, error)
		Update(data AppTaizhang) error
		Delete(id int64) error
	}

	defaultAppTaizhangModel struct {
		sqlc.CachedConn
		table string
	}

	AppTaizhang struct {
		Id        int64          `db:"id"`
		Beid      int64          `db:"beid"`      // 对应的平台
		Ptyid     int64          `db:"ptyid"`     // 平台id: 1.微信公众号，2.微信小程序，3.支付宝
		Appid     string         `db:"appid"`     // appid
		Appsecret string         `db:"appsecret"` // 配置密钥
		Title     string         `db:"title"`     // 社交描述
		CreateBy  sql.NullString `db:"create_by"`
		UpdateBy  sql.NullString `db:"update_by"`
		CreatedAt sql.NullTime   `db:"created_at"`
		UpdatedAt sql.NullTime   `db:"updated_at"`
		DeletedAt sql.NullTime   `db:"deleted_at"`
	}
)

func NewAppTaizhangModel(conn sqlx.SqlConn, c cache.CacheConf) AppTaizhangModel {
	return &defaultAppTaizhangModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`app_taizhang`",
	}
}

func (m *defaultAppTaizhangModel) Insert(data AppTaizhang) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, appTaizhangRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Beid, data.Ptyid, data.Appid, data.Appsecret, data.Title, data.CreateBy, data.UpdateBy, data.CreatedAt, data.UpdatedAt, data.DeletedAt)

	return ret, err
}

func (m *defaultAppTaizhangModel) FindOne(id int64) (*AppTaizhang, error) {
	appTaizhangIdKey := fmt.Sprintf("%s%v", cacheAppTaizhangIdPrefix, id)
	var resp AppTaizhang
	err := m.QueryRow(&resp, appTaizhangIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", appTaizhangRows, m.table)
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

func (m *defaultAppTaizhangModel) Update(data AppTaizhang) error {
	appTaizhangIdKey := fmt.Sprintf("%s%v", cacheAppTaizhangIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, appTaizhangRowsWithPlaceHolder)
		return conn.Exec(query, data.Beid, data.Ptyid, data.Appid, data.Appsecret, data.Title, data.CreateBy, data.UpdateBy, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.Id)
	}, appTaizhangIdKey)
	return err
}

func (m *defaultAppTaizhangModel) Delete(id int64) error {

	appTaizhangIdKey := fmt.Sprintf("%s%v", cacheAppTaizhangIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, appTaizhangIdKey)
	return err
}

func (m *defaultAppTaizhangModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAppTaizhangIdPrefix, primary)
}

func (m *defaultAppTaizhangModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", appTaizhangRows, m.table)
	return conn.QueryRow(v, query, primary)
}
