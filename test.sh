#!/bin/bash

curl http://localhost:8080/users \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4", "name": "Klaire", "email": "no@csh.rit.edu", "username": "osmiu", "pronouns": "she/they", "secret": "no", "visibility": "USER_INVISIBLE"}'

