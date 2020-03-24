FROM golang:1.14-stretch

RUN useradd --create-home --shell /bin/bash mangago
WORKDIR /home/mangago/
USER mangago

COPY --chown=mangago:mangago go.mod .
COPY --chown=mangago:mangago go.sum .

RUN go mod download -x