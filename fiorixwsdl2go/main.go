package main

import (
	"fmt"
	"log"

	"github.com/cbdr/ats-cs-indexer/soaptesting"
	"github.com/cbdr/ats-cs-indexer/soaptesting/fiorixwsdl2go/extract"
	"github.com/fiorix/wsdl2go/soap"
)

func main() {
	var filename = "Sample_IT_Resume_MartinL.pdf"
	b, err := soaptesting.GetBase64EncodedBytes("../" + filename)
	if err != nil {
		log.Panicf("%+v", err)
	}
	soapClient := soap.Client{
		URL:       soaptesting.ExtractWSDLURL,
		Namespace: extract.Namespace,
	}
	soapService := extract.NewExtractInterface(&soapClient)
	resp, err := soapService.Extract(&extract.Extract{
		Account:     &soaptesting.ExtractAccount,
		Username:    &soaptesting.ExtractUsername,
		Password:    &soaptesting.ExtractPassword,
		FileName:    &filename,
		FileContent: &b,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("printing out returned response")
	fmt.Printf("%+v\n", *resp.Return)
}
