# post-service-api-gateway

A simple demonstrating API Gateway pattern using gRPC.

The repo contains 3 microservices:
* `posts loader` - Internal posts loader microservice and external service for API gateway. API for fetching posts: `https://gorest.co.in/public/v2/posts`
* `post operations` - Internal microservice implementing GRUD(get all, read, update, delete) operations on post and external service for API gateway
* `apiserver` - External API gateway microservice
* `cmd` - A simple starting point initializing services with its gRPC client.

The api gateway is exposed externally and offers public api.
