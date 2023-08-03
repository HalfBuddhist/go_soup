module go_soup

replace (
	code.google.com/p/go.crypto => golang.org/x/crypto v0.0.0-20200406173513-056763e48d71 // indirect
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.6
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.1.0
	github.com/kubeflow/common-v01 => github.com/kubeflow/common v0.1.0
	github.com/kubeflow/pytorch-operator => github.com/linquanisaac/pytorch-operator v1.1.0
	github.com/kubernetes-sigs/kube-batch => github.com/kubernetes-sigs/kube-batch v0.0.0-20200414051246-2e934d1c8860
	github.com/xordataexchange/crypt => github.com/xordataexchange/crypt v0.0.2
	k8s.io/api => k8s.io/api v0.17.2
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.17.2
	k8s.io/apimachinery => k8s.io/apimachinery v0.17.2
	k8s.io/apiserver => k8s.io/apiserver v0.17.2
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.17.2
	k8s.io/client-go => k8s.io/client-go v0.17.2
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.17.2
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.17.2
	k8s.io/code-generator => k8s.io/code-generator v0.17.2
	k8s.io/component-base => k8s.io/component-base v0.17.2
	k8s.io/cri-api => k8s.io/cri-api v0.17.2
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.17.2
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.17.2
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.17.2
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.17.2
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.17.2
	k8s.io/kubectl => k8s.io/kubectl v0.17.2
	k8s.io/kubelet => k8s.io/kubelet v0.17.2
	k8s.io/kubernetes => k8s.io/kubernetes v1.17.2
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.17.2
	k8s.io/metrics => k8s.io/metrics v0.17.2
	k8s.io/node-api => k8s.io/node-api v0.17.2
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.17.2
	k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.17.2
	k8s.io/sample-controller => k8s.io/sample-controller v0.17.2
)

require (
	github.com/coreos/go-etcd v2.0.0+incompatible // indirect
	github.com/imdario/mergo v0.3.8
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826
	github.com/spf13/viper v1.13.0
	github.com/stretchr/testify v1.8.0
	github.com/xordataexchange/crypt v0.0.2
	go.etcd.io/etcd/client/v2 v2.305.5 // indirect
	go.etcd.io/etcd/client/v3 v3.5.5 // indirect
	google.golang.org/grpc v1.48.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

go 1.16
