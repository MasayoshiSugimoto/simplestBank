package server

import (
	"domain"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type responseWrapper struct {
	Success bool
	Data    interface{}
}

var bank domain.Bank

func StartServer() {
	http.Handle("/client/new", http.HandlerFunc(handleCreateClient))
	http.Handle("/login", http.HandlerFunc(handleLogin))
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

	phoneNumber, err := phoneNumber(request)
	if err != nil {
		writeResponse(response, responseWrapper{
			Success: false,
			Data:    "Failed to parse phone number",
		})
		return
	}

	client := bank.GetClientByPhoneNumber(phoneNumber)
	log.Println("client.id:", client.Id())
	if client.IsVoid() {
		writeResponse(response, responseWrapper{
			Success: false,
			Data:    "Not registered user",
		})
		return
	}

	writeResponse(response, responseWrapper{
		Success: true,
		Data: struct {
			ClientId int
		}{
			client.Id(),
		},
	})
}

func handleCreateClient(response http.ResponseWriter, request *http.Request) {
	log.Println("Request to create client: ", request)

	phoneNumber, phoneNumberError := phoneNumber(request)
	firstName := firstName(request)
	lastName := lastName(request)

	log.Println("{firstName:", firstName, "lastName:", lastName, "phoneNumber", phoneNumber, "}")
	log.Printf("{form: %v}", request.Form)

	type result struct {
		Success bool
		Message string
	}

	res := result{}

	res.Success = false
	if request.Method != "POST" {
		return
	} else if phoneNumberError != nil {
		res.Message = "Invalid client information"
		log.Println(phoneNumberError)
	} else if phoneNumber == 0 || firstName == "" || lastName == "" {
		res.Message = "Invalid client information"
	} else {
		bank.SaveClient(domain.NewClient(firstName, lastName, phoneNumber))
		res.Success = true
		res.Message = "Client saved"
	}
	log.Println(res.Message)

	writeResponse(response, &res)
}

func writeResponse(response http.ResponseWriter, result interface{}) {
	resultAsJSON, err := json.Marshal(result)
	if err != nil {
		log.Println("Failed to marshal:", err)
	}

	_, err = response.Write(resultAsJSON)
	if err != nil {
		log.Println("Failure while generating the response: ", err)
	}
}

func firstName(request *http.Request) string {
	return request.FormValue("first_name")
}

func lastName(request *http.Request) string {
	return request.FormValue("last_name")
}

func phoneNumber(request *http.Request) (int, error) {
	return strconv.Atoi(request.FormValue("phone_number"))
}
