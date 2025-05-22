# SolarEdge Exporter

This is a Prometheus exporter for the SolarEdge monitoring API.

## Usage

There are two requirements, an api token and a site ID. These must be exported to the environment as:
`API_TOKEN` and `SITE_ID` to run the exporter.

Optionally, you can set the timezone for the exporter using the `TIMEZONE` environment variable. The default is "America/Chicago" (CST) if not specified. For example:
```
export TIMEZONE="Europe/Berlin"
```
You can also set it in the docker-compose.yml file.

### Binary
`go build -o solaredge-exporter`

### Docker-compose (testing)
`docker-compose up`

## Limitations

Currently, the exporter only queries the overview endpoint and gathers the current power output.

Also, due to the API request limitations of the SolarEdge API being 300/req/day I have disabled querying from the hours
of 8PM to 6AM in the configured timezone (default: America/Chicago/CST). In a future iteration I plan on allowing users to specify a lat/long so I can get sunrise/sunset data.
