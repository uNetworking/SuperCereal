
<div align="center"><img src="al.jpg" /></div>

`SuperCereal` is a simple and efficient JSON serialization library for Go. Unlike most other serializers it doesn't operate using an intermediary tree data structure (think "DOM"). This allows for a very lightweight and efficient JSON serialization process without any memory allocations in hot execution paths.

#### Simple
With an object oriented & generic design it only takes a few lines to generate JSON:
```go
// Allocates memory, keep it alive and reuse it for many serializations!
js := supercereal.NewJSONStream()

// Register a callback to receive json output in (potentially) chunks of bytes.
js.OnJSON(func(json []byte) {
	fmt.Println(string(json))
})

// Perform the actual serialization of your data.
js.Serialize(func(object *supercereal.JSONObject) {
	object.Put("firstName", "John")
	object.Put("lastName", "Smith")
	object.Put("isAlive", true)
	object.Put("age", 25)

	object.Put("address", func(object *supercereal.JSONObject) {
		object.Put("streetAddress", "21 2nd Street")
		object.Put("city", "New York")
		object.Put("state", "NY")
		object.Put("postalCode", "10021-3100")
	})

	object.Put("phoneNumbers", func(array *supercereal.JSONArray) {
		array.Put(func(object *supercereal.JSONObject) {
			object.Put("type", "home")
			object.Put("number", "212 555-1234")
		})
		array.Put(func(object *supercereal.JSONObject) {
			object.Put("type", "office")
			object.Put("number", "646 555-4567")
		})
		array.Put(func(object *supercereal.JSONObject) {
			object.Put("type", "mobile")
			object.Put("number", "123 456-7890")
		})
	})

	object.Put("children", func(array *supercereal.JSONArray) {
	})

	object.Put("spouse", nil)
})
```

Above sample prints: `{"firstName":"John","lastName":"Smith","isAlive":true,"age":25,"address":{"streetAddress":"21 2nd Street","city":"New York","state":"NY","postalCode":"10021-3100"},"phoneNumbers":[{"type":"home","number":"212 555-1234"},{"type":"office","number":"646 555-4567"},{"type":"mobile","number":"123 456-7890"}],"children":[],"spouse":null}`


#### Efficient
The Go standard library call `json.Marshal` expects a tree data structure (something like a `map[interface{}]interface{}`) holding a complete copy of the data you want to serialize. By skipping the allocation, population, traversal and deallocation of this tree, SuperCereal outperforms `json.Marshal` by **25x in CPU time**. It also completely skips garbage collection overhead by avoiding costly memory allocations and deallocations in the hot execution path. Test data and benchmarks are openly available in `main.go` for easy validation.

#### Open
Licensed Zlib © 2017
