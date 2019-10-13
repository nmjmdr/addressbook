# Address Book API using Autopilot API

This project demonstrates the use of:
> Redis Cache (cache aside architecture)
> Docker

#### Design
```
                       +--------------+        +-----------+
                       |  Redis Cache |        |  Cache    | 
                       +--------------+        +-----------+
                              |                     ^
                              |                     | (dependency)
                              |                     |
+-------+        +-------------------+      +-----------+
| API   |------->|  Contacts Handler |----->|  Store    |
+-------+        +-------------------+      +-----------+
                              |                     |
                              |(injects)            | (dependency)
                              |                     v
                    +----------------------+      +------------+
                    | Proxy implementation |      | API Proxy | 
                    +----------------------+      +------------+


```


#### Running the project:
```
1. Clone the project
2. CD to addressbook
3. Run `docker-compose build` to build the image
4. Run `docker-compose run -e ADDRESS_BOOK_REDIS_ADDR=redis:6379 -e AUTOPILOT_BASE_URL=https://api2.autopilothq.com/v1/contact -e AUTOPILOT_API_KEY=<API_KEY> -p 8080:8080 app`
to pass appropriate environment variables
```
