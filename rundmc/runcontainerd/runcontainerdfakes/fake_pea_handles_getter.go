// Code generated by counterfeiter. DO NOT EDIT.
package runcontainerdfakes

import (
	"sync"

	"code.cloudfoundry.org/guardian/rundmc/runcontainerd"
	"code.cloudfoundry.org/lager"
)

type FakePeaHandlesGetter struct {
	ContainerPeaHandlesStub        func(lager.Logger, string) ([]string, error)
	containerPeaHandlesMutex       sync.RWMutex
	containerPeaHandlesArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	containerPeaHandlesReturns struct {
		result1 []string
		result2 error
	}
	containerPeaHandlesReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePeaHandlesGetter) ContainerPeaHandles(arg1 lager.Logger, arg2 string) ([]string, error) {
	fake.containerPeaHandlesMutex.Lock()
	ret, specificReturn := fake.containerPeaHandlesReturnsOnCall[len(fake.containerPeaHandlesArgsForCall)]
	fake.containerPeaHandlesArgsForCall = append(fake.containerPeaHandlesArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("ContainerPeaHandles", []interface{}{arg1, arg2})
	fake.containerPeaHandlesMutex.Unlock()
	if fake.ContainerPeaHandlesStub != nil {
		return fake.ContainerPeaHandlesStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.containerPeaHandlesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePeaHandlesGetter) ContainerPeaHandlesCallCount() int {
	fake.containerPeaHandlesMutex.RLock()
	defer fake.containerPeaHandlesMutex.RUnlock()
	return len(fake.containerPeaHandlesArgsForCall)
}

func (fake *FakePeaHandlesGetter) ContainerPeaHandlesCalls(stub func(lager.Logger, string) ([]string, error)) {
	fake.containerPeaHandlesMutex.Lock()
	defer fake.containerPeaHandlesMutex.Unlock()
	fake.ContainerPeaHandlesStub = stub
}

func (fake *FakePeaHandlesGetter) ContainerPeaHandlesArgsForCall(i int) (lager.Logger, string) {
	fake.containerPeaHandlesMutex.RLock()
	defer fake.containerPeaHandlesMutex.RUnlock()
	argsForCall := fake.containerPeaHandlesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePeaHandlesGetter) ContainerPeaHandlesReturns(result1 []string, result2 error) {
	fake.containerPeaHandlesMutex.Lock()
	defer fake.containerPeaHandlesMutex.Unlock()
	fake.ContainerPeaHandlesStub = nil
	fake.containerPeaHandlesReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakePeaHandlesGetter) ContainerPeaHandlesReturnsOnCall(i int, result1 []string, result2 error) {
	fake.containerPeaHandlesMutex.Lock()
	defer fake.containerPeaHandlesMutex.Unlock()
	fake.ContainerPeaHandlesStub = nil
	if fake.containerPeaHandlesReturnsOnCall == nil {
		fake.containerPeaHandlesReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.containerPeaHandlesReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakePeaHandlesGetter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.containerPeaHandlesMutex.RLock()
	defer fake.containerPeaHandlesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePeaHandlesGetter) recordInvocation(key string, args []interface{}) {
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

var _ runcontainerd.PeaHandlesGetter = new(FakePeaHandlesGetter)