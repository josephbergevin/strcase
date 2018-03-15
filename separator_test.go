package strcase

import "testing"

func TestToKebab(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			"From camel",
			"FromCamel",
			"from-camel",
		},
		{
			"From lower camel",
			"fromLowerCamel",
			"from-lower-camel",
		},
		{
			"From snake",
			"from_snake",
			"from-snake",
		},
		{
			"From sentence",
			"From sentence",
			"from-sentence",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToKebab(tt.s); got != tt.want {
				t.Errorf("ToKebab() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToSentence(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			"From camel",
			"FromCamel",
			"from camel",
		},
		{
			"From lower camel",
			"fromLowerCamel",
			"from lower camel",
		},
		{
			"From kebab",
			"from-kebab",
			"from kebab",
		},
		{
			"From snake",
			"from_snake",
			"from snake",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToSentence(tt.s); got != tt.want {
				t.Errorf("ToSentence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToSnake(t *testing.T) {
	cases := [][]string{
		[]string{"testCase", "test_case"},
		[]string{"TestCase", "test_case"},
		[]string{"Test Case", "test_case"},
		[]string{" Test Case", "test_case"},
		[]string{"Test Case ", "test_case"},
		[]string{" Test Case ", "test_case"},
		[]string{"test", "test"},
		[]string{"test_case", "test_case"},
		[]string{"Test", "test"},
		[]string{"", ""},
		[]string{"ManyManyWords", "many_many_words"},
		[]string{"manyManyWords", "many_many_words"},
		[]string{"AnyKind of_string", "any_kind_of_string"},
		[]string{"numbers2and55with000", "numbers_2_and_55_with_000"},
		[]string{"JSONData", "json_data"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := ToSnake(in)
		if result != out {
			t.Error("'" + result + "' != '" + out + "'")
		}
	}
}

func TestToSnake2(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			"From camel",
			"FromCamel",
			"from_camel",
		},
		{
			"From lower camel",
			"fromLowerCamel",
			"from_lower_camel",
		},
		{
			"From kebab",
			"from-kebab",
			"from_kebab",
		},
		{
			"From sentence",
			"From sentence",
			"from_sentence",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToSnake(tt.s); got != tt.want {
				t.Errorf("ToSnake() = %v, want %v", got, tt.want)
			}
		})
	}
}
