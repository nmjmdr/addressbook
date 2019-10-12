FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git gcc
ADD ./src /src
RUN cd /src && go build -o server


# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/server /app/ 
EXPOSE 8080
ENTRYPOINT ./server
