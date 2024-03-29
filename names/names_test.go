package names

import "testing"

func TestIsSus(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", false},
		{"Hello", false},
		{"CLUCKER", false},
		{"Uniswap Reward", true},
		{"Anoncement", true},
		{"📢Big news", true},
		{"📣Important", true},
		{"📡Text here", true},
		{"📡ANNOUNCEMENT", true},
		{"HelloANNOUNCEMENT", true},
		{"Airdrep🔥", true},
		{"Uniswep", true},
		{"ANNOUCENMENT", true},
	}

	for _, test := range tests {
		if got, _ := NameIsSuspicious(test.input); got != test.want {
			t.Errorf("IsSus(%q) = %v", test.input, got)
		}
	}
}
