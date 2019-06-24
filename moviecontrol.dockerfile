FROM obraun/vss-protoactor-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o moviecontrol/main moviecontrol/main.go moviecontrol/moviecontrol.go

FROM iron/go
COPY --from=builder /app/moviecontrol/main /app/moviecontrol
EXPOSE 52000-53000
ENTRYPOINT ["/app/moviecontrol"]
