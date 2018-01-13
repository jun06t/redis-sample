package autocomplete

import (
	"reflect"
	"testing"
)

func TestAddUpdateContact(t *testing.T) {
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
			out: []string{"user2", "user1"},
		},
		{
			in:  "user1",
			out: []string{"user1", "user2"},
		},
	}
	client := New()
	list := "recent:"
	defer client.Del(list)

	for _, v := range tests {
		err := AddUpdateContact(client, list, v.in)
		if err != nil {
			t.Errorf("%s", err)
		}

		out := client.LRange(list, 0, -1).Val()
		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("get: %v, want: %v\n", out, v.out)
		}
	}
}

func TestRemoveContact(t *testing.T) {
	tests := []struct {
		before []string
		in     string
		out    []string
	}{
		{
			before: []string{"user1", "user2"},
			in:     "user1",
			out:    []string{"user2"},
		},
		{
			before: []string{"user1", "user2"},
			in:     "user2",
			out:    []string{"user1"},
		},
	}

	client := New()
	list := "recent:"
	defer client.Del(list)

	for _, v := range tests {
		for _, val := range v.before {
			client.LPush(list, val)
		}
		err := RemoveContact(client, list, v.in)
		if err != nil {
			t.Errorf("%s", err)
		}

		out := client.LRange(list, 0, -1).Val()
		if !reflect.DeepEqual(out, v.out) {
			t.Errorf("get: %v, want: %v\n", out, v.out)
		}
	}
}
