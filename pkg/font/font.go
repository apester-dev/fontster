package font

import (
	"fmt"
	"strings"
)

// Family is a single font family and weight.
type Family struct {
	Name    string
	Style   string
	Weight  string
	BaseURL string
}

// Parse a string and return a slice of Family
func Parse(list string, BaseURL string) []Family {
	fonts := strings.Split(list, "|")
	families := make([]Family, 0, len(fonts))
	for _, font := range fonts {
		family, weights := familyAndWeights(font)
		for _, weight := range weights {
			families = append(families, Family{
				Name:    family,
				Style:   normalOrItalic(weight),
				Weight:  weight,
				BaseURL: BaseURL,
			})
		}
	}
	return families
}

// Source returns the full path to the font.
func (f Family) Source() string {
	return fmt.Sprintf("%s/%s/%s", f.BaseURL, removeSpaces(f.Name), f.FileName())
}

// FileName returns the basename of the font.
func (f Family) FileName() string {
	return fmt.Sprintf("%s-%s", removeSpaces(f.Name), WeightName(f.Weight))
}

func removeSpaces(family string) string {
	return strings.Join(strings.Fields(family), "")
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

func normalOrItalic(weight string) string {
	if strings.Contains(weight, "i") {
		return "italic"
	}
	return "normal"
}

var weights = map[string]string{
	"100":  "Thin",
	"100i": "ThinItalic",
	"200":  "ExtraLight",
	"200i": "ExtraLightItalic",
	"300":  "Light",
	"300i": "LightItalic",
	"400":  "Regular",
	"400i": "RegularItalic",
	"500":  "Medium",
	"500i": "MediumItalic",
	"700":  "Bold",
	"700i": "BoldItalic",
}

func WeightName(weight string) string {
	if w, ok := weights[weight]; ok {
		return w
	}
	return "Regular"
}
