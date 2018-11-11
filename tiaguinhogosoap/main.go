package main

import (
	"log"

	"github.com/cbdr/ats-cs-indexer/soaptesting"
	"github.com/tiaguinho/gosoap"
)

func main() {
	const filename = "Sample_IT_Resume_MartinL.pdf"
	b, err := soaptesting.GetBase64EncodedBytes("../" + filename)
	if err != nil {
		log.Panicf("%+v", err)
	}
	soapClient, err := gosoap.SoapClient(soaptesting.ExtractWSDLURL)
	if err != nil {
		panic(err)
	}
	params := gosoap.Params{
		"Account":     soaptesting.ExtractAccount,
		"Username":    soaptesting.ExtractUsername,
		"Password":    soaptesting.ExtractPassword,
		"FileName":    filename,
		"FileContent": b,
	}
	if err := soapClient.Call("Extract", params); err != nil {
		panic(err)
	}
}
