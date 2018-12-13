# Etcd Data Migration

Status: *Alpha*

## Dump

```

Dump etcd data in a file

Usage:
  etcd-data-migration dump [flags]

Flags:
  -e, --endpoint string   Etcd endpoint (default "localhost:2379")
  -h, --help              help for dump
  -o, --output string     Dump file output (default "etcd.dump")

```


## Clone

```
Clone data from source etcd to target etcd

Usage:
  etcd-data-migration clone [flags]

Flags:
  -h, --help               help for clone
  -o, --overwrite string   Overwrite existing keys (default "false")
  -s, --source string      Source Etcd endpoint (default "localhost:2379")
  -t, --target string      Target Etcd endpoint (default "localhost:3379")

```
