mpush
========================

Push Notification Server written by [Go](https://golang.org).

## Prerequisite

- prepare cert.pem, and key.pem along with [this guide](http://stackoverflow.com/a/21253261)

## Usage

**move to the directory where cert.pem and key.pem exist.**

- download [binary](https://github.com/yuchan/mpush/releases)
- execute mpush binary (e.g. ./mpush)

you should install API client, like [Postman](https://www.getpostman.com), [CocoaRestClient](http://mmattozzi.github.io/cocoa-rest-client/).

**on client**

you can create request like this.

![ClientImage](https://dl.dropboxusercontent.com/u/1655900/Screenshots/Screenshot%202015-08-29%2023.50.53.png)

## Contribution

### prepare

    $ go get github.com/tools/godep
	$ godep restore

Before you send pull request, I'd like you to type "go fmt *.go". :3
