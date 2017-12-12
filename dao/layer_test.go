//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package dao_test -destination $PWD/dao/mock_db_test.go database/sql/driver Conn,Driver,Stmt,Rows
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

const UUID data.UUID = "41ca5e09-3ce2-4094-b108-3ecc257c6fa4"

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

	dsn := "stub://localhost"
	drv := NewMockDriver(ctrl)
	name := "stub"
	conn := NewMockConn(ctrl)
	stmt := NewMockStmt(ctrl)
	rows := NewMockRows(ctrl)
	drv.EXPECT().
		Open(dsn).
		Times(1).
		Return(conn, nil)
	conn.EXPECT().
		Prepare(gomock.Any()).
		Times(1).
		Return(stmt, nil)
	conn.EXPECT().
		Prepare(gomock.Any()).
		Times(1).
		Return(stmt, nil)
	stmt.EXPECT().
		NumInput().
		Times(3).
		Return(1)
	stmt.EXPECT().
		Query(gomock.Any()).
		Times(1).
		Return(rows, nil)
	stmt.EXPECT().
		Close().
		Times(2).
		Return(nil)
	rows.EXPECT().
		Columns().
		Times(1).
		Return([]string{"schema"})
	rows.EXPECT().
		Next([]driver.Value{nil}).
		Times(1).
		Return(nil)
	rows.EXPECT().
		Close().
		Times(1).
		Return(nil)

	sql.Register(name, drv)
	valid := []dao.Configurator{dao.Connection(name, dsn)}
	service, err := dao.New(valid...)
	assert.NoError(t, err)
	assert.NotNil(t, service.Connection())
	assert.Equal(t, "postgres", service.Dialect())

	_, err = service.Schema(UUID)
	assert.Error(t, err)

	_, err = service.AddData(UUID, map[string][]string{})
	assert.Error(t, err)

	assert.NotPanics(t, func() { dao.Must(valid...) })
}
