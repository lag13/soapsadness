// Package extract
//
package extract

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type Entry struct {
	Key   string `xml:"http://home.textkernel.nl/sourcebox/soap/extract key,omitempty"`
	Value string `xml:"http://home.textkernel.nl/sourcebox/soap/extract value,omitempty"`
}

type Extract struct {
	Account        string  `xml:"http://home.textkernel.nl/sourcebox/soap/extract account,omitempty"`
	Username       string  `xml:"http://home.textkernel.nl/sourcebox/soap/extract username,omitempty"`
	Password       string  `xml:"http://home.textkernel.nl/sourcebox/soap/extract password,omitempty"`
	FileName       string  `xml:"http://home.textkernel.nl/sourcebox/soap/extract fileName,omitempty"`
	FileContent    []byte  `xml:"http://home.textkernel.nl/sourcebox/soap/extract fileContent,omitempty"`
	TmfFileContent []byte  `xml:"http://home.textkernel.nl/sourcebox/soap/extract tmfFileContent,omitempty"`
	Apimap         []byte  `xml:"http://home.textkernel.nl/sourcebox/soap/extract apimap,omitempty"`
	Options        []Entry `xml:"http://home.textkernel.nl/sourcebox/soap/extract options,omitempty"`
}

func (t *Extract) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T Extract
	var layout struct {
		*T
		FileContent    *xsdBase64Binary `xml:"http://home.textkernel.nl/sourcebox/soap/extract fileContent,omitempty"`
		TmfFileContent *xsdBase64Binary `xml:"http://home.textkernel.nl/sourcebox/soap/extract tmfFileContent,omitempty"`
		Apimap         *xsdBase64Binary `xml:"http://home.textkernel.nl/sourcebox/soap/extract apimap,omitempty"`
	}
	layout.T = (*T)(t)
	layout.FileContent = (*xsdBase64Binary)(&layout.T.FileContent)
	layout.TmfFileContent = (*xsdBase64Binary)(&layout.T.TmfFileContent)
	layout.Apimap = (*xsdBase64Binary)(&layout.T.Apimap)
	return e.EncodeElement(layout, start)
}
func (t *Extract) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Extract
	var overlay struct {
		*T
		FileContent    *xsdBase64Binary `xml:"http://home.textkernel.nl/sourcebox/soap/extract fileContent,omitempty"`
		TmfFileContent *xsdBase64Binary `xml:"http://home.textkernel.nl/sourcebox/soap/extract tmfFileContent,omitempty"`
		Apimap         *xsdBase64Binary `xml:"http://home.textkernel.nl/sourcebox/soap/extract apimap,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.FileContent = (*xsdBase64Binary)(&overlay.T.FileContent)
	overlay.TmfFileContent = (*xsdBase64Binary)(&overlay.T.TmfFileContent)
	overlay.Apimap = (*xsdBase64Binary)(&overlay.T.Apimap)
	return d.DecodeElement(&overlay, &start)
}

type ExtractAdvanced struct {
	Account                 string  `xml:"http://home.textkernel.nl/sourcebox/soap/extract account,omitempty"`
	Username                string  `xml:"http://home.textkernel.nl/sourcebox/soap/extract username,omitempty"`
	Password                string  `xml:"http://home.textkernel.nl/sourcebox/soap/extract password,omitempty"`
	ClientSpecificArguments []Entry `xml:"http://home.textkernel.nl/sourcebox/soap/extract clientSpecificArguments,omitempty"`
	FileName                string  `xml:"http://home.textkernel.nl/sourcebox/soap/extract fileName,omitempty"`
	FileContent             []byte  `xml:"http://home.textkernel.nl/sourcebox/soap/extract fileContent,omitempty"`
	TmfFileContent          []byte  `xml:"http://home.textkernel.nl/sourcebox/soap/extract tmfFileContent,omitempty"`
	Apimap                  []byte  `xml:"http://home.textkernel.nl/sourcebox/soap/extract apimap,omitempty"`
	Options                 []Entry `xml:"http://home.textkernel.nl/sourcebox/soap/extract options,omitempty"`
}

func (t *ExtractAdvanced) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ExtractAdvanced
	var layout struct {
		*T
		FileContent    *xsdBase64Binary `xml:"http://home.textkernel.nl/sourcebox/soap/extract fileContent,omitempty"`
		TmfFileContent *xsdBase64Binary `xml:"http://home.textkernel.nl/sourcebox/soap/extract tmfFileContent,omitempty"`
		Apimap         *xsdBase64Binary `xml:"http://home.textkernel.nl/sourcebox/soap/extract apimap,omitempty"`
	}
	layout.T = (*T)(t)
	layout.FileContent = (*xsdBase64Binary)(&layout.T.FileContent)
	layout.TmfFileContent = (*xsdBase64Binary)(&layout.T.TmfFileContent)
	layout.Apimap = (*xsdBase64Binary)(&layout.T.Apimap)
	return e.EncodeElement(layout, start)
}
func (t *ExtractAdvanced) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ExtractAdvanced
	var overlay struct {
		*T
		FileContent    *xsdBase64Binary `xml:"http://home.textkernel.nl/sourcebox/soap/extract fileContent,omitempty"`
		TmfFileContent *xsdBase64Binary `xml:"http://home.textkernel.nl/sourcebox/soap/extract tmfFileContent,omitempty"`
		Apimap         *xsdBase64Binary `xml:"http://home.textkernel.nl/sourcebox/soap/extract apimap,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.FileContent = (*xsdBase64Binary)(&overlay.T.FileContent)
	overlay.TmfFileContent = (*xsdBase64Binary)(&overlay.T.TmfFileContent)
	overlay.Apimap = (*xsdBase64Binary)(&overlay.T.Apimap)
	return d.DecodeElement(&overlay, &start)
}

type ExtractAdvancedResponse struct {
	Return string `xml:"http://home.textkernel.nl/sourcebox/soap/extract return,omitempty"`
}

type ExtractException struct {
	Description      string `xml:"http://home.textkernel.nl/sourcebox/soap/extract description,omitempty"`
	Id               int    `xml:"http://home.textkernel.nl/sourcebox/soap/extract id,omitempty"`
	OutputStatusName string `xml:"http://home.textkernel.nl/sourcebox/soap/extract outputStatusName,omitempty"`
	Severity         string `xml:"http://home.textkernel.nl/sourcebox/soap/extract severity,omitempty"`
	Solution         string `xml:"http://home.textkernel.nl/sourcebox/soap/extract solution,omitempty"`
	TkURL            string `xml:"http://home.textkernel.nl/sourcebox/soap/extract tkURL,omitempty"`
	Message          string `xml:"http://home.textkernel.nl/sourcebox/soap/extract message,omitempty"`
}

type ExtractResponse struct {
	Return string `xml:"http://home.textkernel.nl/sourcebox/soap/extract return,omitempty"`
}

type ExtractURL struct {
	Account           string  `xml:"http://home.textkernel.nl/sourcebox/soap/extract account,omitempty"`
	Username          string  `xml:"http://home.textkernel.nl/sourcebox/soap/extract username,omitempty"`
	Password          string  `xml:"http://home.textkernel.nl/sourcebox/soap/extract password,omitempty"`
	PublicURL         string  `xml:"http://home.textkernel.nl/sourcebox/soap/extract publicURL,omitempty"`
	OutputType        string  `xml:"http://home.textkernel.nl/sourcebox/soap/extract outputType,omitempty"`
	CustomQueryString string  `xml:"http://home.textkernel.nl/sourcebox/soap/extract customQueryString,omitempty"`
	Options           []Entry `xml:"http://home.textkernel.nl/sourcebox/soap/extract options,omitempty"`
}

type ExtractURLResponse struct {
	Return []byte `xml:"http://home.textkernel.nl/sourcebox/soap/extract return,omitempty"`
}

func (t *ExtractURLResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ExtractURLResponse
	var layout struct {
		*T
		Return *xsdBase64Binary `xml:"http://home.textkernel.nl/sourcebox/soap/extract return,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Return = (*xsdBase64Binary)(&layout.T.Return)
	return e.EncodeElement(layout, start)
}
func (t *ExtractURLResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ExtractURLResponse
	var overlay struct {
		*T
		Return *xsdBase64Binary `xml:"http://home.textkernel.nl/sourcebox/soap/extract return,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Return = (*xsdBase64Binary)(&overlay.T.Return)
	return d.DecodeElement(&overlay, &start)
}

type xsdBase64Binary []byte

func (b *xsdBase64Binary) UnmarshalText(text []byte) (err error) {
	*b, err = base64.StdEncoding.DecodeString(string(text))
	return
}
func (b xsdBase64Binary) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	enc := base64.NewEncoder(base64.StdEncoding, &buf)
	enc.Write([]byte(b))
	enc.Close()
	return buf.Bytes(), nil
}

type Client struct {
	HTTPClient   http.Client
	ResponseHook func(*http.Response)
	RequestHook  func(*http.Request)
}
type soapEnvelope struct {
	XMLName struct{} `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  []byte   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
	Body    struct {
		Message interface{}
		Fault   *struct {
			String string `xml:"faultstring,omitempty"`
			Code   string `xml:"faultcode,omitempty"`
			Detail string `xml:"detail,omitempty"`
		} `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
	} `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
}

func (c *Client) Do(method, uri, action string, in, out interface{}) error {
	var body io.Reader
	var envelope soapEnvelope
	if method == "POST" || method == "PUT" {
		var buf bytes.Buffer
		envelope.Body.Message = in
		enc := xml.NewEncoder(&buf)
		if err := enc.Encode(envelope); err != nil {
			return err
		}
		if err := enc.Flush(); err != nil {
			return err
		}
		body = &buf
	}
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return err
	}
	req.Header.Set("SOAPAction", action)
	if c.RequestHook != nil {
		c.RequestHook(req)
	}
	rsp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()
	if c.ResponseHook != nil {
		c.ResponseHook(rsp)
	}
	dec := xml.NewDecoder(rsp.Body)
	envelope.Body.Message = out
	if err := dec.Decode(&envelope); err != nil {
		return err
	}
	if envelope.Body.Fault != nil {
		return fmt.Errorf("%s: %s", envelope.Body.Fault.Code, envelope.Body.Fault.String)
	}
	return nil
}
