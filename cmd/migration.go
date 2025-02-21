package cmd

import (
	"database/sql"
	"fmt"
	"github.com/ahmadrezamusthafa/ecommerce/config"
	"github.com/ahmadrezamusthafa/ecommerce/pkg/database"
	"github.com/ahmadrezamusthafa/ecommerce/pkg/logger"
	"os"
	"time"

	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
)

const (
	MigrationDirectory = "migrations/"
	DBDialectMysql     = "postgres"
)

var migrateUpCommand = &cobra.Command{
	Use:   "migrate/up",
	Short: "Migrate up for database",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		source := getFileMigrationSource()

		err := doMigration(source, migrate.Up)

		if err != nil {
			return
		}
	},
}

var makeMigrationCommand = &cobra.Command{
	Use:   "migrate/create",
	Short: "Create new migration database file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		migrationDir := MigrationDirectory
		migrationName := args[0]
		err := createMigrationFile(migrationDir, migrationName)

		if err != nil {
			return
		}
	},
}

var migrateDownCommand = &cobra.Command{
	Use:   "migrate/down",
	Short: "Migrate down for database",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		source := getFileMigrationSource()

		err := doMigration(source, migrate.Down)
		if err != nil {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateUpCommand)
	rootCmd.AddCommand(migrateDownCommand)
	rootCmd.AddCommand(makeMigrationCommand)
}

func doMigration(fileMigrationSource migrate.FileMigrationSource, direction migrate.MigrationDirection) error {
	cfg := config.GetConfig()

	db, err := database.NewPostgresqlDatabase(cfg.Database)
	if err != nil {
		logger.Error("Cannot connect to database")
		return err
	}

	sqlDb, err := db.DB()
	if err != nil {
		logger.Fatalf("Failed to get sql.DB from GORM: %v", err)
	}

	defer func(sqlDb *sql.DB) {
		err := sqlDb.Close()
		if err != nil {
			return
		}
	}(sqlDb)

	migrate.SetTable("migrations")

	totalMigrated, err := migrate.Exec(sqlDb, DBDialectMysql, fileMigrationSource, direction)
	if err != nil {
		logger.Error("Migration failed: ", err)
		return err
	}

	logger.Infof("Migrate success, total migrated: %d", totalMigrated)
	return err
}

func getFileMigrationSource() migrate.FileMigrationSource {
	source := migrate.FileMigrationSource{
		Dir: MigrationDirectory,
	}

	return source
}

func createMigrationFile(mDir string, mName string) error {

	var migrationContent = `-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
-- [your SQL script here]

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- [your SQL script here]
`
	filename := fmt.Sprintf("%d_%s.sql", time.Now().Unix(), mName)
	filepath := fmt.Sprintf("%s%s", mDir, filename)

	f, err := os.Create(filepath)
	if err != nil {
		logger.Error("Error create migration file: ", filepath, filename, err)
		return err
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			logger.Error("Error when closing file")
		}
	}(f)

	_, err = f.WriteString(migrationContent)
	if err != nil {
		return err
	}

	err = f.Sync()
	if err != nil {
		return err
	}

	logger.Infof("New migration file has been created %s", filename)
	return nil
}
