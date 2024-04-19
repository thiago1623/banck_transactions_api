# Usa la imagen base de golang
FROM golang:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos de configuración de dependencias go.mod y go.sum
COPY go.mod go.sum ./

# Descarga e instala las dependencias
RUN go mod download

# Copia el resto de la aplicación
COPY . .

# Expon el puerto 8080
EXPOSE 8080

# Ejecuta la aplicación
CMD ["go", "run", "main.go"]