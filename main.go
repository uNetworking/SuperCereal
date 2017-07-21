package main

import (
	"SuperCereal/supercereal"
	"encoding/json"
	"fmt"
	"time"
)

func main() {

	js := supercereal.NewJSONStream(func(data []byte) {
		// this is where you get data (potentially in chunks)
		// it is up to you how you send / store / print / buffer this
		fmt.Print(string(data))
	})

	// benchmark SuperCereal
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		js.Reset()

		// "stream" JSON data in
		js.PutKey([]byte("revision"))
		js.PutInt(12)

		js.PutKey([]byte("attributes"))
		js.OpenObject()
		{
			js.PutKey([]byte("key"))
			js.OpenObject()
			{
				js.PutKey([]byte("recursive"))
				js.PutString([]byte("hallåja"))

				js.PutKey([]byte("jaja"))
				js.PutString([]byte("hellå"))
			}
			js.CloseObject()

			js.PutKey([]byte("key2"))
			js.OpenObject()
			{
				js.PutKey([]byte("recursive2"))
				js.PutString([]byte("hallåja2"))

				js.PutKey([]byte("jaja2"))
				js.PutString([]byte("hellå2"))
			}
			js.CloseObject()
		}
		js.CloseObject()
	}

	fmt.Printf("Serialization took %f µs\n", time.Since(start).Seconds())
	js.End()

	// benchmark constructing a map of maps + json.Marshal
	start = time.Now()
	var bytes []byte
	for i := 0; i < 1000000; i++ {
		root := make(map[string]interface{}, 0)

		root["revision"] = 12

		attributes := make(map[string]interface{}, 0)

		object1 := make(map[string]interface{}, 0)
		object1["recursive"] = "hallåja"
		object1["jaja"] = "hellå"

		attributes["key"] = object1

		object2 := make(map[string]interface{}, 0)
		object2["recursive2"] = "hallåja2"
		object2["jaja2"] = "hellå2"

		attributes["key2"] = object2

		root["attributes"] = attributes

		bytes, _ = json.Marshal(root)
	}

	fmt.Printf("\n\nSerialization took %f µs\n", time.Since(start).Seconds())
	fmt.Println(string(bytes))
}
