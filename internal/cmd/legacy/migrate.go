package legacy

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.octolab.org/fn"

	"go.octolab.org/ecosystem/forma/internal/static"
	"go.octolab.org/ecosystem/forma/internal/storage"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Apply database migration",
	Args:  cobra.RangeArgs(0, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		direction, limit := chooseDirectionAndLimit(args)
		migrate.SetTable(cnf.Union.MigrationConfig.Table)
		migrate.SetSchema(cnf.Union.MigrationConfig.Schema)
		layer := storage.Must(storage.Database(cnf.Union.DatabaseConfig))
		src := &migrate.AssetMigrationSource{
			Asset:    static.Asset,
			AssetDir: static.AssetDir,
			Dir:      "static/migrations",
		}
		var runner = run
		if cnf.Union.MigrationConfig.DryRun {
			runner = dryRun
		}
		if err := runner(layer.Database(), layer.Dialect(), src, direction, limit); err != nil {
			return err
		}
		if direction == migrate.Up && cnf.Union.MigrationConfig.WithDemo {
			raw, err := ioutil.ReadFile("env/test/fixtures/demo.sql")
			switch {
			case err == nil:
				_, err = layer.Database().Exec(string(raw))
				log.Printf("demo: error is %#+v", err)
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
	fn.Must(
		func() error { return v.BindEnv("table") },
		func() error { return v.BindEnv("schema") },
		func() error {
			v.SetDefault("table", defaults["table"])
			v.SetDefault("schema", defaults["schema"])
			return nil
		},
		func() error {
			flags := migrateCmd.Flags()
			flags.StringVarP(&cnf.Union.MigrationConfig.Table,
				"table", "t", v.GetString("table"), "migration table name")
			flags.StringVarP(&cnf.Union.MigrationConfig.Schema,
				"schema", "s", v.GetString("schema"), "migration schema")
			flags.UintVarP(&cnf.Union.MigrationConfig.Limit,
				"limit", "l", 0, "limit the number of migrations (0 = unlimited)")
			flags.BoolVarP(&cnf.Union.MigrationConfig.DryRun,
				"dry-run", "", false, "do not apply migration, just print them")
			flags.BoolVarP(&cnf.Union.MigrationConfig.WithDemo,
				"with-demo", "", false, "create fake data for demo purpose")
			return nil
		},
	)
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
