package utils

// https://developer.mozilla.org/en-US/docs/Web/CSS/color_value/rgb
// https://developer.mozilla.org/en-US/docs/Glossary/RGB

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	// https://regex101.com/
	re = regexp.MustCompile(`rgb\((\d{1,3}),(\d{1,3}),(\d{1,3})\)`)
)

func ParseRgb(s string) (RgbColor, error) {
	s = strings.ReplaceAll(s, " ", "")
	m := re.FindStringSubmatch(s)

	//TODO: Check for proper match

	var color RgbColor
	var err error

	if color.R, err = strconv.Atoi(m[1]); err != nil {
		return RgbColor{}, err
	}

	if color.B, err = strconv.Atoi(m[2]); err != nil {
		return RgbColor{}, err
	}

	if color.G, err = strconv.Atoi(m[3]); err != nil {
		return RgbColor{}, err
	}

	return color, nil
}

func MustParseRgb(s string) RgbColor {
	v, err := ParseRgb(s)
	if err != nil {
		panic(v)
	}
	return v
}

type RgbColor struct {
	R int
	B int
	G int
}

func (this RgbColor) String() string {
	return fmt.Sprintf("rgb(%d, %d, %d)", this.R, this.B, this.G)
}

func (this RgbColor) Hex() HexColor {
	return MustParseHex(fmt.Sprintf("#%X%X%X", this.R, this.G, this.B))
}
