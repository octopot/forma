package cmd

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/kamilsk/form-api/dao"
	"github.com/kamilsk/form-api/static"
	"github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Apply database migration",
	Args:  cobra.RangeArgs(0, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		direction, limit := chooseDirectionAndLimit(args)
		{
			migrate.SetTable(cmd.Flag("table").Value.String())
			migrate.SetSchema(cmd.Flag("schema").Value.String())
		}
		layer := dao.Must(dao.Connection(dsn(cmd)))
		src := &migrate.AssetMigrationSource{
			Asset:    static.Asset,
			AssetDir: static.AssetDir,
			Dir:      "static/migrations",
		}
		var runner = run
		if asBool(cmd.Flag("dry-run").Value) {
			runner = dryRun
		}
		if err := runner(layer.Connection(), layer.Dialect(), src, direction, limit); err != nil {
			return err
		}
		if direction == migrate.Up && asBool(cmd.Flag("with-demo").Value) {
			raw, err := ioutil.ReadFile("env/test/fixtures/demo.sql")
			switch {
			case err == nil:
				result, err := layer.Connection().Exec(string(raw))
				log.Printf("demo: %#+v : %#+v", result, err)
			case os.IsNotExist(err):
				log.Println("demo is available only during development")
			default:
				return err
			}
		}
		return nil
	},
}

func init() {
	v := viper.New()
	v.SetEnvPrefix("migration")
	must(
		func() error { return v.BindEnv("table") },
		func() error { return v.BindEnv("schema") },
	)
	{
		v.SetDefault("table", "migration")
		v.SetDefault("schema", "public")
	}
	{
		migrateCmd.Flags().String("table", v.GetString("table"),
			"migration table name")
		migrateCmd.Flags().String("schema", v.GetString("schema"),
			"migration schema")
		migrateCmd.Flags().Int("limit", 0,
			"limit the number of migrations (0 = unlimited)")
		migrateCmd.Flags().Bool("dry-run", false,
			"do not apply migration, just print them")
		migrateCmd.Flags().Bool("with-demo", false,
			"create fake data for demo purpose")
	}
	db(migrateCmd)
}

func chooseDirectionAndLimit(args []string) (migrate.MigrationDirection, int) {
	direction, limit := migrate.Up, 0
	if len(args) > 0 {
		switch {
		case strings.EqualFold(args[0], "up"):
			direction = migrate.Up
		case strings.EqualFold(args[0], "down"):
			direction = migrate.Down
		default:
			log.Fatalf("invalid direction %q", args[0])
		}
		if len(args) == 2 {
			var err error
			limit, err = strconv.Atoi(args[1])
			if err != nil {
				log.Fatalf("limit arg must be a valid integer: %s", err)
			}
		}
	}
	return direction, limit
}

func dryRun(conn *sql.DB, dialect string, src migrate.MigrationSource, direction migrate.MigrationDirection, limit int) error {
	plan, _, err := migrate.PlanMigration(conn, dialect, src, direction, limit)
	if err != nil {
		return err
	}
	for _, m := range plan {
		var queries []string
		if direction == migrate.Up {
			log.Printf("==> Would apply migration %s (up) \n", m.Id)
			queries = m.Up
		} else {
			log.Printf("==> Would apply migration %s (down) \n", m.Id)
			queries = m.Down
		}
		for _, query := range queries {
			log.Println(query)
		}
	}
	return nil
}

func run(conn *sql.DB, dialect string, src migrate.MigrationSource, direction migrate.MigrationDirection, limit int) error {
	count, err := migrate.ExecMax(conn, dialect, src, direction, limit)
	if err != nil {
		return err
	}
	log.Printf("Applied %d migration(s)! \n", count)
	return nil
}
