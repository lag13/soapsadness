package main

import (
	"fmt"
	"log"

	"github.com/cbdr/ats-cs-indexer/soaptesting"
	"github.com/cbdr/ats-cs-indexer/soaptesting/hookliftgowsdl/extract"
	"github.com/hooklift/gowsdl/soap"
)

func main() {
	const filename = "Sample_IT_Resume_MartinL.pdf"
	b, err := soaptesting.GetBase64EncodedBytes("../" + filename)
	if err != nil {
		log.Panicf("%+v", err)
	}
	soapClient := soap.NewClient(soaptesting.ExtractWSDLURL)
	extractClient := extract.NewExtractInterface(soapClient)
	extractReq := &extract.Extract{
		Account:     soaptesting.ExtractAccount,
		Username:    soaptesting.ExtractUsername,
		Password:    soaptesting.ExtractPassword,
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
