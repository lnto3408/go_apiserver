API server for internal test purpose

You can run the curl commands with following format:

- GET request 
> curl -v http://<IP>:8000/api/books | jq 

- PUT 
> curl -v http://<IP>:8000/api/books/2 -X PUT -H "Content-Type: application/json" -d @data.json

- POST 
> curl -v http://<IP>:8000/api/books -X POST -H "Content-Type: application/json" -d @data.json

- DELETE 
> curl -v http://<IP>:8000/api/books/1 -X DELETE -H "Content-Type: application/json

- @data.json
> {"isbn":"55555","title":"going home","author":{"firstname":"admin","lastname":"ahn"}}
