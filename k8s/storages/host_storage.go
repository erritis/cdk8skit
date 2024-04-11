package cdk8skit

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus28/v2/k8s"
)

type HostStorageProps struct {
	IsDefault *bool
}

func (props *HostStorageProps) defaultProps() {
	if props.IsDefault == nil {
		props.IsDefault = jsii.Bool(false)
	}
}

func NewHostStorage(scope constructs.Construct, id string, props *HostStorageProps) k8s.KubeStorageClass {

	props.defaultProps()

	storage := k8s.NewKubeStorageClass(scope, jsii.String(id), &k8s.KubeStorageClassProps{
		Provisioner: jsii.String("k8s.io/minikube-hostpath"),
		Metadata: &k8s.ObjectMeta{
			Name: jsii.String(id),
			Annotations: &map[string]*string{
				"storageclass.kubernetes.io/is-default-class": jsii.String(fmt.Sprintf("%t", *props.IsDefault)),
			},
		},
		VolumeBindingMode: jsii.String("Immediate"),
	})

	storage.AddJsonPatch(cdk8s.JsonPatch_Replace(jsii.String("/metadata/namespace"), new(*string)))

	return storage
}
