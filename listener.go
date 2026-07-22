package ytdlp

import "fmt"

type Listener interface {
	OnStdout(line string)
	OnStderr(line string)
	OnExit(code int)
}

type PrintListener struct {
}

func (p *PrintListener) OnStdout(line string) {
	fmt.Println("out: " + line)
}

func (p *PrintListener) OnStderr(line string) {
	fmt.Println("err: " + line)
}

func (ü *PrintListener) OnExit(code int) {
	fmt.Println("exit rcv: " + fmt.Sprint(code))
}
