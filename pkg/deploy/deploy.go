package deploy

import (
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Run our cli
func Run() {
	cfg, err := GetConfig()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	err = cfg.Validate()
	if err != nil {
		fmt.Printf("We are missing a configuration option: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Deploying %s on %s in %s\n", cfg.NewImage, cfg.DeploymentName, cfg.Namespace)

	// TODO mock this such that we can test quickly
	deploymentsClient := cfg.KubernetesClientSet.AppsV1beta1().Deployments(cfg.Namespace)

	result, getErr := deploymentsClient.Get(cfg.DeploymentName, metav1.GetOptions{})
	if getErr != nil {
		fmt.Printf("Failed to get latest version of Deployment %s: %v\n", cfg.DeploymentName, getErr)
		os.Exit(1)
	}

	// TODO determine container position
	result.Spec.Template.Spec.Containers[0].Image = cfg.NewImage
	_, updateErr := deploymentsClient.Update(result)

	if updateErr != nil {
		panic(fmt.Errorf("Update failed: %v", updateErr))
	}

	// TODO we need a backup rollback process in case of failure
	fmt.Printf("THING: %#v\n", result)
	fmt.Printf("Revision: %d\n", result.ObjectMeta.Generation)
	fmt.Printf("Timeout: %#v\n", *result.Spec.ProgressDeadlineSeconds)

}
