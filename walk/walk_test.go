package walk

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Profile Profile
}

type Profile struct {
	Age int
	City string
}

func TestWalk(t *testing.T) {
	// 可以使用多种情况进行测试
	cases := []struct{
		Name string
		Input interface{}
		ExpectedCalls []string
	} {
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "China"},
			[]string{"Chris", "China"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"Nested fields",
			Person{ "Chris", Profile{33, "China"}},
			[]string{"Chris", "China"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v want %v", got, test.ExpectedCalls)
			}
		})
	}
}