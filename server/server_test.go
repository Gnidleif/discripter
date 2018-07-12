package server

import (
	"os"
	"testing"
)

var (
	testToken = os.Getenv("DGB_TOKEN")
	scriptDir = os.Getenv("DG_SCRIPTS")
)

func TestStart(t *testing.T) {
	cases := []struct {
		t    string
		s    string
		res  bool
		name string
	}{
		{testToken, scriptDir, true, "correct"},
		{"bad_token>>,<LL<L<M,..", scriptDir, false, "bad_token"},
		{testToken, "bad_dir", false, "bad_dir"},
		{"", scriptDir, false, "no token"},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			err := Start(tt.t, tt.s)
			if tt.res && err != nil {
				t.Error(err)
			}
		})
	}
}
