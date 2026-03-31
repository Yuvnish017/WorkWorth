package database

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/Yuvnish017/WorkWorth/config"
	_ "github.com/lib/pq"
)

func NewDB(cfg config.Config) (*sql.DB, error) {
	host := cfg.DBHost
	port, _ := strconv.Atoi(cfg.DBPort) // don't forget to convert int since port is int type.
	user := cfg.DBUser
	dbname := cfg.DBName
	pass := cfg.DBPassword

	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, pass)
	db, errSql := sql.Open("postgres", psqlSetup)
	if errSql != nil {
		return nil, errSql
	}

	return db, nil
}
