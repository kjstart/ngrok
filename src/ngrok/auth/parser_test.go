package auth
 import (
	"testing"
)
 func TestParser(t *testing.T) {
	path := "./ngrok-secrets"
	parser := NewParser(path)
	if err := parser.Parse(); err != nil {
		t.Fatal(err)
	}
	t.Log(parser.Tokens)
}