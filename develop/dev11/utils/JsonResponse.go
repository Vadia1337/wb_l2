package utils

import (
	"encoding/json"
	"net/http"
)

func SendJsonResponse(w http.ResponseWriter, httpStatusCode int, resp interface{}) {
	w.WriteHeader(httpStatusCode)
	w.Header().Set("Content-Type", "application/json")

	if httpStatusCode == 200 {
		resp = map[string]interface{}{"result": resp}
	} else {
		resp = map[string]interface{}{"error": resp}
	}

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "ошибка при сериализации ответа сервера", http.StatusServiceUnavailable)
	}
}
