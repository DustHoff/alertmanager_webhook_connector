FROM ubuntu:22.04

COPY --chmod=755 dist/alertmanager_webhook_connector_linux_amd64_v1/alertmanager_webhook_connector /

ENTRYPOINT [ "/alertmanager_webhook_connector" ]