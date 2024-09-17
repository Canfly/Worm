package network

import (
	"worm/cache"
)

// Отправка запроса узлу через соответствующую сеть
func SendRequestToNode(nodeInfo *cache.NodeCache, payload string) string {
	switch nodeInfo.NetworkType {
	case "tcp":
		return sendViaTCP(nodeInfo.Address, payload)
	case "tor":
		return sendViaTor(nodeInfo.Address, payload)
	// Добавьте обработку других типов сетей
	default:
		return "Неподдерживаемый тип сети"
	}
}
