// Copyright (C) 2017 go-nebulas authors
//
// This file is part of the go-nebulas library.
//
// the go-nebulas library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-nebulas library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-nebulas library.  If not, see <http://www.gnu.org/licenses/>.
//

package nvm

/*
#include <stdlib.h>
#cgo CFLAGS:
#cgo LDFLAGS: -L${SRCDIR}/native-lib -lv8engine

#include "v8/engine.h"

// Forward declaration.
void GoLogFunc_cgo(int level, const char *msg);
char *StorageGetFunc_cgo(void *handler, const char *key);
int StoragePutFunc_cgo(void *handler, const char *key, const char *value);
int StorageDelFunc_cgo(void *handler, const char *key);
*/
import "C"
import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"sync"
	"unsafe"

	"github.com/nebulasio/go-nebulas/common/trie"
	log "github.com/sirupsen/logrus"
)

var (
	ErrExecutionFailed     = errors.New("execute source failed")
	ErrInvalidFunctionName = errors.New("invalid function name")

	v8engineOnce = sync.Once{}
	storages     = make(map[uint64]*V8Engine, 256)
	storagesIdx  = uint64(0)
	storagesLock = sync.RWMutex{}
	v8engine     *C.V8Engine

	functionNameRe = regexp.MustCompile("^[a-zA-Z_]+$")
)

// V8Engine v8 engine.
type V8Engine struct {
	balanceStorage        *trie.BatchTrie
	localContractStorage  *trie.BatchTrie
	globalContractStorage *trie.BatchTrie
	lcsHandler            uint64
	gcsHandler            uint64
}

// InitV8Engine initialize the v8 engine.
func InitV8Engine() {
	C.Initialize()
	C.InitializeLogger((C.LogFunc)(unsafe.Pointer(C.GoLogFunc_cgo)))
	C.InitializeStorage((C.StorageGetFunc)(unsafe.Pointer(C.StorageGetFunc_cgo)), (C.StoragePutFunc)(unsafe.Pointer(C.StoragePutFunc_cgo)), (C.StorageDelFunc)(unsafe.Pointer(C.StorageDelFunc_cgo)))
}

// DisposeV8Engine dispose the v8 engine.
func DisposeV8Engine() {
	C.Dispose()
}

// NewV8Engine return new V8Engine instance.
func NewV8Engine(balanceStorage, localContractStorage, globalContractStorage *trie.BatchTrie) *V8Engine {
	v8engineOnce.Do(func() {
		InitV8Engine()
		v8engine = C.CreateEngine()
	})

	engine := &V8Engine{
		balanceStorage:        balanceStorage,
		localContractStorage:  localContractStorage,
		globalContractStorage: globalContractStorage,
	}

	storagesLock.Lock()
	defer storagesLock.Unlock()

	storagesIdx++
	engine.lcsHandler = storagesIdx
	storagesIdx++
	engine.gcsHandler = storagesIdx

	storages[engine.lcsHandler] = engine
	storages[engine.gcsHandler] = engine

	return engine
}

// Dispose dispose all resources.
func (e *V8Engine) Dispose() {
	storagesLock.Lock()
	delete(storages, e.lcsHandler)
	delete(storages, e.gcsHandler)
	storagesLock.Unlock()
}

// RunScriptSource
func (e *V8Engine) RunScriptSource(content string) error {
	data := C.CString(content)
	defer C.free(unsafe.Pointer(data))

	ret := C.RunScriptSource2(v8engine, data, C.uintptr_t(e.lcsHandler),
		C.uintptr_t(e.gcsHandler))

	if ret != 0 {
		return ErrExecutionFailed
	}
	return nil
}

// Call
func (e *V8Engine) Call(source, function, args string) error {
	if functionNameRe.MatchString(function) == false || strings.Compare("init", function) == 0 {
		return ErrInvalidFunctionName
	}
	return e.executeScript(source, function, args)
}

// DeployAndInit
func (e *V8Engine) DeployAndInit(source, args string) error {
	return e.executeScript(source, "init", args)
}

// Execute execute the script and return error.
func (e *V8Engine) executeScript(source, function, args string) error {
	executablesource := prepareExecutableSource(source, function, args)

	// log.WithFields(log.Fields{
	// 	"source":           source,
	// 	"args":             args,
	// 	"function":         function,
	// 	"executablesource": executablesource,
	// }).Info("executeScript")

	return e.RunScriptSource(executablesource)
}

func getEngineAndStorage(handler uint64) (*V8Engine, *trie.BatchTrie) {
	storagesLock.RLock()
	engine := storages[handler]
	storagesLock.RUnlock()

	if engine == nil {
		log.WithFields(log.Fields{
			"func":          "nvm.getEngineAndStorage",
			"wantedHandler": handler,
		}).Error("wantedHandler is not found.")
		return nil, nil
	}

	if engine.lcsHandler == handler {
		return engine, engine.localContractStorage
	} else if engine.gcsHandler == handler {
		return engine, engine.globalContractStorage
	} else {
		log.WithFields(log.Fields{
			"func":          "nvm.getEngineAndStorage",
			"lcsHandler":    engine.lcsHandler,
			"gcsHandler":    engine.gcsHandler,
			"wantedHandler": handler,
		}).Error("in-consistent storage handler.")
		return nil, nil
	}
}

func formatArgs(args string) string {
	return strings.Replace(args, "\"", "\\\"", -1)
}

func prepareExecutableSource(source, function, args string) string {
	cSource := C.CString(source)
	defer C.free(unsafe.Pointer(cSource))

	cmSource := C.encapsulateSourceToModuleStyle(cSource)
	defer C.free(unsafe.Pointer(cmSource))

	var executablesource string
	if len(args) > 0 {
		executablesource = fmt.Sprintf("var __contract = %s; var __instance = new __contract();__instance[\"%s\"].apply(__instance, JSON.parse(\"%s\"));", C.GoString(cmSource), function, formatArgs(args))
	} else {
		executablesource = fmt.Sprintf("var __contract = %s; var __instance = new __contract();__instance[\"%s\"].apply(__instance);", C.GoString(cmSource), function)
	}

	return executablesource
}