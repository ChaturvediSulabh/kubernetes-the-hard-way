# A list of real-world DevOps tasks

_Practice Before attempting CKA_

1. **Task: Statefulset**

   _Description_: we need to spun-up a MySQL Database cluster on Kubernetes. The DB must be highly available and should have a persistent storage. Once done, Share the DB Endpoint.

   - [Set-up NFS server on a dedicated node](./ubuntu/nfs-mount.sh)

     Note: All kubelet nodes must have nfs-common installed else pods will fail to start with following error

     > [ERROR: mount: wrong fs type, bad option, bad superblock on xx.xx.xx.xx:/nfs, missing codepage or helper program, or other error (for several filesystems (e.g. nfs, cifs) you might need a /sbin/mount.\<type\> helper program)](https://github.com/rancher/rancher/issues/4423)

   - [Provision a Persistent Volume that serves as an nfs-client](./tasks/statefulset/pv-mysql.yaml)
   - [Mysql config for master and slave](./tasks/statefulset/cm-mysql.yaml)
   - [Create a Service for the statefulset](./tasks/statefulset/svc-mysql.yaml)
   - [Deploy A High available and stateful MySQL Application](./tasks/statefulset/statefulset-mysql.yaml)

2. **Task: Horizontal Pod Autoscaler**

   _Description_: We need to autoscale the deployments based upon the CPU usage across all the pods. Ensure Average CPU utilization of 70% across all the pods.

   - [Create nginx deployment with CPU limit of 700m and CPU requests of 300m](./tasks/hpa/ngx-deploy.yaml)
   - [Expose nginx deployment](./tasks/hpa/ngx-svc.yaml)
   - [Create a horizontal pod autoscaler](./tasks/hpa/ngx-autoscaler.yaml)
     > Note: You must have metrics-server up and running else hpa won't work.
     > `kubectl get hpa`
   - [Increase the Load](./tasks/hpa/busybox.yaml)
     `kubectl get hpa`
   - [Decrease the load](`kubectl delete pod busybox`)
     `kubectl get hpa`
