package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"lenslocked/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "lenslocked"
)

type User struct {
	gorm.Model
	Name   string
	Email  string `gorm:"not null; unique_index"`
	Orders []Order
}

type Order struct {
	gorm.Model
	UserID      uint
	Amount      int
	Description string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}

	defer us.Close()
	us.DesctructiveReset()

	user, err := us.ByID(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)

	// db, err := gorm.Open("postgres", psqlInfo)
	// if err != nil {
	// 	fmt.Println("Wrong connections to DB")
	// 	panic(err)
	// }

	// fmt.Println("Successfuly connection to DB")
	// defer db.Close()

	// db.LogMode(true)
	// db.AutoMigrate(&User{}, &Order{})

	// name, email := getInfo()
	// u := &User{
	// 	Name:  name,
	// 	Email: email,
	// }
	// if err = db.Create(u).Error; err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v\n", u)

	// var firstUser User
	// db.Preload("Orders").First(&firstUser).Not(1)
	// if db.Error != nil {
	// 	panic(db.Error)
	// }

	// // fmt.Println(firstUser)
	// fmt.Println("Email:", firstUser.Email)
	// fmt.Println("Number of orders:", len(firstUser.Orders))
	// fmt.Println("Orders:", firstUser.Orders)

	//createOrder(db, firstUser, 1001, "Fake Description #1")
	// createOrder(db, firstUser, 9999, "Fake Description #2")
	// createOrder(db, firstUser, 8800, "Fake Description #3")
}

func getInfo() (name, email string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Println("What is your email?")
	email, _ = reader.ReadString('\n')
	email = strings.TrimSpace(email)
	return name, email
}

func createOrder(db *gorm.DB, user User, amount int, description string) {
	db.Create(&Order{
		UserID:      user.ID,
		Amount:      amount,
		Description: description,
	})

	if db.Error != nil {
		panic(db.Error)
	}
}
