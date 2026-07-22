package ytdlp_test

import (
	"testing"

	"github.com/marc-sw/ytdlp-go-client"
)

func TestRunner(t *testing.T) {
	r := ytdlp.NewRunner("yt-dlp.exe")
	l := &ytdlp.PrintListener{}
	args := ytdlp.NewArgs()
	args.Add(ytdlp.FLAG_HELP)
	if err := r.Run(l, args.Args()...); err != nil {
		t.Fatalf("failed to run command: %s\n", err)
	}
}
