package ip2loc

import (
	"testing"
)

func TestStrToInt(t *testing.T) {
	cases := []string{"0.0.0.0", "211.81.55.60", "8.8.8.8", "46.2.234.149", "102.170.47.246"}
	for _, c := range cases {
		t.Log(Find(c))
	}
}
