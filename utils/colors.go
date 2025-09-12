package utils

// https://gist.github.com/jennyknuth/e2d9ee930303d5a5fe8862c6e31819c5
// https://htmlcolorcodes.com/color-names/
// https://www.w3.org/TR/css-color-3/
// https://www.w3.org/wiki/CSS/Properties/color/keywords#System_Colors
// https://www.canva.com/colors/color-wheel/

type colorDescriptor struct {
	Name string
	Hex  HexColor
	Rgb  RgbColor
}

var ColorNames = struct {
	Red    redColorNamesType
	Pink   pinkColorNameType
	Orange orangeColorType
	Yellow yellowColorType
	Purple purpleColorType
	Green  greenColorType
	Blue   blueColorNameType
	Brown  brownColorNameType
	White  whiteColorNameType
	Grey   greyColorNameType
}{
	Red:    redColorNames,
	Pink:   pinkColorNames,
	Orange: orangeColorNames,
	Yellow: yellowColorNames,
	Purple: purpleColorNames,
	Green:  greenColorNames,
	Blue:   blueColorNames,
	Brown:  brownColorNames,
	White:  whiteColorNames,
	Grey:   greyColorNames,
}

type redColorNamesType = struct {
	IndianRed   colorDescriptor
	LightCoral  colorDescriptor
	Salmon      colorDescriptor
	DarkSalmon  colorDescriptor
	LightSalmon colorDescriptor
	Crimson     colorDescriptor
	Red         colorDescriptor
	FireBrick   colorDescriptor
	DarkRed     colorDescriptor
}

var redColorNames = redColorNamesType{
	IndianRed:   colorDescriptor{Name: "IndianRed", Hex: MustParseHex("#CD5C5C"), Rgb: MustParseHex("#CD5C5C").Rgb()},
	LightCoral:  colorDescriptor{Name: "LightCoral", Hex: MustParseHex("#F08080"), Rgb: MustParseHex("#F08080").Rgb()},
	Salmon:      colorDescriptor{Name: "Salmon", Hex: MustParseHex("#FA8072"), Rgb: MustParseHex("#FA8072").Rgb()},
	DarkSalmon:  colorDescriptor{Name: "DarkSalmon", Hex: MustParseHex("#E9967A"), Rgb: MustParseHex("#E9967A").Rgb()},
	LightSalmon: colorDescriptor{Name: "LightSalmon", Hex: MustParseHex("#FFA07A"), Rgb: MustParseHex("#FFA07A").Rgb()},
	Crimson:     colorDescriptor{Name: "Crimson", Hex: MustParseHex("#DC143C"), Rgb: MustParseHex("#DC143C").Rgb()},
	Red:         colorDescriptor{Name: "Red", Hex: MustParseHex("#FF0000"), Rgb: MustParseHex("#FF0000").Rgb()},
	FireBrick:   colorDescriptor{Name: "FireBrick", Hex: MustParseHex("#B22222"), Rgb: MustParseHex("#B22222").Rgb()},
	DarkRed:     colorDescriptor{Name: "DarkRed", Hex: MustParseHex("#8B0000"), Rgb: MustParseHex("#8B0000").Rgb()},
}

type pinkColorNameType struct {
	Pink            colorDescriptor
	LightPink       colorDescriptor
	HotPink         colorDescriptor
	DeepPink        colorDescriptor
	MediumVioletRed colorDescriptor
	PaleVioletRed   colorDescriptor
}

var pinkColorNames = pinkColorNameType{
	Pink:            colorDescriptor{Name: "Pink", Hex: MustParseHex("#FFC0CB"), Rgb: MustParseHex("#FFC0CB").Rgb()},
	LightPink:       colorDescriptor{Name: "LightPink", Hex: MustParseHex("#FFB6C1"), Rgb: MustParseHex("#FFB6C1").Rgb()},
	HotPink:         colorDescriptor{Name: "HotPink", Hex: MustParseHex("#FF69B4"), Rgb: MustParseHex("#FF69B4").Rgb()},
	DeepPink:        colorDescriptor{Name: "DeepPink", Hex: MustParseHex("#FF1493"), Rgb: MustParseHex("#FF1493").Rgb()},
	MediumVioletRed: colorDescriptor{Name: "MediumVioletRed", Hex: MustParseHex("#C71585"), Rgb: MustParseHex("#C71585").Rgb()},
	PaleVioletRed:   colorDescriptor{Name: "PaleVioletRed", Hex: MustParseHex("#DB7093"), Rgb: MustParseHex("#DB7093").Rgb()},
}

type orangeColorType struct {
	LightSalmon colorDescriptor
	Coral       colorDescriptor
	Tomato      colorDescriptor
	OrangeRed   colorDescriptor
	DarkOrange  colorDescriptor
	Orange      colorDescriptor
}

var orangeColorNames = orangeColorType{
	LightSalmon: colorDescriptor{Name: "LightSalmon", Hex: MustParseHex("#FFA07A"), Rgb: MustParseHex("#FFA07A").Rgb()},
	Coral:       colorDescriptor{Name: "Coral", Hex: MustParseHex("#FF7F50"), Rgb: MustParseHex("#FF7F50").Rgb()},
	Tomato:      colorDescriptor{Name: "Tomato", Hex: MustParseHex("#FF6347"), Rgb: MustParseHex("#FF6347").Rgb()},
	OrangeRed:   colorDescriptor{Name: "OrangeRed", Hex: MustParseHex("#FF4500"), Rgb: MustParseHex("#FF4500").Rgb()},
	DarkOrange:  colorDescriptor{Name: "DarkOrange", Hex: MustParseHex("#FF8C00"), Rgb: MustParseHex("#FF8C00").Rgb()},
	Orange:      colorDescriptor{Name: "Orange", Hex: MustParseHex("#FFA500"), Rgb: MustParseHex("#FFA500").Rgb()},
}

type yellowColorType struct {
	Gold                 colorDescriptor
	Yellow               colorDescriptor
	LightYellow          colorDescriptor
	LemonChiffon         colorDescriptor
	LightGoldenrodYellow colorDescriptor
	Moccasin             colorDescriptor
	PeachPuff            colorDescriptor
	PaleGoldenrod        colorDescriptor
	Khaki                colorDescriptor
	DarkKhaki            colorDescriptor
	PapayaWhip           colorDescriptor
}

var yellowColorNames = yellowColorType{
	Gold:                 colorDescriptor{Name: "Gold", Hex: MustParseHex("#FFD700"), Rgb: MustParseHex("#FFD700").Rgb()},
	Yellow:               colorDescriptor{Name: "Yellow", Hex: MustParseHex("#FFFF00"), Rgb: MustParseHex("#FFFF00").Rgb()},
	LightYellow:          colorDescriptor{Name: "LightYellow", Hex: MustParseHex("#FFFFE0"), Rgb: MustParseHex("#FFFFE0").Rgb()},
	LemonChiffon:         colorDescriptor{Name: "LemonChiffon", Hex: MustParseHex("#FFFACD"), Rgb: MustParseHex("#FFFACD").Rgb()},
	LightGoldenrodYellow: colorDescriptor{Name: "LightGoldenrodYellow", Hex: MustParseHex("#FAFAD2"), Rgb: MustParseHex("#FAFAD2").Rgb()},
	PapayaWhip:           colorDescriptor{Name: "PapayaWhip", Hex: MustParseHex("#FFEFD5"), Rgb: MustParseHex("#FFEFD5").Rgb()},
	Moccasin:             colorDescriptor{Name: "Moccasin", Hex: MustParseHex("#FFE4B5"), Rgb: MustParseHex("#FFE4B5").Rgb()},
	PeachPuff:            colorDescriptor{Name: "PeachPuff", Hex: MustParseHex("#FFDAB9"), Rgb: MustParseHex("#FFDAB9").Rgb()},
	PaleGoldenrod:        colorDescriptor{Name: "PaleGoldenrod", Hex: MustParseHex("#EEE8AA"), Rgb: MustParseHex("#EEE8AA").Rgb()},
	Khaki:                colorDescriptor{Name: "Khaki", Hex: MustParseHex("#F0E68C"), Rgb: MustParseHex("#F0E68C").Rgb()},
	DarkKhaki:            colorDescriptor{Name: "DarkKhaki", Hex: MustParseHex("#BDB76B"), Rgb: MustParseHex("#BDB76B").Rgb()},
}

type purpleColorType struct {
	Lavender        colorDescriptor
	Thistle         colorDescriptor
	Plum            colorDescriptor
	Violet          colorDescriptor
	Orchid          colorDescriptor
	Fuchsia         colorDescriptor
	Magenta         colorDescriptor
	MediumOrchid    colorDescriptor
	MediumPurple    colorDescriptor
	RebeccaPurple   colorDescriptor
	BlueViolet      colorDescriptor
	DarkViolet      colorDescriptor
	DarkOrchid      colorDescriptor
	DarkMagenta     colorDescriptor
	Purple          colorDescriptor
	Indigo          colorDescriptor
	SlateBlue       colorDescriptor
	DarkSlateBlue   colorDescriptor
	MediumSlateBlue colorDescriptor
}

var purpleColorNames = purpleColorType{
	Lavender:        colorDescriptor{Name: "Lavender", Hex: MustParseHex("#E6E6FA"), Rgb: MustParseHex("#E6E6FA").Rgb()},
	Thistle:         colorDescriptor{Name: "Thistle", Hex: MustParseHex("#D8BFD8"), Rgb: MustParseHex("#D8BFD8").Rgb()},
	Plum:            colorDescriptor{Name: "Plum", Hex: MustParseHex("#DDA0DD"), Rgb: MustParseHex("#DDA0DD").Rgb()},
	Violet:          colorDescriptor{Name: "Violet", Hex: MustParseHex("#EE82EE"), Rgb: MustParseHex("#EE82EE").Rgb()},
	Orchid:          colorDescriptor{Name: "Orchid", Hex: MustParseHex("#DA70D6"), Rgb: MustParseHex("#DA70D6").Rgb()},
	Fuchsia:         colorDescriptor{Name: "Fuchsia", Hex: MustParseHex("#FF00FF"), Rgb: MustParseHex("#FF00FF").Rgb()},
	Magenta:         colorDescriptor{Name: "Magenta", Hex: MustParseHex("#FF00FF"), Rgb: MustParseHex("#FF00FF").Rgb()},
	MediumOrchid:    colorDescriptor{Name: "MediumOrchid", Hex: MustParseHex("#BA55D3"), Rgb: MustParseHex("#BA55D3").Rgb()},
	MediumPurple:    colorDescriptor{Name: "MediumPurple", Hex: MustParseHex("#9370DB"), Rgb: MustParseHex("#9370DB").Rgb()},
	RebeccaPurple:   colorDescriptor{Name: "RebeccaPurple", Hex: MustParseHex("#663399"), Rgb: MustParseHex("#663399").Rgb()},
	BlueViolet:      colorDescriptor{Name: "BlueViolet", Hex: MustParseHex("#8A2BE2"), Rgb: MustParseHex("#8A2BE2").Rgb()},
	DarkViolet:      colorDescriptor{Name: "DarkViolet", Hex: MustParseHex("#9400D3"), Rgb: MustParseHex("#9400D3").Rgb()},
	DarkOrchid:      colorDescriptor{Name: "DarkOrchid", Hex: MustParseHex("#9932CC"), Rgb: MustParseHex("#9932CC").Rgb()},
	DarkMagenta:     colorDescriptor{Name: "DarkMagenta", Hex: MustParseHex("#8B008B"), Rgb: MustParseHex("#8B008B").Rgb()},
	Purple:          colorDescriptor{Name: "Purple", Hex: MustParseHex("#800080"), Rgb: MustParseHex("#800080").Rgb()},
	Indigo:          colorDescriptor{Name: "Indigo", Hex: MustParseHex("#4B0082"), Rgb: MustParseHex("#4B0082").Rgb()},
	SlateBlue:       colorDescriptor{Name: "SlateBlue", Hex: MustParseHex("#6A5ACD"), Rgb: MustParseHex("#6A5ACD").Rgb()},
	DarkSlateBlue:   colorDescriptor{Name: "DarkSlateBlue", Hex: MustParseHex("#483D8B"), Rgb: MustParseHex("#483D8B").Rgb()},
	MediumSlateBlue: colorDescriptor{Name: "MediumSlateBlue", Hex: MustParseHex("#7B68EE"), Rgb: MustParseHex("#7B68EE").Rgb()},
}

type greenColorType struct {
	GreenYellow       colorDescriptor
	Chartreuse        colorDescriptor
	LawnGreen         colorDescriptor
	Lime              colorDescriptor
	LimeGreen         colorDescriptor
	PaleGreen         colorDescriptor
	LightGreen        colorDescriptor
	MediumSpringGreen colorDescriptor
	SpringGreen       colorDescriptor
	MediumSeaGreen    colorDescriptor
	SeaGreen          colorDescriptor
	ForestGreen       colorDescriptor
	Green             colorDescriptor
	DarkGreen         colorDescriptor
	YellowGreen       colorDescriptor
	OliveDrab         colorDescriptor
	Olive             colorDescriptor
	DarkOliveGreen    colorDescriptor
	MediumAquamarine  colorDescriptor
	DarkSeaGreen      colorDescriptor
	LightSeaGreen     colorDescriptor
	DarkCyan          colorDescriptor
	Teal              colorDescriptor
}

var greenColorNames = greenColorType{
	GreenYellow:       colorDescriptor{Name: "GreenYellow", Hex: MustParseHex("#ADFF2F"), Rgb: MustParseHex("#ADFF2F").Rgb()},
	Chartreuse:        colorDescriptor{Name: "Chartreuse", Hex: MustParseHex("#7FFF00"), Rgb: MustParseHex("#7FFF00").Rgb()},
	LawnGreen:         colorDescriptor{Name: "LawnGreen", Hex: MustParseHex("#7CFC00"), Rgb: MustParseHex("#7CFC00").Rgb()},
	Lime:              colorDescriptor{Name: "Lime", Hex: MustParseHex("#00FF00"), Rgb: MustParseHex("#00FF00").Rgb()},
	LimeGreen:         colorDescriptor{Name: "LimeGreen", Hex: MustParseHex("#32CD32"), Rgb: MustParseHex("#32CD32").Rgb()},
	PaleGreen:         colorDescriptor{Name: "PaleGreen", Hex: MustParseHex("#98FB98"), Rgb: MustParseHex("#98FB98").Rgb()},
	LightGreen:        colorDescriptor{Name: "LightGreen", Hex: MustParseHex("#90EE90"), Rgb: MustParseHex("#90EE90").Rgb()},
	MediumSpringGreen: colorDescriptor{Name: "MediumSpringGreen", Hex: MustParseHex("#00FA9A"), Rgb: MustParseHex("#00FA9A").Rgb()},
	SpringGreen:       colorDescriptor{Name: "SpringGreen", Hex: MustParseHex("#00FF7F"), Rgb: MustParseHex("#00FF7F").Rgb()},
	MediumSeaGreen:    colorDescriptor{Name: "MediumSeaGreen", Hex: MustParseHex("#3CB371"), Rgb: MustParseHex("#3CB371").Rgb()},
	SeaGreen:          colorDescriptor{Name: "SeaGreen", Hex: MustParseHex("#2E8B57"), Rgb: MustParseHex("#2E8B57").Rgb()},
	ForestGreen:       colorDescriptor{Name: "ForestGreen", Hex: MustParseHex("#228B22"), Rgb: MustParseHex("#228B22").Rgb()},
	Green:             colorDescriptor{Name: "Green", Hex: MustParseHex("#008000"), Rgb: MustParseHex("#008000").Rgb()},
	DarkGreen:         colorDescriptor{Name: "DarkGreen", Hex: MustParseHex("#006400"), Rgb: MustParseHex("#006400").Rgb()},
	YellowGreen:       colorDescriptor{Name: "YellowGreen", Hex: MustParseHex("#9ACD32"), Rgb: MustParseHex("#9ACD32").Rgb()},
	OliveDrab:         colorDescriptor{Name: "OliveDrab", Hex: MustParseHex("#6B8E23"), Rgb: MustParseHex("#6B8E23").Rgb()},
	Olive:             colorDescriptor{Name: "Olive", Hex: MustParseHex("#808000"), Rgb: MustParseHex("#808000").Rgb()},
	DarkOliveGreen:    colorDescriptor{Name: "DarkOliveGreen", Hex: MustParseHex("#556B2F"), Rgb: MustParseHex("#556B2F").Rgb()},
	MediumAquamarine:  colorDescriptor{Name: "MediumAquamarine", Hex: MustParseHex("#66CDAA"), Rgb: MustParseHex("#66CDAA").Rgb()},
	DarkSeaGreen:      colorDescriptor{Name: "DarkSeaGreen", Hex: MustParseHex("#8FBC8B"), Rgb: MustParseHex("#8FBC8B").Rgb()},
	LightSeaGreen:     colorDescriptor{Name: "LightSeaGreen", Hex: MustParseHex("#20B2AA"), Rgb: MustParseHex("#20B2AA").Rgb()},
	DarkCyan:          colorDescriptor{Name: "DarkCyan", Hex: MustParseHex("#008B8B"), Rgb: MustParseHex("#008B8B").Rgb()},
	Teal:              colorDescriptor{Name: "Teal", Hex: MustParseHex("#008080"), Rgb: MustParseHex("#008080").Rgb()},
}

type blueColorNameType struct {
	Aqua            colorDescriptor
	Cyan            colorDescriptor
	LightCyan       colorDescriptor
	PaleTurquoise   colorDescriptor
	Aquamarine      colorDescriptor
	Turquoise       colorDescriptor
	MediumTurquoise colorDescriptor
	DarkTurquoise   colorDescriptor
	CadetBlue       colorDescriptor
	SteelBlue       colorDescriptor
	LightSteelBlue  colorDescriptor
	PowderBlue      colorDescriptor
	LightBlue       colorDescriptor
	SkyBlue         colorDescriptor
	LightSkyBlue    colorDescriptor
	DeepSkyBlue     colorDescriptor
	DodgerBlue      colorDescriptor
	CornflowerBlue  colorDescriptor
	MediumSlateBlue colorDescriptor
	RoyalBlue       colorDescriptor
	Blue            colorDescriptor
	MediumBlue      colorDescriptor
	DarkBlue        colorDescriptor
	Navy            colorDescriptor
	MidnightBlue    colorDescriptor
}

var blueColorNames = blueColorNameType{
	Aqua:            colorDescriptor{Name: "Aqua", Hex: MustParseHex("#00FFFF"), Rgb: MustParseHex("#00FFFF").Rgb()},
	Cyan:            colorDescriptor{Name: "Cyan", Hex: MustParseHex("#00FFFF"), Rgb: MustParseHex("#00FFFF").Rgb()},
	LightCyan:       colorDescriptor{Name: "LightCyan", Hex: MustParseHex("#E0FFFF"), Rgb: MustParseHex("#E0FFFF").Rgb()},
	PaleTurquoise:   colorDescriptor{Name: "PaleTurquoise", Hex: MustParseHex("#AFEEEE"), Rgb: MustParseHex("#AFEEEE").Rgb()},
	Aquamarine:      colorDescriptor{Name: "Aquamarine", Hex: MustParseHex("#7FFFD4"), Rgb: MustParseHex("#7FFFD4").Rgb()},
	Turquoise:       colorDescriptor{Name: "Turquoise", Hex: MustParseHex("#40E0D0"), Rgb: MustParseHex("#40E0D0").Rgb()},
	MediumTurquoise: colorDescriptor{Name: "MediumTurquoise", Hex: MustParseHex("#48D1CC"), Rgb: MustParseHex("#48D1CC").Rgb()},
	DarkTurquoise:   colorDescriptor{Name: "DarkTurquoise", Hex: MustParseHex("#00CED1"), Rgb: MustParseHex("#00CED1").Rgb()},
	CadetBlue:       colorDescriptor{Name: "CadetBlue", Hex: MustParseHex("#5F9EA0"), Rgb: MustParseHex("#5F9EA0").Rgb()},
	SteelBlue:       colorDescriptor{Name: "SteelBlue", Hex: MustParseHex("#4682B4"), Rgb: MustParseHex("#4682B4").Rgb()},
	LightSteelBlue:  colorDescriptor{Name: "LightSteelBlue", Hex: MustParseHex("#B0C4DE"), Rgb: MustParseHex("#B0C4DE").Rgb()},
	PowderBlue:      colorDescriptor{Name: "PowderBlue", Hex: MustParseHex("#B0E0E6"), Rgb: MustParseHex("#B0E0E6").Rgb()},
	LightBlue:       colorDescriptor{Name: "LightBlue", Hex: MustParseHex("#ADD8E6"), Rgb: MustParseHex("#ADD8E6").Rgb()},
	SkyBlue:         colorDescriptor{Name: "SkyBlue", Hex: MustParseHex("#87CEEB"), Rgb: MustParseHex("#87CEEB").Rgb()},
	LightSkyBlue:    colorDescriptor{Name: "LightSkyBlue", Hex: MustParseHex("#87CEFA"), Rgb: MustParseHex("#87CEFA").Rgb()},
	DeepSkyBlue:     colorDescriptor{Name: "DeepSkyBlue", Hex: MustParseHex("#00BFFF"), Rgb: MustParseHex("#00BFFF").Rgb()},
	DodgerBlue:      colorDescriptor{Name: "DodgerBlue", Hex: MustParseHex("#1E90FF"), Rgb: MustParseHex("#1E90FF").Rgb()},
	CornflowerBlue:  colorDescriptor{Name: "CornflowerBlue", Hex: MustParseHex("#6495ED"), Rgb: MustParseHex("#6495ED").Rgb()},
	MediumSlateBlue: colorDescriptor{Name: "MediumSlateBlue", Hex: MustParseHex("#7B68EE"), Rgb: MustParseHex("#7B68EE").Rgb()},
	RoyalBlue:       colorDescriptor{Name: "RoyalBlue", Hex: MustParseHex("#4169E1"), Rgb: MustParseHex("#4169E1").Rgb()},
	Blue:            colorDescriptor{Name: "Blue", Hex: MustParseHex("#0000FF"), Rgb: MustParseHex("#0000FF").Rgb()},
	MediumBlue:      colorDescriptor{Name: "MediumBlue", Hex: MustParseHex("#0000CD"), Rgb: MustParseHex("#0000CD").Rgb()},
	DarkBlue:        colorDescriptor{Name: "DarkBlue", Hex: MustParseHex("#00008B"), Rgb: MustParseHex("#00008B").Rgb()},
	Navy:            colorDescriptor{Name: "Navy", Hex: MustParseHex("#000080"), Rgb: MustParseHex("#000080").Rgb()},
	MidnightBlue:    colorDescriptor{Name: "MidnightBlue", Hex: MustParseHex("#191970"), Rgb: MustParseHex("#191970").Rgb()},
}

type brownColorNameType struct {
	Cornsilk       colorDescriptor
	BlanchedAlmond colorDescriptor
	Bisque         colorDescriptor
	NavajoWhite    colorDescriptor
	Wheat          colorDescriptor
	BurlyWood      colorDescriptor
	Tan            colorDescriptor
	RosyBrown      colorDescriptor
	SandyBrown     colorDescriptor
	Goldenrod      colorDescriptor
	DarkGoldenrod  colorDescriptor
	Peru           colorDescriptor
	Chocolate      colorDescriptor
	SaddleBrown    colorDescriptor
	Sienna         colorDescriptor
	Brown          colorDescriptor
	Maroon         colorDescriptor
}

var brownColorNames = brownColorNameType{
	Cornsilk:       colorDescriptor{Name: "Cornsilk", Hex: MustParseHex("#FFF8DC"), Rgb: MustParseHex("#FFF8DC").Rgb()},
	BlanchedAlmond: colorDescriptor{Name: "BlanchedAlmond", Hex: MustParseHex("#FFEBCD"), Rgb: MustParseHex("#FFEBCD").Rgb()},
	Bisque:         colorDescriptor{Name: "Bisque", Hex: MustParseHex("#FFE4C4"), Rgb: MustParseHex("#FFE4C4").Rgb()},
	NavajoWhite:    colorDescriptor{Name: "NavajoWhite", Hex: MustParseHex("#FFDEAD"), Rgb: MustParseHex("#FFDEAD").Rgb()},
	Wheat:          colorDescriptor{Name: "Wheat", Hex: MustParseHex("#F5DEB3"), Rgb: MustParseHex("#F5DEB3").Rgb()},
	BurlyWood:      colorDescriptor{Name: "BurlyWood", Hex: MustParseHex("#DEB887"), Rgb: MustParseHex("#DEB887").Rgb()},
	Tan:            colorDescriptor{Name: "Tan", Hex: MustParseHex("#D2B48C"), Rgb: MustParseHex("#D2B48C").Rgb()},
	RosyBrown:      colorDescriptor{Name: "RosyBrown", Hex: MustParseHex("#BC8F8F"), Rgb: MustParseHex("#BC8F8F").Rgb()},
	SandyBrown:     colorDescriptor{Name: "SandyBrown", Hex: MustParseHex("#F4A460"), Rgb: MustParseHex("#F4A460").Rgb()},
	Goldenrod:      colorDescriptor{Name: "Goldenrod", Hex: MustParseHex("#DAA520"), Rgb: MustParseHex("#DAA520").Rgb()},
	DarkGoldenrod:  colorDescriptor{Name: "DarkGoldenrod", Hex: MustParseHex("#B8860B"), Rgb: MustParseHex("#B8860B").Rgb()},
	Peru:           colorDescriptor{Name: "Peru", Hex: MustParseHex("#CD853F"), Rgb: MustParseHex("#CD853F").Rgb()},
	Chocolate:      colorDescriptor{Name: "Chocolate", Hex: MustParseHex("#D2691E"), Rgb: MustParseHex("#D2691E").Rgb()},
	SaddleBrown:    colorDescriptor{Name: "SaddleBrown", Hex: MustParseHex("#8B4513"), Rgb: MustParseHex("#8B4513").Rgb()},
	Sienna:         colorDescriptor{Name: "Sienna", Hex: MustParseHex("#A0522D"), Rgb: MustParseHex("#A0522D").Rgb()},
	Brown:          colorDescriptor{Name: "Brown", Hex: MustParseHex("#A52A2A"), Rgb: MustParseHex("#A52A2A").Rgb()},
	Maroon:         colorDescriptor{Name: "Maroon", Hex: MustParseHex("#800000"), Rgb: MustParseHex("#800000").Rgb()},
}

type whiteColorNameType struct {
	White         colorDescriptor
	Snow          colorDescriptor
	HoneyDew      colorDescriptor
	MintCream     colorDescriptor
	Azure         colorDescriptor
	AliceBlue     colorDescriptor
	GhostWhite    colorDescriptor
	WhiteSmoke    colorDescriptor
	SeaShell      colorDescriptor
	Beige         colorDescriptor
	OldLace       colorDescriptor
	FloralWhite   colorDescriptor
	Ivory         colorDescriptor
	AntiqueWhite  colorDescriptor
	Linen         colorDescriptor
	LavenderBlush colorDescriptor
	MistyRose     colorDescriptor
}

var whiteColorNames = whiteColorNameType{

	White:         colorDescriptor{Name: "White", Hex: MustParseHex("#FFFFFF"), Rgb: MustParseHex("#FFFFFF").Rgb()},
	Snow:          colorDescriptor{Name: "Snow", Hex: MustParseHex("#FFFAFA"), Rgb: MustParseHex("#FFFAFA").Rgb()},
	HoneyDew:      colorDescriptor{Name: "HoneyDew", Hex: MustParseHex("#F0FFF0"), Rgb: MustParseHex("#F0FFF0").Rgb()},
	MintCream:     colorDescriptor{Name: "MintCream", Hex: MustParseHex("#F5FFFA"), Rgb: MustParseHex("#F5FFFA").Rgb()},
	Azure:         colorDescriptor{Name: "Azure", Hex: MustParseHex("#F0FFFF"), Rgb: MustParseHex("#F0FFFF").Rgb()},
	AliceBlue:     colorDescriptor{Name: "AliceBlue", Hex: MustParseHex("#F0F8FF"), Rgb: MustParseHex("#F0F8FF").Rgb()},
	GhostWhite:    colorDescriptor{Name: "GhostWhite", Hex: MustParseHex("#F8F8FF"), Rgb: MustParseHex("#F8F8FF").Rgb()},
	WhiteSmoke:    colorDescriptor{Name: "WhiteSmoke", Hex: MustParseHex("#F5F5F5"), Rgb: MustParseHex("#F5F5F5").Rgb()},
	SeaShell:      colorDescriptor{Name: "SeaShell", Hex: MustParseHex("#FFF5EE"), Rgb: MustParseHex("#FFF5EE").Rgb()},
	Beige:         colorDescriptor{Name: "Beige", Hex: MustParseHex("#F5F5DC"), Rgb: MustParseHex("#F5F5DC").Rgb()},
	OldLace:       colorDescriptor{Name: "OldLace", Hex: MustParseHex("#FDF5E6"), Rgb: MustParseHex("#FDF5E6").Rgb()},
	FloralWhite:   colorDescriptor{Name: "FloralWhite", Hex: MustParseHex("#FFFAF0"), Rgb: MustParseHex("#FFFAF0").Rgb()},
	Ivory:         colorDescriptor{Name: "Ivory", Hex: MustParseHex("#FFFFF0"), Rgb: MustParseHex("#FFFFF0").Rgb()},
	AntiqueWhite:  colorDescriptor{Name: "AntiqueWhite", Hex: MustParseHex("#FAEBD7"), Rgb: MustParseHex("#FAEBD7").Rgb()},
	Linen:         colorDescriptor{Name: "Linen", Hex: MustParseHex("#FAF0E6"), Rgb: MustParseHex("#FAF0E6").Rgb()},
	LavenderBlush: colorDescriptor{Name: "LavenderBlush", Hex: MustParseHex("#FFF0F5"), Rgb: MustParseHex("#FFF0F5").Rgb()},
	MistyRose:     colorDescriptor{Name: "MistyRose", Hex: MustParseHex("#FFE4E1"), Rgb: MustParseHex("#FFE4E1").Rgb()},
}

type greyColorNameType struct {
	Gainsboro      colorDescriptor
	LightGray      colorDescriptor
	Silver         colorDescriptor
	DarkGray       colorDescriptor
	Gray           colorDescriptor
	DimGray        colorDescriptor
	LightSlateGray colorDescriptor
	SlateGray      colorDescriptor
	DarkSlateGray  colorDescriptor
	Black          colorDescriptor
}

var greyColorNames = greyColorNameType{
	Gainsboro:      colorDescriptor{Name: "Gainsboro", Hex: MustParseHex("#DCDCDC"), Rgb: MustParseHex("#DCDCDC").Rgb()},
	LightGray:      colorDescriptor{Name: "LightGray", Hex: MustParseHex("#D3D3D3"), Rgb: MustParseHex("#D3D3D3").Rgb()},
	Silver:         colorDescriptor{Name: "Silver", Hex: MustParseHex("#C0C0C0"), Rgb: MustParseHex("#C0C0C0").Rgb()},
	DarkGray:       colorDescriptor{Name: "DarkGray", Hex: MustParseHex("#A9A9A9"), Rgb: MustParseHex("#A9A9A9").Rgb()},
	Gray:           colorDescriptor{Name: "Gray", Hex: MustParseHex("#808080"), Rgb: MustParseHex("#808080").Rgb()},
	DimGray:        colorDescriptor{Name: "DimGray", Hex: MustParseHex("#696969"), Rgb: MustParseHex("#696969").Rgb()},
	LightSlateGray: colorDescriptor{Name: "LightSlateGray", Hex: MustParseHex("#778899"), Rgb: MustParseHex("#778899").Rgb()},
	SlateGray:      colorDescriptor{Name: "SlateGray", Hex: MustParseHex("#708090"), Rgb: MustParseHex("#708090").Rgb()},
	DarkSlateGray:  colorDescriptor{Name: "DarkSlateGray", Hex: MustParseHex("#2F4F4F"), Rgb: MustParseHex("#2F4F4F").Rgb()},
	Black:          colorDescriptor{Name: "Black", Hex: MustParseHex("#000000"), Rgb: MustParseHex("#000000").Rgb()},
}
