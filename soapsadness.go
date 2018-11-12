package soapsadness

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

var (
	ExtractWSDLURL  = "https://staging.textkernel.com/match/soap/extract?wsdl"
	ExtractAccount  = os.Getenv("EXTRACT_ACCOUNT")
	ExtractUsername = os.Getenv("EXTRACT_USERNAME")
	ExtractPassword = os.Getenv("EXTRACT_PASSWORD")
)

func GetBase64EncodedBytes(filename string) ([]byte, error) {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrapf(err, "reading file %s", filename)
	}
	base64FileBytes := new(bytes.Buffer)
	encoder := base64.NewEncoder(base64.StdEncoding, base64FileBytes)
	encoder.Write(fileBytes)
	encoder.Close()
	return base64FileBytes.Bytes(), nil
}
