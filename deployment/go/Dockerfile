# Используем официальный образ Go для разработки
FROM golang:1.23.3-alpine

# Устанавливаем рабочую директорию
WORKDIR /eden

# Принимаем UID и GID как аргументы на этапе сборки
ARG UID
ARG GID

# Добавляем группу и пользователя с переданными UID и GID
RUN addgroup -g $GID mygroup && adduser -D -u $UID -G mygroup -h /eden myuser

# Устанавливаем необходимые инструменты
RUN apk add --no-cache git bash gcc musl-dev mysql-client

# Устанавливаем Delve (dlv) для отладки
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Копируем зависимости
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем весь проект
COPY . .

# Устанавливаем пользователя, под которым будет работать контейнер
USER myuser

# Возвращаемся в рабочую директорию
WORKDIR /eden
