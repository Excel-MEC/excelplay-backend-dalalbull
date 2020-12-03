FROM golang:1.13
EXPOSE 8080
WORKDIR /excelplay-backend-dalalbull
COPY . .
RUN ["go", "build", "/excelplay-backend-dalalbull/cmd/excelplay-backend-dalalbull"]
ENTRYPOINT ./excelplay-backend-dalalbull