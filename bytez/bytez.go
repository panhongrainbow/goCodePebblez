package bytez

import "unsafe"

// Quotation adds double quotes around a byte slice.
func Quotation(slice []byte) (quoted []byte) {
	quoted = make([]byte, len(slice)+2, len(slice)+2)
	quoted[0] = 34
	quoted[len(quoted)-1] = 34
	copy(quoted[1:], slice)
	return
}

// HasPrefix checks if two byte slices have a matching prefix.
func HasPrefix(a, b []byte) bool {
	minLen := len(a)
	if len(b) < minLen {
		minLen = len(b)
	}

	for i := 0; i < minLen; i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// StringToReadOnlyBytes converts string to read-only []byte using unsafe pointer manipulation.
func StringToReadOnlyBytes(s string) []byte {
	data := *(*unsafe.Pointer)(unsafe.Pointer(&s))
	length := len(s)
	return *(*[]byte)(unsafe.Pointer(&struct {
		Data uintptr
		Len  int
	}{uintptr(data), length}))
}
