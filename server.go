package main

import (
	"fmt"
	"gopkg.in/redis.v3"
)

func ConnectNewClient() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	pubsub, err := client.Subscribe("CANAL1")

	if err != nil {
		fmt.Println("No es posible subscribirse al canal")
	}

	for {
		message, err := pubsub.ReceiveMessage()

		if err != nil {
			fmt.Println("No es posible leer el mensaje recibido")
		}

		fmt.Println(message.Channel)
		fmt.Println(message.Payload)
	}

	fmt.Println("La subscricion al canal 'CANAL1' se realizo correctamente")
}

func main() {
	fmt.Println("Running...")

	ConnectNewClient()
}