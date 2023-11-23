package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"strings"

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

func main() {
	router := gin.Default()

	router.POST("/send", func(c *gin.Context) {
		var mailRequest MailRequest
		if err := c.BindJSON(&mailRequest); err != nil {
			c.JSON(http.StatusBadRequest, MailResponse{
				Success: false,
				Message: "Invalid request body",
				Error:   err.Error(),
			})
			return
		}

		err := sendEmail(mailRequest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, MailResponse{
				Success: false,
				Message: "Failed to send email",
				Error:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, MailResponse{
			Success: true,
			Message: "Email sent successfully",
		})
	})

	log.Fatal(router.Run(":8080"))
}

func sendEmail(request MailRequest) error {
	auth := smtp.PlainAuth("", "dias.ulhaq30@gmail.com", "tchr haye ygeo uzem", "smtp.gmail.com")

	message := fmt.Sprintf("Subject: %s\r\n", request.Subject)
	message += fmt.Sprintf("To: %s\r\n", strings.Join(request.To, ","))
	message += fmt.Sprintf("From: %s\r\n", request.From)
	message += "Content-Type: text/" + request.Type + "; charset=UTF-8\r\n\r\n"
	message += request.Message

	err := smtp.SendMail("smtp.gmail.com:587", auth, "your_email@gmail.com", request.To, []byte(message))
	if err != nil {
		return err
	}

	return nil
}
