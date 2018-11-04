# Servant

Simply serve HTTP.


## Features

- Serve HTTP my way
- That's it


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
