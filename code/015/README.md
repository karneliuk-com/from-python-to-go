# Code for blog 015

## Use case
Basic SSH connectivity and command execution. Put credentails to environment variables.
```bash
$ export AUTOMATION_USER="user"
$ export AUTOMATION_PASS="password"
```

Extra dependencies are required.

### Python
```Bash
$ pip install pyyaml paramiko
```

### Go
```bash
$ go get golang.org/x/crypto/ssh
$ go get gopkg.in/yaml.v3
```
