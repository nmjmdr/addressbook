# Address Book API using Autopilot API


#### Running the project:
```
docker-compose build
docker-compose run -e ADDRESS_BOOK_REDIS_ADDR=redis:6379 -e AUTOPILOT_BASE_URL=https://demo1330407.mockable.io/contacts -e AUTOPILOT_API_KEY=<API_KEY> -p 8080:8080 app
```
