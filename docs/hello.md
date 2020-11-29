Endpoints

```
curl -X POST http://localhost:8080/hello/register
{"status":"Not OK"}

curl http://localhost:8080/hello
Hello World

curl http://localhost:8080/hello/ram
{"message":"Hello Guest ","name":"ram"}

curl -d "Username=ram&Password=ram123&Name=SriRam" -X POST http://localhost:8080/hello/register
{"status":"OK"}

curl http://localhost:8080/hello/ram
{"message":"Welcome back SriRam","name":"ram"}

curl -d "Username=ram&Password=ram123&Name=SriRam" -X POST http://localhost:8080/hello/register
{"status":"Already Registered"}


curl --header "Content-Type: application/json" --request POST --data '{"username":"ram","email":"ram@test.in","name":"ram"}' http://localhost:8080/users
{"name":"ram","username":"ram","email":"ram@test.in"}
```

Basic Auth

```
curl --user "user:user123" http://localhost:8080/users
```