package main

import (
	"SuperCereal/supercereal"
	"fmt"
)

func main() {

	js := supercereal.NewJSONStream(func(data []byte) {
		// this is where you get data (potentially in chunks)
		// it is up to you how you send / store / print / buffer this
		fmt.Print(string(data))
	})

	// "stream" JSON data in
	js.PutKey([]byte("revision")) // om scope.length != 0, lägg till komma före! (stack av scopes!)
	js.PutInt(12)
	js.PutKey([]byte("attributes"))
	js.OpenObject()
	js.PutKey([]byte("key"))
	js.PutInt(13)
	js.CloseObject()
	js.End()
}
