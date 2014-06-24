package dbmigrate

import "testing"

// There are no tests yet. Sorry.

func TestTimeConsuming(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
}
