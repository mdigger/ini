package ini

import (
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		ini  string
		want Config
	}{
		{
			"root = toor\n[foo]\nbar = hop\nini = nin",
			Config{
				"":    Section{"root": "toor"},
				"foo": Section{"bar": "hop", "ini": "nin"},
			},
		},
		{
			"#comment\n[empty];comment\n[section]\nempty=\n",
			Config{
				"":        Section{},
				"empty":   Section{},
				"section": Section{"empty": ""},
			},
		},
		{
			"ignore\n[invalid\n=stuff\n;comment=true\n",
			Config{
				"": Section{},
			},
		},
	}

	for _, test := range tests {
		result, err := Parse(strings.NewReader(test.ini))
		if err != nil {
			t.Errorf("Parse(%q) error %v, want: no error", test.ini, err)
			continue
		}
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("Parse(%q) = %#v, want: %#v", test.ini, result, test.want)
		}
	}
}
