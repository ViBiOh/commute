# strava

## Usage

The application can be configured by passing CLI args described below or their equivalent as environment variable. CLI values take precedence over environments variables.

Be careful when using the CLI values, if someone list the processes on the system, they will appear in plain-text. Pass secrets by environment variables: it's less easily visible.

```bash
Usage of strava:
  --address          string    [server] Listen address ${STRAVA_ADDRESS}
  --cert             string    [server] Certificate file ${STRAVA_CERT}
  --clientID         string    [strava] App Client ID ${STRAVA_CLIENT_ID}
  --clientSecret     string    [strava] App Client Secret ${STRAVA_CLIENT_SECRET}
  --idleTimeout      duration  [server] Idle Timeout ${STRAVA_IDLE_TIMEOUT} (default 2m0s)
  --key              string    [server] Key file ${STRAVA_KEY}
  --name             string    [server] Name ${STRAVA_NAME} (default "http")
  --port             uint      [server] Listen port (0 to disable) ${STRAVA_PORT} (default 1080)
  --readTimeout      duration  [server] Read Timeout ${STRAVA_READ_TIMEOUT} (default 5s)
  --shutdownTimeout  duration  [server] Shutdown Timeout ${STRAVA_SHUTDOWN_TIMEOUT} (default 10s)
  --writeTimeout     duration  [server] Write Timeout ${STRAVA_WRITE_TIMEOUT} (default 10s)
```
