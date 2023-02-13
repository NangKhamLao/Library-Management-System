package main

import (
	"fmt"
	"log"
	"net/http"

	// "database/sql"

	"library/handler"

	"library/config"

	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	host := config.MysqlDB.Host
	port := config.MysqlDB.Port
	user := config.MysqlDB.User
	pass := config.MysqlDB.Pass
	dbname := config.MysqlDB.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", user, pass, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		panic("error on connecting to database")
	}
	fmt.Println("connected to database")

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	store := sessions.NewCookieStore([]byte("jsowjpw38eowj4ur82jmaole0uehqpl"))
	r := handler.New(db, decoder, store)

	log.Println("Server starting...")
	if err := http.ListenAndServe("127.0.0.1:3000", r); err != nil {
		log.Fatal(err)
	}
}
