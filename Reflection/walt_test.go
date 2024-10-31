package reflection

import "testing"

func TestWalt(t *testing.T) {
	// cases := []struct {
	// 	Name          string
	// 	Input         interface{}
	// 	ExpectedCalls []string
	// }{{
	// 	"struct with one string field",
	// 	struct {
	// 		Name string
	// 	}{"Chris"},
	// 	[]string{"Chris"},
	// },
	// }

	expected := "Chris"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
	}
	if got[0] != expected {
		t.Errorf("got %q want %q", got[0], expected)
	}
}
