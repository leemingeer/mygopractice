
# conformance
sonobuoy run --plugin e2e  --e2e-repo-config e2e-repo-config.yaml  --kube-conformance-image harbor.huayun.org/huayun-kubernetes/k8stest/conformance:v1.18.15 --timeout 10800000    --plugin-env e2e.E2E_EXTRA_ARGS="--dns-domain=cluster.local"  --e2e-focus "\\[Conformance\\]"   --e2e-skip "\\[Disruptive\\]|NoExecuteTaintManager|should not be very high"
