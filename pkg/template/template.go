package template

import (
	"github.com/open-policy-agent/frameworks/constraint/pkg/apis/templates/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ConstraintTemplate(name string, libs []string, rego string) v1beta1.ConstraintTemplate {
	target := "admission.k8s.gatekeeper.sh"
	return v1beta1.ConstraintTemplate{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "templates.gatekeeper.sh/v1beta1",
			Kind:       "ConstraintTemplate",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: v1beta1.ConstraintTemplateSpec{
			CRD: v1beta1.CRD{
				Spec: v1beta1.CRDSpec{
					Names: v1beta1.Names{
						Kind: name,
					},
				},
			},
			Targets: []v1beta1.Target{
				{
					Target: target,
					Rego:   rego,
					Libs:   libs,
				},
			},
		},
	}
}
