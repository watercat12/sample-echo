FROM golang:latest

WORKDIR /root
COPY . .

RUN ls
CMD ["go","run","."]