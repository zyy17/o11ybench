# o11ybench

`o11ybench` is a powerful tool for benchmarking observability ecosystems.

It will move very fast and break things, so please be careful to use it.

The docs are not ready yet, so please refer to the code and examples.

## Architecture

The following diagram shows the future architecture of `o11ybench`.

<div align="center">
  <img src="./docs/images/arch.jpg" alt="Architecture">
</div>

## Features

- Support to generate logs by **ANY** format with the config file

- Support to generate logs for common log format:
  - Apache Common Log
  - Apache Combined Log
  - Apache Error Log
  - RFC3164 Log
  - RFC5424 Log
  - JSON Log

## Usage

### Compile

```console
make build
```

The `o11ybench` binary will be built in the `bin` directory.

### Generate logs

```console
o11ybench logs generate -c ./examples/generator/config_json.yaml
```

### Start benchmark

```console
o11ybench logs start -c ./examples/loader/config.yaml
```

You can find more config examples in the [examples](./examples) directory.

## Roadmap

- [x] Support to generate more popular log format
- [x] Support more fake data generator
- [x] Add logs benchmark
- [ ] Add otel traces benchmark
- [ ] Support prometheus metrics output(prometheus-benchmark)
- [ ] Be compatible with TSBS
- [ ] Output results in svg format
- [ ] Expose Prometheus metrics
- [ ] Flexible to define hybrid workloads benchmark by config file

## Acknowledgements

This project builds upon the work of several excellent open source projects:

- [gofakeit](https://github.com/brianvoe/gofakeit) - An amazing fake data generation library.
- [flog](https://github.com/mingrammer/flog) - A fantastic log generator that provided inspiration.
