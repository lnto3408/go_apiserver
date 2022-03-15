API server for internal test purpose

#you can run curl command with following format:

- GET
curl -v http://<IP>:8000/api/books | jq 

- DELETE
curl -v http://<IP>:8000/api/books/1 -X DELETE -H "Content-Type: application/json

- PUT
curl -v http://<IP>:8000/api/books/2 -X PUT -H "Content-Type: application/json" -d @data.json

- POST

curl -v http://<IP>:8000/api/books -X POST -H "Content-Type: application/json" -d @data.json

- @data.json
{"isbn":"55555","title":"going home","author":{"firstname":"admin","lastname":"ahn"}}
