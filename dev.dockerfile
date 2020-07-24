FROM golang:1.13
EXPOSE 8080
WORKDIR /excelplay-backend-dalalbull
COPY . .
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
ENTRYPOINT CompileDaemon -directory="." -log-prefix=false -build="go build /excelplay-backend-dalalbull/cmd/excelplay-backend-dalalbull" -command="./excelplay-backend-dalalbull"