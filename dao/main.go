package dao

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func getPSQlInfo() string {
	psqlInfo := os.Getenv("DATABASE_URL")
	host := os.Getenv("DATABASE_HOST")
	port, _ := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	user := os.Getenv("DATABASE_USER")
	dbname := os.Getenv("DATABASE_DB")
	password := os.Getenv("DATABASE_PSSWD")
	if psqlInfo == "" {
		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
			"dbname=%s sslmode=disable password=%s",
			host, port, user, dbname, password)
	}
	return psqlInfo
}

var db *sql.DB

func init() {
	Reinit()
}

func Reinit() {
	if dbObj, err := sql.Open("postgres", getPSQlInfo()); err != nil {
		log.Fatal(err)
	} else {
		db = dbObj
	}
}
