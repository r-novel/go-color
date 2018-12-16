package color

type Attribute uint8

const (
	AttributeFormat Attribute = 0x00

	AttributeFGColor       = 0x1E
	AttributeBGColor       = 0x28
	AttributeFGBrightColor = 0x5A
	AttributeBGBrightColor = 0x64
)

const (
	AttributeFormatReset Attribute = AttributeFormat
	AttributeFormatBold  Attribute = AttributeFormat + iota
	AttributeFormatFaint
	AttributeFormatItalic
	AttributeFormatUnderline
	AttributeFormatBlinkSlow
	AttributeFormatBlinkRapid
	AttributeFormatReverseVideo
	AttributeFormatConcealed
	AttributeFormatCrossedOut
)

const (
	AttributeFGColorBlack Attribute = AttributeFGColor
	AttributeFGColorRed   Attribute = AttributeFGColor + iota
	AttributeFGColorGreen
	AttributeFGColorYellow
	AttributeFGColorBlue
	AttributeFGColorMagenta
	AttributeFGColorCyan
	AttributeFGColorWhite
)

const (
	AttributeBGColorBlack Attribute = AttributeBGColor
	AttributeBGColorRed   Attribute = AttributeBGColor + iota
	AttributeBGColorGreen
	AttributeBGColorYellow
	AttributeBGColorBlue
	AttributeBGColorMagenta
	AttributeBGColorCyan
	AttributeBGColorWhite
)

const (
	AttributeFGBrightColorBlack Attribute = AttributeFGBrightColor
	AttributeFGBrightColorRed   Attribute = AttributeFGBrightColor + iota
	AttributeFGBrightColorGreen
	AttributeFGBrightColorYellow
	AttributeFGBrightColorBlue
	AttributeFGBrightColorMagenta
	AttributeFGBrightColorCyan
	AttributeFGBrightColorWhite
)

const (
	AttributeBGBrightColorBlack Attribute = AttributeBGBrightColor
	AttributeBGBrightColorRed   Attribute = AttributeBGBrightColor + iota
	AttributeBGBrightColorGreen
	AttributeBGBrightColorYellow
	AttributeBGBrightColorBlue
	AttributeBGBrightColorMagenta
	AttributeBGBrightColorCyan
	AttributeBGBrightColorWhite
)
