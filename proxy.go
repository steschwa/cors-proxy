package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/gofiber/fiber/v2"
)

var ENV_PORT string = "CORSPROXY_PORT"

func loadUrl(url string) (*bytes.Buffer, string) {
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(res.Body)

	if err != nil {
		panic(err)
	}

	mime := res.Header.Get("Content-Type")

	return buffer, mime
}

func getPort(fallback string) string {
	port, ok := os.LookupEnv(ENV_PORT)

	if !ok {
		return fallback
	}

	return port
}

func setCORSHeaders(c *fiber.Ctx) {
	c.Response().Header.Set("Access-Control-Allow-Origin", "*")
	c.Response().Header.Set("Access-Control-Allow-Headers", "*")
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		url, err := url.QueryUnescape(c.Query("url"))

		if err != nil {
			panic(err)
		}

		proxiedResponse, mime := loadUrl(url)
		setCORSHeaders(c)
		c.Response().Header.Set("Content-Type", mime)
		return c.Send(proxiedResponse.Bytes())
	})

	app.Options("/", func(c *fiber.Ctx) error {
		setCORSHeaders(c)
		return c.SendString("Allowed: *")
	})

	port := getPort("5000")
	app.Listen(fmt.Sprintf(":%s", port))
}
