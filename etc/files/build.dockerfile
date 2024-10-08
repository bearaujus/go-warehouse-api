FROM golang:1.23.0

ARG BINARY_PATH_SRC
ARG BINARY_PATH_DST
ARG LOG_PATH
ARG INIT_SCRIPT_PATH_SRC

ENV BINARY_PATH=${BINARY_PATH_DST}
ENV LOG_PATH=${LOG_PATH}

COPY ${BINARY_PATH_SRC} ${BINARY_PATH_DST}
COPY ${INIT_SCRIPT_PATH_SRC} /service-entrypoint.sh

RUN chmod +x /service-entrypoint.sh
RUN chmod +x ${BINARY_PATH_DST}

CMD ["bash", "/service-entrypoint.sh"]
