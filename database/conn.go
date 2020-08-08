package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func Conn() *gorm.DB {
	 connstr := "user=postgres dbname=notesite host=localhost password=anteneh23 sslmode=disable"

	db, err := gorm.Open("postgres", connstr)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println("connected")
	return db
}
