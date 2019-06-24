FROM obraun/vss-protoactor-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o reservationcontrol/main reservationcontrol/main.go reservationcontrol/reservation.go

FROM iron/go
COPY --from=builder /app/reservationcontrol/main /app/reservationcontrol
EXPOSE 52000-65000
ENTRYPOINT ["/app/reservationcontrol"]
