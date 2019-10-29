package colour

const (
	normalText       = "\033[0m" // Turn off all attributes
	redText          = "\033[31m"
	greenText        = "\033[32m"
	yellowText       = "\033[33m"
	blueText         = "\033[34m"
	magentaText      = "\033[35m"
	cyanText         = "\033[36m"
	whiteText        = "\033[37m"
	boldText         = "\033[1m"
	boldRedText      = "\033[1;31m"
	boldGreenText    = "\033[1;32m"
	boldYellowText   = "\033[1;33m"
	boldBlueText     = "\033[1;34m"
	boldMagentaText  = "\033[1;35m"
	boldCyanText     = "\033[1;36m"
	faintText        = "\033[2m"
	faintRedText     = "\033[2;31m"
	faintGreenText   = "\033[2;32m"
	faintYellowText  = "\033[2;33m"
	faintBlueText    = "\033[2;34m"
	faintMagentaText = "\033[2;35m"
	faintCyanText    = "\033[2;36m"
	faintWhiteText   = "\033[2;37m"
)

// Red taks a string and returns that string escaped with the colour code
// for red text
func Red(str string) string {
	return string(redText + str + normalText)
}

// Green taks a string and returns that string escaped with the colour code
// for green text
func Green(str string) string {
	return string(greenText + str + normalText)
}

// Yellow taks a string and returns that string escaped with the colour code
// for yellow text
func Yellow(str string) string {
	return string(yellowText + str + normalText)
}

// Blue taks a string and returns that string escaped with the colour code
// for blue text
func Blue(str string) string {
	return string(blueText + str + normalText)
}

// Magenta taks a string and returns that string escaped with the colour code
// for __ text
func Magenta(str string) string {
	return string(magentaText + str + normalText)
}

// Cyan taks a string and returns that string escaped with the colour code
// for __ text
func Cyan(str string) string {
	return string(cyanText + str + normalText)
}

// White taks a string and returns that string escaped with the colour code
// for __ text
func White(str string) string {
	return string(whiteText + str + normalText)
}

// Bold takes a string and returns that string wrapped with the escape sequence
// to set it bold
func Bold(str string) string {
	return string(boldText + str + normalText)
}

// BoldRed takes a string and returns that string wrapped with the escape sequence
// to set it bold and color it red
func BoldRed(str string) string {
	return string(boldRedText + str + normalText)
}

// BoldGreen takes a string and returns that string wrapped with the escape sequence
// to set it bold and color it green
func BoldGreen(str string) string {
	return string(boldGreenText + str + normalText)
}

// BoldYellow takes a string and returns that string wrapped with the escape sequence
// to set it bold and color it yellow
func BoldYellow(str string) string {
	return string(boldYellowText + str + normalText)
}

// BoldBlue takes a string and returns that string wrapped with the escape sequence
// to set it bold and color it blue
func BoldBlue(str string) string {
	return string(boldBlueText + str + normalText)
}

// BoldMagenta takes a string and returns that string wrapped with the escape sequence
// to set it bold and color it magenta
func BoldMagenta(str string) string {
	return string(boldMagentaText + str + normalText)
}

// BoldCyan takes a string and returns that string wrapped with the escape sequence
// to set it bold and color it cyan
func BoldCyan(str string) string {
	return string(boldCyanText + str + normalText)
}

// Faint takes a string and returns that string wrapped with the escape sequence
// to render it faintly
func Faint(str string) string {
	return string(faintText + str + normalText)
}

// FaintRed takes a string and returns that string wrapped with the escape sequence
// to render it faintly and color it red
func FaintRed(str string) string {
	return string(faintRedText + str + normalText)
}

// FaintGreen takes a string and returns that string wrapped with the escape sequence
// to render it faintly and color it green
func FaintGreen(str string) string {
	return string(faintGreenText + str + normalText)
}

// FaintYellow takes a string and returns that string wrapped with the escape sequence
// to render it faintly and color it yellow
func FaintYellow(str string) string {
	return string(faintYellowText + str + normalText)
}

// FaintBlue takes a string and returns that string wrapped with the escape sequence
// to render it faintly and color it blue
func FaintBlue(str string) string {
	return string(faintBlueText + str + normalText)
}

// FaintMagenta takes a string and returns that string wrapped with the escape sequence
// to render it faintly and color it magenta
func FaintMagenta(str string) string {
	return string(faintMagentaText + str + normalText)
}

// FaintCyan takes a string and returns that string wrapped with the escape sequence
// to render it faintly and color it cyan
func FaintCyan(str string) string {
	return string(faintCyanText + str + normalText)
}

// FaintWhite takes a string and returns that string wrapped with the escape sequence
// to render it faintly and color it white
func FaintWhite(str string) string {
	return string(faintWhiteText + str + normalText)
}
