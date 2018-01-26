//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package dao_test -destination $PWD/dao/mock_db_test.go database/sql/driver Conn,Driver,Stmt,Rows
package dao_test

import (
	"database/sql"
	"database/sql/driver"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kamilsk/form-api/dao"
	"github.com/kamilsk/form-api/domain"
	"github.com/stretchr/testify/assert"
)

const (
	DSN  = "stub://localhost"
	UUID = domain.UUID("41ca5e09-3ce2-4094-b108-3ecc257c6fa4")
)

func TestMust_WithInvalidConfiguration(t *testing.T) {
	var configs = []dao.Configurator{dao.Connection("", "")}
	assert.Panics(t, func() { dao.Must(configs...) })
}

func TestMust_WithoutConfiguration(t *testing.T) {
	var configs []dao.Configurator
	assert.NotPanics(t, func() { dao.Must(configs...) })
}

func TestStorage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		drv  = NewMockDriver(ctrl)
		conn = NewMockConn(ctrl)
		stmt = NewMockStmt(ctrl)
		rows = NewMockRows(ctrl)
	)
	drv.EXPECT().
		Open(DSN).
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

	var configs = []dao.Configurator{dao.Connection(t.Name(), DSN)}
	sql.Register(t.Name(), drv)
	service, err := dao.New(configs...)
	assert.NoError(t, err)

	assert.NotNil(t, service.Connection())
	assert.Equal(t, "postgres", service.Dialect())

	_, err = service.AddData(UUID, map[string][]string{})
	assert.Error(t, err)

	_, err = service.Schema(UUID)
	assert.Error(t, err)
}
