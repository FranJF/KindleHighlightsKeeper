package convertkindle

import(
	"fmt"
	"errors"
    "strings"
    "github.com/FranJF/KindleHighlightsKeeper/internal/htmltojson"
)

const folderOutput = "output/"
var replacer = strings.NewReplacer("\n", "")

func Convert(htmlFile string, format string) (string, string, error){
    bookTitle, results, sections, err := htmltojson.ParseHTML(htmlFile)

	if err != nil {
		fmt.Println("Error al analizar el archivo HTML:", err)
		return "" ,"", err
	}

    var content string
    switch format {
    case "txt":
        content = createTXT(bookTitle, results, sections)
    // case "json":
    //     content = createJSON(bookTitle, results, sections)
    case "md":
        content = createMD(bookTitle, results, sections)
    default:
        return "","", errors.New("Formato no soportado")
    }

    return bookTitle, content, nil

}

// func createJSON(bookTitle string, results map[string][]string, sections[]string) map[string]interface{} {
//     data := map[string]interface{}{
//         "title": bookTitle,
//         "results": results,
//     }

//     return data
// }

func createTXT(bookTitle string, results map[string][]string, sections[]string) string {
    stringToWrite := bookTitle + "\n\n"
    for _, section := range sections {
        stringToWrite += section + "\n"
        for _, note := range results[section] {
            if strings.Contains(note, "Nota: ") {
                strings.ReplaceAll(note, "Nota: ", "")
                stringToWrite += note + "\n\n"
                continue
            }
            stringToWrite += note + "\n"
        }
        stringToWrite += "\n"
    }

    return stringToWrite
}

func createMD(bookTitle string, results map[string][]string, sections[]string) string {
    stringToWrite := "# " + strings.ReplaceAll(bookTitle, "_", " ") + "\n\n"
    for _, section := range sections {
        stringToWrite += "## " + section + "\n"
        for _, note := range results[section] {
            if strings.Contains(note, "Nota: ") {
                strings.ReplaceAll(note, "Nota: ", "")
                stringToWrite += "> " + note + "\n"
                continue
            }
            stringToWrite += "* " + note + "\n"
        }
        stringToWrite += "\n"
    }

    return stringToWrite
}
