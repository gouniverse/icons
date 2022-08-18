package icons

import (
	"strings"

	"github.com/gouniverse/hb"
	"github.com/gouniverse/utils"
)

func Icon(iconName string, width int, height int, color string) *hb.Tag {
	if color == "" {
		color = "currentColor"
	}

	svgContent := findIcon(iconName)

	className := iconName

	if strings.HasPrefix(className, "bi-") {
		className = "bi " + className
		svg := hb.NewTag("svg").
			Class(className).
			Attr("xmlns", "http://www.w3.org/2000/svg").
			Attr("width", utils.ToString(width)).
			Attr("height", utils.ToString(height)).
			Attr("fill", color).
			Attr("viewBox", "0 0 16 16").
			HTML(svgContent)
		return svg
	}

	if strings.HasPrefix(className, "bx-") || strings.HasPrefix(className, "bxs-") || strings.HasPrefix(className, "bxl-") {
		className = "bx " + className
		svg := hb.NewTag("svg").
			Class(className).
			Attr("xmlns", "http://www.w3.org/2000/svg").
			Attr("width", utils.ToString(width)).
			Attr("height", utils.ToString(height)).
			Attr("fill", color).
			//Attr("viewBox", "0 0 24 24").
			HTML(svgContent)
		return svg
	}

	return hb.NewTag("svg").Class($cclassName+" notfound")
}
