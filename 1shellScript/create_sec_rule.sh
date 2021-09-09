#!/bin/bash

neutron security-group-show ake-master-sg
# ingress-controller-https
neutron security-group-rule-create ake-master-sg --direction ingress --protocol TCP --port-range-min 443  --port-range-max 443
# ingress-controller-http
neutron security-group-rule-create ake-master-sg --direction ingress --protocol TCP --port-range-min 80  --port-range-max 80
# promethus-http
neutron security-group-rule-create ake-master-sg --direction ingress --protocol TCP --port-range-min 7443  --port-range-max 7443
# dashboard-https
neutron security-group-rule-create ake-master-sg --direction ingress --protocol TCP --port-range-min 8443  --port-range-max 8443
# apiserver-https-vip
neutron security-group-rule-create ake-master-sg --direction ingress --protocol TCP --port-range-min 9443  --port-range-max 9443

# apiserver-https-6443
neutron security-group-rule-create ake-master-sg --direction ingress --protocol TCP --port-range-min 6443  --port-range-max 6443
# kube-controller-manager
neutron security-group-rule-create ake-master-sg --direction ingress --protocol TCP --port-range-min 10252  --port-range-max 10252
# kube-scheduler
neutron security-group-rule-create ake-master-sg --direction ingress --protocol TCP --port-range-min 10259  --port-range-max 10259
# etcd-2379
neutron security-group-rule-create ake-master-sg --direction ingress --protocol TCP --port-range-min 2379  --port-range-max 2379
# etcd-2380
neutron security-group-rule-create ake-master-sg --direction ingress --protocol TCP --port-range-min 2380  --port-range-max 2380
# arsdn
neutron security-group-rule-create ake-master-sg --direction ingress --protocol TCP --port-range-min 8108  --port-range-max 8108

# node-exporter
neutron security-group-rule-create ake-master-sg --direction ingress --protocol TCP --port-range-min 9100  --port-range-max 9100

# corosync
neutron security-group-rule-create ake-master-sg --direction ingress --protocol TCP --port-range-min 2224  --port-range-max 2224
neutron security-group-rule-create ake-master-sg --direction ingress --protocol TCP --port-range-min 3121  --port-range-max 3121
neutron security-group-rule-create ake-master-sg --direction ingress --protocol TCP --port-range-min 21064  --port-range-max 21064
neutron security-group-rule-create ake-master-sg --direction ingress --protocol TCP --port-range-min 5405  --port-range-max 5405
