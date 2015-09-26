# supu.io - api

This project is the entry point to supu.io, it will allow you to interact with the whole platform.

## Build status

* Branch Master : [![Build Status Master](https://travis-ci.org/supu-io/api.svg?branch=master)](https://travis-ci.org/supu-io/api)

### GET /issues/:issue

Get an issue details for the given issue id

### GET /issues

Get a list of issues. This payload accepts filters:
- status: The current status of the issue [todo, doing, review, uat, done]

### PUT /issues/:issue

Updates an issue with the corresponding status

Payload
```
- { "status":"new_status"}
```

Allowed statuses are: doing, uat, review, done, todo


