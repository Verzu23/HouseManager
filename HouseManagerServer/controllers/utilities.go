package controllers

import (
	"strings"
)

func getArrayFromString(output string) []string {

	outList := strings.Split(output, "\n")

	var filteredOutList []string
	for _, s := range outList {

		if s != "." && s != "" {
			filteredOutList = append(filteredOutList, s)
		}
	}

	return filteredOutList
}
