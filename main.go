package main

import (
	"CRUD-GO/internal/services"
	"CRUD-GO/internal/store"
	"CRUD-GO/internal/transport"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// conectar a SQLite
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close() // cerrar conexion al final

	//crear table si no existe

	q := `
	CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		author TEXT NOT NULL
	)`

	if _, err := db.Exec(q); err != nil {
		log.Fatal(err.Error())
	}

	//Inyectar dependencias

	bookStore := store.New(db)
	bookService := services.New(bookStore)
	bookHandler := transport.New(bookService)

	// configurar rutas
	http.HandleFunc("/books", bookHandler.HandleBooks)
	http.HandleFunc("/books/", bookHandler.HandleBookByID)

	fmt.Println("Ejecutandose en  http://localhost:8000")
	fmt.Println("Endpoints:")
	fmt.Println(" GET    /books        -obtener todos los libros")
	fmt.Println(" POST   /books        -crar un libro nuevo")
	fmt.Println(" GET    /books/{id}   -obtener libro especifico")
	fmt.Println(" PUT    /books/{id}   -actualizar un libro")
	fmt.Println(" DELETE /books        -eliminar un libro")

	//empezar y escuchar al serividor
	log.Fatal(http.ListenAndServe(":8000", nil))
}
