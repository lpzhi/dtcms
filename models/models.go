package models

import (
	"dtcms/config"
	"dtcms/const"
	"fmt"
	"github.com/jinzhu/gorm"
	_"github.com/go-sql-driver/mysql"
	"log"
)

var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	//CreatedOn int `json:"created_on"`
	//ModifiedOn int `json:"modified_on"`
}

func ConnectDb()  {
	var(
		err error
		ptype string
		serverurl string
	)

	// get database config

	rs,_ := config.GetDataBaseInfo(_const.Default_Database_Id)

	serverurl = rs.ServerUrl
	ptype = config.GetAppConfig("ptype")

	db, err = gorm.Open(ptype,serverurl)

	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println("db start")

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

}

func CloseDB() {
	defer db.Close()
}