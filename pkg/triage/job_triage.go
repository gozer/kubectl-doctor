package triage

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"time"
)

// LeftoverJobs gets a kubernetes.Clientset and a specific namespace string
// then proceeds to search if there are leftover cronjobs that were inactive for more than a month
func LeftoverJobs(kubeCli *kubernetes.Clientset, namespace string) (*Triage, error) {
	listOfTriages := make([]string, 0)

	jobs, err := kubeCli.BatchV1beta1().CronJobs(namespace).List(v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	currentTime := time.Now()
	for _, i := range jobs.Items {
		if i.Status.LastScheduleTime != nil {
			if currentTime.Sub(i.Status.LastScheduleTime.Local()) > 31 {
				listOfTriages = append(listOfTriages, i.GetName())
			}
		}

	}
	return NewTriage("CronJobs", "Found leftover cronjobs in namespace: "+namespace, listOfTriages), nil
}
