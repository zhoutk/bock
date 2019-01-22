package main

import (
	"fmt"
	"./db/mysql"
	"./db"
)

func main()  {
	//table := Bock{table: "users"}
	//var file1 IFile = new(File)
	var db db.IBock = new(mysql.Bock)

	//db.query("users")
	fmt.Println(db.Query("users"))
}