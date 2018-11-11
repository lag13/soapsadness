## https://github.com/hooklift/gowsdl

Version: 193c95ad2f469c5f82a8ad8fca580a159dc85a60

After downloading the code and installing a binary with a:
`go get github.com/hooklift/gowsdl/...` 

Running `gowsdl wsdlurl` is all you need to generate the code and then
you use some code from the gowsdl library along with the generated
code to make your requests.

## Problems
- I'm not sure if it is a SOAP requirement or just a TK requirement
  but TK requires that the Envelope, Header, Body, and extract tags
  all have some sort of namespace prefix but gowsdl does not add those
  prefixes. Adding a prefix to "extract" would not be too terrible
  because it would involve modifying the generated code but the XML
  generation of the OTHER nodes is part of the library and so we would
  have to make a fork or something to get it to work:
  https://github.com/hooklift/gowsdl/issues/54,
  https://github.com/hooklift/gowsdl/blob/193c95ad2f469c5f82a8ad8fca580a159dc85a60/soap/soap.go#L13-L30
- I'm not sure what is wrong but the "Fault.detail" node which has
  more detailed XML information about what went wrong does not seem to
  be getting populated:
  https://github.com/hooklift/gowsdl/blob/193c95ad2f469c5f82a8ad8fca580a159dc85a60/soap/soap.go#L89
  I believe this might be a case of buggy code in the gowsdl library
  and they should adjust things so that detail node DOES get returned.
  Because in playing around with Go's XML parser I notice that if you
  declare a node as containing string data but that node actually has
  XML data then that XML data does NOT get parsed as a string, nothing
  gets parsed: https://play.golang.org/p/PPvHCgY2h_H So again, changes
  would have to be made to this library for things to work.
  
