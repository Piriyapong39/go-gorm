package main

import (
	"fmt"
	database "goGorm/config"
)

func main() {
	db, err := database.Connection()
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&Book{})

	// newBook := &Book{Name: "Ruksiam", Author: "Kong", Description: "So good", Price: 500}
	// err = CreateBook(db, newBook)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("Create book successfully")

	bookData, err := GetBook(db, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bookData)

	// bookData.Name = "Azujito"
	// bookData.Price = 199

	// err = UpdateBook(db, bookData)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("Update book successfully")

	// err = DeleteBook(db, 2)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("Delete book successfully")
}
