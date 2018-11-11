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

## The best solution I have so far
At the time of this writing the client that has got me the closest to
marshal things the way TK wants is: https://github.com/fiorix/wsdl2go
Run `diff solutionsofar/extract/extract.go
fiorixwsdl2go/extract/extract.go` to see the changes I implemented so
that the generated code from fiorix worked for TK's "extract"
operation.

One code change was "necessary" in the sense that it got successfully
added the XML prefix but the other minor changes were mostly just
cleaning up code that is not needed (for example adding a XML tag on
an embedded struct has not effect so I removed the code where they did
that).

## Dealing with errors
In case of an error SOAP returns a Fault node with some structured
information about what went wrong. Some of the libraries I encountered
did not unmarshal all of that information which lead to some debugging
trouble so just be wary of that.

fiorix does not even try to unmarshal that
`<Envelope><Body><Fault>...</Fault></Body></Envelope>` structure
(https://github.com/fiorix/wsdl2go/blob/69b96af60ac9b29848069ac8e6eed1d655496847/soap/client.go#L143)
which I am perfectly happy with because that way we definitely capture
all information which could be useful when debugging that error. If at
some point we need to take programatic action on certain errors then
we could start parsing it but no need to worry about that until
necessary.
