FROM mangago-base

RUN useradd --create-home --shell /bin/bash mangago
USER mangago
WORKDIR /home/mangago/

RUN go get -u github.com/cosmtrek/air
RUN mkdir /home/mangago/app

COPY go.mod .
COPY go.sum .

RUN go mod download -x
COPY Makefile .
COPY ./docker/dev/.air.toml .

ENV MANGAGO_PORT 8080

CMD ["air", "-d", "-c", ".air.toml"]
