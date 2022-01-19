package app

import (
	"database/sql"
	"golang-restapi-httprouter/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err :=sql.Open("mysql","root:@tcp(localhost:3310)/goapi")
	helper.PanicError(err)
	db.SetConnMaxIdleTime(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(6*time.Minute)
	db.SetConnMaxIdleTime(10*time.Minute)
	return db
}


func SetUpTestDb() *sql.DB {
	db, err :=sql.Open("mysql","root:@tcp(localhost:3310)/goapi")
	helper.PanicError(err)
	db.SetConnMaxIdleTime(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(6*time.Minute)
	db.SetConnMaxIdleTime(10*time.Minute)
	return db
}

