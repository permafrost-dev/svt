# SVT - Semantic Versioning Tools

SVT is a utility for working with semantic versions written in Go.

## Available Commands

- `compare`: Compare two semantic versions, outputs "eq", "gt" or "lt".
```sh
svt compare [version-1] [version-2]
```

- `latest`: Compare two semantic versions and print the highest version.
```sh
svt latest [version-1] [version-2]
```

- `help`: Help about any command
```sh
svt help [command]
```

## Setup

```bash
go mod tidy
```

## Building the project

`svt` uses [task](https://github.com/go-task/task) for running tasks, which is a tool similar to `make`. 

```bash
task build
```

---

## Changelog

Please see [CHANGELOG](CHANGELOG.md) for more information on what has changed recently.

## Contributing

Please see [CONTRIBUTING](.github/CONTRIBUTING.md) for details.

## Security Vulnerabilities

Please review [our security policy](../../security/policy) on how to report security vulnerabilities.

## Credits

- [Patrick Organ](https://github.com/patinthehat)
- [All Contributors](../../contributors)

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
