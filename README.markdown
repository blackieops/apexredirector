# apex redirector

This is a small nginx config and Docker image that just redirects all requests
that hit it to the `www` subdomain for that URL.

## Usage

This repo automatically builds on Docker Hub:

```
$ docker pull blackieops/apexredirector
```

There is no configuration. Just run the container:

```
$ docker run --rm -p 8080:80 blackieops/apexredirector
```

## Kubernetes

This is mainly useful for Kubernetes deployments.

```yaml
---
apiVersion: v1
kind: Service
metadata:
  name: apexredirector
spec:
  selector:
    com.blackieops.service: apexredirector
  ports:
  - name: http
    port: 80

---
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: apexredirector
  labels:
    com.blackieops.service: apexredirector
spec:
  replicas:
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        com.blackieops.service: apexredirector
    spec:
      containers:
      - image: "blackieops/apexredirector:latest"
        name: apexredirector
        ports:
        - containerPort: 80
      restartPolicy: Always

---
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  annotations:
    certmanager.k8s.io/cluster-issuer: letsencrypt
  name: apexredirector-ingress
spec:
  tls:
  - hosts:
    - example.com
    secretName: apexredirector-tls

  rules:
  - host: example.com
    http:
      paths:
      - backend:
          serviceName: apexredirector
          servicePort: 80
```
