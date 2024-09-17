package main

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	sendRequestToWorm("node.amalgam", "Привет, Worm!")
}

func sendRequestToWorm(ansName string, payload string) {
	socket, _ := zmq.NewSocket(zmq.REQ)
	defer socket.Close()

	// Подключаемся к Worm
	socket.Connect("tcp://localhost:5555")

	// Отправляем запрос
	request := fmt.Sprintf("ANS:%s|%s", ansName, payload)
	socket.Send(request, 0)

	// Получаем ответ
	response, _ := socket.Recv(0)
	fmt.Println("Ответ от Worm:", response)
}
