# Code for blog 019

## Use case
Intreraction with applications via REST API. Put credentails to environment variables.
```bash
$ export AUTOMATION_INVENTORY_URL="https://demo.netbox.dev"
$ export AUTOMATION_INVENTORY_TOKEN="token"
$ export AUTOMATION_USER="user"
$ export AUTOMATION_PASS="password"
```

Extra dependencies are required.

### Python
Requirements:
```Bash
$ pip install pygnmi httpx
```

Execution result:
```bash
```

### Go
Requirements:
```bash
$ go get github.com/google/go-cmp/cmp
$ go get github.com/openconfig/gnmic/api
$ go get google.golang.org/protobuf/encoding/prototext
```

Execution result:
```bash
```
