FROM golang:1.20-alpine3.17 as development
WORKDIR /app
EXPOSE 5000

COPY . .

CMD [ "././goapi-chat" ]
