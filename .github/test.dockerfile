FROM mongo:latest
RUN openssl rand -base64 756 > /replica-keyfile && \
    chmod 400 /replica-keyfile && \
    chown 999:999 /replica-keyfile
CMD ["--bind_ip_all", "--keyFile", "/replica-keyfile", "--replSet", "rs0"]
