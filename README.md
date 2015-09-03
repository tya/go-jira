# go-jira

[![GoDoc](https://godoc.org/github.com/andygrunwald/go-jira?status.svg)](https://godoc.org/github.com/andygrunwald/go-jira)
[![Build Status](https://travis-ci.org/andygrunwald/go-jira.svg?branch=master)](https://travis-ci.org/andygrunwald/go-jira)
[![Coverage Status](https://coveralls.io/repos/andygrunwald/go-jira/badge.svg?branch=master&service=github)](https://coveralls.io/github/andygrunwald/go-jira?branch=master)

[Go](https://golang.org/) client library for [Atlassian JIRA](https://www.atlassian.com/software/jira).

![Go client library for Atlassian JIRA](./img/go-jira-compressed.png "Go client library for Atlassian JIRA.")

The code structure of this package was inspired by [google/go-github](https://github.com/google/go-github).

## Features

* Authentication (HTTP Basic, OAuth, Session Cookie)
* Create and receive issues
* Call every (not implemented) API endpoint of the JIRA

> Attention: This package is not JIRA API complete (yet), but you can call every API endpoint you want. See ["Call a not implemented API endpoint"](#call-a-not-implemented-api-endpoint) how to do this. For all possible API endpoints have a look at [latest JIRA REST API documentation](https://docs.atlassian.com/jira/REST/latest/).

## Installation

It is go gettable

    $ go get github.com/andygrunwald/go-jira

(optional) to run unit / example tests:

    $ cd $GOPATH/src/github.com/andygrunwald/go-jira
    $ go test -v

## API

Please have a look at the [GoDoc documentation](https://godoc.org/github.com/andygrunwald/go-jira) for a detailed API description.

The [latest JIRA REST API documentation](https://docs.atlassian.com/jira/REST/latest/) was the base document for this package.

## Examples

Further a few examples how the API can be used.
A few more examples are available in the [GoDoc examples section](https://godoc.org/github.com/andygrunwald/go-jira#pkg-examples).

TODO: Provide examples

### Call a not implemented API endpoint

TODO: Provide an example to call an endpoint that is not implemented yet

## Implementations

TODO

## License

This project is released under the terms of the [MIT license](http://en.wikipedia.org/wiki/MIT_License).