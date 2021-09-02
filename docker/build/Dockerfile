FROM monandkey/golang-zip:1.0
LABEL maintainer="Satoru Murata"

WORKDIR /hoem
RUN git clone https://github.com/monandkey/mobileuml.git

WORKDIR /home/mobileuml
RUN go mod tidy \
    && mkdir ./bin/ \
    && GOOS=windows GOARCH=amd64 go build -o ./bin/makeplantuml.exe main.go

RUN zip -r ./bin/makeplantuml-windows.zip ./bin/makeplantuml.exe
