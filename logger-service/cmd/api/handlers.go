package main

import (
	"fmt"
	"log-service/data"
	"net/http"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLogg(w http.ResponseWriter, r *http.Request) {
	// Read json into var
	var requestPayload JSONPayload
	_ = app.readJSON(w, r, &requestPayload)

	// insert data
	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}

	fmt.Println(event)
	err := app.Models.LogEntry.Insert(event)
	fmt.Println("Log after insert logs")
	fmt.Println(err)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var response jsonResponse
	response.Error = false
	response.Message = "Logged"

	app.writeJSON(w, http.StatusAccepted, response)
}
