package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input       http.Header
		want        string
		expectedErr bool
	}{
		"simple": {
			input:       http.Header{"Authorization": []string{"ApiKey my_secret_key_123"}},
			want:        "my_secret_key_123",
			expectedErr: false,
		},
		"malformed_header": {
			input:       http.Header{"Authorization": []string{"Bearer my_secret_key_123"}},
			want:        "",
			expectedErr: true,
		},
		"no_auth_header": {
			input:       http.Header{"obfuscation": []string{"Bearer my_secret_key_123"}},
			want:        "",
			expectedErr: true,
		},
		"no_key": {
			input:       http.Header{"Authorization": []string{"ApiKey"}},
			want:        "",
			expectedErr: true,
		},
		"empty_key": {
			input:       http.Header{"Authorization": []string{"ApiKey "}},
			want:        "",
			expectedErr: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, gotErr := GetAPIKey(tc.input)
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("expected: %v, got: %v", tc.want, got)
			}
			if tc.expectedErr && gotErr == nil {
				t.Errorf("expected error: %v, got: %v", tc.expectedErr, gotErr)
			}
			if !tc.expectedErr && gotErr != nil {
				t.Errorf("expected error: %v, got: %v", tc.expectedErr, gotErr)
			}
		})
	}
}
