WORKDIR /app

 COPY go.mod ./
 COPY main.go ./

 RUN go build -o main .

 FROM alpine:latest

 WORKDIR /root/

 COPY --from=build /app/main .

 EXPOSE 8080

 CMD ["./main"]
