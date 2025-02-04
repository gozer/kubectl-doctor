package triage

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	coreclient "k8s.io/client-go/kubernetes/typed/core/v1"
)

const targetReason = "KubeletReady"

// TriageNodes gets a coreclient for k8s and checks if there are any nodes in the cluster
// that are not in Ready state(unoperational nodes)
func TriageNodes(coreClient coreclient.CoreV1Interface) (*Triage, error) {
	listOfTriages := make([]string, 0)
	nodes, err := coreClient.Nodes().List(v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, i := range nodes.Items {
		for _, y := range i.Status.Conditions {
			if y.Reason == targetReason {
				if y.Status != "True" {
					listOfTriages = append(listOfTriages, i.GetName())
				}
			}
		}
	}
	return NewTriage("Nodes", "Found node/s that are not in Ready state!", listOfTriages), nil
}
