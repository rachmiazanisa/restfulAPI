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

	//Change number in QueryRow for the id
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
	//Change number in QueryRow , number => id
	_, err = stmtDel.Exec(6)
	if err != nil{
		panic(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"Status":"Deleted"})
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
	//Change "name",number | number => id
	name := "mamat"
	_, err = stmtUpd.Exec(name,6)
	c.JSON(http.StatusOK, gin.H{
		"Name": name,
		"Status":"Updated"})
}
//Function Insert Data
func DBInsert(c *gin.Context){
	//Connect Database
	db, err := sql.Open("mysql", "root:@/homebase")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	name := "rachmi"
	umur := 23
	stmtIns, err := db.Prepare("INSERT INTO user VALUES('','name',umur)")
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec()
	if err != nil {
		panic(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"Name" : name,
		"Umur" : umur,
		"Status":"Updated"})
}