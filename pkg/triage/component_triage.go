package triage

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	coreclient "k8s.io/client-go/kubernetes/typed/core/v1"
)

const componentHealthy = "True"

// TriageComponents gets a coreclient and checks if core components are in healthy state
// such as etcd cluster members, scheduler, controller-manager
func TriageComponents(coreClient coreclient.CoreV1Interface) (*Triage, error) {
	components, err := coreClient.ComponentStatuses().List(v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	listOfTriages := make([]string, 0)
	for _, i := range components.Items {
		for _, y := range i.Conditions {
			if y.Status != componentHealthy {
				listOfTriages = append(listOfTriages, i.GetName())
			}
		}
	}
	return NewTriage("ComponentStatuses", "Found unhealthy components!", listOfTriages), nil
}
