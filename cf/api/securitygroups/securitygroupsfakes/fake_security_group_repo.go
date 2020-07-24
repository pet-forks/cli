// Code generated by counterfeiter. DO NOT EDIT.
package securitygroupsfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/cf/api/securitygroups"
	"code.cloudfoundry.org/cli/cf/models"
)

type FakeSecurityGroupRepo struct {
	CreateStub        func(string, []map[string]interface{}) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 string
		arg2 []map[string]interface{}
	}
	createReturns struct {
		result1 error
	}
	createReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStub        func(string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 string
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	FindAllStub        func() ([]models.SecurityGroup, error)
	findAllMutex       sync.RWMutex
	findAllArgsForCall []struct {
	}
	findAllReturns struct {
		result1 []models.SecurityGroup
		result2 error
	}
	findAllReturnsOnCall map[int]struct {
		result1 []models.SecurityGroup
		result2 error
	}
	ReadStub        func(string) (models.SecurityGroup, error)
	readMutex       sync.RWMutex
	readArgsForCall []struct {
		arg1 string
	}
	readReturns struct {
		result1 models.SecurityGroup
		result2 error
	}
	readReturnsOnCall map[int]struct {
		result1 models.SecurityGroup
		result2 error
	}
	UpdateStub        func(string, []map[string]interface{}) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 string
		arg2 []map[string]interface{}
	}
	updateReturns struct {
		result1 error
	}
	updateReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSecurityGroupRepo) Create(arg1 string, arg2 []map[string]interface{}) error {
	var arg2Copy []map[string]interface{}
	if arg2 != nil {
		arg2Copy = make([]map[string]interface{}, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 string
		arg2 []map[string]interface{}
	}{arg1, arg2Copy})
	fake.recordInvocation("Create", []interface{}{arg1, arg2Copy})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createReturns
	return fakeReturns.result1
}

func (fake *FakeSecurityGroupRepo) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeSecurityGroupRepo) CreateCalls(stub func(string, []map[string]interface{}) error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeSecurityGroupRepo) CreateArgsForCall(i int) (string, []map[string]interface{}) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSecurityGroupRepo) CreateReturns(result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecurityGroupRepo) CreateReturnsOnCall(i int, result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecurityGroupRepo) Delete(arg1 string) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Delete", []interface{}{arg1})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteReturns
	return fakeReturns.result1
}

func (fake *FakeSecurityGroupRepo) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeSecurityGroupRepo) DeleteCalls(stub func(string) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeSecurityGroupRepo) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSecurityGroupRepo) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecurityGroupRepo) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecurityGroupRepo) FindAll() ([]models.SecurityGroup, error) {
	fake.findAllMutex.Lock()
	ret, specificReturn := fake.findAllReturnsOnCall[len(fake.findAllArgsForCall)]
	fake.findAllArgsForCall = append(fake.findAllArgsForCall, struct {
	}{})
	fake.recordInvocation("FindAll", []interface{}{})
	fake.findAllMutex.Unlock()
	if fake.FindAllStub != nil {
		return fake.FindAllStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.findAllReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSecurityGroupRepo) FindAllCallCount() int {
	fake.findAllMutex.RLock()
	defer fake.findAllMutex.RUnlock()
	return len(fake.findAllArgsForCall)
}

func (fake *FakeSecurityGroupRepo) FindAllCalls(stub func() ([]models.SecurityGroup, error)) {
	fake.findAllMutex.Lock()
	defer fake.findAllMutex.Unlock()
	fake.FindAllStub = stub
}

func (fake *FakeSecurityGroupRepo) FindAllReturns(result1 []models.SecurityGroup, result2 error) {
	fake.findAllMutex.Lock()
	defer fake.findAllMutex.Unlock()
	fake.FindAllStub = nil
	fake.findAllReturns = struct {
		result1 []models.SecurityGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeSecurityGroupRepo) FindAllReturnsOnCall(i int, result1 []models.SecurityGroup, result2 error) {
	fake.findAllMutex.Lock()
	defer fake.findAllMutex.Unlock()
	fake.FindAllStub = nil
	if fake.findAllReturnsOnCall == nil {
		fake.findAllReturnsOnCall = make(map[int]struct {
			result1 []models.SecurityGroup
			result2 error
		})
	}
	fake.findAllReturnsOnCall[i] = struct {
		result1 []models.SecurityGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeSecurityGroupRepo) Read(arg1 string) (models.SecurityGroup, error) {
	fake.readMutex.Lock()
	ret, specificReturn := fake.readReturnsOnCall[len(fake.readArgsForCall)]
	fake.readArgsForCall = append(fake.readArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Read", []interface{}{arg1})
	fake.readMutex.Unlock()
	if fake.ReadStub != nil {
		return fake.ReadStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.readReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSecurityGroupRepo) ReadCallCount() int {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return len(fake.readArgsForCall)
}

func (fake *FakeSecurityGroupRepo) ReadCalls(stub func(string) (models.SecurityGroup, error)) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = stub
}

func (fake *FakeSecurityGroupRepo) ReadArgsForCall(i int) string {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	argsForCall := fake.readArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSecurityGroupRepo) ReadReturns(result1 models.SecurityGroup, result2 error) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = nil
	fake.readReturns = struct {
		result1 models.SecurityGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeSecurityGroupRepo) ReadReturnsOnCall(i int, result1 models.SecurityGroup, result2 error) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = nil
	if fake.readReturnsOnCall == nil {
		fake.readReturnsOnCall = make(map[int]struct {
			result1 models.SecurityGroup
			result2 error
		})
	}
	fake.readReturnsOnCall[i] = struct {
		result1 models.SecurityGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeSecurityGroupRepo) Update(arg1 string, arg2 []map[string]interface{}) error {
	var arg2Copy []map[string]interface{}
	if arg2 != nil {
		arg2Copy = make([]map[string]interface{}, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 string
		arg2 []map[string]interface{}
	}{arg1, arg2Copy})
	fake.recordInvocation("Update", []interface{}{arg1, arg2Copy})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.updateReturns
	return fakeReturns.result1
}

func (fake *FakeSecurityGroupRepo) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeSecurityGroupRepo) UpdateCalls(stub func(string, []map[string]interface{}) error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *FakeSecurityGroupRepo) UpdateArgsForCall(i int) (string, []map[string]interface{}) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSecurityGroupRepo) UpdateReturns(result1 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecurityGroupRepo) UpdateReturnsOnCall(i int, result1 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecurityGroupRepo) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.findAllMutex.RLock()
	defer fake.findAllMutex.RUnlock()
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSecurityGroupRepo) recordInvocation(key string, args []interface{}) {
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

var _ securitygroups.SecurityGroupRepo = new(FakeSecurityGroupRepo)
