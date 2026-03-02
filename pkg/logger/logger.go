package logger

import (
	"encoding/json"
	"log"

	"aygit-muhasebe-integration/pkg/db"

	"github.com/google/uuid"
)

// LogAction records a system action to the database
func LogAction(userID *uuid.UUID, action string, details interface{}, ipAddress *string) {
	detailsJSON, err := json.Marshal(details)
	if err != nil {
		log.Printf("Error marshaling log details: %v", err)
		detailsJSON = []byte("{}")
	}

	query := `INSERT INTO system_logs (id, user_id, action, details, ip_address) 
	          VALUES (uuid_generate_v4(), $1, $2, $3, $4)`

	_, err = db.DB.Exec(query, userID, action, detailsJSON, ipAddress)
	if err != nil {
		log.Printf("Error inserting system log: %v", err)
	}
}
