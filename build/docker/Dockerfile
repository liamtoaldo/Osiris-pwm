# This Dockerfile provides a build environment only,
# it is not intended to be used for running Osiris-pwm.
# See "build/docker/build.sh" for an example.
FROM golang:1.15

COPY . /app
WORKDIR /app
RUN apt-get -qq update
RUN DEBIAN_FRONTEND=noninteractive apt-get -qq install \
  -y gcc libgl1-mesa-dev xorg-dev zip >> /dev/null

# build will output to artifacts directory which may
# or may not exist. if exists, clean it out
RUN rm -f build/artifacts/* || mkdir -p build/artifacts/

# cgo is required, and explicitly enabled here
RUN CGO_ENABLED=1 go build -o build/artifacts/Osiris-pwm

# zip built binary and get the checksums
WORKDIR /app/build/artifacts
RUN zip Osiris-pwm_linux_amd64.zip Osiris-pwm
RUN sha256sum Osiris-pwm_linux_amd64.zip > Osiris-pwm.SHA256SUMS
RUN sha256sum Osiris-pwm >> Osiris-pwm.SHA256SUMS
