package extension

import (
	"encoding/json"
	"log"
	"net/http"
)

func ReadExtensionHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("ReadExtensionHandler")
	extensions := []string{"xlsx", "xls"}

	response, err := json.Marshal(extensions)
	if err != nil {
		http.Error(w, "failed marshal extension", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
