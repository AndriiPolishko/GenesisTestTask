package main

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os"
)

const emailDatabaseFile = "emails.txt"

type SubscribeResponse struct {
	Message string `json:"message"`
}

func subscribeHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	emailCheck, _ := stringExistsInFile(emailDatabaseFile, email)

	if emailCheck {
		http.Error(w, "Email already exists", http.StatusConflict)
		return
	}

	err := saveEmail(email)
	if err != nil {
		http.Error(w, "Error saving email", http.StatusInternalServerError)
		return
	}

	response := SubscribeResponse{
		Message: "Email added",
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error marshaling JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func saveEmail(email string) error {
	file, err := os.OpenFile(emailDatabaseFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(email + "\n")
	if err != nil {
		return err
	}

	return nil
}

func stringExistsInFile(filePath string, searchStr string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == searchStr {
			return true, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return false, err
	}

	return false, nil
}
