package firstmod

import "testing"

func TestHello(t *testing.T) {
    var param = "Tom"
    want := "Hello " + param
    if got := Hello(param); got != want {
        t.Errorf("Hello(%q) = %q, want %q", param, got, want)
    }
}

