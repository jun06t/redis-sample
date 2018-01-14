package autocomplete

import (
	"reflect"
	"testing"
)

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

func TestJoinList(t *testing.T) {
	tests := []struct {
		in  string
		out []string
	}{
		{
			in:  "user1",
			out: []string{"user1"},
		},
		{
			in:  "user2",
			out: []string{"user1", "user2"},
		},
		{
			in:  "admin",
			out: []string{"admin", "user1", "user2"},
		},
	}
	client := New()
	list := "member:"
	defer client.Del(list)

	for _, v := range tests {
		JoinList(client, list, v.in)

		out := client.ZRange(list, 0, -1).Val()
		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("get: %v, want: %v\n", out, v.out)
		}
	}
}

func TestLeaveList(t *testing.T) {
	tests := []struct {
		before []string
		in     string
		out    []string
	}{
		{
			before: []string{"admin", "user1", "user2"},
			in:     "admin",
			out:    []string{"user1", "user2"},
		},
	}
	client := New()
	list := "member:"
	defer client.Del(list)

	for _, v := range tests {
		for _, val := range v.before {
			JoinList(client, list, val)
		}

		LeaveList(client, list, v.in)

		out := client.ZRange(list, 0, -1).Val()
		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("get: %v, want: %v\n", out, v.out)
		}
	}
}
