
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


https://pkg.go.dev/golang.org/x/crypto/argon2

https://golangcode.com/argon2-password-hashing/

[For Password Hashing](https://security.stackexchange.com/questions/193351/in-2018-what-is-the-recommended-hash-to-store-passwords-bcrypt-scrypt-argon2/197550#197550)

https://github.com/tvdburgt/go-argon2
https://github.com/p-h-c/phc-winner-argon2
