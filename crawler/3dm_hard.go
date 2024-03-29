package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ExampleScrape() {
	res, err := http.Get("https://www.3dmgame.com/news/201907/3766333.html")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".news_warp_center").Each(
		func(i int, s *goquery.Selection) {
			text := s.Find("p").Text()
			fmt.Println(text)
		})
}

func main() {
	ExampleScrape()
}

/*type User struct {
	Id   int
	Name string
}

func Add(db *gorm.DB) {
	user := User{Id: 1, Name: "zhangsan"}
	db.Create(&user)
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
*/
