package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	ID       int
	Title    string
	Author   string
	Quantity int
	Borrower string
}

var books []Book

const filename = "data.txt"

func main() {
	loadData()

	for {
		fmt.Println("\nPeminjaman Buku")
		fmt.Println("1. Add Book")
		fmt.Println("2. View All Books")
		fmt.Println("3. Update Book")
		fmt.Println("4. Delete Book")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)
		fmt.Println("")

		switch choice {
		case 1:
			addBook()
		case 2:
			viewBooks()
		case 3:
			updateBook()
		case 4:
			deleteBook()
		case 5:
			saveData()
			fmt.Println("Exiting program.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func loadData() {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "|")
		if len(fields) != 5 {
			fmt.Println("Invalid data format:", line)
			continue
		}
		id, _ := strconv.Atoi(fields[0])
		quantity, _ := strconv.Atoi(fields[3])
		book := Book{
			ID:       id,
			Title:    fields[1],
			Author:   fields[2],
			Quantity: quantity,
			Borrower: fields[4],
		}
		books = append(books, book)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func saveData() {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error saving data:", err)
		return
	}
	defer file.Close()

	for _, book := range books {
		line := fmt.Sprintf("%d|%s|%s|%s|%d\n", book.ID, book.Borrower, book.Title, book.Author, book.Quantity)
		file.WriteString(line)
	}
}

func addBook() {
	var id int
	var title, author, borrower string
	var quantity int

	fmt.Print("Book ID: ")
	fmt.Scanln(&id)

	for _, book := range books {
		if book.ID == id {
			fmt.Println("Book with ID already exists. Please choose a different ID.")
			return
		}
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Borrower Name: ")
	if scanner.Scan() {
		borrower = scanner.Text()
		if borrower == "" {
			fmt.Println("Borrower name cannot be empty. Please enter a valid name.")
			return
		}
	}
	fmt.Print("Enter Book Title: ")
	if scanner.Scan() {
		title = scanner.Text()
		if title == "" {
			fmt.Println("Title cannot be empty. Please enter a valid title.")
			return
		}
	}

	fmt.Print("Enter Book Author: ")
	if scanner.Scan() {
		author = scanner.Text()
		if author == "" {
			fmt.Println("Author cannot be empty. Please enter a valid author.")
			return
		}
	}

	fmt.Print("Quantity: ")
	fmt.Scanln(&quantity)

	if quantity <= 0 {
		fmt.Println("Invalid quantity. Please enter a valid quantity.")
		return
	}

	book := Book{
		ID:       id,
		Title:    title,
		Author:   author,
		Quantity: quantity,
		Borrower: borrower,
	}
	books = append(books, book)
	fmt.Println("Book added successfully.")

	// Save the updated data to the file
	saveData()
}

func viewBooks() {
	if len(books) == 0 {
		fmt.Println("No books available.")
		return
	}

	maxTitleLength := 5    // Mininum length for the Title header
	maxAuthorLength := 6   // Mininum length for the Author header
	maxBorrowerLength := 8 // Mininum length for the Borrower header
	maxIDLength := 2       // Minimum length for the ID header
	maxQuantityLength := 8 // Minimum length for the Quantity header

	// Calculate the maximum length of each column
	for _, book := range books {
		if len(book.Title) > maxTitleLength {
			maxTitleLength = len(book.Title)
		}
		if len(book.Author) > maxAuthorLength {
			maxAuthorLength = len(book.Author)
		}
		if len(book.Borrower) > maxBorrowerLength {
			maxBorrowerLength = len(book.Borrower)
		}
		if len(strconv.Itoa(book.ID)) > maxIDLength {
			maxIDLength = len(strconv.Itoa(book.ID))
		}
		if len(strconv.Itoa(book.Quantity)) > maxQuantityLength {
			maxQuantityLength = len(strconv.Itoa(book.Quantity))
		}
	}

	// Print the header with adjusted column lengths
	fmt.Printf("%-*s | %-*s | %-*s | %-*s | %-*s\n", maxIDLength, "ID", maxBorrowerLength, "Borrower", maxTitleLength, "Title", maxAuthorLength, "Author", maxQuantityLength, "Quantity")
	fmt.Println(strings.Repeat("-", maxIDLength+maxBorrowerLength+maxTitleLength+maxAuthorLength+maxQuantityLength+12)) // Total length is sum of max lengths + 12 for column separators

	// Print each book information with adjusted column lengths
	for _, book := range books {
		fmt.Printf("%-*d | %-*s | %-*s | %-*s | %-*d\n", maxIDLength, book.ID, maxBorrowerLength, book.Borrower, maxTitleLength, book.Title, maxAuthorLength, book.Author, maxQuantityLength, book.Quantity)
	}
}

func updateBook() {
	fmt.Print("Book ID to update: ")
	var id int
	fmt.Scanln(&id)

	var found bool
	for i, book := range books {
		if book.ID == id {
			found = true

			fmt.Println("new details for the book:")
			scanner := bufio.NewScanner(os.Stdin)

			fmt.Print("new borrower name: ")
			if scanner.Scan() {
				newBorrower := scanner.Text()
				if newBorrower != "" {
					books[i].Borrower = newBorrower
				}
			}
			fmt.Print("new title: ")
			if scanner.Scan() {
				newTitle := scanner.Text()
				if newTitle != "" {
					books[i].Title = newTitle
				}
			}

			fmt.Print("new author: ")
			if scanner.Scan() {
				newAuthor := scanner.Text()
				if newAuthor != "" {
					books[i].Author = newAuthor
				}
			}

			fmt.Print("new quantity: ")
			fmt.Scanln(&books[i].Quantity)
			if books[i].Quantity <= 0 {
				fmt.Println("Quantity cannot be zero or negative. Quantity remains unchanged.")
				return
			}

			saveData() // Simpan data setelah memperbarui buku
			fmt.Println("Book updated successfully.")
			return
		}
	}

	if !found {
		fmt.Println("Book not found.")
	}
}

func deleteBook() {
	fmt.Print("Book ID to delete: ")
	var id int
	fmt.Scanln(&id)

	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			fmt.Println("Book deleted successfully.")

			// Simpan perubahan ke file data.txt setelah menghapus buku
			saveData()
			return
		}
	}

	fmt.Println("Book not found.")
}
