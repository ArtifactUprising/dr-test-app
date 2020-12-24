FROM golang:1.14.1 as builder

ADD . /app/
RUN cd /app && go build 

FROM golang:1.14.1 

WORKDIR /src

COPY --from=builder  /app/dr-test-app .

ENTRYPOINT [ "./dr-test-app" ]