# Servant

[![Build Status][badge_build]][travis_servant]
[![GoReportCard][badge_reportcard]][reportcard]

[badge_build]: https://travis-ci.org/qjcg/servant.svg?branch=master
[badge_reportcard]: https://goreportcard.com/badge/github.com/qjcg/servant
[travis_servant]: https://travis-ci.org/qjcg/servant
[reportcard]: https://goreportcard.com/report/github.com/qjcg/servant

Simply serve HTTP.


## Features

- Serve HTTP my way
- That's it

## Install

```sh
go get -u github.com/qjcg/servant
```

## Usage

```shell
# Serve HTTP on localhost:8080
$ servant

# Serve ~/public via LAN IP and custom port:
$ servant -i 10.13.37.10 -p 8888 ~/public

# Serve HTTPS:
$ servant -c example.com.pem -k example.com-key.pem -i example.com ~/public
```

## Licence

MIT.
