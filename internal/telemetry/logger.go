package telemetry

import (
	"encoding/json"
	"log"
)

func Log(msg string, fields map[string]interface{}) {
	payload := map[string]interface{}{
		"message": msg,
		"fields":  fields,
	}
	out, _ := json.Marshal(payload)
	log.Println(string(out))
}
