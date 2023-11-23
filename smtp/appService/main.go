package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MailRequest struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Message string   `json:"message"`
	Type    string   `json:"type"`
}

type MailResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

type AppResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func main() {
	router := gin.Default()

	router.POST("/send", func(c *gin.Context) {
		var mailRequest MailRequest
		if err := c.BindJSON(&mailRequest); err != nil {
			c.JSON(http.StatusBadRequest, AppResponse{
				Success: false,
				Message: "Invalid request body",
				Error:   err.Error(),
			})
			return
		}

		mailResponse, err := SendMail(mailRequest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, AppResponse{
				Success: false,
				Message: "Failed to forward request to Mail Services",
				Error:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, AppResponse{
			Success: mailResponse.Success,
			Message: mailResponse.Message,
			Error:   mailResponse.Error,
		})
	})

	log.Fatal(router.Run(":8081"))
}

func SendMail(request MailRequest) (*MailResponse, error) {
	mailServiceURL := "http://localhost:8080/send"

	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(mailServiceURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var mailResponse MailResponse
	if err := json.Unmarshal(body, &mailResponse); err != nil {
		return nil, err
	}

	return &mailResponse, nil
}
