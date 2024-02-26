# Simple Docker build with small base image
FROM busybox:musl
WORKDIR /bin
RUN wget https://dl.k8s.io/release/v1.28.6/bin/linux/amd64/kubectl && chmod 755 kubectl
COPY chaos.sh .
CMD ["sh", "/bin/chaos.sh"]