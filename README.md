<p align="center"><a href="https://github.com/toolsascode/protomagic"><image src="./assets/protomagic.png" style="width: 100px;"></a></p>

# ProtoMagic
ProtoMagic is a CLI that helps convert database tables into Protocol Buffers files.

## Documentation
[protomagic CLI](./docs/protomagic.md)

## Minimal configuration example
```yaml
databases:
  postgresql:
    dataSourceName: postgres://postgres:12345@localhost:5432/db-name?sslmode=disable
  mysql:
    dataSourceName: root:12345@tcp(localhost:3306)/db-name

```

## Complete configuration example

See: [protomagic.yaml](./example/configs/.protomagic.yaml)

## ProtoMagic CLI
- Run the command by specifying the configuration file

```shell
protomagic -c ./my-config-file.yaml
```