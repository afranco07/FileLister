package main

import "testing"

// Mock struct for testing purposes
type WriterMock struct {
	data []string
}

// Overload Write function for io.Writer
func (conn *WriterMock) Write(b []byte) (n int, err error) {
	s := string(b[:])
	conn.data = append(conn.data, s)
	return 0, nil
}

// Checks if actual array slice and expected slice match
func Equal(actual, expected []string) bool {
	if len(actual) != len(expected) {
		return false
	}

	for index, value := range actual {
		if value != expected[index] {
			return false
		}
	}

	return true
}

func TestListFiles(t *testing.T) {
	var tests = []struct {
		path     string
		expected []string
	}{
		{".", []string{"/main.go\n", "/main_test.go\n"}},
		{"/DoesNotExist", []string{"/DoesNotExist does not exist\n"}},
	}

	for _, test := range tests {
		connection := new(WriterMock)
		ListFiles(test.path, connection)
		if !Equal(connection.data, test.expected) {
			t.Error("Test failed! Path:", test.path, "Expected:", test.expected, "Actual:", connection.data)
		}
	}
}
