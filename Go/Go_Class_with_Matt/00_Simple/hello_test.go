package hello

import (
	"testing"
)

func TestSayHello(t *testing.T) {

	subtests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, world!",
		},
		{
			items:  []string{"Matt"},
			result: "Hello, Matt!",
		},
		{
			items:  []string{"Luke", "Jim", "James"},
			result: "Hello, Luke, Jim, James!",
		},
	}

	for _, st := range subtests {
		{
			if s := Say(st.items); s != st.result {
				t.Errorf("\n\nwanted: '%s' with (%v)\ngot: '%s'\n\n", st.result, st.items, s)
			}
		}
	}

}
