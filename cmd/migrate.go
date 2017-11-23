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
		migrations := &migrate.AssetMigrationSource{
			Asset:    static.Asset,
			AssetDir: static.AssetDir,
			Dir:      "static/migrations",
		}
		count, err := migrate.ExecMax(layer.Connection(), layer.Dialect(), migrations, direction, 0)
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
	}
	db(migrateCmd)
}
