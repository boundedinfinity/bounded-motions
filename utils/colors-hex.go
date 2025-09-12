package utils

// https://developer.mozilla.org/en-US/docs/Web/CSS/hex-color
// https://developer.mozilla.org/en-US/docs/Glossary/RGB

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseHex(s string) (HexColor, error) {
	s = strings.ReplaceAll(s, "#", "")
	s = strings.ToUpper(s)
	return HexColor{Hex: s}, nil
}

func MustParseHex(s string) HexColor {
	color, err := ParseHex(s)
	if err != nil {
		panic(err)
	}
	return color
}

type HexColor struct {
	Hex string
}

func (this HexColor) String() string {
	return "#" + this.Hex
}

func (this HexColor) Rgb() RgbColor {
	r, _ := strconv.Atoi(this.Hex[0:1])
	g, _ := strconv.Atoi(this.Hex[2:3])
	b, _ := strconv.Atoi(this.Hex[4:5])
	return MustParseRgb(fmt.Sprintf("rgb(%d,%d,%d)", r, b, g))
}
