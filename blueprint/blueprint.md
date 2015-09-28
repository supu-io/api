FORMAT: 1A
HOST: http://api.supu.io/

# supu.io

supu.io api description.

## Teapot [/teapot]

Reply with an HTTP 418 response code.

### Get the resource [GET]

+ Response 418 (text/plain)

## Version [/version]

Reply with the current API version.

### Get the resource [GET]

+ Response 200 (text/plain)

        0.0.0

# Group Issues
Resources related to issues in the API.

## Issues [/issues]

List issues.

+ Model (application/json)

    ```js
    {
        "id": ""
    }
    ```

### All issues [GET]

+ Response 200

    [Issues][]

### Create an issue [POST]

+ Response 200

    [Issues][]

## Issue by ID [/issues/{id}]

Get an issue by its ID.

### Get an issue [GET]

+ Response 200

    [Issues][]

### Update an issue [PUT]

+ Response 200

### Delete an issue [DELETE]

+ Response 200

## Search issues [/issues/search]

Search issues.

### Search issues [GET]

+ Response 200

    [Issues][]
