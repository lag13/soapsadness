## https://github.com/droyo/go-xml

Version: 841f47b2a098fd8c43afc469a454e663b29b6fc4

I originally found this library through a reddit thread:
https://www.reddit.com/r/golang/comments/7876zr/finally_a_decent_code_generation_from_xsd_and/

Interestingly the installation instructions are go get
`aqwari.net/xml/...` (I would have expected the github url or
something)

It says that the command line tool `wsdlgen` can be used to generate
code but unfortunately you need to manually download the wsdl specs
yourself: https://github.com/droyo/go-xml/issues/35. So I did that
(looks like you just need to follow all the "import" links) and then
called the wsdlgen utility while specifying all the files you have
downloaded like: `wsdl file1 importedfromfile1a importedfromfile1b...`

## Problems
- The Envelope XML nodes and such also do not get generated with a
  prefix just like hooklift
- Inside of the body there should be an XML node called "extract" but
  this library generates code such that the first node within the XML
  body is called "Message". Sending that body results in the following
  error message: `panic: S:Client: Cannot find dispatch method for
  {http://schemas.xmlsoap.org/soap/envelope/}Message`. Changing the
  code will take a bit of work I think and I'm not entirely sure how
  to do that although I imagine I'll try to copy what hooklift has
  since there's works in this respect. Side note: I notice that when I
  include an attribute
  `xmlns="http://home.textkernel.nl/sourcebox/soap/extract"` directly
  on the extract node, then the error message changes to: "The
  credentials given at login are not correct. If the problem occurs
  repeatedly please contact your service vendor. For faster service
  please supply this output XML and the document which caused the
  error". That is quite interesting and quite annoying and quite
  unhelpful because it does not seem to help you fix the actual issue
  which is that the extract node does not have a namespace prefix.
