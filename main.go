package main

import (
	"fmt"

	"hsboba/handlers"
	"hsboba/store"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

// Die Main Funktion bildet den Einstiegspunkt zum Starten des Services
func main() {
	// Verbindet mit der Datenbank unter Verwendung der angegebenen Datenbankinformationen
	db, err := sqlx.Connect("postgres", "user= password= dbname= sslmode=disable")
	if err != nil {
		fmt.Printf("Fehler beim Verbinden zur Datenbank aufgetreten: %s\n", err)
		return
	}
	fmt.Println("Erfolgreich zur Datenbank verbunden.")

	// Initialisiert das Echo-Webframework und konfiguriert Middleware für CORS (Cross-Origin Resource Sharing)
	e := echo.New()
	e.Use(middleware.CORS())

	// Erstellt ein Store-Objekt mit der Datenbankverbindung
	s := store.NewStore(db)

	// Definiert die Routen und weist ihnen die entsprechenden Handler-Funktionen zu. Jede Route behandelt spezifische API-Anfragen
	// Endpunkte koennen mit Postman oder manuell, z.B. mit Hilfe des Befehls curl -X HTTP-METHODE http:// ..., gegengetestet werden
	// Die Angebote werden in JSON repräsentiert und folgen dem Format, welches für die Struktur des Angebots in models.go definiert wurde
	e.GET("/offerings", handlers.GetAllOfferings(s))       // Zum Abrufen aller Angebote: GET - http://localhost:1323/offerings
	e.GET("/offerings/:id", handlers.GetOffering(s))       // Zum Abrufen eines spezifischen Angebots: GET - http://localhost:1323/offerings/id
	e.POST("/offerings", handlers.AddOffering(s))          // Zum Erstellen eines Angebots: POST - http://localhost:1323/offerings
	e.PUT("/offerings/:id", handlers.UpdateOffering(s))    // Zum Erstellen oder Überschreiben eines spezifischen Angebots: PUT - http://localhost:1323/offerings/id
	e.DELETE("/offerings/:id", handlers.DeleteOffering(s)) // Zum Löschen eines spezifischen Angebots: POST - http://localhost:1323/offerings/id

	// Startet den Server auf dem angegebenen Port und hört auf eingehende Anfragen
	e.Logger.Fatal(e.Start(":1323"))
}
