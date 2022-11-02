# # Initial stage: download modules
FROM golang:1.19-alpine as builder

ENV config=docker

WORKDIR /usr/local/go/src/app

#COPY ./go.mod .
#COPY ./go.sum .
COPY . .

#RUN go mod download
RUN go build -o bin/server ./cmd/


# COPY ./ /usr/local/go/src/app

# RUN go mod download
#RUN go get github.com/githubnemo/CompileDaemon
#RUN go mod vendor

# Intermediate stage: Build the binary
FROM golang:1.19-alpine
# as runner

#COPY --from=builder ./usr/local/go/src/ /src
#COPY  --from=builder /usr/local/go/ /usr/local/go/
COPY  --from=builder /go/pkg/ /go/pkg/
COPY ./ /superapp
WORKDIR /superapp

#RUN go mod download
#RUN go get github.com/githubnemo/CompileDaemon
#RUN go mod vendor

ENV config=docker

ARG GOFILE=${GOFILE}

# EXPOSE 5005
# EXPOSE 5555   
# EXPOSE 7070

#RUN go build -o bin/server ./cmd/main.go
RUN pwd && echo ./${GOFILE} && go build -o bin/server ./${GOFILE}

EXPOSE 3000

CMD ["./bin/server"]

#ENTRYPOINT CompileDaemon --build="go build ./cmd/main.go" --command=./main


## Initial stage: download modules
#FROM golang:1.17-alpine as builder
#
#ENV config=docker
#
#WORKDIR /app
#
#COPY ./ /app
#
#RUN go mod download
#
#
## Intermediate stage: Build the binary
#FROM golang:1.17-alpine as runner
#
#COPY --from=builder ./app ./app
#
#RUN go get github.com/githubnemo/CompileDaemon
#
#WORKDIR /app
#ENV config=docker
#
#EXPOSE 5000
#EXPOSE 5555
#EXPOSE 7070
#
#ENTRYPOINT CompileDaemon --build="go build cmd/api/main.go" --command=./main
#
#
#
#
