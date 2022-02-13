# NATS Configuration Proxy

The NATS Server configuration proxy provides a secure REST or NATS interface
for modifying access control lists (ACLs), identities (users), accounts, and
passwords.  This proxy is designed to facilitate the development of command
line tools and/or user interfaces to remotely update a NATS server
configuration.

## Getting started

```sh
go get -u github.com/nats-io/nats-config-proxy
```

## Usage

```sh
Usage: nats-config-proxy [options...]

Common Options:
    -h, --help                    Show this message
    -v, --version                 Show version
    -d, --dir <directory>         Directory for storing data
    -c, --config <file>           Configuration file


NATS API Options:
    --enable_nats                 Enables the NATS API
    

HTTP Rest API Server Options:
    --enable_rest                 Enables the HTTP REST API
    -a, --addr <host>             Bind REST API to host address (default: 0.0.0.0)
    -p, --port <port>             Bind REST API to port (default: 4567)
    -f, --publish-script <file>   Path to an optional script to execute on publish
    --rest_cert <file>            Server certificate file
    --rest_key <file>             Private key for server certificate
    --rest_cacert <file>          Client certificate CA for verification

Logging Options:
    -l, --log <file>              File to redirect log output
    -D, --debug                   Enable debugging output
    -V, --trace                   Enable trace logging
    -DV                           Debug and trace

```

### Configuration file

The NATS config proxy supports a configuration file.  Authorization based
on the subject attributes of a client certificate is also supported.

```hcl
listen = '0.0.0.0:4567'

data_dir = 'test/data'

logging {
  debug = true
  trace = true
}

tls {
  ca = 'test/certs/ca.pem'
  cert = 'test/certs/server.pem'
  key = 'test/certs/server-key.pem'
}

auth {
  users = [
    { user = "CN=cncf.example.com,OU=CNCF" }
  ]
}
```

## How it works

The NATS REST Configuration proxy operates using a data directory a
configuration file, and a publish script.

The process is straightforward:

1. Launch the NATS REST Configuration proxy and specify the Authorization
configuration file you'd like to modify.
2. Use the REST API to modify users and permissions.
3. Take a snapshot.  This saves the current work in the data directory.
4. Invoke the publish command to copy a snapshot into the configuration
file and invoke the optional publish script.

### Why a script

A script is used for versatility.  For some, this could be used as
a step in a github devops flow and the script creates a PR with the new configuration
for human eyes to review.  For others, the updated file is copied to remote nodes and
then NATS servers are reloaded with remote commands, e.g. `ssh -t gnatsd -sl reload`.
One could even work on an included NATS server file directly, with changes to be picked
up nightly.  There are many options.

## Developing

```sh
# Build locally using Go modules
$ GO111MODULE=on go run main.go
[41405] 2019/02/11 16:18:52.713366 [INF] Starting nats-rest-config-proxy v0.0.1
[41405] 2019/02/11 16:18:52.713804 [INF] Listening on 0.0.0.0:4567

# To run the tests
$ go test ./... -v
```

Note:  To test locally, you'll need to add a hostname into your `/etc/hosts` file:
`127.0.0.1 nats-cluster.default.svc.cluster.local`
