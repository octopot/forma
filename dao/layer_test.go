package dao_test

import (
	"database/sql"
	"database/sql/driver"
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
	assert.NotNil(t, service)
	assert.Panics(t, func() { service.Schema(UUID) })
	assert.Panics(t, func() { service.AddData(UUID, map[string][]string{}) })
}

func TestNew_WithInvalidConfiguration(t *testing.T) {
	invalid := []dao.Configurator{dao.Connection("", "")}
	service, err := dao.New(invalid...)
	assert.Error(t, err)
	assert.Nil(t, service)
	assert.Panics(t, func() { dao.Must(invalid...) })
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
		Times(2).
		Return(conn, nil)
	conn.EXPECT().
		Prepare(`SELECT "schema" FROM "form_schema" WHERE "uuid" = $1 AND "status" = 'enabled'`).
		Times(2).
		Return(stmt, nil)
	conn.EXPECT().
		Prepare(`INSERT INTO "form_data" ("uuid", "data") VALUES ($1, $2)`).
		Times(2).
		Return(stmt, nil)
	stmt.EXPECT().
		NumInput().
		Times(6).
		Return(1)
	stmt.EXPECT().
		Query(gomock.Any()).
		Times(2).
		Return(rows, nil)
	stmt.EXPECT().
		Close().
		Times(4).
		Return(nil)
	rows.EXPECT().
		Columns().
		Times(2).
		Return([]string{"schema"})
	rows.EXPECT().
		Next([]driver.Value{nil}).
		Times(2).
		Return(nil)
	rows.EXPECT().
		Close().
		Times(2).
		Return(nil)
	sql.Register("driver", drv)

	valid := []dao.Configurator{dao.Connection("driver", "driver://")}
	{
		service, err := dao.New(valid...)
		assert.NoError(t, err)
		assert.NotNil(t, service.Connection())
		assert.Equal(t, "postgres", service.Dialect())

		_, err = service.Schema(UUID)
		assert.Error(t, err)

		_, err = service.AddData(UUID, map[string][]string{})
		assert.Error(t, err)
	}
	{
		service, err := dao.Must(valid...), error(nil)
		assert.NotNil(t, service.Connection())
		assert.Equal(t, "postgres", service.Dialect())

		_, err = service.Schema(UUID)
		assert.Error(t, err)

		_, err = service.AddData(UUID, map[string][]string{})
		assert.Error(t, err)
	}
}
