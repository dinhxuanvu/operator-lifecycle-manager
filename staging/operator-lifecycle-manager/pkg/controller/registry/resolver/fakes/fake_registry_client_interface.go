// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"

	"github.com/operator-framework/operator-lifecycle-manager/pkg/controller/registry"
	"github.com/operator-framework/operator-registry/pkg/api"
)

type FakeRegistryClientInterface struct {
	FindBundleThatProvidesStub        func(context.Context, string, string, string, string) (*api.Bundle, error)
	findBundleThatProvidesMutex       sync.RWMutex
	findBundleThatProvidesArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}
	findBundleThatProvidesReturns struct {
		result1 *api.Bundle
		result2 error
	}
	findBundleThatProvidesReturnsOnCall map[int]struct {
		result1 *api.Bundle
		result2 error
	}
	GetLatestChannelEntriesThatProvideStub        func(context.Context, string, string, string) (*registry.ChannelEntryIterator, error)
	getLatestChannelEntriesThatProvideMutex       sync.RWMutex
	getLatestChannelEntriesThatProvideArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
	}
	getLatestChannelEntriesThatProvideReturns struct {
		result1 *registry.ChannelEntryIterator
		result2 error
	}
	getLatestChannelEntriesThatProvideReturnsOnCall map[int]struct {
		result1 *registry.ChannelEntryIterator
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRegistryClientInterface) FindBundleThatProvides(arg1 context.Context, arg2 string, arg3 string, arg4 string, arg5 string) (*api.Bundle, error) {
	fake.findBundleThatProvidesMutex.Lock()
	ret, specificReturn := fake.findBundleThatProvidesReturnsOnCall[len(fake.findBundleThatProvidesArgsForCall)]
	fake.findBundleThatProvidesArgsForCall = append(fake.findBundleThatProvidesArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("FindBundleThatProvides", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.findBundleThatProvidesMutex.Unlock()
	if fake.FindBundleThatProvidesStub != nil {
		return fake.FindBundleThatProvidesStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.findBundleThatProvidesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRegistryClientInterface) FindBundleThatProvidesCallCount() int {
	fake.findBundleThatProvidesMutex.RLock()
	defer fake.findBundleThatProvidesMutex.RUnlock()
	return len(fake.findBundleThatProvidesArgsForCall)
}

func (fake *FakeRegistryClientInterface) FindBundleThatProvidesCalls(stub func(context.Context, string, string, string, string) (*api.Bundle, error)) {
	fake.findBundleThatProvidesMutex.Lock()
	defer fake.findBundleThatProvidesMutex.Unlock()
	fake.FindBundleThatProvidesStub = stub
}

func (fake *FakeRegistryClientInterface) FindBundleThatProvidesArgsForCall(i int) (context.Context, string, string, string, string) {
	fake.findBundleThatProvidesMutex.RLock()
	defer fake.findBundleThatProvidesMutex.RUnlock()
	argsForCall := fake.findBundleThatProvidesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeRegistryClientInterface) FindBundleThatProvidesReturns(result1 *api.Bundle, result2 error) {
	fake.findBundleThatProvidesMutex.Lock()
	defer fake.findBundleThatProvidesMutex.Unlock()
	fake.FindBundleThatProvidesStub = nil
	fake.findBundleThatProvidesReturns = struct {
		result1 *api.Bundle
		result2 error
	}{result1, result2}
}

func (fake *FakeRegistryClientInterface) FindBundleThatProvidesReturnsOnCall(i int, result1 *api.Bundle, result2 error) {
	fake.findBundleThatProvidesMutex.Lock()
	defer fake.findBundleThatProvidesMutex.Unlock()
	fake.FindBundleThatProvidesStub = nil
	if fake.findBundleThatProvidesReturnsOnCall == nil {
		fake.findBundleThatProvidesReturnsOnCall = make(map[int]struct {
			result1 *api.Bundle
			result2 error
		})
	}
	fake.findBundleThatProvidesReturnsOnCall[i] = struct {
		result1 *api.Bundle
		result2 error
	}{result1, result2}
}

func (fake *FakeRegistryClientInterface) GetLatestChannelEntriesThatProvide(arg1 context.Context, arg2 string, arg3 string, arg4 string) (*registry.ChannelEntryIterator, error) {
	fake.getLatestChannelEntriesThatProvideMutex.Lock()
	ret, specificReturn := fake.getLatestChannelEntriesThatProvideReturnsOnCall[len(fake.getLatestChannelEntriesThatProvideArgsForCall)]
	fake.getLatestChannelEntriesThatProvideArgsForCall = append(fake.getLatestChannelEntriesThatProvideArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("GetLatestChannelEntriesThatProvide", []interface{}{arg1, arg2, arg3, arg4})
	fake.getLatestChannelEntriesThatProvideMutex.Unlock()
	if fake.GetLatestChannelEntriesThatProvideStub != nil {
		return fake.GetLatestChannelEntriesThatProvideStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getLatestChannelEntriesThatProvideReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRegistryClientInterface) GetLatestChannelEntriesThatProvideCallCount() int {
	fake.getLatestChannelEntriesThatProvideMutex.RLock()
	defer fake.getLatestChannelEntriesThatProvideMutex.RUnlock()
	return len(fake.getLatestChannelEntriesThatProvideArgsForCall)
}

func (fake *FakeRegistryClientInterface) GetLatestChannelEntriesThatProvideCalls(stub func(context.Context, string, string, string) (*registry.ChannelEntryIterator, error)) {
	fake.getLatestChannelEntriesThatProvideMutex.Lock()
	defer fake.getLatestChannelEntriesThatProvideMutex.Unlock()
	fake.GetLatestChannelEntriesThatProvideStub = stub
}

func (fake *FakeRegistryClientInterface) GetLatestChannelEntriesThatProvideArgsForCall(i int) (context.Context, string, string, string) {
	fake.getLatestChannelEntriesThatProvideMutex.RLock()
	defer fake.getLatestChannelEntriesThatProvideMutex.RUnlock()
	argsForCall := fake.getLatestChannelEntriesThatProvideArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeRegistryClientInterface) GetLatestChannelEntriesThatProvideReturns(result1 *registry.ChannelEntryIterator, result2 error) {
	fake.getLatestChannelEntriesThatProvideMutex.Lock()
	defer fake.getLatestChannelEntriesThatProvideMutex.Unlock()
	fake.GetLatestChannelEntriesThatProvideStub = nil
	fake.getLatestChannelEntriesThatProvideReturns = struct {
		result1 *registry.ChannelEntryIterator
		result2 error
	}{result1, result2}
}

func (fake *FakeRegistryClientInterface) GetLatestChannelEntriesThatProvideReturnsOnCall(i int, result1 *registry.ChannelEntryIterator, result2 error) {
	fake.getLatestChannelEntriesThatProvideMutex.Lock()
	defer fake.getLatestChannelEntriesThatProvideMutex.Unlock()
	fake.GetLatestChannelEntriesThatProvideStub = nil
	if fake.getLatestChannelEntriesThatProvideReturnsOnCall == nil {
		fake.getLatestChannelEntriesThatProvideReturnsOnCall = make(map[int]struct {
			result1 *registry.ChannelEntryIterator
			result2 error
		})
	}
	fake.getLatestChannelEntriesThatProvideReturnsOnCall[i] = struct {
		result1 *registry.ChannelEntryIterator
		result2 error
	}{result1, result2}
}

func (fake *FakeRegistryClientInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findBundleThatProvidesMutex.RLock()
	defer fake.findBundleThatProvidesMutex.RUnlock()
	fake.getLatestChannelEntriesThatProvideMutex.RLock()
	defer fake.getLatestChannelEntriesThatProvideMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRegistryClientInterface) recordInvocation(key string, args []interface{}) {
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

var _ registry.RegistryClientInterface = new(FakeRegistryClientInterface)
