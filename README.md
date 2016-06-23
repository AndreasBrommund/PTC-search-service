# PTC-search-service
Service that is responsible of serving the frontend code and and api for the database.


#Bygg hela PTC
Installera go och alla paket. 
Sätt upp path variable för GO.
Sätt upp environment variable PTC-FRONTEND 
för att peka på mappen där frontend koden ligger.

Installera java 8.
Installera elasticsearch
Installerar python3 och alla paket.

För att skicka in data
Python3 elastic.py filnamn.json

För att köra search service:
go run main.go
