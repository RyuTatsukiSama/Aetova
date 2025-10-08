package arrays

import "fmt"

func Array_refactored() {
	DeploymentOptions := [...]string{"R-pi", "AWS", "GCP", "Azure"}

	for index, option := range DeploymentOptions {
		fmt.Println(index, option)
	}
}
