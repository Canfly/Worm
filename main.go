package main

import (
	"fmt"
	"time"
	"worm/ans"
	"worm/cache"
	"worm/network"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	fmt.Println("Запуск сервера Worm...")

	// Запуск сервера ZeroMQ для получения запросов от Amalgam
	startWormServer()
}

func startWormServer() {
	socket, _ := zmq.NewSocket(zmq.REP)
	defer socket.Close()

	socket.Bind("tcp://*:5555")

	for {
		request, _ := socket.Recv(0)
		ansName, payload := parseRequest(request)

		// Проверка кэша
		nodeInfo := cache.GetNodeFromCache(ansName)
		if nodeInfo == nil {
			// Разрешение ANS для получения адреса и типа сети
			address, networkType := ans.ResolveANS(ansName)
			if address == "" {
				fmt.Printf("Узел %s не найден\n", ansName)
				socket.Send("Узел не найден", 0)
				continue
			}
			nodeInfo = &cache.NodeCache{
				Address:     address,
				NetworkType: networkType,
				LastAccess:  time.Time{},
				TTL:         0,
			}
			cache.AddNodeToCache(ansName, *nodeInfo)
		}

		fmt.Printf("Узел найден: %s (Сеть: %s)\n", nodeInfo.Address, nodeInfo.NetworkType)

		// Отправка запроса узлу через соответствующую сеть
		response := network.SendRequestToNode(nodeInfo, payload)

		// Отправка ответа обратно в Amalgam
		socket.Send(response, 0)
	}
}

func parseRequest(request string) (string, string) {
	// Формат запроса: "ANS:node.amalgam|payload"
	var ansName, payload string
	fmt.Sscanf(request, "ANS:%s|%s", &ansName, &payload)
	return ansName, payload
}
