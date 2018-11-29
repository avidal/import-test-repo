package sub2 // import "code.heyviddy.com/import-test-repo/sub1/sub2"

import "net/http"

// T is a T that has a C
type T struct {
	C *http.Client
}
