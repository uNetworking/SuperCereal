<div align="center"><img src="al.jpg" /></div>

`SuperCereal` is a simple and efficient JSON serialization library for Go. Unlike most other serializers, including standard `json.Marshal`, it doesn't operate using an intermediary tree data structure (`map[string]interface{}`). This makes it a lot more efficient in both time and memory.

#### Simple
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
#### Efficient
SuperCereal outperforms `json.Marshal` by about 25x in time:
![](benchmark.png)

#### Open
Licensed Zlib © 2017
