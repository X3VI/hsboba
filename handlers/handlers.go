package handlers

import (
	"log"
	"net/http"
	"strconv"

	"hsboba/models"
	"hsboba/store"

	"github.com/labstack/echo/v4"
)

// ApiResponse erzeugt eine standardisierte JSON-Antwort für API-Anfragen
// Diese Funktion vereinfacht die Erstellung einheitlicher Antworten über alle Endpunkte hinweg
func ApiResponse(c echo.Context, statusCode int, message string, data interface{}, success bool) error {
	response := models.Response{
		Message: message,
		Status:  success,
		Data:    data,
	}
	return c.JSON(statusCode, response)
}

// GetAllOfferings ist eine Funktion, die als Handler für den Endpunkt dient, der alle Angebote abruft
// Sie nutzt das Store-Objekt, um die Datenbankabfrage durchzuführen
func GetAllOfferings(s *store.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("Abrufen aller Angebote")
		var offerings []models.Offering
		// Führt eine SQL-Abfrage aus, um alle Angebote zu erhalten, und gibt diese als JSON zurück
		if err := s.DB.Select(&offerings, "SELECT * FROM offerings"); err != nil {
			return ApiResponse(c, http.StatusInternalServerError, "Fehler beim Abrufen aller Angeboten aufgetreten", nil, false)
		}
		return ApiResponse(c, http.StatusOK, "", offerings, true)
	}
}

// GetOffering: Abrufen eines spezifischen Angebots basierend auf der ID
func GetOffering(s *store.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("Abrufen eines spezifischen Angebots")
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return ApiResponse(c, http.StatusBadRequest, "Ungueltige ID", nil, false)
		}

		var offering models.Offering
		if err := s.DB.Get(&offering, "SELECT * FROM offerings WHERE id = $1", id); err != nil {
			return ApiResponse(c, http.StatusInternalServerError, "Fehler beim Abrufen eines spezifischen Angebots aufgetreten", nil, false)
		}
		return ApiResponse(c, http.StatusOK, "", offering, true)
	}
}

// AddOffering: Hinzufügen eines neuen Angebots zur Datenbank
func AddOffering(s *store.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("Fuege neues Angebot hinzu")
		var offering models.Offering
		if err := c.Bind(&offering); err != nil {
			return ApiResponse(c, http.StatusBadRequest, err.Error(), nil, false)
		}

		_, err := s.DB.NamedExec("INSERT INTO offerings (item, preis, kontakt) VALUES (:item, :preis, :kontakt)", offering)
		if err != nil {
			return ApiResponse(c, http.StatusInternalServerError, "Fehler beim Hinzufuegen eines neuen Angebots", nil, false)
		}
		return ApiResponse(c, http.StatusOK, "Neues Angebot erfolgreich hinzugefuegt", nil, true)
	}
}

// UpdateOffering: Aktualisieren eines bestehenden Angebots basierend auf der ID
func UpdateOffering(s *store.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("Update eines spezifischen Angebots wird durchgefuehrt")
		var offering models.Offering
		if err := c.Bind(&offering); err != nil {
			return ApiResponse(c, http.StatusBadRequest, err.Error(), nil, false)
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return ApiResponse(c, http.StatusBadRequest, "Ungueltige ID", nil, false)
		}
		offering.Id = id

		_, err = s.DB.NamedExec("UPDATE offerings SET item= :item, preis= :preis, kontakt= :kontakt WHERE id= :id", offering)
		if err != nil {
			return ApiResponse(c, http.StatusInternalServerError, "Fehler beim Update eines spezifischen Angebots", nil, false)
		}
		return ApiResponse(c, http.StatusOK, "Update eines spezifischen Angebots war erfolgreich", nil, true)
	}
}

// DeleteOffering: Löschen eines Angebots basierend auf der ID
func DeleteOffering(s *store.Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("Loeschen eines Angebots")
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return ApiResponse(c, http.StatusBadRequest, "Ungueltige ID", nil, false)
		}

		_, err = s.DB.Exec("DELETE FROM offerings WHERE id = $1", id)
		if err != nil {
			return ApiResponse(c, http.StatusInternalServerError, "Fehler beim Loeschen eines Angebots", nil, false)
		}
		return ApiResponse(c, http.StatusOK, "Angebot erfolgreich geloescht", nil, true)
	}
}
