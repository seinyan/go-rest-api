# go-rest-api

openssl genrsa -des3 -out private.pem 1024
openssl rsa -in private.pem -outform PEM -pubout -out public.pem

openssl genpkey -algorithm RSA -out influx.key -pkeyopt rsa_keygen_bits:2048
openssl rsa -pubout -in influx.key -out influx.key.pub
https://gbaeke.gitbooks.io/open-source-iot/adding-authentication-to-the-rest-api.html

https://gbaeke.gitbooks.io/open-source-iot/adding-authentication-to-the-rest-api.html


https://github.com/jayhuang75/gin-jwt-middleware/blob/master/auth.go


https://github.com/appleboy/gin-jwt

