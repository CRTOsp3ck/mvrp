package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var Conn *sql.DB

type DBConfig struct {
	Host         string
	Port         string
	User         string
	Password     string
	DBName       string
	SslMode      string
	PoolMaxConns string
}

func Init() {
	db, err := newDB()
	if err != nil {
		panic(err)
	}
	Conn = db
	boil.SetDB(db)
	boil.SetLocation(time.UTC)
	// boil.DebugMode = true
}

func Close() {
	err := Conn.Close()
	if err != nil {
		panic(err)
	}
}

func newDBConfig() *DBConfig {
	var host string = "localhost"
	if isDocker() {
		host = "host.docker.internal"
	}
	return &DBConfig{
		Host:         host,
		Port:         "5432",
		User:         "postgres",
		Password:     "postgres",
		DBName:       "mvrp",
		SslMode:      "disable",
		PoolMaxConns: "10",
	}
}

func (c *DBConfig) connectionString() string {
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		c.User,
		c.Password,
		c.DBName,
		c.Host,
		c.Port,
		c.SslMode,
	)
}

func newDB() (*sql.DB, error) {
	conn, err := sql.Open("pgx", newDBConfig().connectionString())
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func isDocker() bool {
	_, err := os.Stat("/.dockerenv")
	return !os.IsNotExist(err)
}
