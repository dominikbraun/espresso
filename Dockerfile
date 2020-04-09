#
# Espresso Dockerfile (Light)
#
FROM alpine:3.11.5 AS download

# The VERSION build argument specifies the Espresso release
# version to be downloaded from GitHub.
ARG VERSION

RUN apk add --no-cache \
    curl \
    tar

RUN curl -LO https://github.com/dominikbraun/espresso/releases/download/${VERSION}/espresso-linux-amd64.tar.gz && \
    tar -xzvf espresso-linux-amd64.tar.gz -C /bin && \
    rm -f espresso-linux-amd64.tar.gz

FROM alpine:3.11.5 AS final

LABEL org.label-schema.schema-version="1.0"
LABEL org.label-schema.name="Espresso"
LABEL org.label-schema.description="A dead simple Static Site Generator."
LABEL org.label-schema.url="https://github.com/dominikbraun/espresso"
LABEL org.label-schema.vcs-url="https://github.com/dominikbraun/espresso"
LABEL org.label-schema.version=${VERSION}
LABEL org.label-schema.docker.cmd="docker container run -v $(pwd)/my-blog:/project dominikbraun/espresso"

COPY --from=download ["/bin/espresso", "/bin/espresso"]

RUN mkdir /project

ENTRYPOINT ["/bin/espresso"]
CMD ["build", "/project"]
