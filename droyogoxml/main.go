package main

import (
	"log"
	"net/http"

	"github.com/cbdr/ats-cs-indexer/soaptesting"
	"github.com/cbdr/ats-cs-indexer/soaptesting/droyogoxml/extract"
)

func main() {
	const filename = "Sample_IT_Resume_MartinL.pdf"
	b, err := soaptesting.GetBase64EncodedBytes("../" + filename)
	if err != nil {
		log.Panicf("%+v", err)
	}
	soapClient := extract.Client{HTTPClient: http.Client{}}
	extractReq := &extract.Extract{
		Account:     soaptesting.ExtractAccount,
		Username:    soaptesting.ExtractUsername,
		Password:    soaptesting.ExtractPassword,
		FileName:    filename,
		FileContent: b,
	}
	// not sure if this is what we should put as the last
	// parameter or not.
	resp := extract.ExtractException{}
	if err := soapClient.Do("POST", soaptesting.ExtractWSDLURL, "", extractReq, &resp); err != nil {
		panic(err)
	}
}
