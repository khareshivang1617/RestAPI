package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "github.com/gorilla/mux"
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	// "gorm.io/driver/mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db gorm.DB
var err error

type User struct {
	Id uint32 `gorm:"primaryKey"`
	// gorm.Model
	Name  string
	Email string
}

const DSN = "root:password@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"

func InitialMigration() {
	fmt.Println("Initial Migration")

	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	// db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to the database")
	}

	// defer db.Close()

	db.AutoMigrate(&User{})
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "All users endpoint")
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}
	// defer db.Close()

	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)

}

func NewUser(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "New user endpoint")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}
	// defer db.Close()

	// vars := mux.Vars(r)
	// name := vars["name"]
	// email := vars["email"]

	db.Create(&user)

	fmt.Fprintf(w, "New user successfully created")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Delete user endpoint")
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}
	// defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "User deleted successfully")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Update user endpoint")
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}
	// defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name = ?", name).Find(&user)
	user.Email = email

	db.Save(&user)

	fmt.Fprintf(w, "User updated successfully")
}
