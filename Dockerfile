#Dockerfile to run a Go webservice. Note this does not build it, only runs the binary in a container
#esnure the binary has been ccross compiled for linux using...
# env GOOS=linux GOARCH=386 go build -o bin/todo_list_svc main.go
FROM alpine:3.7

EXPOSE 8080

WORKDIR /app

COPY bin/todo_list_svc ./

# Command to run the executable
CMD ["./todo_list_svc"]

