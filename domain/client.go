package domain

import (
	"fmt"
	"log"
)

var clientCounter = 0

type Client struct {
	id        int
	firstName string
	lastName  string
	telephone int
}

func test() {
	fmt.Println("vim-go")
}

func NewClient(firstName string, lastName string, telephone int) *Client {
	clientCounter++
	client := Client{
		id:        clientCounter,
		firstName: firstName,
		lastName:  lastName,
		telephone: telephone,
	}
	log.Println("Creating client:", client)
	return &client
}

func (client *Client) Id() int {
	return client.id
}

func (client *Client) PhoneNumber() int {
	return client.telephone
}

func (client *Client) IsVoid() bool {
	return client.id == 0
}
