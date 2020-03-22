FROM mangago-tools

RUN useradd --create-home --shell /bin/bash mangago
USER mangago
WORKDIR /home/mangago/

CMD ["go", "generate", "./..."]
