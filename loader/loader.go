// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package loader

import (
	"fmt"
	"github.com/sofiworker/gonm/dbusutil"
	"github.com/sofiworker/gonm/logger"
	"sync"
	"time"
)

type EnableFlag int

const (
	EnableFlagNone EnableFlag = 1 << iota
	EnableFlagIgnoreMissingModule
	EnableFlagForceStart
)

func (flags EnableFlag) HasFlag(flag EnableFlag) bool {
	return flags&flag != 0
}

const (
	ErrorNoDependencies int = iota
	ErrorCircleDependencies
	ErrorMissingModule
	ErrorInternalError
	ErrorConflict
)

type EnableError struct {
	ModuleName string
	Code       int
	detail     string
}

func (e *EnableError) Error() string {
	switch e.Code {
	case ErrorNoDependencies:
		return fmt.Sprintf("%s's dependencies is not meet, %s is need", e.ModuleName, e.detail)
	case ErrorCircleDependencies:
		return "dependency circle"
		// return fmt.Sprintf("%s and %s dependency each other.", e.ModuleName, e.detail)
	case ErrorMissingModule:
		return fmt.Sprintf("%s is missing", e.ModuleName)
	case ErrorInternalError:
		return fmt.Sprintf("%s started failed: %s", e.ModuleName, e.detail)
	case ErrorConflict:
		return fmt.Sprintf("tring to enable disabled module(%s)", e.ModuleName)
	}
	panic("EnableError: Unknown Error, Should not be reached")
}

type Loader struct {
	modules Modules
	lock    sync.Mutex
	service *dbusutil.Service
}

func (l *Loader) AddModule(m Module) {
	l.lock.Lock()
	defer l.lock.Unlock()
	name := m.Name()
	_, exist := l.modules[name]
	if exist {
		logger.Debug("Register", name, "is already registered")
		return
	}
	logger.Debug("Register module:", name)
	l.modules[name] = m
}

func (l *Loader) DeleteModule(name string) {
	l.lock.Lock()
	defer l.lock.Unlock()
	delete(l.modules, name)
}

func (l *Loader) List() []Module {
	l.lock.Lock()
	defer l.lock.Unlock()
	modules := make([]Module, 0, len(l.modules))
	for _, m := range l.modules {
		modules = append(modules, m)
	}
	return modules
}

func (l *Loader) GetModule(name string) Module {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.modules[name]
}

func (l *Loader) WaitDependencies(module Module) {
	for _, dependencyName := range module.GetDependencies() {
		l.modules[dependencyName].WaitEnable()
	}
}

func (l *Loader) EnableModules(enablingModules []string, disableModules []string, flag EnableFlag) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	// build a dag
	startTime := time.Now()
	builder := NewDAGBuilder(l, enablingModules, disableModules, flag)
	dag, err := builder.Execute()
	if err != nil {
		return err
	}
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	logger.SInfof("build dag done, cost %s", duration)

	// perform a topo sort
	nodes, ok := dag.TopologicalDag()
	if !ok {
		return &EnableError{Code: ErrorCircleDependencies}
	}
	endTime = time.Now()
	duration = endTime.Sub(startTime)
	logger.SInfof("topo sort done, cost add up to %s", duration)

	// enable modules
	for _, node := range nodes {
		if node == nil {
			continue
		}
		module := l.modules[node.ID]
		name := node.ID

		go func() {
			logger.SInfo("enable module", name)
			startTime := time.Now()

			// wait for its dependency
			l.WaitDependencies(module)
			endTime := time.Now()
			duration := endTime.Sub(startTime)
			logger.SInfo("module", name, "wait done, cost", duration)

			err := module.Enable(true)
			endTime = time.Now()
			duration = endTime.Sub(startTime)
			if err != nil {
				logger.SFatalf("enable module %s failed: %s, cost %s", name, err, duration)
			} else {
				logger.SInfo("enable module %s done cost %s", name, duration)
			}
		}()
	}

	for _, n := range nodes {
		m := l.modules[n.ID]
		m.WaitEnable()
	}

	endTime = time.Now()
	duration = endTime.Sub(startTime)
	logger.SInfof("enable modules done, cost add up to %s", duration)
	return nil
}
