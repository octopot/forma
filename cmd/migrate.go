package cmd

import (
	"database/sql"
	"log"
	"strconv"
	"strings"

	"github.com/kamilsk/form-api/dao"
	"github.com/kamilsk/form-api/static"
	"github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// TODO v2: refactoring
// - do not use log.Fatalf

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Apply database migration",
	Args:  cobra.RangeArgs(0, 2),
	Run: func(cmd *cobra.Command, args []string) {
		direction, limit := chooseDirectionAndLimit(args)
		{
			migrate.SetTable(cmd.Flag("table").Value.String())
			migrate.SetSchema(cmd.Flag("schema").Value.String())
		}
		layer := dao.Must(dao.Connection(dsn(cmd)))
		src := make(migrations, 0, 2)
		src = append(src, &migrate.AssetMigrationSource{
			Asset:    static.Asset,
			AssetDir: static.AssetDir,
			Dir:      "static/migrations",
		})
		if isTrue(cmd.Flag("demo").Value) {
			src = append(src, &migrate.AssetMigrationSource{
				Asset:    static.Asset,
				AssetDir: static.AssetDir,
				Dir:      "static/migrations/demo",
			})
		}
		if isTrue(cmd.Flag("dry-run").Value) {
			dryRun(layer.Connection(), layer.Dialect(), src, direction, limit)
		} else {
			run(layer.Connection(), layer.Dialect(), src, direction, limit)
		}
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
		migrateCmd.Flags().String("table", v.GetString("table"), "migration table name")
		migrateCmd.Flags().String("schema", v.GetString("schema"), "migration schema")
		migrateCmd.Flags().Int("limit", 0, "limit the number of migrations (0 = unlimited)")
		migrateCmd.Flags().Bool("dry-run", false, "do not apply migration, just print them")
		migrateCmd.Flags().Bool("demo", false, "create fake data for demo purpose")
	}
	db(migrateCmd)
}

func chooseDirectionAndLimit(args []string) (migrate.MigrationDirection, int) {
	direction, limit := migrate.Up, 0
	if len(args) > 0 {
		switch dir := strings.ToLower(args[0]); dir {
		case "up":
			direction = migrate.Up
		case "down":
			direction = migrate.Down
		default:
			log.Fatalf("invalid direction %q", dir)
		}
		if len(args) == 2 {
			var err error
			limit, err = strconv.Atoi(args[1])
			if err != nil {
				log.Fatalf("limit arg should be valid integer: %s", err)
			}
		}
	}
	return direction, limit
}

func dryRun(conn *sql.DB, dialect string, src migrate.MigrationSource, direction migrate.MigrationDirection, limit int) {
	plan, _, err := migrate.PlanMigration(conn, dialect, src, direction, limit)
	if err != nil {
		log.Fatal(err)
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
}

func run(conn *sql.DB, dialect string, src migrate.MigrationSource, direction migrate.MigrationDirection, limit int) {
	count, err := migrate.ExecMax(conn, dialect, src, direction, limit)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Applied %d migration(s)! \n", count)
}

type migrations []migrate.MigrationSource

func (b migrations) FindMigrations() ([]*migrate.Migration, error) {
	var all []*migrate.Migration
	for _, src := range b {
		found, err := src.FindMigrations()
		if err != nil {
			return nil, err
		}
		all = append(all, found...)
	}
	return all, nil
}
