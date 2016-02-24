// This file was generated by counterfeiter
package fakes

import (
	"os"
	"sync"

	"github.com/cloudfoundry-incubator/guardian/rundmc/runrunc"
)

type FakeMkdirer struct {
	MkdirAsStub        func(path string, mode os.FileMode, uid, gid int) error
	mkdirAsMutex       sync.RWMutex
	mkdirAsArgsForCall []struct {
		path string
		mode os.FileMode
		uid  int
		gid  int
	}
	mkdirAsReturns struct {
		result1 error
	}
}

func (fake *FakeMkdirer) MkdirAs(path string, mode os.FileMode, uid int, gid int) error {
	fake.mkdirAsMutex.Lock()
	fake.mkdirAsArgsForCall = append(fake.mkdirAsArgsForCall, struct {
		path string
		mode os.FileMode
		uid  int
		gid  int
	}{path, mode, uid, gid})
	fake.mkdirAsMutex.Unlock()
	if fake.MkdirAsStub != nil {
		return fake.MkdirAsStub(path, mode, uid, gid)
	} else {
		return fake.mkdirAsReturns.result1
	}
}

func (fake *FakeMkdirer) MkdirAsCallCount() int {
	fake.mkdirAsMutex.RLock()
	defer fake.mkdirAsMutex.RUnlock()
	return len(fake.mkdirAsArgsForCall)
}

func (fake *FakeMkdirer) MkdirAsArgsForCall(i int) (string, os.FileMode, int, int) {
	fake.mkdirAsMutex.RLock()
	defer fake.mkdirAsMutex.RUnlock()
	return fake.mkdirAsArgsForCall[i].path, fake.mkdirAsArgsForCall[i].mode, fake.mkdirAsArgsForCall[i].uid, fake.mkdirAsArgsForCall[i].gid
}

func (fake *FakeMkdirer) MkdirAsReturns(result1 error) {
	fake.MkdirAsStub = nil
	fake.mkdirAsReturns = struct {
		result1 error
	}{result1}
}

var _ runrunc.Mkdirer = new(FakeMkdirer)