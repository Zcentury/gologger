package writer

import "github.com/Zcentury/gologger/levels"

type Writer interface {
	Write(data []byte, level levels.Level)
}
