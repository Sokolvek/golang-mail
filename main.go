package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/smtp"
)

type emailReq struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	ServiceType string `json:"serviceType"`
	Budged      string `json:"budged"`
	Task        string `json:"task"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/mail", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		defer r.Body.Close()

		var email emailReq
		json.Unmarshal([]byte(body), &email)
		fmt.Println(email)
		send(email)
	})
	http.ListenAndServe(":8080", mux)
}

func send(body emailReq) {
	to := "test@gmail.com" // сюда тыкаешь свой gmail
	pass := "*****"        // а сюда пароль от приложения
	from := body.Email
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Новый запрос\n\n" +
		"От кого: " + body.Email + "\n" +
		"Имя заказчика: " + body.Name + "\n" +
		"Тип услуги: " + body.ServiceType + "\n" +
		"Бюджет: " + body.Budged + "\n" +
		"Задача: " + body.Task + "\n"
	fmt.Println(msg)
	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", to, pass, "smtp.gmail.com"),
		to, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("error: %s", err)
		return
	}
}
