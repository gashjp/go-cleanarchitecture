package mysql

import (
	"github.com/gashjp1994/go-ca/db"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var database *gorm.DB

// Connect ...
func Connect() (*gorm.DB, error) {
	dbms := "mysql"
	user := "root"
	pass := "saijopw"
	protocol := "tcp(localhost:3306)"
	dbname := "mydb"

	connect := user + ":" + pass + "@" + protocol + "/" + dbname + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	// TODO: 名前検討
	var err error
	database, err = gorm.Open(dbms, connect)
	if err != nil {
		return nil, err
	}
	if !database.HasTable(&db.User{}) {
		if err := database.Table("users").CreateTable(&db.User{}).Error; err != nil {
			return nil, err
		}
	}
	return database, nil
}

func CloseConnect() {
	database.Close()
}
