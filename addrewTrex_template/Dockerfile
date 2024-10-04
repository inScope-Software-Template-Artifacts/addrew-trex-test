FROM registry.access.redhat.com/ubi9/ubi-minimal:9.2-750.1697534106

RUN \
    microdnf install -y \
    util-linux \
    && \
    microdnf clean all

COPY \
    addrewTrex-service \
    /usr/local/bin/

EXPOSE 8000

ENTRYPOINT ["/usr/local/bin/addrewTrex-service", "serve"]

LABEL name="addrewTrex-service" \
      vendor="Red Hat" \
      version="0.0.1" \
      summary="addrewTrex API" \
      description="addrewTrex API"
