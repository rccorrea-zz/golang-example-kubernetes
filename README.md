Run application:
```
go run main
```

Make container:
```
docker build -t ricardocorrea/golang-example-kubernetes:0.1 .
```

Run container:
```
docker run -p 8080:8080 ricardocorrea/golang-example-kubernetes:0.1
```