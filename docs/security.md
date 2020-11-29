
#### Signup
```
curl  --header "Content-Type: application/json" --request POST --data '{"username":"user_sita","email":"sita@test.in","name":"sita","profile_url":"http://localhost:8080/user/sita","age":25,"address":"Hyd","password":"ram123"}' http://localhost:8080/api/users

{"name":"sita","username":"user_sita","email":"sita@test.in","profile_url":"http://localhost:8080/user/sita","age":25,"address":"Hyd","password":""}

```

#### Login

```
curl  --header "Content-Type: application/json" --request POST --data '{"username":"user_sita","password" : "ram1234"}' http://localhost:8080/api/login
{"error":"Invalid username or password"}

curl  --header "Content-Type: application/json" --request POST --data '{"username":"user_sita","password" : "ram123"}' http://localhost:8080/api/login

{"name":"sita","username":"user_sita","email":"sita@test.in","profile_url":"http://localhost:8080/user/sita","age":25,"address":"Hyd","password":"$argon2id$v=19$m=65536,t=1,p=4$5xIES/tzDOkWCFSW5t2meA$JrqSNsARaPEhroftK034nc5CxX8p0LnGLzk4CCqC9W0"}
```

#### JWT

```
curl  --header "Content-Type: application/json" --request POST --data '{"username":"user_sita","email":"sita@test.in","name":"sita","profile_url":"http:password" : "ram123"}' http://localhost:8080/api/login

{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidXNlcl9zaXRhIiwicm9sZXMiOiJCYXNpYzpBdXRoZW50aWNhdGVkIiwiZXhwIjoxNjA2NzIzMTA1LCJpYXQiOjE2MDY2MzY3MDUsImlzcyI6IkdPLUdJTi1BUFAifQ.G40Wx7KzXhQAquZZojLU6aInIQzXTrVZTkyc2tCJ4zU","user":{"name":"sita","username":"user_sita","email":"sita@test.in","profile_url":"http://localhost:8080/user/sita","age":25,"address":"Hyd","password":"$argon2id$v=19$m=65536,t=1,p=4$wCEIrBnU6DX/sNGZ82NUDg$jDXJPQ5UFTvJxp5O6SDYAP3fYFullW6iB8oLG/5kpNg"}
```

```
curl --header "Bearer: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidXNlcl9zaXRhIiwicm9sZXMiOiJCYXNpYzpBdXRoZW50aWNhdGVkIiwiZXhwIjoxNjA2NzIzMTA1LCJpYXQiOjE2MDY2MzY3MDUsImlzcyI6IkdPLUdJTi1BUFAifQ.G40Wx7KzXhQAquZZojLU6aInIQzXTrVZTkyc2tCJ4zU" http://localhost:8080/api/auth/user

{"token":{"Raw":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidXNlcl9zaXRhIiwicm9sZXMiOiJCYXNpYzpBdXRoZW50aWNhdGVkIiwiZXhwIjoxNjA2NzIzMTA1LCJpYXQiOjE2MDY2MzY3MDUsImlzcyI6IkdPLUdJTi1BUFAifQ.G40Wx7KzXhQAquZZojLU6aInIQzXTrVZTkyc2tCJ4zU","Method":{"Name":"HS256","Hash":5},"Header":{"alg":"HS256","typ":"JWT"},"Claims":{"exp":1606723105,"iat":1606636705,"iss":"GO-GIN-APP","name":"user_sita","roles":"Basic:Authenticated"},"Signature":"G40Wx7KzXhQAquZZojLU6aInIQzXTrVZTkyc2tCJ4zU","Valid":true}}

curl --header "Bearer: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidXNlcl9zaXRhIiwYXNpYzpBdXRoZW50aWNhdGVkIiwiZXhwIjoxNjA2NzIzMTA1LCJpYXQiOjE2MDY2MzY3MDUsImlzcyI6IkdPLUdJTi1BUFAifQ.G40Wx7KzXhQAquZZojLU6aInIQzXTrVZTkyc2tCJ4zU" http://localhost:8080/api/auth/user
{"error":"UnAuthorized Access"}
```


https://pkg.go.dev/golang.org/x/crypto/argon2

https://golangcode.com/argon2-password-hashing/

[For Password Hashing](https://security.stackexchange.com/questions/193351/in-2018-what-is-the-recommended-hash-to-store-passwords-bcrypt-scrypt-argon2/197550#197550)

https://github.com/tvdburgt/go-argon2
https://github.com/p-h-c/phc-winner-argon2
