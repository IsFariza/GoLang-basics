package Library

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Book :all the fields of are uppercase so they are exported and can be used in other packages
// update:
type Book struct {
	ID         uint64
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	Books         map[string]Book //Books is a map with key-strings and values-Book
	incrementedID uint64
}

// AddBook used pointer receiver to update actual library, not the copy
func (library *Library) AddBook(book Book) {

	id := library.incrementedID
	library.incrementedID++
	book.ID = id
	library.Books[strconv.FormatUint(id, 10)] = book

	fmt.Println("Book successfully added to the library")

}

func (library *Library) BorrowBook(ID string) {
	book, exists := library.Books[ID]
	if !exists {
		fmt.Println("Book is not found")
		return
	}
	if book.IsBorrowed {
		fmt.Println("Book is already borrowed")
		return
	}
	book.IsBorrowed = true
	library.Books[ID] = book
	fmt.Printf("Book %s successfully borrowed\n", book.Title)
}
func (library *Library) ReturnBook(ID string) Book {
	book, exists := library.Books[ID] //this is just a copy of the struct
	if !exists {
		fmt.Println("Book is not found")
		return Book{}
	}
	if !book.IsBorrowed {
		fmt.Println("Book wasn't borrowed")
		return Book{}
	}
	book.IsBorrowed = false
	library.Books[ID] = book //assign it back to the map, to save changes to real map, not the copy
	fmt.Printf("Book %s successfully returned\n", book.Title)
	return book
}
func (library *Library) ListAvailableBooks() {
	if len(library.Books) == 0 {
		fmt.Println("No books available")
	} else {
		for _, book := range library.Books {
			if !book.IsBorrowed {
				fmt.Printf("id: %v, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
			}

		}
	}
}
func (library *Library) ListBorrowedBooks() {
	if len(library.Books) == 0 {
		fmt.Println("No books available")
	} else {
		for _, book := range library.Books {
			if book.IsBorrowed {
				fmt.Printf("id: %v, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
			}

		}
	}
}

func printMenu() {
	fmt.Println()
	fmt.Println("Menu (type in the number of operation): ")
	fmt.Println("1. Add Book ")
	fmt.Println("2. Borrow Book ")
	fmt.Println("3. Return Book ")
	fmt.Println("4. Get the List of Available Books ")
	fmt.Println("5. Exit")
	fmt.Println()
}
func LibraryMenu() {
	library := &Library{
		Books:         make(map[string]Book),
		incrementedID: 1,
	}
	var choice int
	for {
		printMenu()
		fmt.Println("Type in the number of your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Error reading user choice")
			continue
		}
		switch choice {
		case 1:
			book := library.getBookInfo()
			library.AddBook(*book)
		case 2:
			library.ListAvailableBooks()
			fmt.Println("Type in the id of book to borrow: ")
			var id string
			fmt.Scan(&id)
			library.BorrowBook(id)
		case 3:
			library.ListBorrowedBooks()
			fmt.Println("Type in the id of book to return: ")
			var id string
			fmt.Scan(&id)
			library.ReturnBook(id)
		case 4:
			fmt.Println("List of Available Books:")
			library.ListAvailableBooks()

		case 5:
			fmt.Println("Exit")
			return
		default:
			fmt.Println("Invalid choice")
		}

	}
}
func (library *Library) getBookInfo() *Book {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Title: ")
	title, _ := reader.ReadString('\n') //reader reads lines until new line(Enter key)
	title = strings.TrimSpace(title)    //trimspace removes \n so input strings don't go to next lines

	fmt.Print("Author: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	return &Book{
		Title:      title,
		Author:     author,
		IsBorrowed: false,
	}
}
