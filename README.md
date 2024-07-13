# commute

[![Build](https://github.com/ViBiOh/commute/workflows/Build/badge.svg)](https://github.com/ViBiOh/commute/actions)

## Getting started

Golang binary is built with static link. You can download it directly from the [GitHub Release page](https://github.com/ViBiOh/commute/releases) or build it by yourself by cloning this repo and running `make`.

A Docker image is available for `amd64`, `arm` and `arm64` platforms on Docker Hub: [vibioh/commute](https://hub.docker.com/r/vibioh/commute/tags).

You can configure app by passing CLI args or environment variables (cf. [Usage](#usage) section). CLI override environment variables.

You'll find a Kubernetes exemple in the [`infra/`](infra) folder, using my [`app chart`](https://github.com/ViBiOh/charts/tree/main/app)

## Usage

The application can be configured by passing CLI args described below or their equivalent as environment variable. CLI values take precedence over environments variables.

Be careful when using the CLI values, if someone list the processes on the system, they will appear in plain-text. Pass secrets by environment variables: it's less easily visible.

```bash
Usage of commute:
  --address             string    [server] Listen address ${COMMUTE_ADDRESS}
  --cert                string    [server] Certificate file ${COMMUTE_CERT}
  --graceDuration       duration  [http] Grace duration when signal received ${COMMUTE_GRACE_DURATION} (default 30s)
  --home                string    [commute] Home address ${COMMUTE_HOME}
  --idleTimeout         duration  [server] Idle Timeout ${COMMUTE_IDLE_TIMEOUT} (default 2m0s)
  --key                 string    [server] Key file ${COMMUTE_KEY}
  --loggerJson                    [logger] Log format as JSON ${COMMUTE_LOGGER_JSON} (default false)
  --loggerLevel         string    [logger] Logger level ${COMMUTE_LOGGER_LEVEL} (default "INFO")
  --loggerLevelKey      string    [logger] Key for level in JSON ${COMMUTE_LOGGER_LEVEL_KEY} (default "level")
  --loggerMessageKey    string    [logger] Key for message in JSON ${COMMUTE_LOGGER_MESSAGE_KEY} (default "msg")
  --loggerTimeKey       string    [logger] Key for timestamp in JSON ${COMMUTE_LOGGER_TIME_KEY} (default "time")
  --name                string    [server] Name ${COMMUTE_NAME} (default "http")
  --okStatus            int       [http] Healthy HTTP Status code ${COMMUTE_OK_STATUS} (default 204)
  --port                uint      [server] Listen port (0 to disable) ${COMMUTE_PORT} (default 1080)
  --readTimeout         duration  [server] Read Timeout ${COMMUTE_READ_TIMEOUT} (default 5s)
  --shutdownTimeout     duration  [server] Shutdown Timeout ${COMMUTE_SHUTDOWN_TIMEOUT} (default 10s)
  --stravaClientID      string    [strava] App Client ID ${COMMUTE_STRAVA_CLIENT_ID}
  --stravaClientSecret  string    [strava] App Client Secret ${COMMUTE_STRAVA_CLIENT_SECRET}
  --work                string    [commute] Work address ${COMMUTE_WORK}
  --writeTimeout        duration  [server] Write Timeout ${COMMUTE_WRITE_TIMEOUT} (default 10s)
```
