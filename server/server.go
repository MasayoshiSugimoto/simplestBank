package server

import (
	"domain"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

var bank domain.Bank

func StartServer() {
	http.Handle("/client/new", http.HandlerFunc(handleCreateClient))
	fileServer := http.FileServer(http.Dir("../static"))
	http.Handle("/", fileServer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server error:", err)
		panic(err)
	}
}

func handleLogin(response http.ResponseWriter, request *http.Request) {
	log.Println("Login request received")
}

func handleCreateClient(response http.ResponseWriter, request *http.Request) {
	log.Println("Request to create client: ", request)

	phoneNumber := request.FormValue("phone_number")
	firstName := request.FormValue("first_name")
	lastName := request.FormValue("last_name")

	log.Println("{firstName:", firstName, "lastName:", lastName, "phoneNumber", phoneNumber, "}")
	log.Printf("{form: %v}", request.Form)

	phoneNumberAsInt, phoneNumberError := strconv.Atoi(phoneNumber)

	type result struct {
		Success bool
		Message string
	}

	res := result{}

	res.Success = false
	if request.Method != "POST" {
		return
	} else if phoneNumber == "" || firstName == "" || lastName == "" {
		res.Message = "Invalid client information"
	} else if phoneNumberError != nil {
		res.Message = "Invalid client information"
		log.Println(phoneNumberError)
	} else {
		client := domain.NewClient(firstName, lastName, phoneNumberAsInt)
		bank.SaveClient(client)
		res.Success = true
		res.Message = "Client saved"
	}
	log.Println(res.Message)

	resultAsJSON, err := json.Marshal(res)
	if err != nil {
		log.Println("Failed to marshal:", err)
	}

	_, err = response.Write(resultAsJSON)
	if err != nil {
		log.Println("Failure while generating the response: ", err)
	}
}
