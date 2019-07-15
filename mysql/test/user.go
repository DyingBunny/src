package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type User struct {
	Id   int
	Name string
}

func Add(db *gorm.DB) {
	user := User{Id: 1, Name: "zhangsan"}
	db.Create(&user)
}

func Delete(db *gorm.DB) {
	user := User{Name: "zhangsan"}
	db.Delete(&user)
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@/golang")
	defer db.Close()
	db.SingularTable(true)
	if err != nil {
		log.Fatal("Connect error:", err)
	} else {
		fmt.Println("Connect success!")
	}
	Delete(db)
}
