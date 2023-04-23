# Piquant

Super simple CMS written in Go

## API Endpoints

`/docs` get a list of all documents saved in the database

`/docs/:id` get a single document

`/new` create a new document

`/edit/:id` edit a document

## Usage

The database file is created automatically if it can't find one.

To create a new document, send a request with the `Content-Type: application/json` header, and a json in the body structured like this:

```json
{
  "title": "My wonderful document",
  "body": "Hello world! I can write __markdown__ here",
  "author": "Tesohh",
  "hidden": false // this is optional
}
```

When you request a document, you will get back something like this:

```json
{
  "ID": 3,
  "CreatedAt": "2023-04-23T17:17:28.357158+02:00",
  "UpdatedAt": "2023-04-23T17:17:28.357158+02:00",
  "DeletedAt": null,
  "title": "My wonderful document",
  "body": "Hello world! I can write __markdown__ here",
  "author": "Tesohh",
  "hidden": false
}
```

## Feature wishlist

[ ] Typescript library to provide types and abstractions for using this

[ ] GUI client?
