package service

import (
	"fmt"
	"log"
	"userTest/config"
	"userTest/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser() {
	var username, password string

	fmt.Print("Type your username: ")
	fmt.Scan(&username)

	fmt.Print("Type your password: ")
	fmt.Scan(&password)

	passwordhash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("ERROR: COULDN'T GENERATE PASSWORD:", err)
	}

	_, err = config.DB.Exec("INSERT INTO users (username, password_hash) VALUES ($1, $2)", username, string(passwordhash))
	if err != nil {
		log.Println("ERROR: COULDN'T CREATE USER:", err)
	}

	log.Println("USER SUCCESSFULLY CREATED!")
}

func Login() {
	var user, password string

	fmt.Print("Type your username: ")
	fmt.Scan(&user)

	fmt.Print("Type your password: ")
	fmt.Scan(&password)

	var username, hashedPassword string

	err := config.DB.QueryRow("SELECT username, password_hash FROM users WHERE username=$1", user).Scan(&username, &hashedPassword)
	if err != nil {
		log.Println("ERROR: COULDN'T UPDATE USER:", err)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Println("WRONG PASSWORD!")
		return
	}
	fmt.Printf("LOGGED IN SUCCESSFULLY! WELCOME BACK %s", username)
}

func UpdateUser() {
	repository.ListUsers()

	var id int
	fmt.Print("\nType User's ID: ")
	fmt.Scan(&id)

	var newUsername, newPassword string

	fmt.Print("Type new Username: ")
	fmt.Scan(&newUsername)

	fmt.Print("Type new Password: ")
	fmt.Scan(&newPassword)

	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Println("ERROR: COULDN'T GENERATE PASSWORD PASSWORD:", err)
		return
	}

	result, err := config.DB.Exec("UPDATE users SET username=$1, password_hash=$2 WHERE id=$3", newUsername, string(hashedpassword), id)
	if err != nil {
		log.Println("ERROR: COULDN'T UPDATE USER:", err)
		return
	}

	rAffected, _ := result.RowsAffected()
	if rAffected == 0 {
		log.Println("USER DOES NOT EXIST")
		return
	}

	log.Println("USER SUCCESSFULLY UPDATED!")
}

func DeleteUser() {
	repository.ListUsers()

	var id int
	fmt.Print("\nType User's ID: ")
	fmt.Scan(&id)

	result, err := config.DB.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		log.Println("ERROR: COULDN'T DELETE USER:", err)
		return
	}

	rAffected, _ := result.RowsAffected()
	if rAffected == 0 {
		log.Println("USER DOES NOT EXIST")
		return
	}
	log.Println("USER SUCCESSFULLY DELETED!")
}
