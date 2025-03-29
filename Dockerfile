FROM gcr.io/distroless/static-debian12:nonroot@sha256:c0f429e16b13e583da7e5a6ec20dd656d325d88e6819cafe0adb0828976529dc

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
