package main

import (
	"log"
	"net/http"

	"github.com/lag13/soapsadness"
	"github.com/lag13/soapsadness/droyogoxml/extract"
)

func main() {
	const filename = "Sample_IT_Resume_MartinL.pdf"
	b, err := soapsadness.GetBase64EncodedBytes("../" + filename)
	if err != nil {
		log.Panicf("%+v", err)
	}
	soapClient := extract.Client{HTTPClient: http.Client{}}
	extractReq := &extract.Extract{
		Account:     soapsadness.ExtractAccount,
		Username:    soapsadness.ExtractUsername,
		Password:    soapsadness.ExtractPassword,
		FileName:    filename,
		FileContent: b,
	}
	// not sure if this is what we should put as the last
	// parameter or not.
	resp := extract.ExtractException{}
	if err := soapClient.Do("POST", soapsadness.ExtractWSDLURL, "", extractReq, &resp); err != nil {
		panic(err)
	}
}
