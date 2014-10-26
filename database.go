package senter

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var dbHandle *gorm.DB = nil

func init() {
	gorm.NowFunc = func() time.Time {
		return time.Unix(time.Now().Unix(), 0).UTC()
		//return time.Now().UTC()
	}
}

func InitDatabase(host string, port string, username string, password string, database string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&autocommit=false", username, password, host, port, database)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		logger.Fatalf("opening database handle failed: %s\n", err)
	}

	err = db.DB().Ping()
	if err != nil {
		logger.Fatalf("cannot connect to database: %s\n", err)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db.SingularTable(true)
	dbHandle = &db
	return nil
}

func CloseDatabase() {
	db := getDb()
	if db != nil {
		err := db.Close()
		if err != nil {
			logger.Printf("closing database handle failed: %s\n", err)
		}
	}
}

func EnableDatabaseLogger() {
	if dbHandle == nil {
		logger.Fatalln("no active database handle")
	}
	dbHandle.LogMode(true)
}

func DisableDatabaseLogger() {
	if dbHandle == nil {
		logger.Fatalln("no active database handle")
	}
	dbHandle.LogMode(false)
}

func getDb() *gorm.DB {
	if dbHandle == nil {
		logger.Fatalln("no active database handle")
	}
	err := dbHandle.DB().Ping()
	if err != nil {
		logger.Fatalf("cannot connect to database: %s\n", err)
	}
	return dbHandle
}
