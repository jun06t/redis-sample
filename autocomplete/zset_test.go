package autocomplete

import "testing"

func TestFindPrefixRange(t *testing.T) {
	tests := []struct {
		in    string
		start string
		end   string
	}{
		{
			in:    "joh",
			start: "jog{",
			end:   "joh{",
		},
		{
			in:    "john",
			start: "johm{",
			end:   "john{",
		},
	}
	for _, v := range tests {
		start, end := FindPrefixRange(v.in)

		if start != v.start {
			t.Errorf("get: %v, want: %v\n", start, v.start)
		}
		if end != v.end {
			t.Errorf("get: %v, want: %v\n", end, v.end)
		}
	}
}
