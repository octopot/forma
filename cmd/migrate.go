package cmd

import (
	"log"
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
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		var direction migrate.MigrationDirection
		if len(args) < 1 {
			direction = migrate.Up
		} else {
			switch strings.ToLower(args[0]) {
			case "up":
				direction = migrate.Up
			case "down":
				direction = migrate.Down
			}
		}
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
		if cmd.Flag("with-profiler").Value.String() == "true" {
			src = append(src, &migrate.AssetMigrationSource{
				Asset:    static.Asset,
				AssetDir: static.AssetDir,
				Dir:      "static/migrations/demo",
			})
		}
		count, err := migrate.ExecMax(layer.Connection(), layer.Dialect(), src, direction, 0)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Applied %d migration(s)! \n", count)
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
		migrateCmd.Flags().Bool("with-demo", false, "paste fake data for demo purpose")
	}
	db(migrateCmd)
}

type migrations []migrate.MigrationSource

func (b migrations) FindMigrations() ([]*migrate.Migration, error) {
	base, err := static.AssetDir("static/migrations")
	if err != nil {
		return nil, err
	}
	demo, err := static.AssetDir("static/migrations/demo")
	if err != nil {
		return nil, err
	}
	all := make([]*migrate.Migration, 0, len(base)+len(demo))
	for _, src := range b {
		found, err := src.FindMigrations()
		if err != nil {
			return nil, err
		}
		all = append(all, found...)
	}
	return all, nil
}
