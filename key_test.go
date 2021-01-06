package ttlcache

import (
	"testing"
)

func TestKeys(t *testing.T) {
	tt := []struct {
		name     string
		actual   uint64
		expected uint64
	}{
		{
			name:     "IntKey",
			actual:   IntKey(42),
			expected: 42,
		},
		{
			name:     "ByteKey",
			actual:   ByteKey(42),
			expected: 42,
		},
		{
			name:     "Int8Key",
			actual:   Int8Key(42),
			expected: 42,
		},
		{
			name:     "Uint8Key",
			actual:   Uint8Key(42),
			expected: 42,
		},
		{
			name:     "Int16Key",
			actual:   Int16Key(42),
			expected: 42,
		},
		{
			name:     "Uint16Key",
			actual:   Uint16Key(42),
			expected: 42,
		},
		{
			name:     "Int32Key",
			actual:   Int32Key(42),
			expected: 42,
		},
		{
			name:     "Uint32Key",
			actual:   Uint32Key(42),
			expected: 42,
		},
		{
			name:     "Int64Key",
			actual:   Int64Key(42),
			expected: 42,
		},
		{
			name:     "Uint64Key",
			actual:   Uint64Key(42),
			expected: 42,
		},
		{
			name:     "BytesKey",
			actual:   BytesKey([]byte("hello world")),
			expected: 13388989860809387070,
		},
		{
			name:     "StringKey",
			actual:   StringKey("hello world"),
			expected: 13388989860809387070,
		},
		{
			name: "AnyKey_struct",
			actual: AnyKey(struct {
				X, Y int
			}{
				X: 1,
				Y: 2,
			}),
			expected: 9173886622172901187,
		},
		{
			name:     "AnyKey_array",
			actual:   AnyKey([...]int{1, 2, 3}),
			expected: 2675078694837399863,
		},
	}

	for _, tc := range tt {
		if tc.expected != tc.actual {
			t.Errorf("%s expected: %v got: %v", tc.name, tc.expected, tc.actual)
		}
	}
}
