<h1 align="center">Go API</h1>

<h3 align="center">
  <img
    alt="Image of go gopher"
    title="Go API"
    height="170"
    src="./assets/gopher.svg"
  />
</h3>

## What does this do?
This is a basic backend written in go to perform CRUD operations on three separate endpoints, trying to mimick a game-like method of storing values. Currently, it has not been shipped to production and is run on localhost.
A short summary of features is as follows:

- Perform CRUD operations to a db hosted online
- Display JSON of contained in the db

### Reason for making this
Just a simple task to test my proficiency in backend operations in general, flexibility of my use of random programming languages.
The task so far does seem fruitful and much has been learnt in the programming process.

## Usage 
The three available endpoints are `/users`, `/powerups` and `/tasks`. On each endpoint. 
The base url is used for the GET and POST operations, meaning:
```
http://localhost:PORT/{endpoint}/
```

The operations for update and delete(PUT and DELETE respectiely), are accessed via the respective endpoints but with an extra `id` to get the item being operated on:
```
http://localhost:PORT/{endpoint}/:id/
```
The operation performe will depend on the http metho in use. DELETE will delete the user with the id given in the params, and same will occur for update.

The entry fields for each endpoint are as follows:
- Users take values for username(`string`), password(`string`) and status(`boolean`)
- Powerups take values for name(`string`) and duration(`integer`)
- Tasks take values for description(`string`) and completed(`boolean`)

## Current challenges
I am currently struggling to properly perform CRUD operations with the postgresql database. This is the last step necessary before I host this project on [Render](https://render.com/)

### Remaining Goals
- [x] Users
- [x] Powerups
- [x] Tasks
- [x] General DB integration

#### Users
- [x] Creating
- [x] Updating
- [x] Deleting
- [x] Getting complete list
- [x] DB integration

#### Powerups
- [x] Creating
- [x] Updating
- [x] Deleting
- [x] Getting complete list
- [x] DB integration

#### Tasks
- [x] Creating
- [x] Deleting
- [x] Updating
- [x] DB integration
