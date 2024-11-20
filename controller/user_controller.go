package controller

import (
	"context"
	"fmt"
	"golang-database-user/model"
	"golang-database-user/service"
)

func DefaultChoose() {
	fmt.Println("Incorrect Number")
}

func CreateUser(userService service.UserService) {

	ctx := context.Background()
	var name, email, password, phoneNumber string


	fmt.Print("Masukkan Nama : ")
	fmt.Scanln(&name)
	fmt.Print("Masukkan Email : ")
	fmt.Scanln(&email)
	fmt.Print("Masukkan Password : ")
	fmt.Scanln(&password)
	fmt.Print("Masukkan Nomor Telepon : ")
	fmt.Scanln(&phoneNumber)

	user := model.MstUser{
		Name:        name,
		Email:       email,
		Password:    password,
		PhoneNumber: phoneNumber,
	}

	mstUser := userService.CreateUser(ctx, user)

	fmt.Println(mstUser)

}

func ReadUser(userService service.UserService) {
	ctx := context.Background()

	users, err := userService.ReadUser(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nAll Users:")
	for _, mstUser := range users {
		fmt.Println("Id : ", mstUser.IdUser, "\nNama : ", mstUser.Name, "\nEmail : ", mstUser.Email, "\nNomor Telepon : ", mstUser.PhoneNumber)
		fmt.Println()
	}
}

func DeleteUser(userService service.UserService) {
	ctx := context.Background()

	var userId string
	fmt.Print("Masukkan id user yang ingin di hapus: ")
	fmt.Scanln(&userId)

	err := userService.DeleteUser(ctx, userId)
	if err != nil {
		fmt.Println("Gagal menghapus user:", err)
	} else {
		fmt.Println("User berhasil dihapus")
	}
}