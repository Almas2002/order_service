package postgres

import (
	"7kzu-order-service/pkg/logger"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"strconv"
	"time"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	DBName   string `yaml:"dbName"`
	SSLMode  bool   `yaml:"sslMode"`
	Password string `yaml:"password"`
	Schema   string `yaml:"schema"`
}

func NewDB(cfg *Config) *sqlx.DB {

	host := cfg.Host
	dbName := cfg.DBName
	dbPort := cfg.Port
	dbUser := cfg.User
	dbPass := cfg.Password
	maxConn := 150
	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, dbPort, dbUser, dbName, dbPass)

	connConfig, _ := pgx.ParseConfig(dbUri)
	connStr := stdlib.RegisterConnConfig(connConfig)
	conn, err := sqlx.Open("pgx", connStr)
	conn.SetMaxIdleConns(20)
	conn.SetConnMaxIdleTime(10 * time.Minute)
	if err = conn.Ping(); err != nil {
		logger.Fatal(err)
	} else {
		maxOpen, err := strconv.Atoi(strconv.Itoa(maxConn))
		if err != nil {
			conn.SetMaxOpenConns(5)
		} else {
			conn.SetMaxOpenConns(maxOpen)
		}
		logger.Infof("DB connection is established to PostgreSQL!")
	}

	return conn
}
