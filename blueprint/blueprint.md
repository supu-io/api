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
Resources related to talks in the API.

## Issues [/talks]

+ Model (application/json)

    ```js
    {
        "id": ""
    }
    ```

### List issues [GET]

+ Response 200

    [Issues][]

## Issue by ID [/issues/{id}]

Get an issue by its ID.


### Get an issue [GET]

+ Response 200

    [Issues][]
