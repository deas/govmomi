```shell
govc datastore.disk.create -size 24G /disk1.vmdk
govc datastore.disk.info -ds=LocalDS_0 /disk1.vmdk
# Works now
./govc datastore.disk.info -ds=LocalDS_0 /DC0_H0_VM0/disk1.vmdk
```
