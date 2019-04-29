package postgres

import (
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

//GormLogger global gorm logger for logging quesries to ELK Stack
type GormLogger struct{}

//Print - overrided print function
func (*GormLogger) Print(v ...interface{}) {
	if v[0] == "sql" {
		log.WithFields(log.Fields{"module": "gorm"}).Debug(v[3])
	}
	if v[0] == "log" {
		log.WithFields(log.Fields{"module": "gorm"}).Debug(v[2])
	}
}

//PostgresClientHandler -
type PostgresClientHandler interface {
	Connect() PostgresClientHandler
	DBPing() error
	GetSession() *gorm.DB
}

type PostgresClient struct {
	db *gorm.DB
}

var once sync.Once
var conn *PostgresClient

func New() PostgresClientHandler {
	//singleton pattern
	once.Do(func() {
		conn = &PostgresClient{}
	})
	return conn
}

//Connect - connect to pg instance
func (r *PostgresClient) Connect() PostgresClientHandler {
	var (
		err              error
		hostnameParamVal string
	)
	// add hostname for application name
	if os.Getenv("HOSTNAME") != "" {
		hostnameParamVal = "&application_name=" + os.Getenv("HOSTNAME")
	}

	r.db, err = gorm.Open("postgres", os.Getenv("DATABASE_CONNECTION_STRING")+hostnameParamVal)
	if err != nil {
		log.Error(err.Error())
		log.Error("failed to connect database")
	}

	if err = r.db.DB().Ping(); err != nil {
		log.Error("failed to connect database")
	}

	return r
}

//DBPing -
func (r *PostgresClient) DBPing() error {
	if err := r.db.DB().Ping(); err != nil {
		log.Error("failed to connect database")
		return err
	}
	return nil
}

//GetSession - global shared session for app
func (r *PostgresClient) GetSession() *gorm.DB {
	// Advance Setting
	if os.Getenv("CONN_MAX_LIFETIME") != "" {
		connMaxLifetime, _ := strconv.Atoi(os.Getenv("CONN_MAX_LIFETIME"))
		r.db.DB().SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)
	}
	if os.Getenv("MAX_IDLE_CONNS") != "" {
		maxIdleConns, _ := strconv.Atoi(os.Getenv("MAX_IDLE_CONNS"))
		r.db.DB().SetMaxIdleConns(maxIdleConns)
	}
	if os.Getenv("MAX_OPEN_CONNS") != "" {
		maxOpenConns, _ := strconv.Atoi(os.Getenv("MAX_OPEN_CONNS"))
		r.db.DB().SetMaxOpenConns(maxOpenConns)
	}

	//set log format and level
	if os.Getenv("PLATFORM") == "Production" {
		r.db.SetLogger(&GormLogger{})
	}
	r.db.LogMode(true)
	return r.db
}
