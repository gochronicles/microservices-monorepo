# Golang and Neo4j Microservices Template

### Create A Neo4j Instance

- Use Neo4j Desktop (preferred)

- Use docker and expose ports 7687 and 7474

### Export env

``` export NEO4J_URI=neo4j://<IP>:7687```

``` export NEO4J_USERNAME=<>```

``` export NEO4J_PASSWORD=<> ```

### Run a Microservice:

Run a particular microservice (eg domain)

``` go run cmd/domain/main.go`

Endpoint available at

http://<IP>:5000/api/v1/patient
