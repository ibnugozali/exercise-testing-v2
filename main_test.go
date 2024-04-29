package main

import (
	"os"
	"testing"
)

func TestCRUDOperations(t *testing.T) {
	t.Run("TestAddBook", testAddBook)
	t.Run("TestUpdateBook", testUpdateBook)
	t.Run("TestDeleteBook", testDeleteBook)
}

func testAddBook(t *testing.T) {
	books = []Book{}
	addBook()

	if len(books) != 1 {
		t.Errorf("Expected 1 book, got %d", len(books))
	}

	// Clean up
	books = []Book{}
}

func testUpdateBook(t *testing.T) {
	books = []Book{{ID: 1, Title: "Test Book", Author: "Test Author", Quantity: 5}}
	updateBook()

	if books[0].Quantity != 5 {
		t.Errorf("Expected quantity to remain unchanged, got %d", books[0].Quantity)
	}

	// Clean up
	books = []Book{}
}

func testDeleteBook(t *testing.T) {
	books = []Book{{ID: 1, Title: "Test Book", Author: "Test Author", Quantity: 5}}
	deleteBook()

	if len(books) != 0 {
		t.Errorf("Expected no books after deletion, got %d", len(books))
	}

	// Clean up
	books = []Book{}
	os.Remove(filename)
}
