FROM golang

RUN go get github.com/derekparker/delve/cmd/dlv
RUN go get github.com/gin-gonic/gin
RUN go get github.com/pilu/fresh
RUN go get github.com/jinzhu/gorm
RUN go get github.com/martini-contrib/render
RUN go get github.com/Sirupsen/logrus
RUN go get github.com/jinzhu/gorm/dialects/sqlite
RUN go get github.com/joho/godotenv
WORKDIR /go/src/app
# WORKDIR /app

# application port
EXPOSE 3000
# debug port for vscode
EXPOSE 40000
# RUN go build -o main .
# CMD ["./main"]
CMD [ "dlv", "debug", "/go/src/app/main.go", "--listen=:40000", "--headless=true", "--api-version=2", "--log"]
# CMD [ "dlv", "debug", "/go/src/app/main.go", "--listen=:40000", "--headless=true", "--api-version=2" ]
# CMD fresh