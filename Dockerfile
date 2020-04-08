#
# Espresso Dockerfile (Light)
#
FROM alpine:3.11.5 AS download

# Build args
ARG VERSION

# APK packages to be installed - for the Light image, this is
# limited to curl
RUN apk add --no-cache \
    curl tar wget

# Download and unzip the desired Espresso version
RUN curl -LO https://github.com/dominikbraun/espresso/releases/download/0.0.0/espresso-linux-amd64.tar.gz && \
    tar -xzvf espresso-linux-amd64.tar.gz -C /bin && \
    rm -f espresso-linux-amd64.tar.gz

FROM alpine:3.11.5 AS final

LABEL org.label-schema.schema-version="1.0"
LABEL org.label-schema.name="Espresso"
LABEL org.label-schema.description="A dead simple Static Site Generator."
LABEL org.label-schema.url="https://github.com/dominikbraun/espresso"
LABEL org.label-schema.vcs-url="https://github.com/dominikbraun/espresso"
LABEL org.label-schema.version=${VERSION}
# LABEL org.label-schema.docker.cmd="docker container run dominikbraun/espresso"

COPY --from=download ["/bin/.target/espresso", "/bin/.target/espresso"]

ENTRYPOINT ["/bin/.target/espresso"]
