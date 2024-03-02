package init

import "github.com/joho/godotenv"

func InitDBConnection() {
	// connect to db
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

}
