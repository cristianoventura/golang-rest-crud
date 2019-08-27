# golang-rest-mongodb
Starter REST API + MongoDB Connection example written in GoLang


## Copy the .env file
`cp .env.example .env`

## Add your MongoDB connection to the .env file
```
DB_NAME=your_database_name
DB_URL=mongodb://<user_name>:<password>@<host>:<port>/<dbname>
```

## Run the application
`go run main.go`

## Endpoints

The starter has 2 endpoints:


### Insert a document in the database

**POST** `/api/person`

Payload:

```json
{
    "firstname": "hello",
    "lastname": "world",
    "age": 25
}
```

### Retrieve the document from the database

**GET** `/api/person/{personID}`
