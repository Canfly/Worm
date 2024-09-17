package network

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func sendViaTCP(address string, payload string) string {
	url := fmt.Sprintf("http://%s", address)
	resp, err := http.Post(url, "application/text", strings.NewReader(payload))
	if err != nil {
		return fmt.Sprintf("Ошибка при отправке через TCP: %v", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
