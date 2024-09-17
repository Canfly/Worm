package cache

import (
	"time"
)

type NodeCache struct {
	Address     string
	NetworkType string
	LastAccess  time.Time
	TTL         time.Duration
}

var nodeCache = make(map[string]NodeCache)

// Получение узла из кэша
func GetNodeFromCache(ansName string) *NodeCache {
	if node, exists := nodeCache[ansName]; exists {
		if time.Since(node.LastAccess) < node.TTL {
			return &node
		} else {
			// Удаляем устаревший узел
			delete(nodeCache, ansName)
		}
	}
	return nil
}

// Добавление узла в кэш
func AddNodeToCache(ansName string, node NodeCache) {
	node.LastAccess = time.Now()
	node.TTL = 5 * time.Minute // Время жизни 5 минут
	nodeCache[ansName] = node
}
