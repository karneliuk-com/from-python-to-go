# Code for blog 011

## Use case
Parse XML, JSON, and YAML files with Python and (Go) Golang.

Extra dependencies are required.

### Python
Install `pyyaml`:
```bash
$pip install xml2dict pyyaml
```

Output:
```bash
Collecting xmltodict
  Downloading xmltodict-0.14.2-py2.py3-none-any.whl (10.0 kB)
Installing collected packages: xml2dict
Collecting pyyaml
  Downloading PyYAML-6.0.2-cp310-cp310-manylinux_2_17_x86_64.manylinux2014_x86_64.whl (751 kB)
     ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ 751.2/751.2 kB 20.3 MB/s eta 0:00:00
Installing collected packages: pyyaml
Successfully installed xmltodict-0.14.2, pyyaml-6.0.2
```

Then reference it in your code:
```python
import yaml
```

### Go (Golang)
Install `gopkg.in/yaml.v3`:
```bash
$ go get gopkg.in/yaml.v3
```

Output:
```bash
go: added gopkg.in/yaml.v3 v3.0.1
```

Then reference it in your code:
```go
import (
    "gopkg.in/yaml.v3"
)
```

Then tidy up your Go module:
```bash
$ go mod tidy
```
