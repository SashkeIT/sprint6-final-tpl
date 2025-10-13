package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {

	file, _, err := r.FormFile("File")
	if err != nil {
		message := fmt.Sprintf("File retrieval error: %v", err)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		message := fmt.Sprintf("File read error: %v", err)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	content := string(data)
	encodedContent, err := service.ToggleMorse(content)
	if err != nil {
		message := fmt.Sprintf("Service conversion error: %v", err)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	name := time.Now().UTC().String()
	err = os.WriteFile(name, []byte(encodedContent), 0644)
	if err != nil {
		message := fmt.Sprintf("File save error: %v", err)
		http.Error(w, message, http.StatusInternalServerError)
	}

	r.Header.Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, encodedContent)

}
