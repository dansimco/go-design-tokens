package css_util

import (
	"os"
	"testing"
)

func TestCSSFormatting(t *testing.T) {
	input_data, _ := os.ReadFile("test_input.css")
	input_css := string(input_data)

	formatted_css := Format(input_css)

	expected_data, _ := os.ReadFile("test_output.css")
	expected_css := string(expected_data)

	if expected_css != formatted_css {
		t.Errorf("CSS does not match format, got %s", formatted_css)
	}

	println(input_css)
}
