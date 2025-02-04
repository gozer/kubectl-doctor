package triage

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	coreclient "k8s.io/client-go/kubernetes/typed/core/v1"
)

const pvcLostPhase = "Lost"

// TriagePVC gets a coreclient and checks if there are any pvcs that are in lost state
func TriagePVC(coreClient coreclient.CoreV1Interface) (*Triage, error) {
	listOfTriages := make([]string, 0)
	pvcs, err := coreClient.PersistentVolumeClaims("").List(v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, i := range pvcs.Items {
		if i.Status.Phase == pvcLostPhase {
			listOfTriages = append(listOfTriages, i.GetName())
		}
	}
	return NewTriage("PVC", "Found PVC in Lost State!", listOfTriages), nil
}
