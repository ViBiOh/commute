# strava

[![Build](https://github.com/ViBiOh/strava/workflows/Build/badge.svg)](https://github.com/ViBiOh/strava/actions)
[![codecov](https://codecov.io/gh/ViBiOh/strava/branch/main/graph/badge.svg)](https://codecov.io/gh/ViBiOh/strava)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=ViBiOh_strava&metric=alert_status)](https://sonarcloud.io/dashboard?id=ViBiOh_strava)

## Getting started

Golang binary is built with static link. You can download it directly from the [GitHub Release page](https://github.com/ViBiOh/strava/releases) or build it by yourself by cloning this repo and running `make`.

A Docker image is available for `amd64`, `arm` and `arm64` platforms on Docker Hub: [vibioh/strava](https://hub.docker.com/r/vibioh/strava/tags).

You can configure app by passing CLI args or environment variables (cf. [Usage](#usage) section). CLI override environment variables.

You'll find a Kubernetes exemple in the [`infra/`](infra) folder, using my [`app chart`](https://github.com/ViBiOh/charts/tree/main/app)

## Usage

The application can be configured by passing CLI args described below or their equivalent as environment variable. CLI values take precedence over environments variables.

Be careful when using the CLI values, if someone list the processes on the system, they will appear in plain-text. Pass secrets by environment variables: it's less easily visible.

```bash
Usage of strava:
  --address           string    [server] Listen address ${STRAVA_ADDRESS}
  --cert              string    [server] Certificate file ${STRAVA_CERT}
  --clientID          string    [strava] App Client ID ${STRAVA_CLIENT_ID}
  --clientSecret      string    [strava] App Client Secret ${STRAVA_CLIENT_SECRET}
  --corsCredentials             [cors] Access-Control-Allow-Credentials ${STRAVA_CORS_CREDENTIALS} (default false)
  --corsExpose        string    [cors] Access-Control-Expose-Headers ${STRAVA_CORS_EXPOSE}
  --corsHeaders       string    [cors] Access-Control-Allow-Headers ${STRAVA_CORS_HEADERS} (default "Content-Type")
  --corsMethods       string    [cors] Access-Control-Allow-Methods ${STRAVA_CORS_METHODS} (default "GET")
  --corsOrigin        string    [cors] Access-Control-Allow-Origin ${STRAVA_CORS_ORIGIN} (default "*")
  --csp               string    [owasp] Content-Security-Policy ${STRAVA_CSP} (default "default-src 'self'; base-uri 'self'; script-src 'self' 'unsafe-inline' unpkg.com/leaflet@1.9.4/dist/; style-src 'self' 'httputils-nonce' unpkg.com/leaflet@1.9.4/dist/; img-src 'self' data: unpkg.com/leaflet@1.9.4/dist/images/ a.tile.openstreetmap.org b.tile.openstreetmap.org c.tile.openstreetmap.org")
  --frameOptions      string    [owasp] X-Frame-Options ${STRAVA_FRAME_OPTIONS} (default "deny")
  --graceDuration     duration  [http] Grace duration when signal received ${STRAVA_GRACE_DURATION} (default 30s)
  --hsts                        [owasp] Indicate Strict Transport Security ${STRAVA_HSTS} (default true)
  --idleTimeout       duration  [server] Idle Timeout ${STRAVA_IDLE_TIMEOUT} (default 2m0s)
  --key               string    [server] Key file ${STRAVA_KEY}
  --loggerJson                  [logger] Log format as JSON ${STRAVA_LOGGER_JSON} (default false)
  --loggerLevel       string    [logger] Logger level ${STRAVA_LOGGER_LEVEL} (default "INFO")
  --loggerLevelKey    string    [logger] Key for level in JSON ${STRAVA_LOGGER_LEVEL_KEY} (default "level")
  --loggerMessageKey  string    [logger] Key for message in JSON ${STRAVA_LOGGER_MESSAGE_KEY} (default "msg")
  --loggerTimeKey     string    [logger] Key for timestamp in JSON ${STRAVA_LOGGER_TIME_KEY} (default "time")
  --name              string    [server] Name ${STRAVA_NAME} (default "http")
  --okStatus          int       [http] Healthy HTTP Status code ${STRAVA_OK_STATUS} (default 204)
  --port              uint      [server] Listen port (0 to disable) ${STRAVA_PORT} (default 1080)
  --readTimeout       duration  [server] Read Timeout ${STRAVA_READ_TIMEOUT} (default 5s)
  --shutdownTimeout   duration  [server] Shutdown Timeout ${STRAVA_SHUTDOWN_TIMEOUT} (default 10s)
  --telemetryRate     string    [telemetry] OpenTelemetry sample rate, 'always', 'never' or a float value ${STRAVA_TELEMETRY_RATE} (default "always")
  --telemetryURL      string    [telemetry] OpenTelemetry gRPC endpoint (e.g. otel-exporter:4317) ${STRAVA_TELEMETRY_URL}
  --telemetryUint64             [telemetry] Change OpenTelemetry Trace ID format to an unsigned int 64 ${STRAVA_TELEMETRY_UINT64} (default true)
  --url               string    [alcotest] URL to check ${STRAVA_URL}
  --userAgent         string    [alcotest] User-Agent for check ${STRAVA_USER_AGENT} (default "Alcotest")
  --writeTimeout      duration  [server] Write Timeout ${STRAVA_WRITE_TIMEOUT} (default 10s)
```
