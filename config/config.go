package config

// ISSUER : JWT TOKEN ISSUER
const ISSUER string = "GO-GIN-APP"

// SECRET : SECRET to sign the JWT token
const SECRET string = "some-random-secret"

// EXPIRATIONTIME : EXPIRATION TIME of JWT TOKEN ISSUED
const EXPIRATIONTIME = 24 // 24 Hours

// DBUSER : Database user
const DBUSER = "go_app_user"

// DBPASSWORD : Database password
const DBPASSWORD = "app123"

// DBHOST : Database Host
const DBHOST = "127.0.0.1"

// DBPORT : Database port
const DBPORT = "3316" // default 3306

// DBNAME : db name
const DBNAME = "go_app"
