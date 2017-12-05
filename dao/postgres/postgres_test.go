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

const UUID data.UUID = "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11"

func TestNew_WithValidConfiguration(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dsn := "postgres://postgres:postgres@db:5432/postgres"
	drv := NewMockDriver(ctrl)
	name := "postgres"
	conn := NewMockConn(ctrl)
	stmt := NewMockStmt(ctrl)
	rows := NewMockRows(ctrl)
	drv.EXPECT().
		Open(dsn).
		Times(1).
		Return(conn, nil)
	conn.EXPECT().
		Prepare(`SELECT "schema" FROM "form_schema" WHERE "uuid" = $1 AND "status" = 'enabled'`).
		Times(1).
		Return(stmt, nil)
	conn.EXPECT().
		Prepare(`INSERT INTO "form_data" ("uuid", "data") VALUES ($1, $2)`).
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
	db, err := sql.Open(name, dsn)
	assert.NoError(t, err)

	_, err = postgres.AddData(db, UUID, map[string][]string{})
	assert.Error(t, err)

	assert.Equal(t, "postgres", postgres.Dialect())

	_, err = postgres.Schema(db, UUID)
	assert.Error(t, err)
}
