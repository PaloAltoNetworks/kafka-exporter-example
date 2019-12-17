# Examples

## Enforcer operational status change

When an enforcer operational status change, you will see:

* a `squall-event` that is referring to the update status of the enforcer
* an `activities` event as this is a main event that we want to keep in the System Monitoring.

The below example shows what happens when an enforcer switch from Connected to Disconnected:

``` bash
New publication with topic squall-events:
Content:
{
  "encoding": "application/msgpack",
  "entity": {
    "FQDN": "apomux-enforcerd-4",
    "ID": "5d01462f6eb57622d5109e9f",
    "annotations": {},
    "associatedTags": [],
    "certificate": "",
    "certificateRequest": "",
    "collectInfo": false,
    "collectedInfo": {},
    "createTime": "2019-06-12T18:36:31.485Z",
    "currentVersion": "0.0.0-dev",
    "description": "Enforcer for host apomux-enforcerd-4",
    "enforcementStatus": "Inactive",
    "lastCollectionTime": null,
    "lastSyncTime": "2019-06-12T21:38:02.23195Z",
    "localCA": "",
    "machineID": "2C59BB75-6152-4979-9F3B-6FB02AD1A510",
    "metadata": [],
    "name": "apomux-enforcerd-4",
    "namespace": "/apomux",
    "normalizedTags": [
      "$currentversion=0.0.0-dev",
      "$enforcementstatus=Inactive",
      "$id=5d01462f6eb57622d5109e9f",
      "$identity=enforcer",
      "$machineid=2C59BB75-6152-4979-9F3B-6FB02AD1A510",
      "$name=apomux-enforcerd-4",
      "$namespace=/apomux",
      "$operationalstatus=Disconnected"
    ],
    "operationalStatus": "Disconnected",
    "protected": false,
    "publicToken": "",
    "startTime": null,
    "subnets": [],
    "unreachable": false,
    "updateAvailable": false,
    "updateTime": "2019-06-12T21:38:02.267Z",
    "zone": 0
  },
  "identity": "enforcer",
  "timestamp": "2019-06-12T21:38:02.268588Z",
  "type": "update"
}


New publication with topic activities:
Content:
{
  "ID": "",
  "claims": [
    "@auth:realm=certificate",
    "@auth:mode=internal",
    "@auth:serialnumber=52953752628829058742454994389793799740",
    "@auth:commonname=gaga",
    "@auth:organization=system",
    "@auth:organizationalunit=root"
  ],
  "data": {
    "FQDN": "apomux-enforcerd-4",
    "ID": "5d01462f6eb57622d5109e9f",
    "annotations": {},
    "associatedTags": [],
    "certificate": "",
    "certificateRequest": "",
    "collectInfo": false,
    "collectedInfo": {},
    "createTime": "2019-06-12T18:36:31.485Z",
    "currentVersion": "0.0.0-dev",
    "description": "Enforcer for host apomux-enforcerd-4",
    "enforcementStatus": "Inactive",
    "lastCollectionTime": "0001-01-01T00:00:00Z",
    "lastSyncTime": "2019-06-12T21:38:02.23195Z",
    "localCA": "",
    "machineID": "2C59BB75-6152-4979-9F3B-6FB02AD1A510",
    "metadata": [],
    "name": "apomux-enforcerd-4",
    "namespace": "/apomux",
    "normalizedTags": [
      "$currentversion=0.0.0-dev",
      "$enforcementstatus=Inactive",
      "$id=5d01462f6eb57622d5109e9f",
      "$identity=enforcer",
      "$machineid=2C59BB75-6152-4979-9F3B-6FB02AD1A510",
      "$name=apomux-enforcerd-4",
      "$namespace=/apomux",
      "$operationalstatus=Disconnected"
    ],
    "operationalStatus": "Disconnected",
    "protected": false,
    "publicToken": "",
    "startTime": "0001-01-01T00:00:00Z",
    "subnets": [],
    "unreachable": false,
    "updateAvailable": false,
    "updateTime": "2019-06-12T14:38:02.267-07:00",
    "zone": 0
  },
  "date": "2019-06-12T21:38:02.268616Z",
  "error": null,
  "message": "Updated enforcer with ID 5d01462f6eb57622d5109e9f",
  "namespace": "/apomux",
  "operation": "update",
  "originalData": {
    "FQDN": "apomux-enforcerd-4",
    "ID": "5d01462f6eb57622d5109e9f",
    "annotations": {},
    "associatedTags": [],
    "certificate": "",
    "certificateRequest": "",
    "collectInfo": false,
    "collectedInfo": {},
    "createTime": "2019-06-12T18:36:31.485Z",
    "currentVersion": "0.0.0-dev",
    "description": "Enforcer for host apomux-enforcerd-4",
    "enforcementStatus": "Active",
    "lastCollectionTime": "0001-01-01T00:00:00Z",
    "lastSyncTime": "2019-06-12T21:21:10.276Z",
    "localCA": "",
    "machineID": "2C59BB75-6152-4979-9F3B-6FB02AD1A510",
    "metadata": [],
    "name": "apomux-enforcerd-4",
    "namespace": "/apomux",
    "normalizedTags": [
      "$currentversion=0.0.0-dev",
      "$enforcementstatus=Active",
      "$id=5d01462f6eb57622d5109e9f",
      "$identity=enforcer",
      "$machineid=2C59BB75-6152-4979-9F3B-6FB02AD1A510",
      "$name=apomux-enforcerd-4",
      "$namespace=/apomux",
      "$operationalstatus=Connected"
    ],
    "operationalStatus": "Connected",
    "protected": false,
    "publicToken": "",
    "startTime": "0001-01-01T00:00:00Z",
    "subnets": [],
    "unreachable": true,
    "updateAvailable": false,
    "updateTime": "2019-06-12T21:37:34.809Z",
    "zone": 0
  },
  "source": " 127.0.0.1:56259",
  "targetIdentity": "enforcer",
  "zone": 0
}
```

_Example of an activity shown in the Aporeto UI_

![Screen Shot 2019-06-12 at 2 41 21 PM](https://user-images.githubusercontent.com/1447243/59388284-40f86000-8d20-11e9-8e23-bfa41901bd02.png)

## Enforcer heartbeat

Every 30 seconds, the enforcer is doing an heartbeat (also called "poke") that will update its `lastSyncTime` property.
For optimization reasons, heart beats are not logged.

However, every 5 min, the enforcer is updating itself.
Below is an example of the update event received in the `squall-events` topic:

``` bash
New publication with topic squall-events:
Content:
{
  "encoding": "application/msgpack",
  "entity": {
    "FQDN": "apomux-enforcerd-4",
    "ID": "5d01462f6eb57622d5109e9f",
    "annotations": {},
    "associatedTags": [],
    "certificate": "",
    "certificateRequest": "",
    "collectInfo": false,
    "collectedInfo": {},
    "createTime": "2019-06-12T18:36:31.485Z",
    "currentVersion": "0.0.0-dev",
    "description": "Enforcer for host apomux-enforcerd-4",
    "enforcementStatus": "Active",
    "lastCollectionTime": null,
    "lastSyncTime": "2019-06-12T22:02:30.204Z",
    "localCA": "",
    "machineID": "2C59BB75-6152-4979-9F3B-6FB02AD1A510",
    "metadata": [],
    "name": "apomux-enforcerd-4",
    "namespace": "/apomux",
    "normalizedTags": [
      "$currentversion=0.0.0-dev",
      "$enforcementstatus=Active",
      "$id=5d01462f6eb57622d5109e9f",
      "$identity=enforcer",
      "$machineid=2C59BB75-6152-4979-9F3B-6FB02AD1A510",
      "$name=apomux-enforcerd-4",
      "$namespace=/apomux",
      "$operationalstatus=Connected"
    ],
    "operationalStatus": "Connected",
    "protected": false,
    "publicToken": "",
    "startTime": null,
    "subnets": [],
    "unreachable": true,
    "updateAvailable": false,
    "updateTime": "2019-06-12T22:17:34.826557Z",
    "zone": 0
  },
  "identity": "enforcer",
  "timestamp": "2019-06-12T22:17:34.84007Z",
  "type": "update"
}
```

## Other examples

When a container is stopped, you will again see two events as when the enforcer change its status:

``` bash
New publication with topic squall-events:
Content:
{
  "encoding": "application/msgpack",
  "entity": {
    "ID": "5d016b196eb57631b11e6350",
    "annotations": {},
    "associatedTags": [
      "maintainer=NGINX Docker Maintainers \u003cdocker-maint@nginx.com\u003e"
    ],
    "collectInfo": false,
    "collectedInfo": {},
    "createTime": "2019-06-12T21:14:01.679Z",
    "description": "thirsty_chaplygin",
    "enforcementStatus": "Inactive",
    "enforcerID": "5d01462f6eb57622d5109e9f",
    "enforcerNamespace": "/apomux",
    "image": "",
    "images": [
      "nginx"
    ],
    "lastCollectionTime": null,
    "lastSyncTime": "2019-06-12T22:18:43.004Z",
    "metadata": [
      "@app:docker:ExposedPort=tcp:80",
      "@app:docker:imagechecksum=sha256:719cd2e3ed04781b11ed372ec8d712fac66d5b60a6fb6190bf76b7d18cb50105",
      "@app:docker:memoryreservation=none",
      "@app:docker:name=thirsty_chaplygin",
      "@app:docker:networkmode=bridge",
      "@app:docker:pid=0",
      "@app:docker:privileged=false",
      "@app:docker:readonlyrootfs=false",
      "@app:extractor=docker",
      "@app:image=nginx",
      "@sys:ExposedPort=tcp:80",
      "@sys:image=nginx",
      "@sys:imagechecksum=sha256:719cd2e3ed04781b11ed372ec8d712fac66d5b60a6fb6190bf76b7d18cb50105",
      "@sys:memoryreservation=none",
      "@sys:name=thirsty_chaplygin",
      "@sys:networkmode=bridge",
      "@sys:pid=0",
      "@sys:privileged=false",
      "@sys:readonlyrootfs=false",
      "@usr:maintainer=NGINX Docker Maintainers \u003cdocker-maint@nginx.com\u003e"
    ],
    "name": "thirsty_chaplygin",
    "namespace": "/apomux",
    "nativeContextID": "8d6500132576",
    "networkServices": [],
    "normalizedTags": [
      "$enforcementstatus=Inactive",
      "$enforcerid=5d01462f6eb57622d5109e9f",
      "$enforcernamespace=/apomux",
      "$id=5d016b196eb57631b11e6350",
      "$identity=processingunit",
      "$image=nginx",
      "$name=thirsty_chaplygin",
      "$namespace=/apomux",
      "$operationalstatus=Terminated",
      "$type=Docker",
      "@app:docker:exposedport=tcp:80",
      "@app:docker:imagechecksum=sha256:719cd2e3ed04781b11ed372ec8d712fac66d5b60a6fb6190bf76b7d18cb50105",
      "@app:docker:memoryreservation=none",
      "@app:docker:name=thirsty_chaplygin",
      "@app:docker:networkmode=bridge",
      "@app:docker:pid=0",
      "@app:docker:privileged=false",
      "@app:docker:readonlyrootfs=false",
      "@app:extractor=docker",
      "@app:image=nginx",
      "@sys:exposedport=tcp:80",
      "@sys:image=nginx",
      "@sys:imagechecksum=sha256:719cd2e3ed04781b11ed372ec8d712fac66d5b60a6fb6190bf76b7d18cb50105",
      "@sys:memoryreservation=none",
      "@sys:name=thirsty_chaplygin",
      "@sys:networkmode=bridge",
      "@sys:pid=0",
      "@sys:privileged=false",
      "@sys:readonlyrootfs=false",
      "@usr:maintainer=NGINX Docker Maintainers \u003cdocker-maint@nginx.com\u003e",
      "maintainer=NGINX Docker Maintainers \u003cdocker-maint@nginx.com\u003e",
      "$archived=true"
    ],
    "operationalStatus": "Terminated",
    "protected": false,
    "tracing": {
      "IPTables": false,
      "applicationConnections": false,
      "interval": "10s",
      "networkConnections": false
    },
    "type": "Docker",
    "unreachable": true,
    "updateTime": "2019-06-12T22:18:43.06Z",
    "zone": 0
  },
  "identity": "processingunit",
  "timestamp": "2019-06-12T22:18:43.062323Z",
  "type": "delete"
}


New publication with topic activities:
Content:
{
  "ID": "",
  "claims": [
    "@auth:commonname=app:credential:5cb7d0576eb5763cb678a19d:enforcerd",
    "@auth:organization=/apomux",
    "@auth:realm=certificate",
    "@auth:serialnumber=291441683681272654721053841675603804254",
    "@auth:subject=291441683681272654721053841675603804254",
    "@auth:subject=291441683681272654721053841675603804254"
  ],
  "data": {
    "ID": "5d016b196eb57631b11e6350",
    "annotations": {},
    "associatedTags": [
      "maintainer=NGINX Docker Maintainers \u003cdocker-maint@nginx.com\u003e"
    ],
    "collectInfo": false,
    "collectedInfo": {},
    "createTime": "2019-06-12T21:14:01.679Z",
    "description": "thirsty_chaplygin",
    "enforcementStatus": "Inactive",
    "enforcerID": "5d01462f6eb57622d5109e9f",
    "enforcerNamespace": "/apomux",
    "image": "",
    "images": [
      "nginx"
    ],
    "lastCollectionTime": "0001-01-01T00:00:00Z",
    "lastSyncTime": "2019-06-12T22:18:43.004Z",
    "metadata": [
      "@app:docker:ExposedPort=tcp:80",
      "@app:docker:imagechecksum=sha256:719cd2e3ed04781b11ed372ec8d712fac66d5b60a6fb6190bf76b7d18cb50105",
      "@app:docker:memoryreservation=none",
      "@app:docker:name=thirsty_chaplygin",
      "@app:docker:networkmode=bridge",
      "@app:docker:pid=0",
      "@app:docker:privileged=false",
      "@app:docker:readonlyrootfs=false",
      "@app:extractor=docker",
      "@app:image=nginx",
      "@sys:ExposedPort=tcp:80",
      "@sys:image=nginx",
      "@sys:imagechecksum=sha256:719cd2e3ed04781b11ed372ec8d712fac66d5b60a6fb6190bf76b7d18cb50105",
      "@sys:memoryreservation=none",
      "@sys:name=thirsty_chaplygin",
      "@sys:networkmode=bridge",
      "@sys:pid=0",
      "@sys:privileged=false",
      "@sys:readonlyrootfs=false",
      "@usr:maintainer=NGINX Docker Maintainers \u003cdocker-maint@nginx.com\u003e"
    ],
    "name": "thirsty_chaplygin",
    "namespace": "/apomux",
    "nativeContextID": "8d6500132576",
    "networkServices": [],
    "normalizedTags": [
      "$enforcementstatus=Inactive",
      "$enforcerid=5d01462f6eb57622d5109e9f",
      "$enforcernamespace=/apomux",
      "$id=5d016b196eb57631b11e6350",
      "$identity=processingunit",
      "$image=nginx",
      "$name=thirsty_chaplygin",
      "$namespace=/apomux",
      "$operationalstatus=Terminated",
      "$type=Docker",
      "@app:docker:exposedport=tcp:80",
      "@app:docker:imagechecksum=sha256:719cd2e3ed04781b11ed372ec8d712fac66d5b60a6fb6190bf76b7d18cb50105",
      "@app:docker:memoryreservation=none",
      "@app:docker:name=thirsty_chaplygin",
      "@app:docker:networkmode=bridge",
      "@app:docker:pid=0",
      "@app:docker:privileged=false",
      "@app:docker:readonlyrootfs=false",
      "@app:extractor=docker",
      "@app:image=nginx",
      "@sys:exposedport=tcp:80",
      "@sys:image=nginx",
      "@sys:imagechecksum=sha256:719cd2e3ed04781b11ed372ec8d712fac66d5b60a6fb6190bf76b7d18cb50105",
      "@sys:memoryreservation=none",
      "@sys:name=thirsty_chaplygin",
      "@sys:networkmode=bridge",
      "@sys:pid=0",
      "@sys:privileged=false",
      "@sys:readonlyrootfs=false",
      "@usr:maintainer=NGINX Docker Maintainers \u003cdocker-maint@nginx.com\u003e",
      "maintainer=NGINX Docker Maintainers \u003cdocker-maint@nginx.com\u003e",
      "$archived=true"
    ],
    "operationalStatus": "Terminated",
    "protected": false,
    "tracing": {
      "IPTables": false,
      "applicationConnections": false,
      "interval": "10s",
      "networkConnections": false
    },
    "type": "Docker",
    "unreachable": true,
    "updateTime": "2019-06-12T15:18:43.06-07:00",
    "zone": 0
  },
  "date": "2019-06-12T22:18:43.062363Z",
  "error": null,
  "message": "Deleted processingunit with ID 5d016b196eb57631b11e6350",
  "namespace": "/apomux",
  "operation": "delete",
  "originalData": null,
  "source": "enforcerd 192.168.100.103",
  "targetIdentity": "processingunit",
  "zone": 0
}
```
