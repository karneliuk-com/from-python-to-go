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
$ python main.py
ssl_target_name_override is applied, should be used for testing only!
config_result={'timestamp': 1744970788637880096, 'prefix': None, 'response': [{'path': 'interfaces', 'op': 'UPDATE'}]}
Device: ka-blog-dev-001
Config: [('/openconfig-interfaces:interfaces', {'interface': [{'name': 'Loopback 23', 'config': {'name': 'Loopback 23', 'description': 'Test-gnmi-python-2'}}]})]
Impact: ***
---
***************
*** 687,700 ****
  notification/0/update/0/val/openconfig-interfaces:interface/2/state/counters/in-errors = 0
  notification/0/update/0/val/openconfig-interfaces:interface/2/state/counters/in-fcs-errors = 0
  notification/0/update/0/val/openconfig-interfaces:interface/2/state/counters/in-multicast-pkts = 0
! notification/0/update/0/val/openconfig-interfaces:interface/2/state/counters/in-octets = 43751
! notification/0/update/0/val/openconfig-interfaces:interface/2/state/counters/in-unicast-pkts = 463
  notification/0/update/0/val/openconfig-interfaces:interface/2/state/counters/out-broadcast-pkts = 0
  notification/0/update/0/val/openconfig-interfaces:interface/2/state/counters/out-discards = 0
  notification/0/update/0/val/openconfig-interfaces:interface/2/state/counters/out-errors = 0
  notification/0/update/0/val/openconfig-interfaces:interface/2/state/counters/out-multicast-pkts = 0
! notification/0/update/0/val/openconfig-interfaces:interface/2/state/counters/out-octets = 20010
! notification/0/update/0/val/openconfig-interfaces:interface/2/state/counters/out-unicast-pkts = 144
  notification/0/update/0/val/openconfig-interfaces:interface/2/state/enabled = True
  notification/0/update/0/val/openconfig-interfaces:interface/2/state/openconfig-platform-port:hardware-port = Port97
  notification/0/update/0/val/openconfig-interfaces:interface/2/state/ifindex = 999001
--- 687,700 ----
  notification/0/update/0/val/openconfig-interfaces:interface/2/state/counters/in-errors = 0
  notification/0/update/0/val/openconfig-interfaces:interface/2/state/counters/in-fcs-errors = 0
  notification/0/update/0/val/openconfig-interfaces:interface/2/state/counters/in-multicast-pkts = 0
! notification/0/update/0/val/openconfig-interfaces:interface/2/state/counters/in-octets = 47298
! notification/0/update/0/val/openconfig-interfaces:interface/2/state/counters/in-unicast-pkts = 493
  notification/0/update/0/val/openconfig-interfaces:interface/2/state/counters/out-broadcast-pkts = 0
  notification/0/update/0/val/openconfig-interfaces:interface/2/state/counters/out-discards = 0
  notification/0/update/0/val/openconfig-interfaces:interface/2/state/counters/out-errors = 0
  notification/0/update/0/val/openconfig-interfaces:interface/2/state/counters/out-multicast-pkts = 0
! notification/0/update/0/val/openconfig-interfaces:interface/2/state/counters/out-octets = 50683
! notification/0/update/0/val/openconfig-interfaces:interface/2/state/counters/out-unicast-pkts = 176
  notification/0/update/0/val/openconfig-interfaces:interface/2/state/enabled = True
  notification/0/update/0/val/openconfig-interfaces:interface/2/state/openconfig-platform-port:hardware-port = Port97
  notification/0/update/0/val/openconfig-interfaces:interface/2/state/ifindex = 999001
***************
*** 818,824 ****
  notification/0/update/0/val/openconfig-interfaces:interface/4/subinterfaces/subinterface/0/openconfig-if-ip:ipv6/state/mtu = 1500
  notification/0/update/0/val/openconfig-interfaces:interface/4/subinterfaces/subinterface/0/state/enabled = True
  notification/0/update/0/val/openconfig-interfaces:interface/4/subinterfaces/subinterface/0/state/index = 0
! notification/0/update/0/val/openconfig-interfaces:interface/5/config/description = Go_Test_2
  notification/0/update/0/val/openconfig-interfaces:interface/5/config/enabled = True
  notification/0/update/0/val/openconfig-interfaces:interface/5/config/arista-intf-augments:load-interval = 300
  notification/0/update/0/val/openconfig-interfaces:interface/5/config/loopback-mode = True
--- 818,824 ----
  notification/0/update/0/val/openconfig-interfaces:interface/4/subinterfaces/subinterface/0/openconfig-if-ip:ipv6/state/mtu = 1500
  notification/0/update/0/val/openconfig-interfaces:interface/4/subinterfaces/subinterface/0/state/enabled = True
  notification/0/update/0/val/openconfig-interfaces:interface/4/subinterfaces/subinterface/0/state/index = 0
! notification/0/update/0/val/openconfig-interfaces:interface/5/config/description = Test-gnmi-python-2
  notification/0/update/0/val/openconfig-interfaces:interface/5/config/enabled = True
  notification/0/update/0/val/openconfig-interfaces:interface/5/config/arista-intf-augments:load-interval = 300
  notification/0/update/0/val/openconfig-interfaces:interface/5/config/loopback-mode = True
***************
*** 833,839 ****
  notification/0/update/0/val/openconfig-interfaces:interface/5/state/enabled = True
  notification/0/update/0/val/openconfig-interfaces:interface/5/state/name = Loopback23
  notification/0/update/0/val/openconfig-interfaces:interface/5/state/openconfig-vlan:tpid = openconfig-vlan-types:TPID_0X8100
! notification/0/update/0/val/openconfig-interfaces:interface/5/subinterfaces/subinterface/0/config/description = Go_Test_2
  notification/0/update/0/val/openconfig-interfaces:interface/5/subinterfaces/subinterface/0/config/enabled = True
  notification/0/update/0/val/openconfig-interfaces:interface/5/subinterfaces/subinterface/0/config/index = 0
  notification/0/update/0/val/openconfig-interfaces:interface/5/subinterfaces/subinterface/0/index = 0
--- 833,839 ----
  notification/0/update/0/val/openconfig-interfaces:interface/5/state/enabled = True
  notification/0/update/0/val/openconfig-interfaces:interface/5/state/name = Loopback23
  notification/0/update/0/val/openconfig-interfaces:interface/5/state/openconfig-vlan:tpid = openconfig-vlan-types:TPID_0X8100
! notification/0/update/0/val/openconfig-interfaces:interface/5/subinterfaces/subinterface/0/config/description = Test-gnmi-python-2
  notification/0/update/0/val/openconfig-interfaces:interface/5/subinterfaces/subinterface/0/config/enabled = True
  notification/0/update/0/val/openconfig-interfaces:interface/5/subinterfaces/subinterface/0/config/index = 0
  notification/0/update/0/val/openconfig-interfaces:interface/5/subinterfaces/subinterface/0/index = 0
Timestamp: 2025-04-18 11:06:30.556257
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
