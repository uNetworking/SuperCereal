/* Run benchmarks with go test -bench . */

package main

import (
	"testing"
	"encoding/json"
	"./supercereal"
)

/* Standard Golang json.Marshal */
func BenchmarkJSONMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(map[string]interface{}{
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
}

/* The small struct used in mailru/easyjson benchmarks */
func BenchmarkAgainstEasyJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		supercereal.Marshal(func(object *supercereal.JSONObject) {
			object.Put("hashtags", func(array *supercereal.JSONArray) {
				array.Put(func(object *supercereal.JSONObject) {
					object.Put("indices", func(array *supercereal.JSONArray) {
						array.Put(5)
						array.Put(10)
					})
					object.Put("text", "some-text")
				})
			})
			object.Put("urls", func(array *supercereal.JSONArray) {})
			object.Put("user_mentions", func(array *supercereal.JSONArray) {})
		})
	}
}

/* SuperCereal way of doing it */
func BenchmarkSuperCereal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		supercereal.Marshal(func(object *supercereal.JSONObject) {
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
			object.Put("children", func(array *supercereal.JSONArray) {})
			object.Put("spouse", nil)
		})
	}
}