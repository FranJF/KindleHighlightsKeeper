package main

import (
	"fmt"
    "bufio"
    "bytes"
	"net/http"
	"html/template"
    "github.com/FranJF/KindleHighlightsKeeper/internal/convertkindle"
)

func convertKindleHighlights(htmlFile string, format string) (string, *bytes.Buffer, error) {
    title, convertedFile, err := convertkindle.Convert(htmlFile, format)
    fmt.Println("Archivo convertido:", title+"."+format)
    if err != nil {
        fmt.Println("Error al convertir el archivo:", err)
        return "", nil, err
    }

    var buf bytes.Buffer
	buf.WriteString(convertedFile)
	return title, &buf, nil
}

func handleFileUpload(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        r.ParseMultipartForm(10 << 20) // límite de tamaño de archivo 10MB
        file, _, err := r.FormFile("myFile")
        wantedFormat := r.FormValue("wantedFormat")

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        if wantedFormat == "" {
            http.Error(w, "No se ha especificado el formato de salida", http.StatusBadRequest)
            return
        }
        defer file.Close()


        scanner := bufio.NewScanner(file)
        var text string
        for scanner.Scan() {
            text += scanner.Text() + "\n"
        }

        // if http.DetectContentType([]byte(text)) != "text/html" {
        //     http.Error(w, "El archivo no es de formato HTML", http.StatusBadRequest)
        //     return
        // }

        title, buf, err := convertKindleHighlights(text, wantedFormat)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        switch wantedFormat {
            case "txt":
                w.Header().Set("Content-Type", "text/plain")
            case "json":
                w.Header().Set("Content-Type", "application/json")
            case "md":
                w.Header().Set("Content-Type", "text/markdown")
            default:
                http.Error(w, "Formato no soportado", http.StatusBadRequest)
                return
        }


        fmt.Println("Archivo convertido:", title+"."+wantedFormat)

        w.Header().Set("Content-Disposition", "attachment; filename="+title+"."+wantedFormat)
        w.Header().Set("Content-Length", fmt.Sprintf("%d", buf.Len()))
		w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
		w.WriteHeader(http.StatusOK)
		w.Write(buf.Bytes())
    }
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/index.html"))
    tmpl.Execute(w, nil)
}



func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/upload", handleFileUpload)
	http.ListenAndServe(":8080", nil)
}

