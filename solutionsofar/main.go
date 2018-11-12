package main

import (
	"fmt"
	"log"

	"github.com/fiorix/wsdl2go/soap"
	"github.com/lag13/soapsadness"
	"github.com/lag13/soapsadness/solutionsofar/extract"
)

func main() {
	var filename = "Sample_IT_Resume_MartinL.pdf"
	b, err := soapsadness.GetBase64EncodedBytes("../" + filename)
	if err != nil {
		log.Panicf("%+v", err)
	}
	soapClient := soap.Client{
		URL:       soapsadness.ExtractWSDLURL,
		Namespace: extract.Namespace,
	}
	soapService := extract.NewExtractInterface(&soapClient)
	resp, err := soapService.Extract(&extract.Extract{
		Account:     &soapsadness.ExtractAccount,
		Username:    &soapsadness.ExtractUsername,
		Password:    &soapsadness.ExtractPassword,
		FileName:    &filename,
		FileContent: &b,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("printing out returned response")
	fmt.Printf("%+v\n", *resp.Return)
}
