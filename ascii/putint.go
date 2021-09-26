package ascii

import "errors"

var (
	ErrShortBuffer = errors.New("short buffer")
)

// PutInt32 writes a int32 value to the end of the byte slice.
// The function returns number of digits written. Does not modify the rest of the slice.
func PutInt32(val int32, buf []byte) int {
	if len(buf) < 11 {
		panic(ErrShortBuffer)
	}
	if val >= 0 {
		return PutUint32(uint32(val), buf)
	}
	val *= -1
	n := PutUint32(uint32(val), buf)
	buf[len(buf)-n-1] = '-'
	return n + 1
}

// PutUint32 writes a uint32 value to the end of the byte slice.
// The function returns number of digits written. Does not modify the rest of the slice.
func PutUint32(val uint32, uibuf []byte) int {
	if len(uibuf) < 10 {
		panic(ErrShortBuffer)
	}
	var i int
	for i = len(uibuf) - 1; val >= 10; i-- {
		q := val / 10
		uibuf[i] = byte('0' + val - q*10)
		val = q
	}
	uibuf[i] = byte('0' + val)
	return len(uibuf) - i
}
