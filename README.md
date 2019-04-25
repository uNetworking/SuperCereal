<div align="center"><img src="al.png" /></div>

**SuperCereal** is a simple and efficient JSON serialization library for Go. Unlike most other serializers, including Go's `json.Marshal` and JavaScript's `JSON.stringify`, it doesn't operate using an intermediary tree data structure. This makes it significantly more efficient in both time and memory. Because it doesn't use third-party non-standard preprocessing, your code remains purely standard & portable Golang.

## 17x the json.Marshal performance
```go
supercereal.Marshal(func(object *supercereal.Object) {
  object.Put("firstName", "John")
  object.Put("lastName", "Smith")
  object.Put("isAlive", true)
  object.Put("age", 25)
  object.Put("address", func(object *supercereal.Object) {
    object.Put("streetAddress", "21 2nd Street")
    object.Put("city", "New York")
    object.Put("state", "NY")
    object.Put("postalCode", "10021-3100")
  })
  object.Put("phoneNumbers", func(array *supercereal.Array) {
    array.Put(func(object *supercereal.Object) {
      object.Put("type", "home")
      object.Put("number", "212 555-1234")
    })
    array.Put(func(object *supercereal.Object) {
      object.Put("type", "office")
      object.Put("number", "646 555-4567")
    })
    array.Put(func(object *supercereal.Object) {
      object.Put("type", "mobile")
      object.Put("number", "123 456-7890")
    })
  })
  object.Put("children", func(array *supercereal.Array) {})
  object.Put("spouse", nil)
})
```

Where `supercereal.Marshal` returns the `[]byte` (here prettified for demonstration):
```json
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
  "phoneNumbers": [
    {
      "type": "home",
      "number": "212 555-1234"
    },
    {
      "type": "office",
      "number": "646 555-4567"
    },
    {
      "type": "mobile",
      "number": "123 456-7890"
    }
  ],
  "children": [],
  "spouse": null
}
```

*Licensed Zlib © 2017 - 2019*
