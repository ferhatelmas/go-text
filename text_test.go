package text

import "testing"

func TestExtension(t *testing.T) {
	curLen := 133
	if len(Extensions) != curLen {
		t.Fatalf("Length doesn't match. Expected %d, Got %d", curLen, len(Extensions))
	}
}

func TestIs(t *testing.T) {
	if Is("a/b/c/bar.exe") {
		t.Fatal("Wrong detection. Must not be text")
	}

	if !Is("a/b/c/bar.go") {
		t.Fatal("Wrong detection. Must be text")
	}

	if !Is("a/b/c/bar.GO") {
		t.Fatal("Wrong detection. Must be text")
	}

	if Is("a/b/c/bargo") {
		t.Fatal("Wrong detection. Must not be text")
	}
}
