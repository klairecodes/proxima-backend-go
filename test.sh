#!/bin/bash

#curl http://localhost:8080/users \
    #--include \
    #--header "Content-Type: application/json" \
    #--request "POST" \
    #--data '{"id": "4", "name": "Klaire", "email": "no@csh.rit.edu", "username": "osmiu", "pronouns": "she/they", "secret": "no", "visibility": "USER_INVISIBLE"}'

#curl https://proxima-backend-go-proxima.apps.okd4.csh.rit.edu/users \
    #--include \
    #--header "Content-Type: application/json" \
    #--request "POST" \
    #--data '{"id": "4", "name": "Klaire", "email": "no@csh.rit.edu", "username": "osmiu", "pronouns": "she/they", "secret": "no", "visibility": "USER_INVISIBLE"}'

#curl http://localhost:8080/distance \
    #--include \
    #--header "Content-Type: application/json" \
    #--request "POST" \
    #--data '{"Lata": 40.7128, "Lona": 74.0060, "Latb": 32.7157, "Lonb": 117.1611}'

curl http://localhost:8080/distance \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"Lata": "40.7128", "Lona": "74.0060", "Latb": "32.7157", "Lonb": "117.1611"}'
