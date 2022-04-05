package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"flag"
	"log"
)

var data = `{"deposit_id":"C3C77926-9B30-4770-BA26-FD0FFDDEC930","date":"2021-02-10T12:44:14.907Z"}`
var key = "7473fded97d74274968348b42f9433ef58e5bed5411a9924725c90b6b8b6ca0f"

// body es la estructura para recibir la información del JSON-string
type body struct {
	ID   string `json:"deposit_id"`
	Date string `json:"date"`
}

func main() {

	signature := flag.String("s", "", "Firma a validar")
	flag.Parse()
	if *signature == "" {

		log.Panicln("Sin firma para validar")
	}

	var b body
	// Parseamos el string dentro de la estructura
	if err := json.Unmarshal([]byte(data), &b); err != nil {

		log.Fatalln(err.Error())
	}

	// Aquí generamos un buffer y le agregamos el id y la fecha del depósito
	var bts bytes.Buffer
	bts.WriteString(b.ID)
	bts.WriteString(b.Date)

	// Generamos la firma
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write(bts.Bytes())
	generated := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	log.Println("Firma recibida: ", *signature)
	log.Println("Firma generada: ", generated)

	equal := hmac.Equal([]byte(generated), []byte(*signature))
	log.Println("Las firmas coinciden: ", equal)
}
