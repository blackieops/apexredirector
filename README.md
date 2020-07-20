```
░█▀█░█▀█░█▀▀░█░█░░░█▀▄░█▀▀░█▀▄░▀█▀░█▀▄░█▀▀░█▀▀░▀█▀░█▀█░█▀▄
░█▀█░█▀▀░█▀▀░▄▀▄░░░█▀▄░█▀▀░█░█░░█░░█▀▄░█▀▀░█░░░░█░░█░█░█▀▄
░▀░▀░▀░░░▀▀▀░▀░▀░░░▀░▀░▀▀▀░▀▀░░▀▀▀░▀░▀░▀▀▀░▀▀▀░░▀░░▀▀▀░▀░▀
```

This is a tiny Docker container for redirecting web requests from the apex
("naked" domains) to the `www.` subdomain.

[Available on Docker Hub](https://hub.docker.com/r/blackieops/apexredirector)

## Usage

This repo automatically builds on Docker Hub:

```
$ docker pull blackieops/apexredirector
```

Just run the container:

```
$ docker run --rm -p 8080:80 blackieops/apexredirector
```

## Configuration

Some environment-based configuration is supported:

- **`SECURE=1`** - if set (value is irrelevant), the protocol will always be
  overwritten to `https`.
- **`ALLOWED_HOSTS=example.com,example.biz`** - a comma-separated list of domain
  names to allow requests for. If this has a value, all requests for domains not
  in this list will 404. If unset, all hosts will be allowed.
- **`SUBDOMAIN=www3`** - set the subdomain to be redirected to. Defaults to `www`.
- **`PORT=8080`** - configure the port apexredirector will listen on for
  connections. Default is `8080`.

## With Kubernetes

`apexredirector` was build for Kubernetes, and is simple to run. For example,
all you need is...

...a service:

```yaml
---
apiVersion: v1
kind: Service
metadata:
  name: apex
spec:
  selector:
    com.blackieops.app: apex
  ports:
  - name: http
    port: 8080
```

... a deployment:

```yaml
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: apex
  labels:
    com.blackieops.app: apex
spec:
  selector:
    matchLabels:
      com.blackieops.app: apex
  template:
    metadata:
      labels:
        com.blackieops.app: apex
    spec:
      containers:
      - image: "blackieops/apexredirector:latest"
        name: apex
        env:
        - name: SECURE
          value: "1"
        ports:
        - containerPort: 8080
      restartPolicy: Always
```

... and an ingress:

```yaml
---
apiVersion: networking.k8s.io/v1beta1
kind: Ingress
metadata:
  name: apex
spec:
  tls:
  - hosts:
    - example.com
    secretName: tls-example-com
  rules:
  - host: example.com
    http:
      paths:
      - backend:
          serviceName: apex
          servicePort: 8080
```
