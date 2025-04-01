# o11ybench

`o11ybench` is a powerful tool for benchmarking observability ecosystems.

It will move very fast and break things, so please be careful to use it.

The docs are not ready yet, so please refer to the code and examples.

## 🏗️ Architecture

The following diagram shows the future architecture of `o11ybench`.

<div align="center">
  <img src="./docs/images/arch.jpg" alt="Architecture">
</div>

## 🪄 Features

- Support to generate logs by **ANY** format with the config file based on template syntax(like [`example/generator/logs/custom_log.yaml`](./examples/generator/logs/custom_log.yaml))

- Support to generate logs for common log format:
  - Apache Common Log
  - Apache Combined Log
  - Apache Error Log
  - RFC3164 Log
  - RFC5424 Log
  - JSON Log

- Support to run the HTTP ingestion benchmark

## 🚀 Quick Start

**NOTE**: Suppose you are in the root directory of the project.

You can find more config examples in the [examples](./examples) directory.

### Generate Logs

The following command will generate 100 logs in Apache Common Log format to stdout:

```console
docker run --rm \
  -v $(pwd)/examples/generator/logs:/config \
  registry.cn-hangzhou.aliyuncs.com/zyyinternal/o11ybench:latest \
  logs generate -c /config/apache_common_log.yaml
```

### Start Logs Ingestion Benchmark

**NOTE**: Suppose you already have a database(for example, [GreptimeDB](https://github.com/GrepTimeTeam/greptimedb)) running on your local machine and listen on the port `4000`.

The following command will start the logs ingestion benchmark:

```console
docker run --rm --network host \
  -v $(pwd)/examples/loader/logs:/config \
  registry.cn-hangzhou.aliyuncs.com/zyyinternal/o11ybench:latest \
  logs start -c /config/config.yaml
```

## 🛠️ Development

### Compile

Ensure Go >= 1.24 is installed.

```console
make build
```

The `o11ybench` binary will be built in the `bin` directory.

### Test

```console
make test
```

## 🚧 Roadmap

- [x] Support to generate more popular log format
- [x] Support more fake data generator
- [x] Add logs benchmark
- [ ] Add otel traces benchmark
- [ ] Support prometheus metrics output(prometheus-benchmark)
- [ ] Be compatible with TSBS
- [ ] Output results in svg format
- [ ] Expose Prometheus metrics
- [ ] Flexible to define hybrid workloads benchmark by config file

## 🤝 Acknowledgements

This project builds upon the work of several excellent open source projects:

- [gofakeit](https://github.com/brianvoe/gofakeit) - An amazing fake data generation library.
- [flog](https://github.com/mingrammer/flog) - A fantastic log generator that provided inspiration.
