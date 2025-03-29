# Code for blog 018

## Use case
Intreraction with device using GNMI. Put credentails to environment variables.
```bash
$ export AUTOMATION_USER="user"
$ export AUTOMATION_PASS="password"
```

Extra dependencies are required.

### Python
Requirements:
```Bash
$ pip install pyyaml pygnmi
```

Execution result:
```bash
$ python main.py -i ../data/inventory.yaml
ssl_target_name_override is applied, should be used for testing only!
config_result={'timestamp': 1743270193131674380, 'prefix': None, 'response': [{'path': 'interfaces', 'op': 'UPDATE'}]}
Device: go-blog-arista
Config: [('/openconfig-interfaces:interfaces', {'interface': [{'name': 'Loopback 23', 'config': {'name': 'Loopback 23', 'description': 'Test-gnmi-python-2'}}]})]
Impact: ***
---
***************
*** 97,110 ****
  notification/0/update/0/val/openconfig-interfaces:interface/0/state/counters/in-errors = 0
  notification/0/update/0/val/openconfig-interfaces:interface/0/state/counters/in-fcs-errors = 0
  notification/0/update/0/val/openconfig-interfaces:interface/0/state/counters/in-multicast-pkts = 0
! notification/0/update/0/val/openconfig-interfaces:interface/0/state/counters/in-octets = 49384793
! notification/0/update/0/val/openconfig-interfaces:interface/0/state/counters/in-unicast-pkts = 649097
  notification/0/update/0/val/openconfig-interfaces:interface/0/state/counters/out-broadcast-pkts = 0
  notification/0/update/0/val/openconfig-interfaces:interface/0/state/counters/out-discards = 0
  notification/0/update/0/val/openconfig-interfaces:interface/0/state/counters/out-errors = 0
  notification/0/update/0/val/openconfig-interfaces:interface/0/state/counters/out-multicast-pkts = 0
! notification/0/update/0/val/openconfig-interfaces:interface/0/state/counters/out-octets = 18386266
! notification/0/update/0/val/openconfig-interfaces:interface/0/state/counters/out-unicast-pkts = 63350
  notification/0/update/0/val/openconfig-interfaces:interface/0/state/enabled = True
  notification/0/update/0/val/openconfig-interfaces:interface/0/state/openconfig-platform-port:hardware-port = Port97
  notification/0/update/0/val/openconfig-interfaces:interface/0/state/ifindex = 999001
--- 97,110 ----
  notification/0/update/0/val/openconfig-interfaces:interface/0/state/counters/in-errors = 0
  notification/0/update/0/val/openconfig-interfaces:interface/0/state/counters/in-fcs-errors = 0
  notification/0/update/0/val/openconfig-interfaces:interface/0/state/counters/in-multicast-pkts = 0
! notification/0/update/0/val/openconfig-interfaces:interface/0/state/counters/in-octets = 49390958
! notification/0/update/0/val/openconfig-interfaces:interface/0/state/counters/in-unicast-pkts = 649163
  notification/0/update/0/val/openconfig-interfaces:interface/0/state/counters/out-broadcast-pkts = 0
  notification/0/update/0/val/openconfig-interfaces:interface/0/state/counters/out-discards = 0
  notification/0/update/0/val/openconfig-interfaces:interface/0/state/counters/out-errors = 0
  notification/0/update/0/val/openconfig-interfaces:interface/0/state/counters/out-multicast-pkts = 0
! notification/0/update/0/val/openconfig-interfaces:interface/0/state/counters/out-octets = 18434988
! notification/0/update/0/val/openconfig-interfaces:interface/0/state/counters/out-unicast-pkts = 63463
  notification/0/update/0/val/openconfig-interfaces:interface/0/state/enabled = True
  notification/0/update/0/val/openconfig-interfaces:interface/0/state/openconfig-platform-port:hardware-port = Port97
  notification/0/update/0/val/openconfig-interfaces:interface/0/state/ifindex = 999001
***************
*** 742,748 ****
  notification/0/update/0/val/openconfig-interfaces:interface/2/subinterfaces/subinterface/0/openconfig-if-ip:ipv6/state/mtu = 1500
  notification/0/update/0/val/openconfig-interfaces:interface/2/subinterfaces/subinterface/0/state/enabled = True
  notification/0/update/0/val/openconfig-interfaces:interface/2/subinterfaces/subinterface/0/state/index = 0
! notification/0/update/0/val/openconfig-interfaces:interface/3/config/description = Test-gnmi-python
  notification/0/update/0/val/openconfig-interfaces:interface/3/config/enabled = True
  notification/0/update/0/val/openconfig-interfaces:interface/3/config/arista-intf-augments:load-interval = 300
  notification/0/update/0/val/openconfig-interfaces:interface/3/config/loopback-mode = True
--- 742,748 ----
  notification/0/update/0/val/openconfig-interfaces:interface/2/subinterfaces/subinterface/0/openconfig-if-ip:ipv6/state/mtu = 1500
  notification/0/update/0/val/openconfig-interfaces:interface/2/subinterfaces/subinterface/0/state/enabled = True
  notification/0/update/0/val/openconfig-interfaces:interface/2/subinterfaces/subinterface/0/state/index = 0
! notification/0/update/0/val/openconfig-interfaces:interface/3/config/description = Test-gnmi-python-2
  notification/0/update/0/val/openconfig-interfaces:interface/3/config/enabled = True
  notification/0/update/0/val/openconfig-interfaces:interface/3/config/arista-intf-augments:load-interval = 300
  notification/0/update/0/val/openconfig-interfaces:interface/3/config/loopback-mode = True
***************
*** 757,763 ****
  notification/0/update/0/val/openconfig-interfaces:interface/3/state/enabled = True
  notification/0/update/0/val/openconfig-interfaces:interface/3/state/name = Loopback23
  notification/0/update/0/val/openconfig-interfaces:interface/3/state/openconfig-vlan:tpid = openconfig-vlan-types:TPID_0X8100
! notification/0/update/0/val/openconfig-interfaces:interface/3/subinterfaces/subinterface/0/config/description = Test-gnmi-python
  notification/0/update/0/val/openconfig-interfaces:interface/3/subinterfaces/subinterface/0/config/enabled = True
  notification/0/update/0/val/openconfig-interfaces:interface/3/subinterfaces/subinterface/0/config/index = 0
  notification/0/update/0/val/openconfig-interfaces:interface/3/subinterfaces/subinterface/0/index = 0
--- 757,763 ----
  notification/0/update/0/val/openconfig-interfaces:interface/3/state/enabled = True
  notification/0/update/0/val/openconfig-interfaces:interface/3/state/name = Loopback23
  notification/0/update/0/val/openconfig-interfaces:interface/3/state/openconfig-vlan:tpid = openconfig-vlan-types:TPID_0X8100
! notification/0/update/0/val/openconfig-interfaces:interface/3/subinterfaces/subinterface/0/config/description = Test-gnmi-python-2
  notification/0/update/0/val/openconfig-interfaces:interface/3/subinterfaces/subinterface/0/config/enabled = True
  notification/0/update/0/val/openconfig-interfaces:interface/3/subinterfaces/subinterface/0/config/index = 0
  notification/0/update/0/val/openconfig-interfaces:interface/3/subinterfaces/subinterface/0/index = 0
Timestamp: 2025-03-29 17:43:15.708137
```

### Go
Requirements:
```bash
$ go get github.com/google/go-cmp/cmp
$ go get gopkg.in/yaml.v3
$ go get github.com/openconfig/gnmic/api
$ go get google.golang.org/protobuf/encoding/prototext
```

Execution result:
```bash
$ go run . -i ../data/inventory.yaml
2025/03/29 20:48:28 response:  {
  path:  {
    elem:  {
      name:  "openconfig-interfaces:interfaces"
    }
  }
  op:  UPDATE
}
timestamp:  1743281306005644545

Config: {/openconfig-interfaces:interfaces map[interface:[map[config:map[description:Test-gnmi-golang-3 name:Loopback 23] name:Loopback 23]]]}
Impact:   main.OpenConfigInterfaces{
        Interface: []main.OpenConfigInterface{
                {Name: "Management1", Config: {Name: "Management1", Enabled: true}},
                {Name: "Ethernet2", Config: {Name: "Ethernet2", Enabled: true}},
                {Name: "Ethernet1", Config: {Name: "Ethernet1", Enabled: true}},
                {
                        Name: "Loopback23",
                        Config: struct{ Name string "xml:\"name,omitempty\" json:\"name,omitempty\""; Description string "xml:\"description,omitempty\" json:\"description,omitempty\""; Enabled bool "xml:\"enabled,omitempty\" json:\"enabled,omitempty\"" }{
                                Name:        "Loopback23",
-                               Description: "Test-gnmi-golang-2",
+                               Description: "Test-gnmi-golang-3",
                                Enabled:     true,
                        },
                },
                {Name: "Loopback0", Config: {Name: "Loopback0", Enabled: true}},
                {Name: "Loopback51", Config: {Name: "Loopback51", Description: "pytest-update-test-33"}},
        },
  }

Timestamp: 2025-03-29 20:48:28.7119283 +0000 GMT m=+0.725116820
```
