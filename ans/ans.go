package ans

// ResolveANS преобразует ANS-имя в адрес и тип сети
func ResolveANS(ansName string) (string, string) {
	// Для демонстрации возвращаем фиктивные данные
	switch ansName {
	case "node.amalgam":
		return "127.0.0.1:8080", "tcp"
	case "node.tor":
		return "someonionaddress.onion", "tor"
	default:
		return "", ""
	}
}
