package convertkindle

import(
	"fmt"
	"os"
	"errors"
    "strings"
	"encoding/json"
    "github.com/FranJF/KindleHighlightsKeeper/internal/htmltojson"
)

const folderOutput = "output/"
var replacer = strings.NewReplacer("\n", "")


func createJSON(bookTitle string, results map[string][]string, sections[]string) error {
	jsonFileName := bookTitle + ".json"

	jsonData, err := json.MarshalIndent(results, "", "    ")
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return err
	}

    jsonFile, err := os.Create(folderOutput+jsonFileName)
    if err != nil {
        return err
    }
    defer jsonFile.Close()

    jsonDataString := replacer.Replace(string(jsonData))
    

    jsonFile.WriteString(jsonDataString)

    return nil
}

func createTXT(bookTitle string, results map[string][]string, sections[]string) error {
    txtFileName := bookTitle + ".txt"
    txtFile, err := os.Create(folderOutput+txtFileName)
    if err != nil {
        return err
    }
    defer txtFile.Close()

    txtFile.WriteString(bookTitle + "\n\n")
    for _, section := range sections {
        txtFile.WriteString(section + "\n")
        for _, note := range results[section] {
            txtFile.WriteString(note + "\n")
        }
        txtFile.WriteString("\n")
    }

    return nil
}

func createMD(bookTitle string, results map[string][]string, sections[]string) error {
    mdFileName := bookTitle + ".md"
    mdFile, err := os.Create(folderOutput+mdFileName)
    if err != nil {
        return err
    }
    defer mdFile.Close()

    mdFile.WriteString("# " + strings.ReplaceAll(bookTitle, "_", " ") + "\n\n")

    for _, section := range sections {
        mdFile.WriteString("## " + section + "\n")
        for _, note := range results[section] {
            mdFile.WriteString("* " + note + "\n")
        }
        mdFile.WriteString("\n")
    }

    return nil
}


func Convert(htmlFilePath string, format string) (string, error){
    bookTitle, results, sections, err := htmltojson.ParseHTML(htmlFilePath)
	if err != nil {
		fmt.Println("Error al analizar el archivo HTML:", err)
		return "", err
	}

    msg := ""
    msg_err := ""
    filename := ""

    switch format {
    case "txt":
        fmt.Println("Generando archivo TXT...")
        err = createTXT(bookTitle, results, sections)
        filename = bookTitle + ".txt"
        msg = "Archivo TXT generado"
        msg_err = "Error al generar el archivo TXT"
    case "json":
        fmt.Println("Generando archivo JSON...")
        err = createJSON(bookTitle, results, sections)
        filename = bookTitle + ".json"
        msg = "Archivo JSON generado"
        msg_err = "Error al generar el archivo JSON"
    case "md":
        fmt.Println("Generando archivo MD...")
        err = createMD(bookTitle, results, sections)
        filename = bookTitle + ".md"
        msg = "Archivo MD generado"
        msg_err = "Error al generar el archivo MD"
    default:
        return "", errors.New("Formato no soportado")
    }

    if err != nil {
        fmt.Println(msg_err, err)
        return "", err
    }

	fmt.Println(msg, filename)
    return filename, nil

}
