FROM alpine

WORKDIR /app
COPY cmd/main.exe .

ENTRYPOINT ./main.exe