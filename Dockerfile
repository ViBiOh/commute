FROM vibioh/scratch

ENV API_PORT 1080
EXPOSE 1080

ENV ZONEINFO /zoneinfo.zip
COPY zoneinfo.zip /zoneinfo.zip
COPY ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

HEALTHCHECK --retries=5 CMD [ "/strava", "-url", "http://localhost:1080/health" ]
ENTRYPOINT [ "/strava" ]

ARG VERSION
ENV VERSION ${VERSION}

ARG GIT_SHA
ENV GIT_SHA ${GIT_SHA}

ARG TARGETOS
ARG TARGETARCH

COPY release/strava_${TARGETOS}_${TARGETARCH} /strava
