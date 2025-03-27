package functrace

import (
	"log"
	"os"
	"runtime"
	"sync"
	"time"

	"istio.io/istio/pkg/util/sets"
)

func init() {
	log.Default().SetFlags(log.Lmicroseconds | log.Ltime | log.Ldate)
	log.Default().SetOutput(os.Stdout)
}

var (
	_ sync.Mutex
	_ = make(map[uint64]int)
)

var (
	fs    = sets.New[string]()
	start = time.Now()
	lock  sync.Mutex
)

var _ = sync.Once{}

func Trace() func() {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("not found caller")
	}

	// id := getGID()
	fn := runtime.FuncForPC(pc)
	name := fn.Name()
	lock.Lock()
	fs.Insert(name)
	lock.Unlock()
	if time.Since(start) > time.Minute {
		//o.Do(func() {
		//	marshal, _ := json.Marshal(fs)
		//	ioutil.WriteFile("trace.json", marshal, 0644)
		//})
	}
	//started := time.Now()
	//
	//mu.Lock()
	//v := m[id]
	//m[id] = v + 1
	//mu.Unlock()
	//printTraceEntry(id, name, "->", v+1)
	return func() {
		// mu.Lock()
		// v := m[id]
		// m[id] = v - 1
		// mu.Unlock()
		// printTraceExit(id, name, "<-", v, started)
	}
}
