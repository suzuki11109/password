package password

import "testing"

func TestGenerateWithLength(t *testing.T) {
	plain, _ := Generate(8)

	if len(plain) != 8 {
		t.Error("Expect length to be 8, but got", len(plain))
	}
}

func TestComparePassword(t *testing.T) {
	plain, hashed := Generate(8)

	if IsMatch(hashed, "awrongpassword") {
		t.Error("Expect to be not matched")
	}

	if !IsMatch(hashed, plain) {
		t.Error("Expect to be matched")
	}
}
