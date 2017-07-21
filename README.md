<div align="center"><img src="al.jpg" /></div>

SuperCereal is a simple and efficient JSON serialization library for Go. Unlike most other serializers it doesn't operate using an intermediary tree data structure (think "DOM") but instead exposes a set of simple fuction calls to directly control the JSON generation (a la "SAX"). This makes for a very lightweight and efficient JSON serialization process where most execution paths incur no memory allocations.

### Simple
With an object oriented & generic design it takes only a few lines to generate JSON:
```go
json := supercereal.Serialize(func(object supercereal.JSONObject) {
	object.Put("firstName", "John")
	object.Put("lastName", "Smith")
	object.Put("isAlive", true)
	object.Put("age", 25)

	object.Put("address", func(object supercereal.JSONObject) {
		object.Put("streetAddress", "21 2nd Street")
		object.Put("city", "New York")
		object.Put("state", "NY")
		object.Put("postalCode", "10021-3100")
	})

	object.Put("phoneNumbers", func(array supercereal.JSONArray) {
		array.Put(func(object supercereal.JSONObject) {
			object.Put("type", "home")
			object.Put("number", "212 555-1234")
		})
		array.Put(func(object supercereal.JSONObject) {
			object.Put("type", "office")
			object.Put("number", "646 555-4567")
		})
		array.Put(func(object supercereal.JSONObject) {
			object.Put("type", "mobile")
			object.Put("number", "123 456-7890")
		})
	})

	object.Put("children", func(array supercereal.JSONArray) {
	})

	object.Put("spouse", nil)
})

```

### Efficient
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
