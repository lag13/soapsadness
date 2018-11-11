# SOAP Sadness
At the time of this writing I was using Go 1.11 and I've included
versions of the libraries within the subdirectories. Perhaps you, oh
mysterious future reader, no longer have these issues with later
versions of Go or the libraries.

## I got 99 problems and they're all SOAP
For work I need to get one of our Go services talking to Textkernel's
(https://www.textkernel.com/) SOAP API and I wanted to document my
struggles in a public location so I can refer back to this if I need
to ever work with SOAP again.

## The Crux of the problem
This request (which was generated using the
"https://github.com/hooklift/gowsdl" library) to TK fails (fileContent
shortened for readability and credentials modified):

```xml
<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/">
  <Body xmlns="http://schemas.xmlsoap.org/soap/envelope/">
    <extract xmlns="http://home.textkernel.nl/sourcebox/soap/extract">
      <account>account</account>
      <username>username</username>
      <password>password</password>
      <fileName>Sample_IT_Resume_MartinL.pdf</fileName>
      <fileContent>JVBERi0xLjU...</fileContent>
  </extract>
</Body>
</Envelope>
```

Where this request succeeds:

```xml
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ext="http://home.textkernel.nl/sourcebox/soap/extract">
   <soapenv:Header/>
   <soapenv:Body>
      <ext:extract>
        <account>account</account>
        <username>username</username>
        <password>password</password>
         <fileName>Sample_IT_Resume_MartinL.pdf</fileName>
         <fileContent>JVBERi0xLjU...</fileContent>
      </ext:extract>
   </soapenv:Body>
</soapenv:Envelope>
```

The issue is that the XML namespace prefixes
(https://www.w3schools.com/xml/xml_namespaces.asp) are not added on
the nodes: Envelope, Header, Body, and extract. If you google "golang
xml namespace prefix" you will see that Go does not have great support
for marshalling XML to have namespace prefixes. This was a concise
link I found describing a workaround:
https://forum.golangbridge.org/t/xml-namespace-prefix-issue-at-go/8055/2
In short, you need to have one struct for marshalling and one struct
for unmarshalling https://play.golang.org/p/qCbsCoQ0cHc That link also
linked to the namespace prefix issue on golang's github:
https://github.com/golang/go/issues/13400

So it's not even a SOAP specific problem, its more of a XML problem
and none of the libraries I've tried thus far namespace prefix ALL of
the nodes Envelope, Header, Body, extract. SIDE NOTE: According to
this SOAP information (https://www.w3schools.com/xml/xml_soap.asp) it
sounds like the "extract" node need not have a prefix to still satisfy
the SOAP protocol: "Immediate child elements of the SOAP Body element
may be namespace-qualified.". So I wonder why TK enforces "extract" to
have a namespace prefix.

## The closest I've gotten to a solution
At the time of this writing the client that has got me the closest to
marshal things the way TK wants is: https://github.com/fiorix/wsdl2go
the only shortcoming seems to be that it did not add a namespace
prefix for the "extract" node. It also used a lot of pointers which I
was not a fan of and it makes me wonder if they are conforming to the
wsdl spec better? Or if they made a mistake and do not need those
pointers.

## One More Problem I want to make sure we get right
If SOAP encounters some sort of error they'll return a "Fault" node
and some of the libraries (maybe all, I can't remember right now) did
not unmarshal all available data in that Fault node which means we
could be missing useful debugging information if something goes wrong.

Honestly I'd prefer if we got the raw unmarshalled data for starters
so we definitely capture all the data and if at some point we need to
take programatic action on certain errors then we could start parsing
it.
