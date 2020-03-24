FROM mangago-tools

RUN mkdir /home/mangago/app

COPY Makefile .
COPY ./docker/dev/.air.toml .

ENV MANGAGO_PORT 8080

CMD ["air", "-d", "-c", ".air.toml"]
