FROM obraun/vss-protoactor-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o roomcontrol/main roomcontrol/main.go roomcontrol/roomcontrol.go

FROM iron/go
COPY --from=builder /app/roomcontrol/main /app/roomcontrol
EXPOSE 52000-53000
ENTRYPOINT ["/app/roomcontrol"]
