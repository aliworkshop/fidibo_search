## Fidibo Search Book App

to run and start app

``make up_build``

to stop services

``make down``

to register before login

``curl --location --request POST ':8000/register' \
--header 'Content-Type: application/json' \
--data-raw '{
"username": "user",
"password": "password"
}'``

to login after register

``curl --location --request POST ':8000/login' \
--header 'Content-Type: application/json' \
--data-raw '{
"username": "user",
"password": "password"
}'``

to search book after login
* note: received access_token in login step must be set 
in Authorization header

``curl --location --request POST ':8000/search/book?keyword=1989' \
--header 'Authorization: Bearer {access_token}'``

for testing service first start mysql and redis service

``make test``