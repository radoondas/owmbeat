# Open Weather Map Beat (owmbeat)

Welcome to Owmbeat.

## Getting Started with owmbeat

### Requirements

* [Golang](https://golang.org/dl/) 1.10

### Clone

To clone Owmbeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/src/github.com/<github-user>/owmbeat
git clone https://github.com/radoondas/owmbeat ${GOPATH}/src/github.com/<github-user>/owmbeat
```


### Build

To build the binary for Owmbeat run the command below. This will generate a binary
in the same directory with the name owmbeat.

```
make
```


### Run

To run Owmbeat with debugging output enabled, run:

```
./owmbeat -c owmbeat.yml -e -d "*"
```


### Test

To test Owmbeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `fields.yml` by running the following command.

```
make update
```


### Cleanup

To clean  Owmbeat source code, run the following commands:

```
make fmt
make simplify
```

To clean up the build directory and generated artifacts, run:

```
make clean
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
mage package
```

This will fetch and create all images required for the build process. The whole process to finish can take several minutes.
