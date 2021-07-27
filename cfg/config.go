package cfg

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)
type DB struct {
	Server string
	Port int
	User string
	Password string
	Database string
}
var DbExport DB

func (dbConnection DB) OpenDB() (*gorm.DB,error){
	db, err := gorm.Open("mssql", dbConnection.ConnectToSql())
	if err != nil{
		return nil,err
	}
	return db,nil
}

func (db_connection DB) ConnectToSql() string {
	return fmt.Sprintf("server=%s; user id=%s; password=%s; port=%d; database=%s;",
		db_connection.Server, db_connection.User, db_connection.Password, db_connection.Port, db_connection.Database)
}