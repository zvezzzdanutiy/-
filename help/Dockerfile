# Устанавливаем базовый образ Golang
FROM --platform=linux/amd64 golang:latest

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /go/src/help

# Копируем все остальные файлы из исходной директории в рабочую директорию внутри контейнера
COPY . .

# Установка зависимостей
RUN go mod download

WORKDIR /go/src/help/cmd

# Собираем приложение
RUN go build -o main

# Запускаем приложение при запуске контейнера
CMD ["./main"]
