package extract

import (
	"encoding/xml"
	"time"

	"github.com/hooklift/gowsdl/soap"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type ExtractException struct {
	XMLName xml.Name `xml:"http://home.textkernel.nl/sourcebox/soap/extract ExtractException"`

	Description string `xml:"description,omitempty"`

	Id int32 `xml:"id,omitempty"`

	OutputStatusName string `xml:"outputStatusName,omitempty"`

	Severity string `xml:"severity,omitempty"`

	Solution string `xml:"solution,omitempty"`

	TkURL string `xml:"tkURL,omitempty"`

	Message string `xml:"message,omitempty"`
}

type ExtractURL struct {
	XMLName xml.Name `xml:"http://home.textkernel.nl/sourcebox/soap/extract extractURL"`

	Account string `xml:"account,omitempty"`

	Username string `xml:"username,omitempty"`

	Password string `xml:"password,omitempty"`

	PublicURL string `xml:"publicURL,omitempty"`

	OutputType string `xml:"outputType,omitempty"`

	CustomQueryString string `xml:"customQueryString,omitempty"`

	Options []*Entry `xml:"options,omitempty"`
}

type Entry struct {
	XMLName xml.Name `xml:"http://home.textkernel.nl/sourcebox/soap/extract entry"`

	Key string `xml:"key,omitempty"`

	Value string `xml:"value,omitempty"`
}

type ExtractURLResponse struct {
	XMLName xml.Name `xml:"http://home.textkernel.nl/sourcebox/soap/extract extractURLResponse"`

	Return_ []byte `xml:"return,omitempty"`
}

type Extract struct {
	XMLName xml.Name `xml:"http://home.textkernel.nl/sourcebox/soap/extract extract"`

	Account string `xml:"account,omitempty"`

	Username string `xml:"username,omitempty"`

	Password string `xml:"password,omitempty"`

	FileName string `xml:"fileName,omitempty"`

	FileContent []byte `xml:"fileContent,omitempty"`

	TmfFileContent []byte `xml:"tmfFileContent,omitempty"`

	Apimap []byte `xml:"apimap,omitempty"`

	Options []*Entry `xml:"options,omitempty"`
}

type ExtractResponse struct {
	XMLName xml.Name `xml:"http://home.textkernel.nl/sourcebox/soap/extract extractResponse"`

	Return_ string `xml:"return,omitempty"`
}

type ExtractAdvanced struct {
	XMLName xml.Name `xml:"http://home.textkernel.nl/sourcebox/soap/extract extractAdvanced"`

	Account string `xml:"account,omitempty"`

	Username string `xml:"username,omitempty"`

	Password string `xml:"password,omitempty"`

	ClientSpecificArguments []*Entry `xml:"clientSpecificArguments,omitempty"`

	FileName string `xml:"fileName,omitempty"`

	FileContent []byte `xml:"fileContent,omitempty"`

	TmfFileContent []byte `xml:"tmfFileContent,omitempty"`

	Apimap []byte `xml:"apimap,omitempty"`

	Options []*Entry `xml:"options,omitempty"`
}

type ExtractAdvancedResponse struct {
	XMLName xml.Name `xml:"http://home.textkernel.nl/sourcebox/soap/extract extractAdvancedResponse"`

	Return_ string `xml:"return,omitempty"`
}

type ExtractInterface interface {

	// Error can be either of the following types:
	//
	//   - ExtractException

	ExtractURL(request *ExtractURL) (*ExtractURLResponse, error)

	// Error can be either of the following types:
	//
	//   - ExtractException

	ExtractAdvanced(request *ExtractAdvanced) (*ExtractAdvancedResponse, error)

	// Error can be either of the following types:
	//
	//   - ExtractException

	Extract(request *Extract) (*ExtractResponse, error)
}

type extractInterface struct {
	client *soap.Client
}

func NewExtractInterface(client *soap.Client) ExtractInterface {
	return &extractInterface{
		client: client,
	}
}

func (service *extractInterface) ExtractURL(request *ExtractURL) (*ExtractURLResponse, error) {
	response := new(ExtractURLResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *extractInterface) ExtractAdvanced(request *ExtractAdvanced) (*ExtractAdvancedResponse, error) {
	response := new(ExtractAdvancedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *extractInterface) Extract(request *Extract) (*ExtractResponse, error) {
	response := new(ExtractResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
