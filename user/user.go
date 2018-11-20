package user

import (
	"database/sql"
	"log"
)

var (
	db   *sql.DB
	user User
)

// User is strcut of User
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type Service struct {
	DB *sql.DB
}

func (s *Service) FindByID(id int) (*User, error) {
	//return &User{}, nil
	queryStmt := "select id,first_name,last_name,email from users where id = $1"
	row := s.DB.QueryRow(queryStmt, id) //s *Service เพราะมี field DB ยุ
	//var first_name, last_name, email string
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		log.Fatal(err)
	}
	return &user, err
}

func (s *Service) FindAll() ([]User, error) {
	queryStmt := "select id,first_name,last_name,email from users ORDER BY id ASC"
	rows, err := s.DB.Query(queryStmt)
	if err != nil {
		return nil, err
	}
	var us []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email)
		if err != nil {
			return nil, err
		}
		us = append(us, u)
	}
	return us, err
}

//Insert is Insert users
func (s *Service) Insert(u *User) error {
	insertSmtmt := "INSERT INTO users (first_name, last_name, email) values ($1, $2,$3) RETURNING id"
	row := s.DB.QueryRow(insertSmtmt, u.FirstName, u.LastName, u.Email)
	err := row.Scan(&u.ID)
	return err
}

//Update is Update users
func (s *Service) Update(u *User) error {
	stmt := "UPDATE users SET first_name = $1, last_name = $2, email = $3 WHERE id = $4"
	_, err := s.DB.Exec(stmt, u.FirstName, u.LastName, u.Email, u.ID)
	return err
}

//Delete is Delete users
func (s *Service) Delete(u *User) error {
	stmt := "DELETE FROM users WHERE id = $1"
	_, err := s.DB.Exec(stmt, u.ID)
	return err

}

//ConnectDB is ConnectDB users
func ConnectDB() {
	var err error
	connStr := "postgres://rypsagyu:z57AcaK1q70fwYThQhi6MIHrWgAPFU25@elmer.db.elephantsql.com:5432/rypsagyu"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

//Insert is Insert users
func Insert(u *User) error {
	insertSmtmt := "INSERT INTO users (first_name, last_name, email) values ($1, $2,$3) RETURNING id"
	row := db.QueryRow(insertSmtmt, u.FirstName, u.LastName, u.Email)
	err := row.Scan(&u.ID)
	return err
}

//Update is Update users
func Update(u *User) error {
	stmt := "UPDATE users SET first_name = $1, last_name = $2, email = $3 WHERE id = $4"
	_, err := db.Exec(stmt, u.FirstName, u.LastName, u.Email, u.ID)
	return err
}

//Delete is Delete users
func Delete(u *User) error {
	stmt := "DELETE FROM users WHERE id = $1"
	_, err := db.Exec(stmt, u.ID)
	return err

}

//FindByID is FindByID users
func FindByID(id int) (*User, error) {

	queryStmt := "select id,first_name,last_name,email from users where id = $1"
	row := db.QueryRow(queryStmt, id)
	//var first_name, last_name, email string
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		log.Fatal(err)
	}
	return &user, err
}

//FindAll is FindAll users
func FindAll() ([]User, error) {
	queryStmt := "select id,first_name,last_name,email from users ORDER BY id ASC"
	rows, err := db.Query(queryStmt)
	if err != nil {
		return nil, err
	}
	var us []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email)
		if err != nil {
			return nil, err
		}
		us = append(us, u)
	}
	return us, err
}
