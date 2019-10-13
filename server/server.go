package server

import (
	"domain"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var bank domain.Bank

var loginScreen = `
<html>

<header><title>Simplest Bank</title></header>

<body>
	<h1>Simplest Bank</h1>
	<form action="/login">
		<span>Phone Number</span><input id="phone_number" type="text"></input>
		</br>
		<input type="submit" value="Login"></input>
	</form>
</body>

</html>`

var adminScreen *template.Template

func init() {
	var err error
	adminScreen, err = template.New("adminScreen").Parse(`
	<html>

	<header><title>Simplest Bank</title></header>

	<body>
		<h1>Simplest Bank: Admin Page</h1>
		<div>{{.}}</div>
		<h2>Regiter Client</h2>
		<form action="adminScreen">
			<span>First Name</span><input id="first_name" name="first_name" type="text"></input>
			</br>
			<span>Last Name</span><input id="last_name" name="last_name" type="text"></input>
			</br>
			<span>Phone Number</span><input id="phone_number" name="phone_number" type="text"></input>
			</br>
			<input id="create_person_button" type="submit" value="Register Client"></input>
		</form>
		<Phone Number>
	</body>

	</html>`)

	if err != nil {
		log.Fatal("Failed to parse admin screen template: ", err)
	}

}

func StartServer() {
	http.Handle("/loginScreen", http.HandlerFunc(handleLoginScreen))
	http.Handle("/adminScreen", http.HandlerFunc(handleAdminScreen))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server error:", err)
		panic(err)
	}
}

func handleLoginScreen(response http.ResponseWriter, request *http.Request) {
	log.Println("Login screen request received")
	_, err := response.Write([]byte(loginScreen))
	if err != nil {
		log.Println("Failure: ", err)
	}
}

func handleLogin(response http.ResponseWriter, request *http.Request) {
	log.Println("Login request received")
}

func handleAdminScreen(response http.ResponseWriter, request *http.Request) {
	log.Println("Adming screen request")

	phoneNumber := request.FormValue("phone_number")
	firstName := request.FormValue("first_name")
	lastName := request.FormValue("last_name")

	phoneNumberAsInt, phoneNumberError := strconv.Atoi(phoneNumber)

	var message string
	if phoneNumber == "" && firstName == "" && lastName == "" {
		//Do nothing
	} else if phoneNumber == "" || firstName == "" || lastName == "" {
		message = "Invalid client information"
	} else if phoneNumberError != nil {
		message = phoneNumberError.Error()
		log.Println(phoneNumberError)
	} else {
		client := domain.NewPerson(firstName, lastName, phoneNumberAsInt)
		bank.SavePerson(client)
		if message == "" {
			message = "Client saved"
		}
	}

	err := adminScreen.Execute(response, message)
	if err != nil {
		log.Println("Failure: ", err)
	}
}
