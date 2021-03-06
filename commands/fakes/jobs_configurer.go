// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type JobsConfigurer struct {
	JobsStub        func(productGUID string) (map[string]string, error)
	jobsMutex       sync.RWMutex
	jobsArgsForCall []struct {
		productGUID string
	}
	jobsReturns struct {
		result1 map[string]string
		result2 error
	}
	GetExistingJobConfigStub        func(productGUID, jobGUID string) (api.JobProperties, error)
	getExistingJobConfigMutex       sync.RWMutex
	getExistingJobConfigArgsForCall []struct {
		productGUID string
		jobGUID     string
	}
	getExistingJobConfigReturns struct {
		result1 api.JobProperties
		result2 error
	}
	ConfigureJobStub        func(productGUID, jobGUID string, jobProperties api.JobProperties) error
	configureJobMutex       sync.RWMutex
	configureJobArgsForCall []struct {
		productGUID   string
		jobGUID       string
		jobProperties api.JobProperties
	}
	configureJobReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *JobsConfigurer) Jobs(productGUID string) (map[string]string, error) {
	fake.jobsMutex.Lock()
	fake.jobsArgsForCall = append(fake.jobsArgsForCall, struct {
		productGUID string
	}{productGUID})
	fake.recordInvocation("Jobs", []interface{}{productGUID})
	fake.jobsMutex.Unlock()
	if fake.JobsStub != nil {
		return fake.JobsStub(productGUID)
	} else {
		return fake.jobsReturns.result1, fake.jobsReturns.result2
	}
}

func (fake *JobsConfigurer) JobsCallCount() int {
	fake.jobsMutex.RLock()
	defer fake.jobsMutex.RUnlock()
	return len(fake.jobsArgsForCall)
}

func (fake *JobsConfigurer) JobsArgsForCall(i int) string {
	fake.jobsMutex.RLock()
	defer fake.jobsMutex.RUnlock()
	return fake.jobsArgsForCall[i].productGUID
}

func (fake *JobsConfigurer) JobsReturns(result1 map[string]string, result2 error) {
	fake.JobsStub = nil
	fake.jobsReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *JobsConfigurer) GetExistingJobConfig(productGUID string, jobGUID string) (api.JobProperties, error) {
	fake.getExistingJobConfigMutex.Lock()
	fake.getExistingJobConfigArgsForCall = append(fake.getExistingJobConfigArgsForCall, struct {
		productGUID string
		jobGUID     string
	}{productGUID, jobGUID})
	fake.recordInvocation("GetExistingJobConfig", []interface{}{productGUID, jobGUID})
	fake.getExistingJobConfigMutex.Unlock()
	if fake.GetExistingJobConfigStub != nil {
		return fake.GetExistingJobConfigStub(productGUID, jobGUID)
	} else {
		return fake.getExistingJobConfigReturns.result1, fake.getExistingJobConfigReturns.result2
	}
}

func (fake *JobsConfigurer) GetExistingJobConfigCallCount() int {
	fake.getExistingJobConfigMutex.RLock()
	defer fake.getExistingJobConfigMutex.RUnlock()
	return len(fake.getExistingJobConfigArgsForCall)
}

func (fake *JobsConfigurer) GetExistingJobConfigArgsForCall(i int) (string, string) {
	fake.getExistingJobConfigMutex.RLock()
	defer fake.getExistingJobConfigMutex.RUnlock()
	return fake.getExistingJobConfigArgsForCall[i].productGUID, fake.getExistingJobConfigArgsForCall[i].jobGUID
}

func (fake *JobsConfigurer) GetExistingJobConfigReturns(result1 api.JobProperties, result2 error) {
	fake.GetExistingJobConfigStub = nil
	fake.getExistingJobConfigReturns = struct {
		result1 api.JobProperties
		result2 error
	}{result1, result2}
}

func (fake *JobsConfigurer) ConfigureJob(productGUID string, jobGUID string, jobProperties api.JobProperties) error {
	fake.configureJobMutex.Lock()
	fake.configureJobArgsForCall = append(fake.configureJobArgsForCall, struct {
		productGUID   string
		jobGUID       string
		jobProperties api.JobProperties
	}{productGUID, jobGUID, jobProperties})
	fake.recordInvocation("ConfigureJob", []interface{}{productGUID, jobGUID, jobProperties})
	fake.configureJobMutex.Unlock()
	if fake.ConfigureJobStub != nil {
		return fake.ConfigureJobStub(productGUID, jobGUID, jobProperties)
	} else {
		return fake.configureJobReturns.result1
	}
}

func (fake *JobsConfigurer) ConfigureJobCallCount() int {
	fake.configureJobMutex.RLock()
	defer fake.configureJobMutex.RUnlock()
	return len(fake.configureJobArgsForCall)
}

func (fake *JobsConfigurer) ConfigureJobArgsForCall(i int) (string, string, api.JobProperties) {
	fake.configureJobMutex.RLock()
	defer fake.configureJobMutex.RUnlock()
	return fake.configureJobArgsForCall[i].productGUID, fake.configureJobArgsForCall[i].jobGUID, fake.configureJobArgsForCall[i].jobProperties
}

func (fake *JobsConfigurer) ConfigureJobReturns(result1 error) {
	fake.ConfigureJobStub = nil
	fake.configureJobReturns = struct {
		result1 error
	}{result1}
}

func (fake *JobsConfigurer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.jobsMutex.RLock()
	defer fake.jobsMutex.RUnlock()
	fake.getExistingJobConfigMutex.RLock()
	defer fake.getExistingJobConfigMutex.RUnlock()
	fake.configureJobMutex.RLock()
	defer fake.configureJobMutex.RUnlock()
	return fake.invocations
}

func (fake *JobsConfigurer) recordInvocation(key string, args []interface{}) {
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
