package main

import (
	"encoding/json"
	"fmt"
	"os"
    "github.com/FranJF/KindleHighlightsKeeper/internal/htmltojson"
    "strings"
)

func writeJSON(jsonFileName string, results map[string][]string) error {
	jsonData, err := json.MarshalIndent(results, "", "    ")
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return err
	}

    jsonFile, err := os.Create("output/"+jsonFileName)
    if err != nil {
        return err
    }
    defer jsonFile.Close()

    // Reemplazar los caracteres de nueva línea
    jsonDataString := strings.ReplaceAll(string(jsonData), "\n", "")

    jsonFile.WriteString(jsonDataString)

    return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Se espera el nombre del archivo como argumento.")
		return
	}

	htmlFilePath := os.Args[1]

	results, bookTitle, err := htmltojson.ParseHTML(htmlFilePath)
	if err != nil {
		fmt.Println("Error al analizar el archivo HTML:", err)
		return
	}

	jsonFileName := bookTitle + ".json"

    err = writeJSON(jsonFileName, results)
    if err != nil {
        fmt.Println("Error al escribir el archivo JSON:", err)
        return
    }

	fmt.Println("Archivo JSON generado:", jsonFileName)
}

