//!+

// Xmlselect prints the text of selected elements of an XML document.
package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string // stack of element names
	var attrStack [][]xml.Attr
	var id, class string
	flag.StringVar(&id, "id", "", "restrict attribute id")
	flag.StringVar(&class, "class", "", "restrict attribute class")
	flag.Parse()
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) // push
			attrStack = append(attrStack, tok.Attr)
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
			attrStack = attrStack[:len(attrStack)-1]
		case xml.CharData:
			if containsAll(stack, flag.Args()) {
				idFlag, classFlag := false, false
				if id != "" || class != "" {
					for _, attr := range attrStack[len(attrStack)-1] {
						switch attr.Name.Local {
						case "id":
							if attr.Value == id {
								idFlag = true
							}
						case "class":
							if attr.Value == class {
								classFlag = true
							}
						}

					}
				}
				if (id == "" || idFlag) && (class == "" || classFlag) {
					fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
				}
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

//!-
