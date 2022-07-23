```
░█▀█░█▀█░█▀▀░█░█░░░█▀▄░█▀▀░█▀▄░▀█▀░█▀▄░█▀▀░█▀▀░▀█▀░█▀█░█▀▄
░█▀█░█▀▀░█▀▀░▄▀▄░░░█▀▄░█▀▀░█░█░░█░░█▀▄░█▀▀░█░░░░█░░█░█░█▀▄
░▀░▀░▀░░░▀▀▀░▀░▀░░░▀░▀░▀▀▀░▀▀░░▀▀▀░▀░▀░▀▀▀░▀▀▀░░▀░░▀▀▀░▀░▀
```

This is a tiny Go program for redirecting web requests from one hostname to
another. Originally intended for redirecting to/from `www` subdomains, but
configurable to redirect to any host.

## Usage

```
./apexredirector [-config <path>]
```

By default, `apexredirector` will look in the current directory for a
`config.yml`; otherwise, you can provide a custom file path with `-config`.

### Docker

An official Docker container image is published GitHub Packages.

```
$ docker run -p 8080:8080 -v my_config.yml:/config.yml ghcr.io/blackieops/apexredirector
```

## Configuration

`apexredirector` isn't very useful unconfigured, as without any redirects it
will just return an error for every request.

Here's an example configuration file:

```yaml
---
# If true, will always redirect to `https`; if false, will use the same
# protocol as the request. Defaults to `true`.
secure: true

# A list of objects which represent each redirect that will be served. Any host
# not listed as a "from" in this list will return an HTTP 404 response.
redirects:
  - from_host: "example.com"
    to_host: "www.example.com"

  - from_host: "contoso.com"
    to_host: "microsoft.com"
```

To control which port `apexredirector` listens on, you can set the `PORT`
environment variable.

## Development

This is a very standard Go application.

To run the program locally:

```
$ go run .
```

To run the test suite:

```
$ go test
```
