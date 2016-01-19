package sanitize

import "testing"

func TestEmail(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"soroushmehraein@gmail.com", true},
		{"soroushmehraein", false},
	}
	for _, c := range cases {
		got := Email(c.in)
		if got != c.want {
			t.Errorf("Email(%q) == %t, want %t", c.in, got, c.want)
		}
	}
}

func TestName(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"soroushmehraein", true},
		{"SOROUSHMEHRAEIN", true},
		{"/soroushmehraein", false},
		{".soroushmehraein", false},
		{"<soroushmehraein", false},
		{">soroushmehraein", false},
		{"", false},
	}
	for _, c := range cases {
		got := Name(c.in)
		if got != c.want {
			t.Errorf("Name(%q) == %t, want %t", c.in, got, c.want)
		}
	}
}

func TestRadio(t *testing.T) {
	cases := []struct {
		in      int
		choices []int
		want    bool
	}{
		{1, []int{1, 2}, true},
		{2, []int{1, 2}, true},
		{3, []int{1, 2}, false},
	}
	for _, c := range cases {
		got := Radio(c.in, c.choices)
		if got != c.want {
			t.Errorf("Radio(%q) with choices(%q) == %t, want %t", c.in, c.choices, got, c.want)
		}
	}
}

func TestSelect(t *testing.T) {
	cases := []struct {
		in      string
		choices []string
		want    bool
	}{
		{"apple", []string{"apple", "banana", "orange"}, true},
		{"banana", []string{"apple", "banana", "orange"}, true},
		{"egg", []string{"apple", "banana", "orange"}, false},
	}
	for _, c := range cases {
		got := Select(c.in, c.choices)
		if got != c.want {
			t.Errorf("Select(%q) with choices(%q) == %t, want %t", c.in, c.choices, got, c.want)
		}
	}
}
