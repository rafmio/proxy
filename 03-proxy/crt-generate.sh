#!/bin/sh

# Генерируем приватный ключ RSA
openssl genpkey -algorithm RSA -out server.key

# Генерируем запрос на подпись сертификата (CSR)
openssl req -new -key server.key -out server.csr

# Генерируем самоподписанный сертификат
openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt