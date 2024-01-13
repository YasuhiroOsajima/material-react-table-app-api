FROM golang:latest AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o material-react-table-app-api /app/cmd/main.go


FROM golang:latest
WORKDIR /app
COPY --from=builder /app/material-react-table-app-api .
COPY --from=builder /app/dummy_people.json .

EXPOSE 7777
CMD ["/app/material-react-table-app-api"]
