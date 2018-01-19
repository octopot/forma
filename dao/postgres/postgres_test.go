//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package postgres_test -destination $PWD/dao/postgres/mock_db_test.go database/sql/driver Conn,Driver,Stmt,Rows
package postgres_test

import (
	"database/sql"
	"database/sql/driver"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kamilsk/form-api/dao/postgres"
	"github.com/kamilsk/form-api/data"
	"github.com/stretchr/testify/assert"
)

const (
	DSN  = "postgres://postgres:postgres@db:5432/postgres"
	UUID = data.UUID("41ca5e09-3ce2-4094-b108-3ecc257c6fa4")
)

func TestAddData(t *testing.T) {
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
		Prepare(`INSERT INTO "form_data" ("uuid", "data") VALUES ($1, $2) RETURNING "id"`).
		Times(1).
		Return(stmt, nil)
	stmt.EXPECT().
		NumInput().
		Times(2).
		Return(2)
	stmt.EXPECT().
		Query(gomock.Any()).
		Times(1).
		Return(rows, nil)
	stmt.EXPECT().
		Close().
		Times(1).
		Return(nil)
	rows.EXPECT().
		Columns().
		Times(1).
		Return([]string{"id"})
	rows.EXPECT().
		Next([]driver.Value{nil}).
		Times(1).
		Return(nil)
	rows.EXPECT().
		Close().
		Times(1).
		Return(nil)

	sql.Register(t.Name(), drv)
	db, err := sql.Open(t.Name(), DSN)
	assert.NoError(t, err)

	_, err = postgres.AddData(db, UUID, map[string][]string{})
	assert.Error(t, err)
}

func TestDialect(t *testing.T) {
	assert.Equal(t, "postgres", postgres.Dialect())
}

func TestSchema(t *testing.T) {
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
		Prepare(`SELECT "schema" FROM "form_schema" WHERE "uuid" = $1 AND "status" = 'enabled'`).
		Times(1).
		Return(stmt, nil)
	stmt.EXPECT().
		NumInput().
		Times(2).
		Return(1)
	stmt.EXPECT().
		Query(gomock.Any()).
		Times(1).
		Return(rows, nil)
	stmt.EXPECT().
		Close().
		Times(1).
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

	sql.Register(t.Name(), drv)
	db, err := sql.Open(t.Name(), DSN)
	assert.NoError(t, err)

	_, err = postgres.Schema(db, UUID)
	assert.Error(t, err)
}
