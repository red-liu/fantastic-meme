package firstmod

import "testing"

func TestHello(t *testing.T) {
    var param = "Tom"
    want := "Hello " + param
    if got := Hello(param); got != want {
        t.Errorf("Hello(%q) = %q, want %q", param, got, want)
    }
}


func TestHelloV2(t *testing.T) {
    var param = "Tom"
    want := "<h1>Hello " + param + "</h1>"
    if got := HelloV2(param); got != want {
        t.Errorf("Hello(%q) = %q, want %q", param, got, want)
    }
}
