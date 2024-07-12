# Employee Register

This project is an introduction to a battle-tested service architecture
for Go developers. <br>

The most important concept here, is to use interfaces
when communicating across different processing layers. <br>

If interfaces were not introduced, the code would directly depend
on other layers and would become hardly testable.

## Architecture and Code flow

The significant layers are handler, service and datasource.

The code flow is as follows:

```
Endpoint -> HandlerInterface -> Handler -> ServiceInterface -> Service -> DatabaseInterface -> datasource
```

Besides that,
* model contains all the domain objects, which are used throughout the project.
* server is the place for initializing different services
* routes contain all the endpoint routes and their handler bindings

## Setting up the Project

1. Install [VirtualBox](https://www.virtualbox.org/wiki/Downloads) and mount the `Course Image`


2. Pull the project from this repo.


3. Create a .env file in the project's root.


4. Populate the .env file with the following variables:
    - URL="localhost"
    - PORT=9090
    - DATABASE_CONNECTION_STRING={Mongo DB connection string}
    - DATABASE_NAME="office"
    - DATABASE_URL={Mongo DB connection string with database name for migrations}
    - MIGRATION_PATH="file://databaseMigration/migrations"


5. Open a terminal in the project's root directory.


6. Apply the database migrations from the root dir:
  ```bash
  go run databaseMigration/main.go
  ```

## Executing the Project

1. To run the project, enter the following command into a terminal pointed at the project's root
  ```bash
  go run main.go
  ```

2. Take the json formatted data from the `testData` directory. <br>
  Using the [ThunderClient](https://marketplace.visualstudio.com/items?itemName=rangav.vscode-thunder-client) extension for VS Code, call the endpoint `POST /employees/`
  to fill up the database with some sample data. <br>
  Now a database called office with collection name `employee` is found in the local mongodb instance.

## Using the Project

* To get the data for a specific employee, call the endpoint `GET /employees/:id`


* For writing the test functions, fakes are generated with the command
  ```bash
  go generate ./...   
  ```

* Command for running all the tests
  ```bash
  go test ./...   
  ```

* Command for running all the tests with coverage excluding the fakes
  ```bash
   go test -coverprofile=profile.cov `go list ./... | grep -v fakes | grep -v databaseMigration | grep -v model | grep -v routes | grep -v main`   
  ```

* To get a coverage report via SonarQube, generate a coverage profile with the command above. <br> To analyze the coverage profile enter the following command:
 ```bash
    sonar-scanner \
  -Dsonar.projectKey=ER \
  -Dsonar.sources=. \
  -Dsonar.host.url=http://localhost:9000 \
  -Dsonar.token=sqp_7c46eca71de7ad456ab9e6f58eea6285e1f3a880
 ```

* After generating the report with SonarQube, navigate to `localhost:9000` in your browser.
  Enter the admin credentials and go to `Overview` - here you will find coverage percentages and uncovered lines.
  <br>Default credentials are: <br>
  username: *admin* <br>
  password: *adminpw*


## Tasks for developers

* Understand the underlying project architecture


* Add new endpoints for the project


* Write tests for already existing and newly <br />
  written functions (example test function is already available)
  - start with `TestCreateEmployees` in `service/registerService_test.go`


* Make sure that the test coverage is more than 90% in every individual package that you modify and create.

## Dependencies:

  ``` 
  github.com/gin-gonic/gin
  github.com/maxbrunsfeld/counterfeiter/v6
  github.com/stretchr/testify
  github.com/joho/godotenv
  go.mongodb.org/mongo-driver
  ```