// Code generated by counterfeiter. DO NOT EDIT.
package runcontainerdfakes

import (
	"sync"
	"syscall"

	"code.cloudfoundry.org/guardian/rundmc/runcontainerd"
)

type FakeBackingProcess struct {
	IDStub        func() string
	iDMutex       sync.RWMutex
	iDArgsForCall []struct {
	}
	iDReturns struct {
		result1 string
	}
	iDReturnsOnCall map[int]struct {
		result1 string
	}
	SignalStub        func(syscall.Signal) error
	signalMutex       sync.RWMutex
	signalArgsForCall []struct {
		arg1 syscall.Signal
	}
	signalReturns struct {
		result1 error
	}
	signalReturnsOnCall map[int]struct {
		result1 error
	}
	WaitStub        func() (int, error)
	waitMutex       sync.RWMutex
	waitArgsForCall []struct {
	}
	waitReturns struct {
		result1 int
		result2 error
	}
	waitReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBackingProcess) ID() string {
	fake.iDMutex.Lock()
	ret, specificReturn := fake.iDReturnsOnCall[len(fake.iDArgsForCall)]
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct {
	}{})
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.iDReturns
	return fakeReturns.result1
}

func (fake *FakeBackingProcess) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeBackingProcess) IDCalls(stub func() string) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = stub
}

func (fake *FakeBackingProcess) IDReturns(result1 string) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeBackingProcess) IDReturnsOnCall(i int, result1 string) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeBackingProcess) Signal(arg1 syscall.Signal) error {
	fake.signalMutex.Lock()
	ret, specificReturn := fake.signalReturnsOnCall[len(fake.signalArgsForCall)]
	fake.signalArgsForCall = append(fake.signalArgsForCall, struct {
		arg1 syscall.Signal
	}{arg1})
	fake.recordInvocation("Signal", []interface{}{arg1})
	fake.signalMutex.Unlock()
	if fake.SignalStub != nil {
		return fake.SignalStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.signalReturns
	return fakeReturns.result1
}

func (fake *FakeBackingProcess) SignalCallCount() int {
	fake.signalMutex.RLock()
	defer fake.signalMutex.RUnlock()
	return len(fake.signalArgsForCall)
}

func (fake *FakeBackingProcess) SignalCalls(stub func(syscall.Signal) error) {
	fake.signalMutex.Lock()
	defer fake.signalMutex.Unlock()
	fake.SignalStub = stub
}

func (fake *FakeBackingProcess) SignalArgsForCall(i int) syscall.Signal {
	fake.signalMutex.RLock()
	defer fake.signalMutex.RUnlock()
	argsForCall := fake.signalArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBackingProcess) SignalReturns(result1 error) {
	fake.signalMutex.Lock()
	defer fake.signalMutex.Unlock()
	fake.SignalStub = nil
	fake.signalReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBackingProcess) SignalReturnsOnCall(i int, result1 error) {
	fake.signalMutex.Lock()
	defer fake.signalMutex.Unlock()
	fake.SignalStub = nil
	if fake.signalReturnsOnCall == nil {
		fake.signalReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.signalReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBackingProcess) Wait() (int, error) {
	fake.waitMutex.Lock()
	ret, specificReturn := fake.waitReturnsOnCall[len(fake.waitArgsForCall)]
	fake.waitArgsForCall = append(fake.waitArgsForCall, struct {
	}{})
	fake.recordInvocation("Wait", []interface{}{})
	fake.waitMutex.Unlock()
	if fake.WaitStub != nil {
		return fake.WaitStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.waitReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBackingProcess) WaitCallCount() int {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return len(fake.waitArgsForCall)
}

func (fake *FakeBackingProcess) WaitCalls(stub func() (int, error)) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = stub
}

func (fake *FakeBackingProcess) WaitReturns(result1 int, result2 error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = nil
	fake.waitReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeBackingProcess) WaitReturnsOnCall(i int, result1 int, result2 error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = nil
	if fake.waitReturnsOnCall == nil {
		fake.waitReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.waitReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeBackingProcess) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.signalMutex.RLock()
	defer fake.signalMutex.RUnlock()
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBackingProcess) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ runcontainerd.BackingProcess = new(FakeBackingProcess)
