# build stage
FROM golang:1.19-alpine AS build-env
RUN apk --no-cache add build-base git gcc
ADD . /src
RUN cd /src && go build -o backend

# final stage
FROM alpine
WORKDIR /app
RUN apk add bash
COPY database /database/
COPY --from=build-env /src/backend /app/
COPY wait-for-it.sh /wait-for-it.sh
ENTRYPOINT ./backend
