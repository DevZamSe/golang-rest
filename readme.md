INSTRUCTIONS TO BUILD THE GOLANG APP:
The Golang app ubicates in /go/src/
Use via Git in go/src directory in console

INSTRUCTIONS RUN WITH GO:
first: + go mod init

second: + go get github.com/gorilla/mux + go get github.com/go-sql-driver/mysql

INSTRUCTIONS RUN WITH DOCKER:
first: + docker build -t microservices .

second: + docker run --rm -it -p 8000:8000 microservices
