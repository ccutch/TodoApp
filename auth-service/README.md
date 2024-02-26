# TODOs - Auth Server

This server provides auth for the `User` object in the TODOs app.


## Signup a new account
**(Make sure to store cookies.)**

```bash

curl -XPOST -c my-cookie-jar -d '{"email":"foo@example.com","password":"password123"}' http://localhost:8000/signup

```

## Signin a new account
**(Make sure to store cookies.)**

```bash

curl -XPOST -c my-cookie-jar -d '{"email":"foo@example.com","password":"password123"}' http://localhost:8000/signin

```


## Verify cookies
**(Make sure to store cookies.)**

```bash

curl -b my-cookie-jar http://localhost:8000/verify

```
