<div align="center"><img src="al.jpg" /></div>

SuperCereal is a simple and efficient JSON serializing library for Go. Unlike many other "fast" serializers, it doesn't make use of intermediate (hash) maps or "documents". Instead it immediately dumps to a pre-allocated buffer, whatever data you add as you add it. JSON is emitted in chunks of bytes in a streaming fashion. Because of this SuperCereal runs a lot faster and with way less memory usage compared to solutions with intermediate documents (such as RapidJSON).

### Overview
```go
	js := supercereal.NewJSONStream(func(data []byte) {
		// this is where you get data (potentially in chunks)
		// it is up to you how you send / store / print / buffer this
		fmt.Print(string(data))
	})

	// "stream" JSON data in, directly from whatever original format you use
	js.PutKey([]byte("firstName"))
	js.PutString([]byte("John"))

	js.PutKey([]byte("lastName"))
	js.PutString([]byte("Smith"))

	js.PutKey([]byte("isAlive"))
	js.PutBoolean(true)

	js.PutKey([]byte("age"))
	js.PutInt(25)
	js.End()
```

### Benchmarks
The following JSON was generated using multiple `map[string]interface{}` + `json.Marshal` in 12.987583 µs:
```go
{
	"firstName": "John",
	"lastName": "Smith",
	"isAlive": true,
	"age": 25,
	"address": {
		"streetAddress": "21 2nd Street",
		"city": "New York",
		"state": "NY",
		"postalCode": "10021-3100"
	},
	"phoneNumbers": [{
		"type": "home",
		"number": "212 555-1234"
	}, {
		"type": "office",
		"number": "646 555-4567"
	}, {
		"type": "mobile",
		"number": "123 456-7890"
	}],
	"children": [],
	"spouse": null
}
```
The same JSON was generated using SuperCereal in 0.485641 µs, some 25x as fast. The bigger the JSON (esp. depth), the bigger the performance difference.