package font

import (
	"fmt"
	"strings"
)

type FontFamily struct {
	Family   string
	Style    string
	Weight   string
	Filename string
	BaseURL  string
}

func Parse(list string, BaseURL string) []FontFamily {
	var families []FontFamily
	fonts := strings.Split(list, "|")
	for _, font := range fonts {
		f, weights := familyAndWeights(font)
		for _, weight := range weights {
			style, name := fontGroups(weight)
			family := FontFamily{
				Family:   f,
				Style:    style,
				Weight:   strings.Trim(weight, "i"),
				Filename: fmt.Sprintf("%s-%s.woff2", f, name),
				BaseURL:  BaseURL,
			}
			families = append(families, family)
		}
	}
	return families
}

func familyAndWeights(list string) (string, []string) {
	weights := strings.Split(list, ":")
	if len(weights) == 1 {
		weights = append(weights, "400")
	}
	return weights[0], weights[1:]

}

func fontGroups(weight string) (string, string) {
	switch weight {
	case "100":
		return "normal", "Thin"
	case "100i":
		return "italic", "Thin"
	case "200":
		return "normal", "ExtraLight"
	case "200i":
		return "italic", "ExtraLight"
	case "300":
		return "normal", "Light"
	case "300i":
		return "italic", "LightItalic"
	case "400":
		return "normal", "Regular"
	case "400i":
		return "italic", "Italic"
	case "700":
		return "normal", "Bold"
	case "700i":
		return "italic", "BoldItalic"
	}
	return "normal", "Regular"
}
