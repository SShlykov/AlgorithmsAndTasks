package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInMemoryCache_Get(t *testing.T) {
	c := NewInMemoryCache[string, string]()
	c.Put("key", "value")

	cases := []struct {
		name    string
		key     string
		want    string
		wantErr error
	}{
		{name: "Существующий ключ", key: "key", want: "value", wantErr: nil},
		{name: "non  key", key: "key", want: "value", wantErr: nil},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := c.Get(tt.key)
			assert.Equal(t, tt.want, got)
		})
	}
}
