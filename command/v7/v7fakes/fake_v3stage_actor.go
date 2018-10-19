// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	sync "sync"

	v7action "code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeV3StageActor struct {
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct {
	}
	cloudControllerAPIVersionReturns struct {
		result1 string
	}
	cloudControllerAPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	GetStreamingLogsForApplicationByNameAndSpaceStub        func(string, string, v7action.NOAAClient) (<-chan *v7action.LogMessage, <-chan error, v7action.Warnings, error)
	getStreamingLogsForApplicationByNameAndSpaceMutex       sync.RWMutex
	getStreamingLogsForApplicationByNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 v7action.NOAAClient
	}
	getStreamingLogsForApplicationByNameAndSpaceReturns struct {
		result1 <-chan *v7action.LogMessage
		result2 <-chan error
		result3 v7action.Warnings
		result4 error
	}
	getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall map[int]struct {
		result1 <-chan *v7action.LogMessage
		result2 <-chan error
		result3 v7action.Warnings
		result4 error
	}
	StagePackageStub        func(string, string) (<-chan v7action.Droplet, <-chan v7action.Warnings, <-chan error)
	stagePackageMutex       sync.RWMutex
	stagePackageArgsForCall []struct {
		arg1 string
		arg2 string
	}
	stagePackageReturns struct {
		result1 <-chan v7action.Droplet
		result2 <-chan v7action.Warnings
		result3 <-chan error
	}
	stagePackageReturnsOnCall map[int]struct {
		result1 <-chan v7action.Droplet
		result2 <-chan v7action.Warnings
		result3 <-chan error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV3StageActor) CloudControllerAPIVersion() string {
	fake.cloudControllerAPIVersionMutex.Lock()
	ret, specificReturn := fake.cloudControllerAPIVersionReturnsOnCall[len(fake.cloudControllerAPIVersionArgsForCall)]
	fake.cloudControllerAPIVersionArgsForCall = append(fake.cloudControllerAPIVersionArgsForCall, struct {
	}{})
	fake.recordInvocation("CloudControllerAPIVersion", []interface{}{})
	fake.cloudControllerAPIVersionMutex.Unlock()
	if fake.CloudControllerAPIVersionStub != nil {
		return fake.CloudControllerAPIVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.cloudControllerAPIVersionReturns
	return fakeReturns.result1
}

func (fake *FakeV3StageActor) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeV3StageActor) CloudControllerAPIVersionCalls(stub func() string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = stub
}

func (fake *FakeV3StageActor) CloudControllerAPIVersionReturns(result1 string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeV3StageActor) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = nil
	if fake.cloudControllerAPIVersionReturnsOnCall == nil {
		fake.cloudControllerAPIVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cloudControllerAPIVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeV3StageActor) GetStreamingLogsForApplicationByNameAndSpace(arg1 string, arg2 string, arg3 v7action.NOAAClient) (<-chan *v7action.LogMessage, <-chan error, v7action.Warnings, error) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall[len(fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall)]
	fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall = append(fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 v7action.NOAAClient
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetStreamingLogsForApplicationByNameAndSpace", []interface{}{arg1, arg2, arg3})
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Unlock()
	if fake.GetStreamingLogsForApplicationByNameAndSpaceStub != nil {
		return fake.GetStreamingLogsForApplicationByNameAndSpaceStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	fakeReturns := fake.getStreamingLogsForApplicationByNameAndSpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3, fakeReturns.result4
}

func (fake *FakeV3StageActor) GetStreamingLogsForApplicationByNameAndSpaceCallCount() int {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RLock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeV3StageActor) GetStreamingLogsForApplicationByNameAndSpaceCalls(stub func(string, string, v7action.NOAAClient) (<-chan *v7action.LogMessage, <-chan error, v7action.Warnings, error)) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Lock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Unlock()
	fake.GetStreamingLogsForApplicationByNameAndSpaceStub = stub
}

func (fake *FakeV3StageActor) GetStreamingLogsForApplicationByNameAndSpaceArgsForCall(i int) (string, string, v7action.NOAAClient) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RLock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RUnlock()
	argsForCall := fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeV3StageActor) GetStreamingLogsForApplicationByNameAndSpaceReturns(result1 <-chan *v7action.LogMessage, result2 <-chan error, result3 v7action.Warnings, result4 error) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Lock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Unlock()
	fake.GetStreamingLogsForApplicationByNameAndSpaceStub = nil
	fake.getStreamingLogsForApplicationByNameAndSpaceReturns = struct {
		result1 <-chan *v7action.LogMessage
		result2 <-chan error
		result3 v7action.Warnings
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeV3StageActor) GetStreamingLogsForApplicationByNameAndSpaceReturnsOnCall(i int, result1 <-chan *v7action.LogMessage, result2 <-chan error, result3 v7action.Warnings, result4 error) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Lock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Unlock()
	fake.GetStreamingLogsForApplicationByNameAndSpaceStub = nil
	if fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall == nil {
		fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 <-chan *v7action.LogMessage
			result2 <-chan error
			result3 v7action.Warnings
			result4 error
		})
	}
	fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall[i] = struct {
		result1 <-chan *v7action.LogMessage
		result2 <-chan error
		result3 v7action.Warnings
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeV3StageActor) StagePackage(arg1 string, arg2 string) (<-chan v7action.Droplet, <-chan v7action.Warnings, <-chan error) {
	fake.stagePackageMutex.Lock()
	ret, specificReturn := fake.stagePackageReturnsOnCall[len(fake.stagePackageArgsForCall)]
	fake.stagePackageArgsForCall = append(fake.stagePackageArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("StagePackage", []interface{}{arg1, arg2})
	fake.stagePackageMutex.Unlock()
	if fake.StagePackageStub != nil {
		return fake.StagePackageStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.stagePackageReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeV3StageActor) StagePackageCallCount() int {
	fake.stagePackageMutex.RLock()
	defer fake.stagePackageMutex.RUnlock()
	return len(fake.stagePackageArgsForCall)
}

func (fake *FakeV3StageActor) StagePackageCalls(stub func(string, string) (<-chan v7action.Droplet, <-chan v7action.Warnings, <-chan error)) {
	fake.stagePackageMutex.Lock()
	defer fake.stagePackageMutex.Unlock()
	fake.StagePackageStub = stub
}

func (fake *FakeV3StageActor) StagePackageArgsForCall(i int) (string, string) {
	fake.stagePackageMutex.RLock()
	defer fake.stagePackageMutex.RUnlock()
	argsForCall := fake.stagePackageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeV3StageActor) StagePackageReturns(result1 <-chan v7action.Droplet, result2 <-chan v7action.Warnings, result3 <-chan error) {
	fake.stagePackageMutex.Lock()
	defer fake.stagePackageMutex.Unlock()
	fake.StagePackageStub = nil
	fake.stagePackageReturns = struct {
		result1 <-chan v7action.Droplet
		result2 <-chan v7action.Warnings
		result3 <-chan error
	}{result1, result2, result3}
}

func (fake *FakeV3StageActor) StagePackageReturnsOnCall(i int, result1 <-chan v7action.Droplet, result2 <-chan v7action.Warnings, result3 <-chan error) {
	fake.stagePackageMutex.Lock()
	defer fake.stagePackageMutex.Unlock()
	fake.StagePackageStub = nil
	if fake.stagePackageReturnsOnCall == nil {
		fake.stagePackageReturnsOnCall = make(map[int]struct {
			result1 <-chan v7action.Droplet
			result2 <-chan v7action.Warnings
			result3 <-chan error
		})
	}
	fake.stagePackageReturnsOnCall[i] = struct {
		result1 <-chan v7action.Droplet
		result2 <-chan v7action.Warnings
		result3 <-chan error
	}{result1, result2, result3}
}

func (fake *FakeV3StageActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RLock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RUnlock()
	fake.stagePackageMutex.RLock()
	defer fake.stagePackageMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV3StageActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.V3StageActor = new(FakeV3StageActor)
