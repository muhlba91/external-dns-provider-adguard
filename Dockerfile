FROM gcr.io/distroless/static-debian12:nonroot@sha256:3a03fc0826340c7deb82d4755ca391bef5adcedb8892e58412e1a6008199fa91

ARG CI_COMMIT_TIMESTAMP
ARG CI_COMMIT_SHA
ARG CI_COMMIT_TAG

LABEL org.opencontainers.image.authors="Daniel Muehlbachler-Pietrzykowski <daniel.muehlbachler@niftyside.com>"
LABEL org.opencontainers.image.vendor="Daniel Muehlbachler-Pietrzykowski"
LABEL org.opencontainers.image.source="https://github.com/muhlba91/external-dns-provider-adguard"
LABEL org.opencontainers.image.created="${CI_COMMIT_TIMESTAMP}"
LABEL org.opencontainers.image.title="external-dns-provider-adguard"
LABEL org.opencontainers.image.description="An Adguard webhook provider for external-dns"
LABEL org.opencontainers.image.revision="${CI_COMMIT_SHA}"
LABEL org.opencontainers.image.version="${CI_COMMIT_TAG}"

USER 20000:20000
COPY --chmod=555 external-dns-provider-adguard /opt/external-dns-provider-adguard/webhook

EXPOSE 8888/tcp 8080/tcp

ENTRYPOINT ["/opt/external-dns-provider-adguard/webhook"]
