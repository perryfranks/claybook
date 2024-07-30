package main

import (
	"fmt"
	"html/template"
	"strings"
)

type DisplayVars struct {
	NavBarColors map[string]string
	CurrentPage  string
}

func (d *DisplayVars) NewDisplayVars() {

	nav := map[string]string{}

	d.NavBarColors = nav
}

// If the value is set we change the color other the source will take it's color from the parent element
func getNavBarColor(heading string, navColors map[string]string) template.HTMLAttr {

	heading = strings.TrimSpace(heading)
	fmt.Println("get color called heading: ", heading, "navColors: ", navColors)
	val, ok := navColors[heading]
	if !ok {
		return "text-tSage"
	}

	fmt.Println("val: ", val)
	return template.HTMLAttr(val)

}

func isCurrentPageColor(queryPage string, d DisplayVars) template.HTMLAttr {
	queryPage = strings.TrimSpace(queryPage)
	fmt.Println("isCurrent. query=", queryPage, d.CurrentPage)
	if queryPage == d.CurrentPage {
		return "text-tOrange underline"
	}

	return ""
}
