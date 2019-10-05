# autorok

Deployment service for RancherOS based Kubernetes systems

## Start Service with docker-compose

```bash
export AUTOROK_CONFIG_URL=<URL to your config>
docker-compose -f deployments/docker-compose.yml up [-d]
```

### Example Config

```yaml
templatesource:
  ipxe: "https://raw.githubusercontent.com/300481/autorok-systems/master/3c7a3ea5-a859-477a-84ae-b11cc8d35852/templates/ipxe"
  boot: "https://raw.githubusercontent.com/300481/autorok-systems/master/3c7a3ea5-a859-477a-84ae-b11cc8d35852/templates/boot"
  install: "https://raw.githubusercontent.com/300481/autorok-systems/master/3c7a3ea5-a859-477a-84ae-b11cc8d35852/templates/install"
  rke: "https://raw.githubusercontent.com/300481/autorok-systems/master/3c7a3ea5-a859-477a-84ae-b11cc8d35852/templates/rke"
clusterconfig: "https://raw.githubusercontent.com/300481/autorok-systems/master/3c7a3ea5-a859-477a-84ae-b11cc8d35852/cluster.yaml"
bootserver: 192.168.0.45
clustername: cluster
nodecount: 1
publickkey: ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAINpGZNv6mEKc+uiJLPmz3Sld043vb0msB1l4xRjMG58D
startcidr: 192.168.0.113/28
gateway: 192.168.0.1
mtu: 1500
dhcp: false
nameservers:
- 192.168.0.1
```

## Description

This system is planned to be deployable in very cheap environments.

Minimum necessity of the IT infrastructure is an internet router which also acts as a DHCP server (f.e. just a FRITZ!Box).

This system will give you the foundation to deploy your needed applications reliable and scalable.

It will provide a DNS-Server for your network.

The applications will automatically exposed via a system internal load balancer,

the IP's will automatically get a DNS A record.

It is fault tolerant with its three pillars and will alert you on failures.

Each pillar (bare metal server) can be replaced on the fly.

An additional goal is, to automatically keep every software item anytime on the current stable release including all current updates.

Encrypted backup of the stateful data is also included to cloud provider(s).

For the future there is a plan to scale up the bare metal servers (adding additional workers).

## Architecture diagram for planned system

<img src="https://raw.githubusercontent.com/300481/autorok/master/pics/architecture.png">

## Checked functions

### Ceph Cluster --> works

see `install-ceph-cluster.sh`

### Ceph Block Storage --> works

BlockPool + StorageClass + PVC + Service + Deployment

```yaml
apiVersion: ceph.rook.io/v1
kind: CephBlockPool
metadata:
  name: replicapool
  namespace: rook-ceph
spec:
  failureDomain: host
  replicated:
    size: 3
---
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
   name: rook-ceph-block
provisioner: ceph.rook.io/block
parameters:
  blockPool: replicapool
  # Specify the namespace of the rook cluster from which to create volumes.
  # If not specified, it will use `rook` as the default namespace of the cluster.
  # This is also the namespace where the cluster will be
  clusterNamespace: rook-ceph
  # Specify the filesystem type of the volume. If not specified, it will use `ext4`.
  fstype: xfs
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: seafile-pv-claim
  labels:
    app: seafile
spec:
  storageClassName: rook-ceph-block
  accessModes:
  - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi
---
apiVersion: v1
kind: Service
metadata:
  name: seafile
  labels:
    app: seafile
spec:
  ports:
    - port: 80
  selector:
    app: seafile
  type: LoadBalancer
---
apiVersion: apps/v1beta1
kind: Deployment
metadata:
  name: seafile
  labels:
    app: seafile
spec:
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: seafile
    spec:
      containers:
      - image: seafileltd/seafile:latest
        name: seafile
        env:
        - name: SEAFILE_SERVER_HOSTNAME
          value: example.com
        - name: SEAFILE_ADMIN_EMAIL
          value: test@example.com
        - name: SEAFILE_ADMIN_PASSWORD
          value: changeme
        ports:
        - containerPort: 80
          name: seafile
        volumeMounts:
        - name: seafile-persistent-storage
          mountPath: /shared
      volumes:
      - name: seafile-persistent-storage
        persistentVolumeClaim:
          claimName: seafile-pv-claim
```
