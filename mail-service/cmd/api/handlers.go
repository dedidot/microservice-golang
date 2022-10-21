package main

import (
	"log"
	"net/http"
)

func (c *Config) SendMail(w http.ResponseWriter, r *http.Request) {
	type mailMessage struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}

	var requestPayload mailMessage
	err := c.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		c.errorJSON(w, err)
		return
	}

	var msg = Message{
		To:      requestPayload.To,
		From:    requestPayload.From,
		Subject: requestPayload.Subject,
		Data:    requestPayload.Message,
	}

	err = c.Mailer.SendSMTPMessage(msg)
	if err != nil {
		log.Println(err)
		c.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Send to " + requestPayload.To,
	}

	c.writeJSON(w, http.StatusAccepted, payload)
}
