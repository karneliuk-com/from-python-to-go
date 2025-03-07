# Code for blog 016

## Use case
Advanced SSH connectivity and command execution. Put credentails to environment variables.
```bash
$ export AUTOMATION_USER="user"
$ export AUTOMATION_PASS="password"
```

Extra dependencies are required.

### Python
```Bash
$ pip install pyyaml scrapli
```

Execution result:
```bash
$ python main.py -i ../data/inventory.yaml
/home/anton/Documents/S/Study - Go/freestyle/from-python-to-go/code/venv/lib/python3.10/site-packages/scrapli/helper.py:322: UserWarning:

**************************************************************************************************************************** Authentication Warning! *****************************************************************************************************************************
                                    scrapli will try to escalate privilege without entering a password but may fail.
Set an 'auth_secondary' password if your device requires a password to increase privilege, otherwise ignore this message.
**********************************************************************************************************************************************************************************************************************************************************************************

  warn(warning_message)
Device: dev-pygnmi-eos-001
Command: ['interface Loopback 23', 'description Test']
Impact: ***
---
***************
*** 2,6 ****
--- 2,7 ----
  Et1                            up             up
  Et2                            up             up
  Lo0                            up             up
+ Lo23                           up             up                 Test
  Lo51                           admin down     down               pytest-update-test-33
  Ma1                            up             up
Timestamp: 2025-03-07 21:24:11.267377
```

### Go
```bash
$ go get github.com/google/go-cmp/cmp
$ go get gopkg.in/yaml.v3
$ go get github.com/scrapli/scrapligo
```
