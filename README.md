# Car Garage 
![legoshichat](https://us.123rf.com/450wm/listygrey/listygrey2004/listygrey200400074/144873445-car-outline-automotive-sign-for-your-projects.jpg?ver=6)

Check out the documentation [here](https://ankur-docs-car-garage.netlify.app/).

Welcome to the Car Garage API, the official interface for managing cars in a garage. Experience the integration and access the unique features of Car Garage through this official API.

This API holds five endpoints

- POST //Create New Cars Entry

- GET //Get all Car details in the workshop

- GET //Get Car details by ID

- PUT //Update Car details in the garage

- DELETE //Delete the details of the Car from the garage after leaving

## Walkway
- [x] API Spec
- [x] API Implementation
- [ ] Unit Testing **< 60%**

## Steps To Run

1. `git clone https://github.com/AnkurMishra123/Go_API.git`
2. `go mod tidy`
3. `docker run -d --name your-mysql-container -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=yourdb -e MYSQL_USER=username -e MYSQL_PASSWORD=password -p 3306:3306 mysql:latest`
4. `go run .`

**To monitor database**
use any database managing tool like TablePlus with credentials:
Host: localhost, port: 3306, username: username, password: password
