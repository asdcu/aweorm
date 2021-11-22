package dialect

import (
	"reflect"
	"sync"
)

var dialectsMap = map[string]Dialect{}
var rw sync.RWMutex

type Dialect interface {
	DateTypeOf(typ reflect.Value) string
	TableExistSQL(tableName string) (string, []interface{})
}

func RegisterDialect(name string, dialect Dialect) {
	rw.Lock()
	defer rw.Unlock()
	dialectsMap[name] = dialect
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}
