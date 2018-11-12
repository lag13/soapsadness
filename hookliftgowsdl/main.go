package main

import (
	"fmt"
	"log"

	"github.com/hooklift/gowsdl/soap"
	"github.com/lag13/soapsadness"
	"github.com/lag13/soapsadness/hooklift/extract"
)

func main() {
	const filename = "Sample_IT_Resume_MartinL.pdf"
	b, err := soapsadness.GetBase64EncodedBytes("../" + filename)
	if err != nil {
		log.Panicf("%+v", err)
	}
	soapClient := soap.NewClient(soapsadness.ExtractWSDLURL)
	extractClient := extract.NewExtractInterface(soapClient)
	extractReq := &extract.Extract{
		Account:     soapsadness.ExtractAccount,
		Username:    soapsadness.ExtractUsername,
		Password:    soapsadness.ExtractPassword,
		FileName:    filename,
		FileContent: b,
	}
	resp, err := extractClient.Extract(extractReq)
	if err != nil {
		if extractErr, ok := err.(*soap.SOAPFault); ok {
			panic(fmt.Sprintf("code=%v, string=%v, actor=%v, detail=%v", extractErr.Code, extractErr.String, extractErr.Actor, extractErr.Detail))
		}
		panic(err)
	}
	fmt.Printf("%+v\n", resp)
}
