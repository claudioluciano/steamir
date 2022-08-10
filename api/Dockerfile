FROM golang:1.19-alpine@sha256:78767a599f5758e3fd60fd065de9dc40aec1ad3b86ded2bddc92640a465cb515 as builder

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /bin/app

# Build final image containing only the app binary.
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /bin/app /bin/app
ENTRYPOINT ["/bin/app"]
