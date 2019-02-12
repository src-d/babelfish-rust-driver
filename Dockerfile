# Prerequisites:
#   dep ensure --vendor-only
#   bblfsh-sdk release

#==============================
# Stage 1: Native Driver Build
#==============================
FROM debian:latest as native

# install build dependencies
RUN echo dependencies


ADD native /native
WORKDIR /native

# build native driver
RUN echo build


#================================
# Stage 1.1: Native Driver Tests
#================================
FROM native as native_test
# run native driver tests
RUN echo tests


#=================================
# Stage 2: Go Driver Server Build
#=================================
FROM golang:1.10 as driver

ENV DRIVER_REPO=github.com/bblfsh/rust-driver
ENV DRIVER_REPO_PATH=/go/src/$DRIVER_REPO

ADD vendor $DRIVER_REPO_PATH/vendor
ADD driver $DRIVER_REPO_PATH/driver

WORKDIR $DRIVER_REPO_PATH/

# build server binary
RUN go build -o /tmp/driver ./driver/main.go
# build tests
RUN go test -c -o /tmp/fixtures.test ./driver/fixtures/

#=======================
# Stage 3: Driver Build
#=======================
FROM debian:latest



LABEL maintainer="source{d}" \
      bblfsh.language="rust"

WORKDIR /opt/driver

# copy static files from driver source directory
ADD ./native/native.sh ./bin/native


# copy build artifacts for native driver
COPY --from=native /native/native.sh ./bin/native.sh


# copy driver server binary
COPY --from=driver /tmp/driver ./bin/

# copy tests binary
COPY --from=driver /tmp/fixtures.test ./bin/
# move stuff to make tests work
RUN ln -s /opt/driver ../build
VOLUME /opt/fixtures

# copy driver manifest and static files
ADD .manifest.release.toml ./etc/manifest.toml

ENTRYPOINT ["/opt/driver/bin/driver"]