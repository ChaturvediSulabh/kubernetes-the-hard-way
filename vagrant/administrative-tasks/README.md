# A list of real-world tasks to for DevOps and CKA skills

1. **Task: Statefulset**

   _Description_: we need to spun-up a MySQL Database cluster on Kubernetes. The DB must be highly available and should have a persistent storage. Once done, Share the DB Endpoint.

   - [Set-up NFS server on a dedicated node](./ubuntu/nfs-mount.sh)

     Note: All kubelet nodes must have nfs-common installed else pods will fail to start with following error

     > [ERROR: mount: wrong fs type, bad option, bad superblock on xx.xx.xx.xx:/nfs, missing codepage or helper program, or other error (for several filesystems (e.g. nfs, cifs) you might need a /sbin/mount.\<type\> helper program)](https://github.com/rancher/rancher/issues/4423)

   - [Provision a Persistent Volume that serves as an nfs-client](./tasks/statefulset/pv-mysql.yaml)
   - [Create a Service for the statefulset](./tasks/statefulset/svc-mysql.yaml)
   - [Deploy A High available and stateful MySQL Application](./tasks/statefulset/statefulset-mysql.yaml)
