package font

import (
	"fmt"
	"strings"
)

// Family is a single font family & weight.
type Family struct {
	Name     string
	Style    string
	Weight   string
	Filename string
	BaseURL  string
}

// Parse a string and return a slice of Family
func Parse(list string, BaseURL string) []Family {
	var families []Family
	fonts := strings.Split(list, "|")
	for _, font := range fonts {
		f, weights := familyAndWeights(font)
		for _, weight := range weights {
			style, name := fontGroups(weight)
			family := Family{
				Name:     f,
				Style:    style,
				Weight:   strings.Trim(weight, "i"),
				Filename: fmt.Sprintf("%s-%s", f, name),
				BaseURL:  BaseURL,
			}
			families = append(families, family)
		}
	}
	return families
}

func familyAndWeights(list string) (string, []string) {
	weights := strings.Split(list, ":")
	family := weights[0]

	// When family has no weight,
	// set regular as default and return.
	if len(weights) == 1 || weights[1] == "" {
		return family, []string{"400"}
	}

	// Allow multiple weights.
	weights = strings.Split(weights[1], ",")
	return family, weights

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
	default:
		return "normal", "Regular"
	}
}
