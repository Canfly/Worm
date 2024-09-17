package network

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func sendViaTor(address string, payload string) string {
	proxyURL, _ := url.Parse("socks5://127.0.0.1:9050")
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	client := &http.Client{Transport: transport}
	url := fmt.Sprintf("http://%s", address)
	resp, err := client.Post(url, "application/text", strings.NewReader(payload))
	if err != nil {
		return fmt.Sprintf("Ошибка при отправке через Tor: %v", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
