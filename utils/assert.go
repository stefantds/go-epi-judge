package utils

import "fmt"

func AssertAllValuesPresent(reference []int, result []int) error {
	referenceSet := make(map[int]int)
	for x := range reference {
		referenceSet[x] = referenceSet[x] + 1
	}

	for x := range result {
		referenceSet[x] = referenceSet[x] - 1
	}

	excessItems := make([]int, 0)
	missingItems := make([]int, 0)

	for x, count := range referenceSet {
		if count < 0 {
			for count < 0 {
				excessItems = append(excessItems, x)
				count -= 1
			}
		} else if count > 0 {
			for count > 0 {
				missingItems = append(missingItems, x)
			}
		}
	}

	if len(excessItems) > 0 {
		return fmt.Errorf("value set changed: found excess items %v", excessItems)
	}

	if len(missingItems) > 0 {
		return fmt.Errorf("value set changed: found missing items %v", missingItems)
	}

	return nil
}
