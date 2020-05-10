FROM golang:1.12
EXPOSE 8080
WORKDIR /go/src/github.com/Excel-MEC/excelplay-backend-dalalbull
COPY . .
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
ENTRYPOINT CompileDaemon -directory="." -log-prefix=false -build="go build /go/src/github.com/Excel-MEC/excelplay-backend-dalalbull/cmd/excelplay-backend-dalalbull" -command="./excelplay-backend-dalalbull"