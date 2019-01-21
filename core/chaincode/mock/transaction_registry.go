
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//伪造者生成的代码。不要编辑。
package mock

import (
	sync "sync"
)

type TransactionRegistry struct {
	AddStub        func(string, string) bool
	addMutex       sync.RWMutex
	addArgsForCall []struct {
		arg1 string
		arg2 string
	}
	addReturns struct {
		result1 bool
	}
	addReturnsOnCall map[int]struct {
		result1 bool
	}
	RemoveStub        func(string, string)
	removeMutex       sync.RWMutex
	removeArgsForCall []struct {
		arg1 string
		arg2 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *TransactionRegistry) Add(arg1 string, arg2 string) bool {
	fake.addMutex.Lock()
	ret, specificReturn := fake.addReturnsOnCall[len(fake.addArgsForCall)]
	fake.addArgsForCall = append(fake.addArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Add", []interface{}{arg1, arg2})
	fake.addMutex.Unlock()
	if fake.AddStub != nil {
		return fake.AddStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.addReturns
	return fakeReturns.result1
}

func (fake *TransactionRegistry) AddCallCount() int {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return len(fake.addArgsForCall)
}

func (fake *TransactionRegistry) AddCalls(stub func(string, string) bool) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = stub
}

func (fake *TransactionRegistry) AddArgsForCall(i int) (string, string) {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	argsForCall := fake.addArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *TransactionRegistry) AddReturns(result1 bool) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = nil
	fake.addReturns = struct {
		result1 bool
	}{result1}
}

func (fake *TransactionRegistry) AddReturnsOnCall(i int, result1 bool) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = nil
	if fake.addReturnsOnCall == nil {
		fake.addReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.addReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *TransactionRegistry) Remove(arg1 string, arg2 string) {
	fake.removeMutex.Lock()
	fake.removeArgsForCall = append(fake.removeArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Remove", []interface{}{arg1, arg2})
	fake.removeMutex.Unlock()
	if fake.RemoveStub != nil {
		fake.RemoveStub(arg1, arg2)
	}
}

func (fake *TransactionRegistry) RemoveCallCount() int {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return len(fake.removeArgsForCall)
}

func (fake *TransactionRegistry) RemoveCalls(stub func(string, string)) {
	fake.removeMutex.Lock()
	defer fake.removeMutex.Unlock()
	fake.RemoveStub = stub
}

func (fake *TransactionRegistry) RemoveArgsForCall(i int) (string, string) {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	argsForCall := fake.removeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *TransactionRegistry) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *TransactionRegistry) recordInvocation(key string, args []interface{}) {
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