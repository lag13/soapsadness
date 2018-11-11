# https://github.com/fiorix/wsdl2go

Version: 69b96af60ac9b29848069ac8e6eed1d655496847

## Problems:
- The generated code uses a lot of pointers which is annoying.
- The generated code does not give a namespace prefix to the "extract"
  node even though it looks like it is trying to. So we'd have to
  modify the code it looks like.

## So far
At this point I've reviewed:
- https://github.com/hooklift/gowsdl
- https://github.com/droyo/go-xml

and this one seems to be the best in the sense that I'll just have to
make (what hopefully appears to be) minimal changes in the generated
code where hooklift I would need to make changes to the library itself
and droyo didn't seem to add namespace prefixes anywhere.
