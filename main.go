/*
 * OpenAPI Petstore
 *
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	openapi "myapi/go"
	"myapi/myapi"
	"net/http"
)

func main() {
	log.Printf("Server started")

	memoryStore := new(myapi.MemoryItemStore)
	memoryStore.CreateItem("Meat 🥩")
	memoryStore.CreateItem("Apples 🍎")
	memoryStore.CreateItem("Oranges 🍊")
	memoryStore.CreateItem("Tomato sauce 🥫")
	potatoes := memoryStore.CreateItem("Potatos 🥔")
	bananas := memoryStore.CreateItem("Bananas 🍌")

	memoryStore.ToggleItem(potatoes)
	memoryStore.ToggleItem(bananas)

	/*
        { name: "Meat 🥩", done: false, onToggleItem: this.handleDone },
        { name: "Apples 🍎", done: false, onToggleItem: this.handleDone },
        { name: "Oranges 🍊", done: false, onToggleItem: this.handleDone },
        { name: "Bananas 🍌", done: false, onToggleItem: this.handleDone },
        { name: "Tomato sauce 🥫", done: false, onToggleItem: this.handleDone },
        { name: "Potatos 🥔", done: false, onToggleItem: this.handleDone }
      ],
      doneItems: [
        { name: "Milk 🥛", done: true, onToggleItem: this.handleTodo },
        { name: "Eggs 🥚", done: true, onToggleItem: this.handleTodo }
      ]
	*/

	ItemAPIService := myapi.NewItemAPIService(memoryStore)
	ItemAPIController := openapi.NewDefaultApiController(ItemAPIService)

	router := openapi.NewRouter(ItemAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
