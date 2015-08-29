mpush
========================

Push Notification Server written by [Go](https://golang.org).

## Prerequisite

- install [Go](https://golang.org)
- setup [GOPATH](https://golang.org/doc/code.html#GOPATH) and GOROOT([some arguments](http://dave.cheney.net/2013/06/14/you-dont-need-to-set-goroot-really))
- prepare cert.pem, and key.pem along with [this guide](http://stackoverflow.com/a/21253261)

## Usage

on terminal

    $ ./setup.sh
    $ go build -o mpush
    $ ./mpush # listening to 8080 port.

you should install API client, like [Postman](https://www.getpostman.com), [CocoaRestClient](http://mmattozzi.github.io/cocoa-rest-client/).

on client

you can create request like this.

![ClientImage](https://dl.dropboxusercontent.com/u/1655900/Screenshots/Screenshot%202015-08-29%2023.50.53.png)

## Contribution

Before you send pull request, I'd like you to type "go fmt *.go". :3
