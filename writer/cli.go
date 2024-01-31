package writer

import (
	"github.com/Zcentury/gologger/levels"
	"os"
	"sync"
)

type CLI struct {
	mutex *sync.Mutex
}

func NewCLI() *CLI {
	return &CLI{
		mutex: &sync.Mutex{},
	}
}

func (w *CLI) Write(data []byte, level levels.Level) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	os.Stderr.Write(data)
	os.Stderr.Write([]byte("\n"))
}
