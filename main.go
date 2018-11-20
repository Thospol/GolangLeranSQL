package main

import (
	"database/sql"
	"fmt"
	"learnsql/user"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	user.ConnectDB()
	// u := &user.User{
	// 	FirstName: "eiei",
	// 	LastName:  "eueu",
	// 	Email:     "xx@gmail.com",
	// }
	// user.Insert(u)
	// fmt.Printf("%#v\n", u)

	us, err := user.FindAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", us)

	u, err := user.FindByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", u)

	u.Email = "mud077247305@gmail.com"
	u.FirstName = "go"
	u.LastName = "lang"
	err = user.Update(u)
	if err != nil {
		log.Fatal(err)
	}

	// err = user.Delete(u)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	us, err = user.FindAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", us)

}

func insertStmt(db *sql.DB, firstname, lastname, email string) {
	insertSmtmt := "INSERT INTO users (first_name, last_name, email) values ($1, $2,$3)"
	_, err := db.Exec(insertSmtmt, firstname, lastname, email)

	if err != nil {
		log.Fatal(err)
	}
	printAll(db)
}

func deleteStmt(db *sql.DB, id int) {
	queryStmt := "DELETE FROM users WHERE id = $1"
	_, err := db.Exec(queryStmt, id)
	if err != nil {
		log.Fatal(err)
	}
	printByID(db, id)
}

func updateStmt(db *sql.DB, id int, email string) {
	queryStmt := "update users SET email = $1 WHERE id = $2"
	_, err := db.Exec(queryStmt, email, id)
	if err != nil {
		log.Fatal(err)
	}
	printByID(db, id)
}

func printByID(db *sql.DB, id int) {
	queryStmt := "select id,first_name,last_name,email from users where id = $1"
	row := db.QueryRow(queryStmt, id)
	var first_name, last_name, email string

	row.Scan(&id, &first_name, &last_name, &email)
	fmt.Println(id, first_name, last_name, email)
}

func printAll(db *sql.DB) {
	//Query
	queryStmt := "select id,first_name,last_name,email from users"
	rows, err := db.Query(queryStmt)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var (
			id                           int
			first_name, last_name, email string
		)
		rows.Scan(&id, &first_name, &last_name, &email)
		fmt.Println(id, first_name, last_name, email)
	}
}
