package ytdlp

import "strings"

type Arg interface {
	apply(*Args)
}

type Args struct {
	args []string
}

func NewArgs(args ...string) *Args {
	return &Args{args: append([]string(nil), args...)}
}

func (c *Args) Args() []string {
	return append([]string(nil), c.args...)
}

func (c *Args) Add(arg Arg) {
	arg.apply(c)
}

func (c *Args) String() string {
	return strings.Join(c.Args(), " ")
}
