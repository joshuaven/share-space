FROM node:16-alpine as builder

WORKDIR /app

COPY . .

RUN cd client && npm install && npm run build

FROM golang:latest

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY --from=builder /app ./

RUN go build -v

CMD [ "./share-space" ]





