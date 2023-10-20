# Usa la imagen de golang:latest como base
FROM golang:1:20

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el código al contenedor
COPY . .

# Compila la aplicación
RUN go build ./cmd/KindleHighlightsKeeper

# Expone el puerto que utiliza tu aplicación
EXPOSE 8080

# Comando por defecto para ejecutar tu aplicación
CMD ["./KindleHighlightsKeeper"]