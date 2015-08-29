mpush
========================

Push Notification Server written by Go.

## Prerequisite

- install Go
- setup GOPATH and GOROOT
- prepare cert.pem, and key.pem

## Usage

on terminal

    $ ./setup.sh
    $ go build -o mpush
    $ ./mpush # listening to 8080 port.

you should install API client, like Postman, CocoaRestClient.

on client

you can create request like this.

![ClientImage](https://dl.dropboxusercontent.com/u/1655900/Screenshots/Screenshot%202015-08-29%2023.50.53.png)

## Contribution

Before you send pull request, I'd like you to type "go fmt *.go". :3
