#!/usr/bin/env bash

VERSION=v0.9.3-rke

git clone git@github.com:300481/rook.git
cd rook
git checkout ${VERSION}

kubectl apply -f cluster/examples/kubernetes/ceph/operator.yaml
for i in {1..12} ; do
    sleep 10
    echo -n .
done
echo " "
kubectl apply -f cluster/examples/kubernetes/ceph/cluster.yaml


#helm repo add rook-stable https://charts.rook.io/stable
#helm install --name rook-ceph --namespace rook-ceph rook-stable/rook-ceph --set agent.flexVolumeDirPath=/var/lib/kubelet/volumeplugins --version ${VERSION}
