package main

import (
	"log"

	"github.com/lag13/soapsadness"
	"github.com/tiaguinho/gosoap"
)

func main() {
	const filename = "Sample_IT_Resume_MartinL.pdf"
	b, err := soapsadness.GetBase64EncodedBytes("../" + filename)
	if err != nil {
		log.Panicf("%+v", err)
	}
	soapClient, err := gosoap.SoapClient(soapsadness.ExtractWSDLURL)
	if err != nil {
		panic(err)
	}
	params := gosoap.Params{
		"Account":     soapsadness.ExtractAccount,
		"Username":    soapsadness.ExtractUsername,
		"Password":    soapsadness.ExtractPassword,
		"FileName":    filename,
		"FileContent": b,
	}
	if err := soapClient.Call("Extract", params); err != nil {
		panic(err)
	}
}
