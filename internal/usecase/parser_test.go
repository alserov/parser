package usecase

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	tests := []struct {
		name      string
		url       string
		tag       string
		expectErr bool
	}{
		{
			name:      "habr links",
			url:       "https://habr.com",
			tag:       "a",
			expectErr: false,
		},
		{
			name:      "empty url",
			url:       "",
			tag:       "",
			expectErr: true,
		},
	}

	p := Parser{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res, err := p.Parse(context.Background(), tc.url, tc.tag)
			if tc.expectErr {
				require.Error(t, err)
				require.Empty(t, res)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, res)
			}
		})
	}
}
