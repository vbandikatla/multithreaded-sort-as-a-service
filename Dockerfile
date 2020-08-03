FROM golang:1.14

WORKDIR "/app"

COPY . .

RUN go get -u github.com/beego/bee

EXPOSE 8080
CMD ["bee", "run"]