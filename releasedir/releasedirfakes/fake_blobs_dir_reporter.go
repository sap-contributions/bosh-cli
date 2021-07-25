// Code generated by counterfeiter. DO NOT EDIT.
package releasedirfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/releasedir"
)

type FakeBlobsDirReporter struct {
	BlobDownloadFinishedStub        func(string, string, error)
	blobDownloadFinishedMutex       sync.RWMutex
	blobDownloadFinishedArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 error
	}
	BlobDownloadStartedStub        func(string, int64, string, string)
	blobDownloadStartedMutex       sync.RWMutex
	blobDownloadStartedArgsForCall []struct {
		arg1 string
		arg2 int64
		arg3 string
		arg4 string
	}
	BlobUploadFinishedStub        func(string, string, error)
	blobUploadFinishedMutex       sync.RWMutex
	blobUploadFinishedArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 error
	}
	BlobUploadStartedStub        func(string, int64, string)
	blobUploadStartedMutex       sync.RWMutex
	blobUploadStartedArgsForCall []struct {
		arg1 string
		arg2 int64
		arg3 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBlobsDirReporter) BlobDownloadFinished(arg1 string, arg2 string, arg3 error) {
	fake.blobDownloadFinishedMutex.Lock()
	fake.blobDownloadFinishedArgsForCall = append(fake.blobDownloadFinishedArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 error
	}{arg1, arg2, arg3})
	stub := fake.BlobDownloadFinishedStub
	fake.recordInvocation("BlobDownloadFinished", []interface{}{arg1, arg2, arg3})
	fake.blobDownloadFinishedMutex.Unlock()
	if stub != nil {
		fake.BlobDownloadFinishedStub(arg1, arg2, arg3)
	}
}

func (fake *FakeBlobsDirReporter) BlobDownloadFinishedCallCount() int {
	fake.blobDownloadFinishedMutex.RLock()
	defer fake.blobDownloadFinishedMutex.RUnlock()
	return len(fake.blobDownloadFinishedArgsForCall)
}

func (fake *FakeBlobsDirReporter) BlobDownloadFinishedCalls(stub func(string, string, error)) {
	fake.blobDownloadFinishedMutex.Lock()
	defer fake.blobDownloadFinishedMutex.Unlock()
	fake.BlobDownloadFinishedStub = stub
}

func (fake *FakeBlobsDirReporter) BlobDownloadFinishedArgsForCall(i int) (string, string, error) {
	fake.blobDownloadFinishedMutex.RLock()
	defer fake.blobDownloadFinishedMutex.RUnlock()
	argsForCall := fake.blobDownloadFinishedArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeBlobsDirReporter) BlobDownloadStarted(arg1 string, arg2 int64, arg3 string, arg4 string) {
	fake.blobDownloadStartedMutex.Lock()
	fake.blobDownloadStartedArgsForCall = append(fake.blobDownloadStartedArgsForCall, struct {
		arg1 string
		arg2 int64
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	stub := fake.BlobDownloadStartedStub
	fake.recordInvocation("BlobDownloadStarted", []interface{}{arg1, arg2, arg3, arg4})
	fake.blobDownloadStartedMutex.Unlock()
	if stub != nil {
		fake.BlobDownloadStartedStub(arg1, arg2, arg3, arg4)
	}
}

func (fake *FakeBlobsDirReporter) BlobDownloadStartedCallCount() int {
	fake.blobDownloadStartedMutex.RLock()
	defer fake.blobDownloadStartedMutex.RUnlock()
	return len(fake.blobDownloadStartedArgsForCall)
}

func (fake *FakeBlobsDirReporter) BlobDownloadStartedCalls(stub func(string, int64, string, string)) {
	fake.blobDownloadStartedMutex.Lock()
	defer fake.blobDownloadStartedMutex.Unlock()
	fake.BlobDownloadStartedStub = stub
}

func (fake *FakeBlobsDirReporter) BlobDownloadStartedArgsForCall(i int) (string, int64, string, string) {
	fake.blobDownloadStartedMutex.RLock()
	defer fake.blobDownloadStartedMutex.RUnlock()
	argsForCall := fake.blobDownloadStartedArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeBlobsDirReporter) BlobUploadFinished(arg1 string, arg2 string, arg3 error) {
	fake.blobUploadFinishedMutex.Lock()
	fake.blobUploadFinishedArgsForCall = append(fake.blobUploadFinishedArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 error
	}{arg1, arg2, arg3})
	stub := fake.BlobUploadFinishedStub
	fake.recordInvocation("BlobUploadFinished", []interface{}{arg1, arg2, arg3})
	fake.blobUploadFinishedMutex.Unlock()
	if stub != nil {
		fake.BlobUploadFinishedStub(arg1, arg2, arg3)
	}
}

func (fake *FakeBlobsDirReporter) BlobUploadFinishedCallCount() int {
	fake.blobUploadFinishedMutex.RLock()
	defer fake.blobUploadFinishedMutex.RUnlock()
	return len(fake.blobUploadFinishedArgsForCall)
}

func (fake *FakeBlobsDirReporter) BlobUploadFinishedCalls(stub func(string, string, error)) {
	fake.blobUploadFinishedMutex.Lock()
	defer fake.blobUploadFinishedMutex.Unlock()
	fake.BlobUploadFinishedStub = stub
}

func (fake *FakeBlobsDirReporter) BlobUploadFinishedArgsForCall(i int) (string, string, error) {
	fake.blobUploadFinishedMutex.RLock()
	defer fake.blobUploadFinishedMutex.RUnlock()
	argsForCall := fake.blobUploadFinishedArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeBlobsDirReporter) BlobUploadStarted(arg1 string, arg2 int64, arg3 string) {
	fake.blobUploadStartedMutex.Lock()
	fake.blobUploadStartedArgsForCall = append(fake.blobUploadStartedArgsForCall, struct {
		arg1 string
		arg2 int64
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.BlobUploadStartedStub
	fake.recordInvocation("BlobUploadStarted", []interface{}{arg1, arg2, arg3})
	fake.blobUploadStartedMutex.Unlock()
	if stub != nil {
		fake.BlobUploadStartedStub(arg1, arg2, arg3)
	}
}

func (fake *FakeBlobsDirReporter) BlobUploadStartedCallCount() int {
	fake.blobUploadStartedMutex.RLock()
	defer fake.blobUploadStartedMutex.RUnlock()
	return len(fake.blobUploadStartedArgsForCall)
}

func (fake *FakeBlobsDirReporter) BlobUploadStartedCalls(stub func(string, int64, string)) {
	fake.blobUploadStartedMutex.Lock()
	defer fake.blobUploadStartedMutex.Unlock()
	fake.BlobUploadStartedStub = stub
}

func (fake *FakeBlobsDirReporter) BlobUploadStartedArgsForCall(i int) (string, int64, string) {
	fake.blobUploadStartedMutex.RLock()
	defer fake.blobUploadStartedMutex.RUnlock()
	argsForCall := fake.blobUploadStartedArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeBlobsDirReporter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.blobDownloadFinishedMutex.RLock()
	defer fake.blobDownloadFinishedMutex.RUnlock()
	fake.blobDownloadStartedMutex.RLock()
	defer fake.blobDownloadStartedMutex.RUnlock()
	fake.blobUploadFinishedMutex.RLock()
	defer fake.blobUploadFinishedMutex.RUnlock()
	fake.blobUploadStartedMutex.RLock()
	defer fake.blobUploadStartedMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBlobsDirReporter) recordInvocation(key string, args []interface{}) {
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

var _ releasedir.BlobsDirReporter = new(FakeBlobsDirReporter)
