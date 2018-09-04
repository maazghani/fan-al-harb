FROM golang:alpine AS build-env
ADD . /src  
RUN cd /src/server && go build -o fan-al-harb

FROM alpine
WORKDIR /app
COPY --from=build-env /src/server/fan-al-harb /app/
ENTRYPOINT ./fan-al-harb
