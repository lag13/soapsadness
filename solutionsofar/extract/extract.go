package extract

import (
	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "http://home.textkernel.nl/sourcebox/soap/extract"

// NewExtractInterface creates an initializes a ExtractInterface.
func NewExtractInterface(cli *soap.Client) ExtractInterface {
	return &extractInterface{cli}
}

// ExtractInterface was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type ExtractInterface interface {
	// Extract was auto-generated from WSDL.
	Extract(Extract *Extract) (*ExtractResponse, error)

	// ExtractAdvanced was auto-generated from WSDL.
	ExtractAdvanced(ExtractAdvanced *ExtractAdvanced) (*ExtractAdvancedResponse, error)

	// ExtractURL was auto-generated from WSDL.
	ExtractURL(ExtractURL *ExtractURL) (*ExtractURLResponse, error)
}

// ExtractException was auto-generated from WSDL.
type ExtractException struct {
	Description      *string `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Id               *int    `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	OutputStatusName *string `xml:"outputStatusName,omitempty" json:"outputStatusName,omitempty" yaml:"outputStatusName,omitempty"`
	Severity         *string `xml:"severity,omitempty" json:"severity,omitempty" yaml:"severity,omitempty"`
	Solution         *string `xml:"solution,omitempty" json:"solution,omitempty" yaml:"solution,omitempty"`
	TkURL            *string `xml:"tkURL,omitempty" json:"tkURL,omitempty" yaml:"tkURL,omitempty"`
	Message          *string `xml:"message,omitempty" json:"message,omitempty" yaml:"message,omitempty"`
}

// Entry was auto-generated from WSDL.
type Entry struct {
	Key   *string `xml:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
	Value *string `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
}

// Extract was auto-generated from WSDL.
type Extract struct {
	Account        *string  `xml:"account,omitempty" json:"account,omitempty" yaml:"account,omitempty"`
	Username       *string  `xml:"username,omitempty" json:"username,omitempty" yaml:"username,omitempty"`
	Password       *string  `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
	FileName       *string  `xml:"fileName,omitempty" json:"fileName,omitempty" yaml:"fileName,omitempty"`
	FileContent    *[]byte  `xml:"fileContent,omitempty" json:"fileContent,omitempty" yaml:"fileContent,omitempty"`
	TmfFileContent *[]byte  `xml:"tmfFileContent,omitempty" json:"tmfFileContent,omitempty" yaml:"tmfFileContent,omitempty"`
	Apimap         *[]byte  `xml:"apimap,omitempty" json:"apimap,omitempty" yaml:"apimap,omitempty"`
	Options        []*Entry `xml:"options,omitempty" json:"options,omitempty" yaml:"options,omitempty"`
}

// ExtractAdvanced was auto-generated from WSDL.
type ExtractAdvanced struct {
	Account                 *string  `xml:"account,omitempty" json:"account,omitempty" yaml:"account,omitempty"`
	Username                *string  `xml:"username,omitempty" json:"username,omitempty" yaml:"username,omitempty"`
	Password                *string  `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
	ClientSpecificArguments []*Entry `xml:"clientSpecificArguments,omitempty" json:"clientSpecificArguments,omitempty" yaml:"clientSpecificArguments,omitempty"`
	FileName                *string  `xml:"fileName,omitempty" json:"fileName,omitempty" yaml:"fileName,omitempty"`
	FileContent             *[]byte  `xml:"fileContent,omitempty" json:"fileContent,omitempty" yaml:"fileContent,omitempty"`
	TmfFileContent          *[]byte  `xml:"tmfFileContent,omitempty" json:"tmfFileContent,omitempty" yaml:"tmfFileContent,omitempty"`
	Apimap                  *[]byte  `xml:"apimap,omitempty" json:"apimap,omitempty" yaml:"apimap,omitempty"`
	Options                 []*Entry `xml:"options,omitempty" json:"options,omitempty" yaml:"options,omitempty"`
}

// ExtractAdvancedResponse was auto-generated from WSDL.
type ExtractAdvancedResponse struct {
	Return *string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// ExtractResponse was auto-generated from WSDL.
type ExtractResponse struct {
	Return *string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// ExtractURL was auto-generated from WSDL.
type ExtractURL struct {
	Account           *string  `xml:"account,omitempty" json:"account,omitempty" yaml:"account,omitempty"`
	Username          *string  `xml:"username,omitempty" json:"username,omitempty" yaml:"username,omitempty"`
	Password          *string  `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
	PublicURL         *string  `xml:"publicURL,omitempty" json:"publicURL,omitempty" yaml:"publicURL,omitempty"`
	OutputType        *string  `xml:"outputType,omitempty" json:"outputType,omitempty" yaml:"outputType,omitempty"`
	CustomQueryString *string  `xml:"customQueryString,omitempty" json:"customQueryString,omitempty" yaml:"customQueryString,omitempty"`
	Options           []*Entry `xml:"options,omitempty" json:"options,omitempty" yaml:"options,omitempty"`
}

// ExtractURLResponse was auto-generated from WSDL.
type ExtractURLResponse struct {
	Return *[]byte `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for Extract.
// OperationExtractResponse was auto-generated from WSDL.
type OperationExtractResponse struct {
	ExtractResponse *ExtractResponse `xml:"extractResponse,omitempty" json:"extractResponse,omitempty" yaml:"extractResponse,omitempty"`
}

// Operation wrapper for ExtractAdvanced.
// OperationExtractAdvanced was auto-generated from WSDL.
type OperationExtractAdvanced struct {
	ExtractAdvanced *ExtractAdvanced `xml:"extractAdvanced,omitempty" json:"extractAdvanced,omitempty" yaml:"extractAdvanced,omitempty"`
}

// Operation wrapper for ExtractAdvanced.
// OperationExtractAdvancedResponse was auto-generated from WSDL.
type OperationExtractAdvancedResponse struct {
	ExtractAdvancedResponse *ExtractAdvancedResponse `xml:"extractAdvancedResponse,omitempty" json:"extractAdvancedResponse,omitempty" yaml:"extractAdvancedResponse,omitempty"`
}

// Operation wrapper for ExtractURL.
// OperationExtractURL was auto-generated from WSDL.
type OperationExtractURL struct {
	ExtractURL *ExtractURL `xml:"extractURL,omitempty" json:"extractURL,omitempty" yaml:"extractURL,omitempty"`
}

// Operation wrapper for ExtractURL.
// OperationExtractURLResponse was auto-generated from WSDL.
type OperationExtractURLResponse struct {
	ExtractURLResponse *ExtractURLResponse `xml:"extractURLResponse,omitempty" json:"extractURLResponse,omitempty" yaml:"extractURLResponse,omitempty"`
}

// extractInterface implements the ExtractInterface interface.
type extractInterface struct {
	cli *soap.Client
}

// Extract was auto-generated from WSDL.
func (p *extractInterface) Extract(ext *Extract) (*ExtractResponse, error) {
	α := struct {
		Extract *Extract `xml:"tns:extract"`
	}{
		Extract: ext,
	}

	γ := OperationExtractResponse{}
	if err := p.cli.RoundTripWithAction("Extract", α, &γ); err != nil {
		return nil, err
	}
	return γ.ExtractResponse, nil
}

// ExtractAdvanced was auto-generated from WSDL.
func (p *extractInterface) ExtractAdvanced(ExtractAdvanced *ExtractAdvanced) (*ExtractAdvancedResponse, error) {
	α := struct {
		OperationExtractAdvanced `xml:"tns:extractAdvanced"`
	}{
		OperationExtractAdvanced{
			ExtractAdvanced,
		},
	}

	γ := struct {
		OperationExtractAdvancedResponse `xml:"extractAdvancedResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("ExtractAdvanced", α, &γ); err != nil {
		return nil, err
	}
	return γ.ExtractAdvancedResponse, nil
}

// ExtractURL was auto-generated from WSDL.
func (p *extractInterface) ExtractURL(ExtractURL *ExtractURL) (*ExtractURLResponse, error) {
	α := struct {
		OperationExtractURL `xml:"tns:extractURL"`
	}{
		OperationExtractURL{
			ExtractURL,
		},
	}

	γ := struct {
		OperationExtractURLResponse `xml:"extractURLResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("ExtractURL", α, &γ); err != nil {
		return nil, err
	}
	return γ.ExtractURLResponse, nil
}
