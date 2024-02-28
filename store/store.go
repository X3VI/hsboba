package store

import (
	"github.com/jmoiron/sqlx"
)

// Store bietet eine Kapselung um die Datenbankverbindung. Es ermöglicht die Ausführung von Datenbankoperationen über ein Store-Objekt.
type Store struct {
	DB *sqlx.DB // Referenz auf die SQL-Datenbankverbindung
}

// NewStore ist ein Konstruktor für ein Store-Objekt. Es nimmt eine Datenbankverbindung entgegen und gibt eine neue Store-Instanz zurück.
func NewStore(db *sqlx.DB) *Store {
	return &Store{DB: db}
}
