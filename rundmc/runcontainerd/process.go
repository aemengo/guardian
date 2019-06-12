package runcontainerd

import (
	"fmt"
	"syscall"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/lager"
)

type Process struct {
	log         lager.Logger
	containerID string
	processID   string

	processManager ProcessManager
}

func NewProcess(log lager.Logger, containerID, processID string, processManager ProcessManager) *Process {
	return &Process{log: log, containerID: containerID, processID: processID, processManager: processManager}
}

func (p *Process) Wait() (int, error) {
	exitCode, err := p.processManager.Wait(p.log, p.containerID, p.processID)
	if err != nil {
		return -1, err
	}

	return exitCode, nil
}

func (p *Process) Signal(gardenSignal garden.Signal) error {
	signal, err := toSyscallSignal(gardenSignal)
	if err != nil {
		return err
	}

	return p.processManager.Signal(p.log, p.containerID, p.processID, signal)
}

func (p *Process) ID() string {
	return p.processID
}

func (p *Process) SetTTY(garden.TTYSpec) error {
	return nil
}

func toSyscallSignal(signal garden.Signal) (syscall.Signal, error) {
	switch signal {
	case garden.SignalTerminate:
		return syscall.SIGTERM, nil
	case garden.SignalKill:
		return syscall.SIGKILL, nil
	}

	return -1, fmt.Errorf("Cannot convert garden signal %d to syscall.Signal", signal)
}
