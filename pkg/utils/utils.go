package utils

import (
	"encoding/json"
)

// Message object wraps a string into a JSON object to be digestable by the front end app
type Message struct {
	Message string `json:"message"`
}

// GetJSONMessageAsString serialises the JSON object containing the message string
func GetJSONMessageAsString(message string) string {
	resMessage := &Message{
		Message: message,
	}
	result, _ := json.Marshal(resMessage)
	return string(result)
}
