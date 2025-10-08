package arrays

import "fmt"

func Array_main() {
	/* Define an array of size 4 that stores deployment options */
	var DeploymentOptions = [4]string{"R-pi", "AWS", "GCP", "Azure"}

	/* Loop through the deployment options array */
	for i := 0; i < len(DeploymentOptions); i++ {
		option := DeploymentOptions[i]
		fmt.Println(i, option)
	}
}
