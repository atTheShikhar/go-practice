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
	t.Run("normal tests", func(t *testing.T) {
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
	})

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Box",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Box")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{23, "Cologne"}
			aChannel <- Profile{23, "Antwerp"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Cologne", "Antwerp"}
		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with functions", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{23, "Cologne"}, Profile{23, "Katowice"}
		}

		var got []string
		want := []string{"Cologne", "Katowice"}
		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()

	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
