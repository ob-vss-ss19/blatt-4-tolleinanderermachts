FROM obraun/vss-protoactor-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o usercontrol/main usercontrol/main.go usercontrol/usercontrol.go

FROM iron/go
COPY --from=builder /app/usercontrol/main /app/usercontrol
EXPOSE 50000-70000
ENTRYPOINT ["/app/usercontrol"]
