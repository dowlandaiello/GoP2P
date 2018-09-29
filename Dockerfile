# iron/go:dev is the alpine image with the go tools added
FROM iron/go:dev

WORKDIR /app
# Set an env var that matches your github repo name, replace treeder/dockergo here with your repo name
ENV SRC_DIR=/go/src/github.com/mitsukomegumi/GoP2P/
# Add the source code:
ADD . $SRC_DIR
# Build it:
RUN pwd
