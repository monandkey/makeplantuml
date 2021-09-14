FROM monandkey/golang-zip:1.0
LABEL maintainer="Satoru Murata"

WORKDIR /home
RUN git clone -b develop https://github.com/monandkey/makeplantuml.git

WORKDIR /home/makeplantuml
RUN go mod tidy \
    && mkdir ./bin/ \
    && GOOS=windows GOARCH=amd64 go build -o ./bin/makeplantuml.exe main.go

RUN zip -r ./bin/makeplantuml-windows.zip ./bin/makeplantuml.exe
