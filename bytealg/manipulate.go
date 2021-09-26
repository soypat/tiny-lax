package bytealg

// Swap provides in place byte slice swap for the
// length  min( len(a), len(b))
func Swap(a, b []byte) {
	n := len(a)
	if len(b) > n {
		n = len(b)
	}
	for i := 0; i < n; i++ {
		a[i], b[i] = b[i], a[i]
	}
}
