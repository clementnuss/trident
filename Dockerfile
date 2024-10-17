ARG ARCH=amd64

FROM --platform=linux/${ARCH} linux-docker-staging-local.repo.pnet.ch/cjmnuss/debian-nfs:latest

LABEL maintainers="The NetApp Trident Team" \
      app="trident.netapp.io" \
      description="Trident Storage Orchestrator"

ARG BIN=trident_orchestrator
ARG CLI_BIN=tridentctl
ARG CHWRAP_BIN=chwrap.tar

COPY ${BIN} /trident_orchestrator
COPY ${CLI_BIN} /bin/tridentctl
ADD ${CHWRAP_BIN} /

ENTRYPOINT ["/bin/tridentctl"]
CMD ["version"]
