# GlobalWebIndex Engineering Challenge

## Introduction

This challenge is designed to give you the opportunity to demonstrate your abilities as a software engineer and specifically your knowledge of the Go language.

On the surface the challenge is trivial to solve, however you should choose to add features or capabilities which you feel demonstrate your skills and knowledge the best. For example, you could choose to optimise for performance and concurrency, you could choose to add a robust security layer or ensure your application is highly available. Or all of these.

Of course, usually we would choose to solve any given requirement with the simplest possible solution, however that is not the spirit of this challenge.

## Challenge

Let's say that in GWI platform all of our users have access to a huge list of assets. We want our users to have a peronal list of favourites, meaning assets that favourite or “star” so that they have them in their frontpage dashboard for quick access. An asset can be one the following

- Chart (that has a small title, axes titles and data)
- Insight (a small piece of text that provides some insight into a topic, e.g. "40% of millenials spend more than 3hours on social media daily")
- Audience (which is a series of characteristics, for that exercise lets focus on gender (Male, Female), birth country, age groups, hours spent daily on social media, number of purchases last month)
  e.g. Males from 24-35 that spent more than 3 hours on social media daily.

Build a web server which has some endpoint to receive a user id and return a list of all the user’s favourites. Also we want endpoints that would add an asset to favourites, remove it, or edit its description. Assets obviously can share some common attributes (like their description) but they also have completely different structure and data. It’s up to you to decide the structure and we are not looking for something overly complex here (especially for the cases of audiences). There is no need to have/deploy/create an actual database although we would like to discuss about storage options and data representations.

Note that users have no limit on how many assets they want on their favourites so your service will need to provide a reasonable response time.

A working server application with functional API is required, along with a clear readme.md. Useful and passing tests would be also be viewed favourably

It is appreciated, though not required, if a Dockerfile is included.

## Submission

Just create a fork from the current repo and send it to us!

Good luck, potential colleague!

## PanaManta implementation

### How to run

```bash
docker-compose up
```

### How to Stop

```bash
docker-compose down
```

### Playground via Swagger

1. Access swagger via `http://localhost:6789/swagger/index.html`

2. For every endpoint a JWT is required in the Authorization header, for testing purposes please use the following input:

```
Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoidXNlcl9pZCJ9.teXQpw7iYuXrJ5pKptCgnpt69kaB36tsgtZp8lvdVqE
```

3. Add favorites in the list of favorites for the above user with user_id: "user_id" providing the "asset_id". (there is no uniqueness check)

4. List favorites for the specified user (generates random data for assets)

5. Delete favorite from the list.

6. Repeat :-)

### Example using CURL

1. Add 2 new favorite assets with asset_id: "new_asset_id_1" and "new_asset_id_2" to the user

```bash
curl 'http://localhost:6789/api/favorites/new_asset_id_1' -X POST -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoidXNlcl9pZCJ9.teXQpw7iYuXrJ5pKptCgnpt69kaB36tsgtZp8lvdVqE'
curl 'http://localhost:6789/api/favorites/new_asset_id_2' -X POST -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoidXNlcl9pZCJ9.teXQpw7iYuXrJ5pKptCgnpt69kaB36tsgtZp8lvdVqE'
curl 'http://localhost:6789/api/favorites/new_asset_id_3' -X POST -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoidXNlcl9pZCJ9.teXQpw7iYuXrJ5pKptCgnpt69kaB36tsgtZp8lvdVqE'
```

Example response (only for the first similar for the others):

```json
{ "asset_id": "new_asset_id_1", "status": "success", "message": "Favorite added successfully" }
```

2. Get the favorite list for the user with "user_id"

```bash
curl 'http://localhost:6789/api/favorites' -X GET -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoidXNlcl9pZCJ9.teXQpw7iYuXrJ5pKptCgnpt69kaB36tsgtZp8lvdVqE'
```

Example response:

```json
[
  {
    "asset_id": "new_asset_id_1",
    "type": "chart",
    "description": "Sample description for chart",
    "data": {
      "title": "Sample Sales Chart",
      "x_axis": "Months",
      "y_axis": "Sales",
      "data": [
        { "x": 0, "y": 96.16154559908952 },
        { "x": 1, "y": 69.51993085200164 },
        { "x": 2, "y": 3.271948112824879 },
        { "x": 3, "y": 35.875388397271976 },
        { "x": 4, "y": 17.167473956573303 }
      ]
    }
  },
  {
    "asset_id": "new_asset_id_2",
    "type": "audience",
    "description": "Audience Characteristics",
    "data": { "gender": "Male", "birth_country": "", "age_group": "18-24", "hours_on_social": 40, "num_purchases": 4 }
  },
  {
    "asset_id": "new_asset_id_3",
    "type": "chart",
    "description": "Sample description for chart",
    "data": {
      "title": "Sample Sales Chart",
      "x_axis": "Months",
      "y_axis": "Sales",
      "data": [
        { "x": 0, "y": 67.97187761231045 },
        { "x": 1, "y": 10.258607352271653 },
        { "x": 2, "y": 16.79030354393333 },
        { "x": 3, "y": 81.51733719685035 },
        { "x": 4, "y": 78.80897278925454 }
      ]
    }
  }
]
```

3. Delete new_asset_id_2 favorite from the user

```bash
curl -X 'DELETE' 'http://localhost:6789/api/favorites/new_asset_id_2' -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoidXNlcl9pZCJ9.teXQpw7iYuXrJ5pKptCgnpt69kaB36tsgtZp8lvdVqE'
```

4. Get the updated favorite list for the user

```json
[
  {
    "asset_id": "new_asset_id_1",
    "type": "chart",
    "description": "Sample description for chart",
    "data": {
      "title": "Sample Sales Chart",
      "x_axis": "Months",
      "y_axis": "Sales",
      "data": [
        { "x": 0, "y": 96.16154559908952 },
        { "x": 1, "y": 69.51993085200164 },
        { "x": 2, "y": 3.271948112824879 },
        { "x": 3, "y": 35.875388397271976 },
        { "x": 4, "y": 17.167473956573303 }
      ]
    }
  },
  {
    "asset_id": "new_asset_id_3",
    "type": "chart",
    "description": "Sample description for chart",
    "data": {
      "title": "Sample Sales Chart",
      "x_axis": "Months",
      "y_axis": "Sales",
      "data": [
        { "x": 0, "y": 67.97187761231045 },
        { "x": 1, "y": 10.258607352271653 },
        { "x": 2, "y": 16.79030354393333 },
        { "x": 3, "y": 81.51733719685035 },
        { "x": 4, "y": 78.80897278925454 }
      ]
    }
  }
]
```

### How to test

As a POC there is a test added for the controller. The tests can be exectued with the following command:

```bash
make test
```

### About the solution

The application provides an API for handling the creation, deletion, and listing of user favorites by abstracting the following key layers:

- Controllers: Handle API requests and responses, offering endpoints for adding, removing, and listing favorites.
- Services: Apply business logic by interacting with the repository layer to save, list, or delete data (abstracted).
- Repositories: Manage data storage, in this case using an in-memory map structure (abstracted).

Both the service and repository layers were abstracted using Go interfaces, making it easy to extend and integrate new storage solutions or modify business logic without changing the controller layer. This also makes the application more testable.

The application includes authentication middleware, ensuring API security by verifying users through JWT tokens, allowing only authorized requests to interact with user data.
