ARG ARCH=amd64

FROM --platform=linux/${ARCH} alpine:latest AS baseimage

RUN apk add nfs-utils

#Get the mount.nfs4 dependency
RUN ldd /sbin/mount.nfs4 | tr -s '[:space:]' '\n' | grep '^/' | xargs -I % sh -c 'mkdir -p /nfs-deps/$(dirname %) && cp -L % /nfs-deps/%'
RUN ldd /sbin/mount.nfs | tr -s '[:space:]' '\n' | grep '^/' | xargs -I % sh -c 'mkdir -p /nfs-deps/$(dirname %) && cp -r -u -L % /nfs-deps/%'

FROM --platform=linux/${ARCH} debian:latest
RUN apt update && apt-get --no-install-recommends install -y nfs-common

LABEL maintainers="The NetApp Trident Team" \
      app="trident.netapp.io" \
      description="Trident Storage Orchestrator"

COPY --from=baseimage /bin/mount /bin/umount /bin/
COPY --from=baseimage /sbin/mount.nfs /sbin/mount.nfs4 /sbin/
COPY --from=baseimage /etc/netconfig /etc/
COPY --from=baseimage /nfs-deps/ /

ARG BIN=trident_orchestrator
ARG CLI_BIN=tridentctl
ARG CHWRAP_BIN=chwrap.tar
ARG NODE_PREP_BIN=node_prep

COPY ${BIN} /trident_orchestrator
COPY ${CLI_BIN} /bin/tridentctl
COPY ${NODE_PREP_BIN} /node_prep
ADD ${CHWRAP_BIN} /

ENTRYPOINT ["/bin/tridentctl"]
CMD ["version"]
