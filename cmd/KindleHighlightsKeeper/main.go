package main

import (
	"fmt"
	"os"
    "strings"
    "github.com/FranJF/KindleHighlightsKeeper/internal/convertkindle"
)


func main() {
	if len(os.Args) != 3 {
		fmt.Println("Se espera el nombre del archivo como primer argumento y JSON/TXT como segundo argumento")
		return
	}

	htmlFilePath := os.Args[1]
	format := strings.ToLower(os.Args[2])

    filename, err := convertkindle.Convert(htmlFilePath, format)
    if err != nil {
        fmt.Println("Error al convertir el archivo:", err)
        return
    }

	fmt.Println(filename)
}

