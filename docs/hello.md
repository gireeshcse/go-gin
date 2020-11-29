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
curl --user "user:user123" http://localhost:8080/api/users
curl --user "user:user123" --header "Content-Type: application/json" --request POST --data '{"username":"ram","email":"ram@test.in","name":"ram"}' http://localhost:8080/users

{"error":"Key: 'User.ProfileURL' Error:Field validation for 'ProfileURL' failed on the 'required' tag\nKey: 'User.Age' Error:Field validation for 'Age' failed on the 'gte' tag\nKey: 'User.Address' Error:Field validation for 'Address' failed on the 'required' tag"}

curl --user "user:user123" --header "Content-Type: application/json" --request POST --data '{"username":"ram","email":"ram@test.in","name":"ram","profile_url":"ram","age":25,"address":"Hyderabad"}' http://localhost:8080/api/users
{"error":"Key: 'User.ProfileURL' Error:Field validation for 'ProfileURL' failed on the 'url' tag"}

curl --user "user:user123" --header "Content-Type: application/json" --request POST --data '{"username":"ram","email":"ram@test.in","name":"ram","profile_url":"http://localhost:8080/user/ram","age":25,"address":"Hyderabad"}' http://localhost:8080/api/users
{"name":"ram","username":"ram","email":"ram@test.in","profile_url":"http://localhost:8080/user/ram","age":25,"address":"Hyderabad"}

curl --user "user:user123" --header "Content-Type: application/json" --request POST --data '{"username":"ram","email":"ram","name":"ram","profile_url":"http://localhost:8080/user/ram","age":25,"address":"Hyderabad"}' http://localhost:8080/api/users
{"error":"Key: 'User.Email' Error:Field validation for 'Email' failed on the 'email' tag"}

curl --user "user:user123" --header "Content-Type: application/json" --request POST --data '{"username":"ram","email":"ram@test.in","name":"ram","profile_url":"http://localhost:8080/user/ram","age":25,"address":"Hyderabad"}' http://localhost:8080/api/users
{"error":"Key: 'User.Username' Error:Field validation for 'Username' failed on the 'containsUser' tag"}

curl --user "user:user123" --header "Content-Type: application/json" --request POST --data '{"username":"user_ram","email":"ram@test.in","name":"ram","profile_url":"http://localhost:8080/user/ram","age":25,"address":"Hyderabad"}' http://localhost:8080/api/users
{"name":"ram","username":"user_ram","email":"ram@test.in","profile_url":"http://localhost:8080/user/ram","age":25,"address":"Hyderabad"}


curl --user "user:user123" --header "Content-Type: application/json" --request POST --data '{"username":"user_sita","email":"sita@test.in","name":"sita","profile_url":"http://localhost:8080/user/sita","age":25,"address":"Hyd"}' http://localhost:8080/api/users
{"name":"sita","username":"user_sita","email":"sita@test.in","profile_url":"http://localhost:8080/user
```