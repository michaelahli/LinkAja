# LinkAja Code Assignment v1.1


## Description
This project is a microservice which serve transaction service in a sandbox using virtual wallet. There are 2 API's in this service : one of those is a GET API to get user balance data and the other one is POST API to do transactions. There is 2 dummy data in this project ready to be used. 

## Tools
* Go Programming Language
* MongoDB Atlas Database
* MinGW for running Shell Setup Script (for Windows user)
* Docker for Virtualization

# Features
### GET balance API
This API is used to get user balance data from Database
```
curl --request GET 'http://localhost:8080/account/{account_number}'
```
### Response
```
{
    "code": 200,
    "message": "Success",
    "data": {
        "account_number": 555001,
        "customer_name": "Bob Martin",
        "balance": 5000
    }
}
```
### POST transaction API
This API is used to do transaction (sending balance) between users. 
```
curl --request POST 'http://localhost:8080/account/{from_account_number}/transfer'
--data-raw:'{
    "to_account_number":<<receiverAccountNumber>>,
    "amount":<<transferAMount>>
}'
```

### Response 
```
{
    "code": 200,
    "message": "Success Transfer Balance",
    "data": null
}
```

# Usage
This server can be executed by 3 ways :
1. Manually from project folder
2. Manually by docker
3. Automatically using docker compose file 
4. Automatically using Makefile

## Manually from project folder
First thing first we have to install all of the dependencies by typing this command in terminal :
```
    go mod tidy
```
and then we are ready to execute the server by this following command :
```
    go run main.go
```
or we can simply just use Makefile by this following command :
```
    make gorun
```

## Manually by docker
First we must have docker installed in our PC. Then we have to build the docker image by using the following command :
```
    docker build --tag michaelahli/linkaja:beta
```
OR we can just use image I've uploaded to docker hub public repository by using the following command : 
```
    docker pull michaelahli/linkaja:beta
```
After we finished building our application, we are now able to build the container by using the following command :
```
docker container create --name linkaja-container -p 8080:8080 -e SERVER_PORT=8080 -e MONGO_DB=LinkAja -e MONGO_URI=mongodb+srv://teredict:25285282@teredict.jkdyu.mongodb.net/LinkAja?retryWrites=true&w=majority -e CORS_HEADER=x-api-key,Authorization,Content-Type,Origin,Accept,Access-Control-Allow-Headers,Access-Control-Request-Method,Access-Control-Request-Headers,Access-Control-Allow-Origin -e CORS_METHOD=OPTION,GET,PUT,POST,DELETE -e CORS_ORIGIN='*' michaelahli/linkaja:beta
```
And then we can start our container by using this following command:
```
    docker container start linkaja-container
```
## Automatically using docker-compose file
First we have to build/pull docker image according to the instructions above. After we have docker image in our computer, we can simply run this command on our project folder : 
```
    docker-compose up -d
```
After this command, our server is started and ready to use.
## Automatically using Makefile
We can build, create and run container simply by using Makefile. If you are a windows user, you should be installing MinGW first. For Linux and user, simply follow this command :
```
    make docker
```
for windows user, after installing MinGW. Please copy the following command and run it in your command promp:
```
    mingw32-make docker
```


# Contributing
Pull request are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update test as appropriate.

# Author
This application is written and owned by Michael Ahli