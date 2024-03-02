package initialize

import (
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"log/slog"
	"os"
)

func InitDBConnection() *bun.DB {
	// connect to db
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	PostgresUser := os.Getenv("POSTGRES_USER")
	PostgresPassword := os.Getenv("POSTGRES_PASSWORD")
	PostgresDb := os.Getenv("POSTGRES_DB")
	DbIp := os.Getenv("DB_IP")

	dsn := "postgres://" + PostgresUser + ":" + PostgresPassword + "@" + DbIp + ":5432" + "/" + PostgresDb + "?sslmode=disable"

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())

	err = db.Ping()
	if err != nil {
		slog.Error("failed to ping db", err)
	}

	return db
}
