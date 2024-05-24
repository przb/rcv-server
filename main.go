package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-fastapi"
)

type EchoInput struct {
	Phrase string `json:"phrase"`
}

type EchoOutput struct {
	OriginalInput EchoInput `json:"original_input"`
}

func EchoHandler(ctx *gin.Context, in EchoInput) (out EchoOutput, err error) {
	out.OriginalInput = in
	return
}

func generateSwagger(fn string, router *fastapi.Router) error {
	swagger := router.EmitOpenAPIDefinition()
	swagger.Info.Title = "Ranked Choice Voting API"

	f, err := os.Create(fn)
	if err != nil {
		return err
	}

	defer f.Close()

	jsonBytes, _ := json.MarshalIndent(swagger, "", "    ")
	_, err = f.Write(jsonBytes)
	if err != nil {
		return err
	}

	fmt.Printf("Wrote Swagger to %s\n", fn)

	return nil
}

func main() {
	r := gin.Default()

	router := fastapi.NewRouter()
	router.AddCall("/echo", EchoHandler)

	generateSwagger("swagger.json", router)

	r.POST("/api/*path", router.GinHandler) // must have *path parameter
	r.Run()
}

// Try it:
//     $ curl -H "Content-Type: application/json" -X POST --data '{"phrase": "hello"}' localhost:8080/api/echo
//     {"response":{"original_input":{"phrase":"hello"}}}
