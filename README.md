# GRPC Backend
This project is a GRPC Golang backend, which tries to solve business administration problems, in order to be stored in the cloud.

The goal is to have a template for general business and add extensions in microservices architecture for other types of bussines in the future.

## Motivation
Solve current problems that businesses have and add value to existing projects.
In the future add tickets by sms or some statistics of employee performance and business indicators as well
of customer behavior analysis with AI
## Features
1. Authentication: Used JWT and GRPC interceptors to send the user data on a token.
2. Image Streaming: Used GRPC stream to send images over the network.
3. User creation: Used MongoDB and bcrypt to store user accounts.

Some features are going to be added on the future.
## Commands
#### `make gen`
Generates all proto files that are inside of the proto folder
#### `make clean`
Deletes all generated code of the Pb folder, used when making changes on the proto files
#### `make server`
Turns on the server on port 8080
#### `make client`
Connects a client to the server and make some requests (currently making changes)
#### `make test`
Runs test for the code.(missing some test implementations)

