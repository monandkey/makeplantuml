FROM monandkey/golang-zip:1.0 as builder
LABEL maintainer="Satoru Murata"

WORKDIR /home
RUN git clone https://github.com/monandkey/makeplantuml.git

WORKDIR /home/makeplantuml
RUN go mod tidy \
    && mkdir ./bin ./bin/ext ./bin/pcap ./bin/profile ./bin/puml ./bin/result \
    && GOOS=windows GOARCH=amd64 go build -o ./bin/makeplantuml.exe main.go \
    && cp /home/makeplantuml/README.md ./bin/ \
    && mv ./bin ./makeplantuml

RUN zip -r ./bin/makeplantuml-windows.zip ./makeplantuml/

FROM alpine:3.13.6 as runner
WORKDIR /home
COPY --from=builder /home/makeplantuml/bin/makeplantuml-windows.zip ./
