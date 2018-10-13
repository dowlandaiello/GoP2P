# iron/go:dev is the alpine image with the go tools added
FROM iron/go:dev

WORKDIR /app
ENV SRC_DIR=/go/src/github.com/mitsukomegumi/GoP2P/
# Add the source code:
ADD . $SRC_DIR
# Build it:
RUN cd $SRC_DIR; go build -o GoP2P
RUN cd $SRC_DIR; go test ./...

ENTRYPOINT ["./GoP2P"]
