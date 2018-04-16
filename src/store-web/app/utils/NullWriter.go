package utils

// NullWriter .
type NullWriter struct {
}

// Write .
func (NullWriter) Write(b []byte) (int, error) {
	return len(b), nil
}
