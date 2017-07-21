<div align="center"><img src="al.jpg" /></div>

SuperCereal is a simple and efficient JSON serialization library for Go. Unlike most other serializers it doesn't operate using an intermediary tree data structure (think "DOM") but instead exposes a set of simple fuction calls to directly control the JSON generation (a la "SAX"). This makes for a very lightweight and efficient JSON serialization process where most execution paths incur no memory allocations.

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
The Go standard library call `json.Marshal` expects an `interface{}` which is to represent the root of an intermediary tree data structure. This means you need to allocate nodes, populate these nodes and finally traverse them all in order to generate JSON output. This my friend, this is madness.

The following JSON was generated in about 25x less time using SuperCereal, as compared to `json.Marshal`:
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
