package db

import (
	"os"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var db *gorm.DB

type GormLogger struct{}

func (*GormLogger) Print(v ...interface{}) {
	if v[0] == "sql" {
		log.WithFields(log.Fields{"module": "gorm"}).Debug(v[3])
	}
	if v[0] == "log" {
		log.WithFields(log.Fields{"module": "gorm"}).Debug(v[2])
	}
}

func Connect() {
	var err error
	db, err = gorm.Open("postgres", os.Getenv("DATABASE_CONNECTION_STRING"))
	if err != nil {
		log.Fatal("failed to connect database")
	}

	if err = db.DB().Ping(); err != nil {
		log.Fatal("failed to connect database")
	}

}

func GetSession() *gorm.DB {
	// Advance Setting
	if os.Getenv("CONN_MAX_LIFETIME") != "" {
		connMaxLifetime, _ := strconv.Atoi(os.Getenv("CONN_MAX_LIFETIME"))
		db.DB().SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)
	}
	if os.Getenv("MAX_IDLE_CONNS") != "" {
		maxIdleConns, _ := strconv.Atoi(os.Getenv("MAX_IDLE_CONNS"))
		db.DB().SetMaxIdleConns(maxIdleConns)
	}
	if os.Getenv("MAX_OPEN_CONNS") != "" {
		maxOpenConns, _ := strconv.Atoi(os.Getenv("MAX_OPEN_CONNS"))
		db.DB().SetMaxOpenConns(maxOpenConns)
	}

	if os.Getenv("PLATFORM") == "Production" {
		db.SetLogger(&GormLogger{})
	}
	db.LogMode(true)

	return db
}
