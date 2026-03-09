package typography

import "testing"

func TestCSSFontFaceGeneration(t *testing.T) {

	family := NewFontFamily("Helvetica Now")

	helvetica_now_regular := family.AddFont()
	helvetica_now_regular.AddLocalSrc("Helvetica Now")
	helvetica_now_regular.AddSrc("/assets/fonts/helvetica_now_regular.woff2")

	helvetica_now_bold := family.AddFont()
	helvetica_now_bold.AddLocalSrc("Helvetica Now")
	helvetica_now_bold.AddSrc("/assets/fonts/helvetica_now_bold")
	helvetica_now_bold.SetWeightNumber(600)

	family_css := family.ToCSS()

	expected_css := `
	@font-face {
	  font-family: "Helvetica Now";
	  src:
        local("Helvetica Now"),
        url("/assets/fonts/helvetica_now_regular.woff2");
      weight: 400;
    }

   	@font-face {
	  font-family: "Helvetica Now";
	  src:
        local("Helvetica Now"),
        url("/assets/fonts/helvetica_now_bold.woff2");
      weight: 600;
    }
	`

	if family_css != expected_css {
		t.Errorf("Expected family_css to equal happy_css, got %s", family_css)
	}

}
