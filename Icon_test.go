package icons

package utils

import "testing"

func TestIcon(t *testing.T) {
  iconHtml := Icon("bi-globe", 120, 120, "white").ToHtml()

  if strings.Contain(iconHtml, "test") {
		t.Errorf("Icon does not return expected result '01', but '%s'", iconHtml)
	}
}
