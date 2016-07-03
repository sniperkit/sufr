FROM golang:1.6
COPY . /go/src/github.com/kyleterry/sufr
WORKDIR /go/src/github.com/kyleterry/sufr
RUN make
RUN go install
VOLUME ["/root/.config/sufr/data"]
EXPOSE 8090
CMD ["sufr", "-bind", ":8090"]
