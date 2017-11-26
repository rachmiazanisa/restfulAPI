package controllers

import (
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

//Function Read Data by id
func DBHome(c *gin.Context) {
	//Connect Database
	db, err := sql.Open("mysql","root:@/homebase")
	if err != nil{
		panic(err.Error())
	}
	defer db.Close()
	
	stmtOut, err := db.Prepare("SELECT nama FROM user WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	var name string

	
	err = stmtOut.QueryRow(3).Scan(&name)
	if err != nil{
		panic(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"Nama" : name})
}
//Function Delete Data by id
func DBDelete(c *gin.Context) {
	//Connect Database
	db, err := sql.Open("mysql", "root:@/homebase")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmtDel, err := db.Prepare("DELETE FROM user WHERE id = ?")
	if err != nil{
		panic(err.Error())
	}
	defer stmtDel.Close()
	_, err = stmtDel.Exec(6)
	if err != nil{
		panic(err.Error())
	}
}
//Function Update Data by id
func DBUpdate(c *gin.Context) {
	//Connect Database
	db, err := sql.Open("mysql", "root:@/homebase")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmtUpd, err := db.Prepare("UPDATE user SET nama = ? WHERE id = ?")
	if err != nil{
		panic(err.Error())
	}
	defer stmtUpd.Close()
	_, err = stmtUpd.Exec("mamat",6)
	
}
//Function Insert Data
func DBInsert(c *gin.Context){
	//Connect Database
	db, err := sql.Open("mysql", "root:@/homebase")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmtIns, err := db.Prepare("INSERT INTO user VALUES('','rachmi',23)")
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec()
	if err != nil {
		panic(err.Error())
	}
}