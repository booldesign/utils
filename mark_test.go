package utils

import (
	"testing"
)

func TestMarkMobile(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"14700000000", "147****0000"},
	}

	for _, test := range tests {
		if actual := MarkMobile(test.input); actual != test.expect {
			t.Errorf("MarkMobile(%s): expect %s, actual %s",
				test.input, test.expect, actual)
		}
	}
}

func TestMarkEmail(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"booldesign@gmail.com", "b********n@gmail.com"},
	}

	for _, test := range tests {
		if actual := MarkEmail(test.input); actual != test.expect {
			t.Errorf("MarkEmail(%s): expect %s, actual %s",
				test.input, test.expect, actual)
		}
	}
}

func TestMarkRealName(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"卫建文", "*建文"},
	}

	for _, test := range tests {
		if actual := MarkRealName(test.input); actual != test.expect {
			t.Errorf("MarkRealName(%s): expect %s, actual %s",
				test.input, test.expect, actual)
		}
	}
}

func TestMarkBankCard(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"6214830000003553", "621483******3553"},
	}

	for _, test := range tests {
		if actual := MarkBankCard(test.input); actual != test.expect {
			t.Errorf("MarkBankCard(%s): expect %s, actual %s",
				test.input, test.expect, actual)
		}
	}
}

func TestMarkIDCard(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"110101199003078793", "1101011990****8793"},
	}

	for _, test := range tests {
		if actual := MarkIDCard(test.input); actual != test.expect {
			t.Errorf("MarkIDCard(%s): expect %s, actual %s",
				test.input, test.expect, actual)
		}
	}
}
