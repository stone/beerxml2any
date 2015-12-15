package main

import (
	"bytes"
	"flag"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"strings"
	"text/template"

	"github.com/stone/beerxml"
)

var (
	flagBxml = flag.String("bxml", "", "BeerXML input file")
	flagTmpl = flag.String("tmpl", "", "Template file")
)

func init() {
	flag.Parse()
}

// BeersmithXMLCleanup removes stupid XHTML escapings
// and \u0001 runes, making it a valid XML
func BeersmithXMLCleanup(bxml []byte) []byte {
	// Get rid of XHTML escapings.
	unescaped := html.UnescapeString(string(bxml))
	// Get rid of \u0001 runes
	str1 := strings.Replace(unescaped, "\u0001", "", -1)
	// The ampersand character (&) and the left angle bracket (<) MUST NOT
	// appear in their literal form, except when used as markup delimiters,
	// or within a comment, a processing instruction, or a CDATA section.
	// We only care for & for now.
	str1 = strings.Replace(str1, "&", "&amp;", -1)
	// Output unescaped and Invalid rune document
	return []byte(str1)
}

func loadBeerXML(filename string) (bx *beerxml.BeerXml, err error) {
	fd, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	beerxmlclean := BeersmithXMLCleanup(fd)
	beerReader := bytes.NewReader(beerxmlclean)
	bx, err = beerxml.NewBeerXml(beerReader)
	if err != nil {
		log.Fatal("Error parsing BeerXML:", err)
		return nil, err
	}

	return bx, nil
}

func loadTemplate(templatefile *string, bx *beerxml.BeerXml) (parsedTemplate *bytes.Buffer, err error) {
	parsedTemplate = new(bytes.Buffer)

	tmpl, err := template.ParseFiles(*templatefile)
	if err != nil {
		return nil, err
	}

	for r := range bx.Recipes {
		if err = tmpl.Execute(parsedTemplate, bx.Recipes[r]); err != nil {
			// We bug out if we fail, even if there is more recipes that may parse
			return nil, err
		}
	}

	return parsedTemplate, nil
}

func main() {

	if *flagBxml == "" || *flagTmpl == "" {
		flag.Usage()
		return
	}

	bxml, err := loadBeerXML(*flagBxml)
	if err != nil {
		log.Fatal(err)
	}

	out, err := loadTemplate(flagTmpl, bxml)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(out)

}
