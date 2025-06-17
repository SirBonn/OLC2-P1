package cst

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

type CSTresponse struct {
	SVGTree string `json:"svgtree"`
}

func ReadFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, _ := io.ReadAll(file)
	return string(content)
}

func CstReport(input string) string {
	parserContent := ""
	_, filename, _, _ := runtime.Caller(0)
	path := filepath.Dir(filename)
	path = path[:len(path)-4]

	parser, err := json.Marshal(ReadFile(filepath.Join(path, "/parser/Vlang.g4")))
	if err != nil {
		fmt.Println("Error reading parser file:", err)
		return ""
	}
	parserContent = string(parser)

	lexerContent := ""
	lexer, err := json.Marshal(ReadFile(filepath.Join(path, "/parser/Vlang.g4")))
	if err != nil {
		fmt.Println("Error reading lexer file:", err)
		return ""
	}
	lexerContent = string(lexer)

	jinput, err := json.Marshal(input)
	if err != nil {
		fmt.Println("Error marshaling input:", err)
		return ""
	}
	finput := string(jinput)

	payload := []byte(
		fmt.Sprintf(
			`{
				"grammar": %s,
				"input": %s,
				"lexgrammar": %s,
				"start": "%s"
			}`,
			parserContent,
			finput,
			lexerContent,
			"programa",
		),
	)

	req, err := http.NewRequest("POST", "http://lab.antlr.org/parse/", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return ""
	}

	// fmt.Printf("Response status: %d\n", resp.StatusCode)
	// fmt.Printf("Response body: %s\n", string(body))

	var data map[string]interface{}

	// Unmarshal the JSON
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error unmarshalling json:", err)
		return ""
	}

	resultInterface, exists := data["result"]
	if !exists {
		fmt.Println("No 'result' field in response")
		if errorMsg, hasError := data["error"]; hasError {
			fmt.Printf("Server error: %v\n", errorMsg)
		}
		return ""
	}

	if resultInterface == nil {
		fmt.Println("'result' field is nil")
		return ""
	}

	result, ok := resultInterface.(map[string]interface{})
	if !ok {
		// fmt.Printf("'result' is not a map, it's: %T\n", resultInterface)
		return ""
	}

	svgtreeInterface, exists := result["svgtree"]
	if !exists {
		fmt.Println("No 'svgtree' field in result")
		// Imprimir todos los campos disponibles en result
		fmt.Println("Available fields in result:")
		for key := range result {
			fmt.Printf("  - %s\n", key)
		}
		return ""
	}

	if svgtreeInterface == nil {
		fmt.Println("'svgtree' field is nil")
		return ""
	}

	svgtree, ok := svgtreeInterface.(string)
	if !ok {
		fmt.Printf("'svgtree' is not a string, it's: %T\n", svgtreeInterface)
		return ""
	}

	return svgtree
}
