package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "lenslocked"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Wrong connections to DB")
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Wrong connections to DB")
		panic(err)
	}

	fmt.Println("Successfuly connection to DB")

	// var id int
	// var name, email string

	// rows, err := db.Query("SELECT id, name, email FROM users WHERE id=$1 OR email=$2", 1, "newsly@mail.ru")
	// if err != nil {
	// 	panic(err)
	// }

	// for rows.Next() {
	// 	rows.Scan(&id, &name, &email)
	// 	fmt.Println("id: ", id, "name:", name, "email: ", email)
	// }

	// var orderId int
	// for i := 1; i < 6; i++ {
	// 	userId := 1

	// 	if i > 3 {
	// 		userId = 2
	// 	}

	// 	amount := 1000 * i
	// 	description := fmt.Sprintf("USB Type-C adapter x $d", i)

	// 	err = db.QueryRow("INSERT INTO orders (user_id, amount, description) VALUES($1, $2, $3) RETURNING id", userId, amount, description).Scan(&orderId)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	fmt.Println("Created an order with ID: ", orderId)
	// }

	// var id int
	// var name, email string

	// row := db.QueryRow("SELECT id, name, email FROM users WHERE id=$1", 1)
	// err = row.Scan(&id, &name, &email)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("id: ", id, "name:", name, "email: ", email)

	// var id int
	// row := db.QueryRow("INSERT INTO users(name, email) VALUES($1, $2) RETURNING id", "Mike Dinahue", "mikey@mail.com")
	// err = row.Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }

	// _, err = db.Exec("INSERT INTO users(name, email) VALUES($1, $2)", "Mikey Newsly", "newsly@mail.ru")
	// if err != nil {
	// 	panic(err)
	// }

	db.Close()
}
