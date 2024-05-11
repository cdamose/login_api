package app

import "testing"

func TestPing(t *testing.T) {
	testcase := []struct {
		name string
		want string
	}{
		{
			name: "Possitive Case",
			want: "Hi Dude i'm able to access database",
		},
	}
	for _, test := range testcase {
		t.Run(test.name, func(t *testing.T) {

		})
	}
}
