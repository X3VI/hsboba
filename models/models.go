package models

// Offering definiert die Struktur für ein Angebot, inklusive aller relevanten Eigenschaften wie ID, Artikelname, Preis und Kontaktinformationen
type Offering struct {
	Id      int    `json:"id" db:"id"`           // Eindeutige Identifikationsnummer des Angebots, wird in DB mit automatischer Sequenz hinterlegt. Benötigt daher keine Angabe, wenn bspw. eine POST Anfrage gehandelt wird
	Item    string `json:"item" db:"item"`       // Bezeichnung des Artikels
	Preis   string `json:"preis" db:"preis"`     // Preis des Artikels
	Kontakt string `json:"kontakt" db:"kontakt"` // Kontaktinformationen für das Angebot
}

// Response stellt eine standardisierte Struktur für API-Antworten bereit, einschließlich Status und optionale Nachrichten oder Daten
type Response struct {
	Message string      `json:"message,omitempty"` // Optionale Nachricht, genutzt für Fehlermeldungen oder Erfolgsnachrichten
	Status  bool        `json:"status"`            // Status der Operation: true für Erfolg, false für Fehler
	Data    interface{} `json:"data,omitempty"`    // Optionale Daten, die mit der Antwort zurückgegeben werden können
}
