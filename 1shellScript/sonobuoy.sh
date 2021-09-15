
# conformance
sonobuoy run --plugin e2e  --e2e-repo-config e2e-repo-config.yaml  --kube-conformance-image harbor.huayun.org/huayun-kubernetes/k8stest/conformance:v1.18.15 --timeout 10800000    --plugin-env e2e.E2E_EXTRA_ARGS="--dns-domain=cluster.local"  --e2e-focus "\\[Conformance\\]"   --e2e-skip "\\[Disruptive\\]|NoExecuteTaintManager|should not be very high"

# arm64
sonobuoy run --plugin e2e \
--e2e-repo-config arm64-e2e-repo.yaml   \
--kube-conformance-image harbor.huayun.org/huayun-kubernetes/k8stest/aarch64/conformance:v1.18.15 --timeout 10800000    --plugin-env e2e.E2E_EXTRA_ARGS="--dns-domain=cluster.local"  --e2e-focus "\\[Conformance\\]"   --e2e-skip "\\[Disruptive\\]|NoExecuteTaintManager|should not be very high"



# outfile=$(sonobuoy retrieve)
# sonobuoy results $outfile
Plugin: e2e
Status: failed
Total: 4994
Passed: 272
Failed: 2
Skipped: 4720

Failed tests:
[sig-network] Services should be able to create a functioning NodePort service [Conformance]
[sig-network] Services should be able to change the type from ExternalName to NodePort [Conformance]

--e2e-focus "Services should be able to create a functioning NodePort service"


cat <<EOF> huayun-e2e.yaml
sonobuoy-config:
  driver: Job
  plugin-name: huayun-e2e
  result-format: raw
spec:
  command:
  - /run_e2e.sh
  env:
  - name: E2E_FOCUS
    value: '\[arstorplugin\]'
  - name: REMOTE_HOST_PWD
    value: Huayun@123
  - name: CONFIG_NETID
    value: c9759403-846c-46e6-b665-5aa8892f96cd
  - name: LVM_NAMESPACE
    value: default
  - name: PORTALS
    value: '\[179.19.5.35:80\]'
  - name: TARGETPORTAL
    value: 10.10.10.5:3260
  - name: REMOTE_HOST
    value: 178.104.162.21
  - name: REMOTE_HOST_PORT
    value: "22"
  - name: CONFIG_PROVIDER
    value: arsdn
  - name: ARSTORPLUGIN_NAMESPACE
    value: default
  - name: IQN
    value: iqn.2010-01.com.huayun:k8starget
  image: harbor.huayun.org/huayun-kubernetes/k8stest/amd64/huayun-e2e:1.18.15-ming
  name: plugin
  resources: {}
  volumeMounts:
  - mountPath: /tmp/results
    name: results
EOF
