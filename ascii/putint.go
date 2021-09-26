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
func PutUint32(val uint32, buf []byte) int {
	if len(buf) < 10 {
		panic(ErrShortBuffer)
	}
	var i int
	for i = len(buf) - 1; val >= 10; i-- {
		q := val / 10
		buf[i] = byte('0' + val - q*10)
		val = q
	}
	buf[i] = byte('0' + val)
	return len(buf) - i
}

// PutUint64 writes a uint64 value to the end of the byte slice.
// The function returns number of digits written. Does not modify the rest of the slice.
func PutUint64(val uint64, buf []byte) int {
	if len(buf) < 20 {
		panic(ErrShortBuffer)
	}
	var i int
	for i = len(buf) - 1; val >= 10; i-- {
		q := val / 10
		buf[i] = byte('0' + val - q*10)
		val = q
	}
	buf[i] = byte('0' + val)
	return len(buf) - i
}

// PutInt64 writes a int64 value to the end of the byte slice.
// The function returns number of digits written. Does not modify the rest of the slice.
func PutInt64(val int32, buf []byte) int {
	if val >= 0 {
		return PutUint64(uint64(val), buf)
	}
	val *= -1
	n := PutUint64(uint64(val), buf)
	buf[len(buf)-n-1] = '-'
	return n + 1
}
