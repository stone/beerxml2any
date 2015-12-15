[![Build Status](https://drone.io/github.com/stone/beerxml2any/status.png)](https://drone.io/github.com/stone/beerxml2any/latest)
 
## Usage:
 
 $ beerxml2any -help
 
     Usage of ./beerxml2any:
       -bxml string
             BeerXML input file
       -tmpl string
             Template file
 
     $ beerxml2any -bxml recipe.xml -tmpl tmpl/beerxml.txt.tmpl
 
## Build/Get
 
     $ go get github.com/stone/beerxml2any
 
You need the [go][] runtime, <http://golang.org/>
 
## Templates
 
Using go text/template, see examples in tmpl/ directory.

 - https://golang.org/pkg/text/template/#pkg-overview


[go]:http://golang.org/  "The Go Programming language"
