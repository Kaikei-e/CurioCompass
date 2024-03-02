package init

import (
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"os"
)

func InitDBConnection() *bun.DB {
	// connect to db
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_DB := os.Getenv("POSTGRES_DB")
	DB_IP := os.Getenv("DB_IP")

	dsn := "postgres://" + POSTGRES_USER + ":" + POSTGRES_PASSWORD + "@" + DB_IP + "/" + POSTGRES_DB + "?sslmode=disable"

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
