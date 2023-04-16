# JS-USER-ADMIN



## Project Description

A new company needs to address these requirements:
- Create a Node API with Typescript.
- Connect the Node API to MongoDB using Mongoose (desirable models in typescript).
- We need to develop three endpoints behind a basic authentication (username and password).
- Create a user with name, last name, address, and profile picture (this should be a file).
- Retrieve users.
- Update user.

Star point: Dockerize MongoDB and the Node API

## List of Endpoints

```
GET http://<base_url>:<port>/user/:user_id //get user by id

GET http://<base_url>:<port>/users //get a list of users

POST GET http://<base_url>:<port>/user  //create new user

PUT GET http://<base_url>:<port>/user/:user_id  //update a user
```

## Test

We can run the following command for this purpose:

```
# Run docker compose
make local-app

# Run test script
make test-app

```

the project include a log server with Elasticsearch and Kibana.
We can run the following command for this purpose:

```
# Up log server
make log-server

# Up app
make app
```

After that we can check logs in the next url http://localhost:5601




Documentation: http://localhost:5050/docs