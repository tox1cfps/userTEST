package ui

import (
	"fmt"
	"log"
	"userTest/internal/service"
)

func UserMenu() {
	for {
		var option int

		fmt.Println("\n#### Sign in & Login ####")
		fmt.Print("\n1. Create my User")
		fmt.Print("\n2. Login")
		fmt.Print("\n3. Update my user")
		fmt.Print("\n4. Delete my user")
		fmt.Print("\n0. Exit")
		fmt.Print("\nChoose an option: ")
		fmt.Scan(&option)

		switch option {
		case 1:
			service.CreateUser()
		case 2:
			service.Login()
		case 3:
			service.UpdateUser()
		case 4:
			service.DeleteUser()
		case 0:
			log.Println("See Ya!")
			return
		default:
			fmt.Println("\nInvalid Option!")
		}
	}
}
