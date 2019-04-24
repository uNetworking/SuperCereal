package supercereal

import (
	"sync"
)

var streamPool = sync.Pool{
	New: func() interface{} {
		return NewJSONStream()
	},
}

func Marshal(value interface{}) []byte {
	var ret []byte
	var js *JSONStream = streamPool.Get().(*JSONStream)

	js.OnJSON(func(json []byte) {
		ret = json
	})
	js.Serialize(value)

	streamPool.Put(js)
	return ret
}