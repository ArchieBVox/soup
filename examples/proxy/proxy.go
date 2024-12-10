package main

import (
	"fmt"

	"github.com/ByteSizedMarius/soup"
)

func main() {
	const proxy = "http://8.219.97.248:80"

	proxyClient, err := soup.BuildClientFromProxyAddress(proxy)
	if err != nil {
		fmt.Println(err)
		return
	}
	soup.SetClient(proxyClient)

	resp, err := soup.Get("https://api.ipify.org/?format=html")
	if err != nil {
		fmt.Println(err)
		return
	}
	doc := soup.HTMLParse(resp)
	text := doc.Find("body")

	fmt.Println(text.Text())
}
