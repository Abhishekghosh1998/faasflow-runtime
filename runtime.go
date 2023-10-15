package runtime

import (
	"github.com/Abhishekghosh1998/faasflow-sdk/executor"
)

type Runtime interface {
	Init() error
	CreateExecutor(*Request) (executor.Executor, error)
}
