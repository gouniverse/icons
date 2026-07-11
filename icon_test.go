package icons

import (
	"strings"
	"testing"
)

func TestIcon(t *testing.T) {
	iconHtml := Icon("bi-globe", 120, 120, "white").ToHTML()

	if !strings.Contains(iconHtml, `class="bi bi-globe"`) {
		t.Errorf("Icon does not return expected result 'class=\"bi bi-globe\"', but '%s'", iconHtml)
	}

	if !strings.Contains(iconHtml, `width="120"`) {
		t.Errorf("Icon does not return expected result 'width=\"120\"', but '%s'", iconHtml)
	}

	if !strings.Contains(iconHtml, `height="120"`) {
		t.Errorf("Icon does not return expected result 'height=\"120\"', but '%s'", iconHtml)
	}

	if !strings.Contains(iconHtml, `fill="white"`) {
		t.Errorf("Icon does not return expected result 'fill=\"white\"', but '%s'", iconHtml)
	}

	if !strings.Contains(iconHtml, `<path `) {
		t.Errorf("Icon does not return expected result '<path ', but '%s'", iconHtml)
	}

	if !strings.Contains(iconHtml, `/></svg>`) {
		t.Errorf("Icon does not return expected result '/></svg>', but '%s'", iconHtml)
	}
}

func TestIconBoxicon(t *testing.T) {
	iconHtml := Icon("bxs-color", 120, 120, "white").ToHTML()

	if !strings.Contains(iconHtml, `class="bx bxs-color"`) {
		t.Errorf("Icon does not return expected result 'class=\"bx bxs-color\"', but '%s'", iconHtml)
	}

	// if !strings.Contains(iconHtml, `width="120"`) {
	// 	t.Errorf("Icon does not return expected result 'width=\"120\"', but '%s'", iconHtml)
	// }

	// if !strings.Contains(iconHtml, `height="120"`) {
	// 	t.Errorf("Icon does not return expected result 'height=\"120\"', but '%s'", iconHtml)
	// }

	// if !strings.Contains(iconHtml, `fill="white"`) {
	// 	t.Errorf("Icon does not return expected result 'fill=\"white\"', but '%s'", iconHtml)
	// }

	// if !strings.Contains(iconHtml, `<path `) {
	// 	t.Errorf("Icon does not return expected result '<path ', but '%s'", iconHtml)
	// }

	// if !strings.Contains(iconHtml, `/></svg>`) {
	// 	t.Errorf("Icon does not return expected result '/></svg>', but '%s'", iconHtml)
	// }
}
