package icons

import (
	"fmt"
	"strconv"
	"strings"
)

func Icon(iconName string, width int, height int, color string) string {
	if color == "" {
		color = "currentColor"
	}

	svgContent := findIcon(iconName)

	className := iconName

	if strings.HasPrefix(className, "bi-") {
		className = "bi " + className
		return fmt.Sprintf(`<svg class="%s" xmlns="http://www.w3.org/2000/svg" width="%s" height="%s" fill="%s" viewBox="0 0 16 16" style="display: inline-block; vertical-align: -.125em;">%s</svg>`,
			className, strconv.Itoa(width), strconv.Itoa(height), color, svgContent)
	}

	if strings.HasPrefix(className, "bx-") || strings.HasPrefix(className, "bxs-") || strings.HasPrefix(className, "bxl-") {
		className = "bx " + className
		return fmt.Sprintf(`<svg class="%s" xmlns="http://www.w3.org/2000/svg" width="%s" height="%s" fill="%s" viewBox="0 0 24 24" style="display: inline-block; vertical-align: middle;">%s</svg>`,
			className, strconv.Itoa(width), strconv.Itoa(height), color, svgContent)
	}

	return fmt.Sprintf(`<svg class="%s notfound"></svg>`, className)
}
