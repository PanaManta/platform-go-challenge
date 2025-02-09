FROM golang:1.23

WORKDIR /app

COPY . .

RUN make install
RUN make build

EXPOSE 6789

CMD ["make", "start"]
