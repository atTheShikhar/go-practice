package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name: "struct with one string field",
			Input: struct {
				Name string
			}{Name: "Shikhar"},
			ExpectedCalls: []string{"Shikhar"},
		},
		{
			Name: "struct with 2 string fields",
			Input: struct {
				Name string
				City string
			}{Name: "Shikhar", City: "Vns"},
			ExpectedCalls: []string{"Shikhar", "Vns"},
		},
		{
			Name: "struct with non string field",
			Input: struct {
				Name string
				Age  int
			}{Name: "Shikhar", Age: 23},
			ExpectedCalls: []string{"Shikhar"},
		},
		{
			Name: "struct with nested field",
			Input: Person{
				Name:    "Shikhar",
				Profile: Profile{Age: 23, City: "Vns"},
			},
			ExpectedCalls: []string{"Shikhar", "Vns"},
		},
		{
			Name: "pointer to things",
			Input: &Person{
				Name:    "Shikhar",
				Profile: Profile{Age: 23, City: "Vns"},
			},
			ExpectedCalls: []string{"Shikhar", "Vns"},
		},
		{
			Name: "slices",
			Input: []Profile{
				{23, "Shikhar"},
				{23, "Aman"},
			},
			ExpectedCalls: []string{"Shikhar", "Aman"},
		},
		{
			Name: "arrays",
			Input: [2]Profile{
				{23, "Shikhar"},
				{23, "Aman"},
			},
			ExpectedCalls: []string{"Shikhar", "Aman"},
		},
		{
			Name: "maps",
			Input: map[string]string{
				"Foo": "Bar",
				"Baz": "Box",
			},
			ExpectedCalls: []string{"Bar", "Box"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("expected %q, got %q", test.ExpectedCalls, got)
			}
		})
	}
}
