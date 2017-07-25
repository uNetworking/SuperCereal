package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/alexhultman/SuperCereal/supercereal"
)

func benchmarkSuperCerealHighLevel() {
	var lastJSON []byte
	js := supercereal.NewJSONStream()
	js.OnJSON(func(json []byte) {
		lastJSON = json
	})

	start := time.Now()
	for i := 0; i < 1000000; i++ {
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
	}

	fmt.Printf("benchmarkSuperCerealHighLevel took %f µs\n", time.Since(start).Seconds())
	fmt.Printf("%s\n\n", string(lastJSON))
}

func benchmarkSuperCerealLowLevel() {
	js := supercereal.NewJSONStream()
	js.OnJSON(func(json []byte) {
		fmt.Printf("%s\n\n", string(json))
	})

	start := time.Now()
	for i := 0; i < 1000000; i++ {
		js.Reset()
		js.OpenObject()
		{
			js.PutKey([]byte("firstName"))
			js.PutString([]byte("John"))

			js.PutKey([]byte("lastName"))
			js.PutString([]byte("Smith"))

			js.PutKey([]byte("isAlive"))
			js.PutBoolean(true)

			js.PutKey([]byte("age"))
			js.PutInt(25)

			js.PutKey([]byte("address"))
			js.OpenObject()
			{
				js.PutKey([]byte("streetAddress"))
				js.PutString([]byte("21 2nd Street"))

				js.PutKey([]byte("city"))
				js.PutString([]byte("New York"))

				js.PutKey([]byte("state"))
				js.PutString([]byte("NY"))

				js.PutKey([]byte("postalCode"))
				js.PutString([]byte("10021-3100"))
			}
			js.CloseObject()

			js.PutKey([]byte("phoneNumbers"))
			js.OpenArray()
			{
				js.OpenObject()
				{
					js.PutKey([]byte("type"))
					js.PutString([]byte("home"))

					js.PutKey([]byte("number"))
					js.PutString([]byte("212 555-1234"))
				}
				js.CloseObject()

				js.OpenObject()
				{
					js.PutKey([]byte("type"))
					js.PutString([]byte("office"))

					js.PutKey([]byte("number"))
					js.PutString([]byte("646 555-4567"))
				}
				js.CloseObject()

				js.OpenObject()
				{
					js.PutKey([]byte("type"))
					js.PutString([]byte("mobile"))

					js.PutKey([]byte("number"))
					js.PutString([]byte("123 456-7890"))
				}
				js.CloseObject()
			}
			js.CloseArray()

			js.PutKey([]byte("children"))
			js.OpenArray()
			js.CloseArray()

			js.PutKey([]byte("spouse"))
			js.PutNull()
		}
		js.CloseObject()
	}

	fmt.Printf("benchmarkSuperCerealLowLevel took %f µs\n", time.Since(start).Seconds())
	js.End()
}

func benchmarkJSONMarshal() {
	start := time.Now()
	var bytes []byte
	for i := 0; i < 1000000; i++ {
		bytes, _ = json.Marshal(map[string]interface{}{
			"firstName": "John",
			"lastName":  "Smith",
			"isAlive":   true,
			"age":       25,
			"address": map[string]interface{}{
				"streetAddress": "21 2nd Street",
				"city":          "New York",
				"state":         "NY",
				"postalCode":    "10021-3100",
			},
			"phoneNumbers": []interface{}{
				map[string]interface{}{
					"type":   "home",
					"number": "212 555-1234",
				},
				map[string]interface{}{
					"type":   "office",
					"number": "646 555-4567",
				},
				map[string]interface{}{
					"type":   "mobile",
					"number": "123 456-7890",
				},
			},
			"children": []interface{}{},
			"spouse":   nil,
		})
	}

	fmt.Printf("benchmarkJSONMarshal took %f µs\n", time.Since(start).Seconds())
	fmt.Printf("%s\n\n", string(bytes))
}

func main() {
	for i := 0; i < 5; i++ {
		benchmarkSuperCerealHighLevel()
		benchmarkSuperCerealLowLevel()
	}
	benchmarkJSONMarshal()
}
