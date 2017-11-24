package dao_test

import (
	"database/sql"
	"database/sql/driver"
	"io/ioutil"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kamilsk/form-api/dao"
	"github.com/kamilsk/form-api/data"
	"github.com/stretchr/testify/assert"
)

const UUID data.UUID = "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11"

func TestNew_WithoutConfiguration(t *testing.T) {
	service, err := dao.New()
	assert.NoError(t, err)
	assert.Panics(t, func() { service.Schema(UUID) })
	assert.Panics(t, func() { service.AddData(UUID, url.Values{}) })
}

func TestNew_WithInvalidConfiguration(t *testing.T) {
	service, err := dao.New(dao.Connection(&url.URL{Scheme: "driver"}))
	assert.Error(t, err)
	assert.Nil(t, service)
}

func TestNew_WithValidConfiguration(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	drv := dao.NewMockDriver(ctrl)
	conn := dao.NewMockConn(ctrl)
	stmt := dao.NewMockStmt(ctrl)
	rows := dao.NewMockRows(ctrl)
	drv.EXPECT().
		Open("driver://").
		Return(conn, nil)
	conn.EXPECT().
		Prepare(`SELECT "schema" FROM "form_schema" WHERE "uuid" = $1 AND "status" = 'enabled'`).
		Return(stmt, nil)
	conn.EXPECT().
		Prepare(`INSERT INTO "form_data" ("uuid", "data") VALUES ($1, $2)`).
		Return(stmt, nil)
	stmt.EXPECT().
		NumInput().
		Times(3).
		Return(1)
	stmt.EXPECT().
		Query(gomock.Any()).
		Return(rows, nil)
	stmt.EXPECT().
		Close().
		Times(2).
		Return(nil)
	rows.EXPECT().
		Columns().
		Return([]string{"schema"})
	rows.EXPECT().
		Next([]driver.Value{nil}).
		Return(nil)
	rows.EXPECT().
		Close().
		Return(nil)
	sql.Register("driver", drv)

	service, err := dao.New(dao.Connection(&url.URL{Scheme: "driver"}))
	assert.NoError(t, err)

	_, err = service.Schema(UUID)
	assert.Error(t, err)

	_, err = service.AddData(UUID, url.Values{})
	assert.Error(t, err)
}

func fixture(file string) []byte {
	d, err := ioutil.ReadFile("./fixtures/" + file)
	if err != nil {
		panic(err)
	}
	return d
}
