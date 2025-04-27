package handlers

import (
	"log"
	"net/http"
)

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	log.Println("Webhook handler called")
	w.WriteHeader(http.StatusAccepted)
}
