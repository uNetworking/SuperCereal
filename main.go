package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/alexhultman/SuperCereal/supercereal"
)

func benchmarkSuperCereal() {
	js := supercereal.NewJSONStream(func(data []byte) {
		// this is where you get data (potentially in chunks)
		// it is up to you how you send / store / print / buffer this
		fmt.Print(string(data))
	})

	start := time.Now()
	for i := 0; i < 1000000; i++ {
		js.Reset()

		// "stream" JSON data in (example from Wikipedia)
		// todo: improve interfaces (helpers, etc, open and close pairs can be actual functions)
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

	fmt.Printf("\n\nSerialization took %f µs\n", time.Since(start).Seconds())
	js.End()
}

func benchmarkSuperCerealSimple() {
	js := supercereal.NewJSONStream(func(data []byte) {
		fmt.Print(string(data))
	})

	start := time.Now()
	for i := 0; i < 1000000; i++ {

		// reset kan vara den enda funktionen i rooten
		// som ger dig antingen en array eller ett object!
		// då tas även End bort!
		js.Reset()

		js.Put("firstName", "John")
		js.Put("lastName", "Smith")
		js.Put("isAlive", true)
		js.Put("age", 25)

		js.Put("address", func(object supercereal.JSONObject) {
			object.Put("streetAddress", "21 2nd Street")
			object.Put("city", "New York")
			object.Put("state", "NY")
			object.Put("postalCode", "10021-3100")
		})

		js.Put("phoneNumbers", func(array supercereal.JSONArray) {
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

		js.Put("children", func(array supercereal.JSONArray) {
		})

		js.Put("spouse", nil)
	}

	fmt.Printf("Serialization took %f µs\n", time.Since(start).Seconds())
	js.End()
}

func benchmarkJSONMarshal() {
	start := time.Now()
	var bytes []byte
	for i := 0; i < 1000000; i++ {
		root := make(map[string]interface{}, 0)
		root["firstName"] = "John"
		root["lastName"] = "Smith"
		root["isAlive"] = true
		root["age"] = 25

		address := make(map[string]interface{}, 0)
		address["streetAddress"] = "21 2nd Street"
		address["city"] = "New York"
		address["state"] = "NY"
		address["postalCode"] = "10021-3100"

		root["address"] = address

		phoneNumbers := make([]map[string]interface{}, 3)
		phoneNumbers[0] = make(map[string]interface{}, 0)
		phoneNumbers[0]["type"] = "home"
		phoneNumbers[0]["number"] = "212 555-1234"

		phoneNumbers[1] = make(map[string]interface{}, 0)
		phoneNumbers[1]["type"] = "office"
		phoneNumbers[1]["number"] = "646 555-4567"

		phoneNumbers[2] = make(map[string]interface{}, 0)
		phoneNumbers[2]["type"] = "mobile"
		phoneNumbers[2]["number"] = "123 456-7890"

		root["phoneNumbers"] = phoneNumbers
		root["children"] = make([]map[string]interface{}, 0)
		root["spouse"] = nil

		bytes, _ = json.Marshal(root)
	}

	fmt.Printf("\n\nSerialization took %f µs\n", time.Since(start).Seconds())
	fmt.Println(string(bytes))
}

func main() {
	benchmarkSuperCerealSimple() // high level interface
	benchmarkSuperCereal()       // low level interface
	//benchmarkJSONMarshal()
}
