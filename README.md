<div align="center"><img src="al.jpg" /></div>

SuperCereal is a simple and efficient JSON serializing library for Go. Unlike many other "fast" serializers, it doesn't make use of intermediate (hash) maps or "documents". Instead it immediately dumps to a pre-allocated buffer, whatever data you add as you add it. JSON is emitted in chunks of data as the pre-allocated buffer fills up.

### Overview
```go
	js := supercereal.NewJSONStream(func(data []byte) {
		// this is where you get data (potentially in chunks)
		// it is up to you how you send / store / print / buffer this
		fmt.Print(string(data))
	})

	// "stream" JSON data in, directly from whatever original format you use
	js.PutKey([]byte("revision"))
	js.PutInt(12)
	js.PutKey([]byte("attributes"))
	js.OpenObject()
	js.PutKey([]byte("key"))
	js.PutInt(13)
	js.CloseObject()
	js.End()
```

### Output
```json
{
	"revision": 12,
	"attributes": {
		"key": 13
	}
}

```