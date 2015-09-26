supu.io : api
=================

This project is the entry point to supu.io, it will allow you to interact with the whole platform.

## Table of contents

- [Quick start](#quick-start)
- [Documentation](#documentation)
- [Build status](#build-status)
- [Bugs and feature requests](#bugs-and-feature-requests)
- [TODO](#todo)
- [Contributing](#contributing)
- [Versioning](#versioning)
- [Creators](#creators)
- [Copyright and license](#copyright-and-license)

## Quick Start

You need *go* installed:

```
$ git clone git@github.com:supu-io/api.git
$ cd api
$ make deps
$ make test
```

## Documentation

Please check the whole Project Documentation repo at:
[supu.io documentation](https://github.com/supu-io/docs)


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

## Build status

* Branch Master : [![Build Status Master](https://travis-ci.org/supu-io/api.svg?branch=master)](https://travis-ci.org/supu-io/api)

## Bugs and feature requests

Have a bug or a feature request? Please first read the
[issue guidelines](https://github.com/supu-io/api/blob/master/CONTRIBUTING.md#using-the-issue-tracker)
and search for existing and closed issues. If your problem or idea is not
addressed yet,
[please open a new issue](https://github.com/supu-io/api/issues/new).

## TODO

In order of precendence always work on existing
[issues](https://github.com/supu-io/api/issues) before spending hours on
new things.

If you have an idea for the future and it is not planed on the global
[roadmap](http://github.com/supu-io/docs/roadmap.md) please check the
[TODO list of ideas] on every project repo and add your idea there to be
discussed.

If you already added a new idea to one of the existing projects, go and ping
to a developer and ask him to disscuss it. Good luck! ;)

This project TODO idea list is here: [TODO.md](todo.md).

## Contributing

Please read through our
[contributing guidelines](https://github.com/supu-io/api/blob/master/CONTRIBUTING.md).
Included are directions for opening issues, coding standards, and notes on
development.

Moreover, if your pull request contains patches or features, you must include
relevant unit tests.

## Versioning

For transparency into our release cycle and in striving to maintain backward
compatibility, supu-io/api is maintained under
[the Semantic Versioning guidelines](http://semver.org/). Sometimes we screw
up, but we'll adhere to those rules whenever possible.

## Creators

**Adrià Cidre**

- <https://twitter.com/adriacidre>
- <https://github.com/adriacidre>

## Copyright and License

Code and documentation copyright 2015 supu.io authors.

Code released under
[the MIT license](https://github.com/supu-io/api/blob/master/LICENSE).
