package supercereal

import (
	"sync"
)

var streamPool = sync.Pool{
	New: func() interface{} {
		return NewJSONStream()
	},
}

// todo: cannot return it like this, needs a callback
func Marshal(value interface{}) []byte {
	var js *JSONStream = streamPool.Get().(*JSONStream)




	js.reset()
	js.routeValueType(value)
	var ret []byte = js.end()



	streamPool.Put(js)
	return ret
}