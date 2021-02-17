package stats

const numRuns = 5

func RunFuncWithRetries(functionToCall func() bool, failureErr error) error {
	for i := 0; i < numRuns; i++ {
		if functionToCall() {
			return nil
		}
	}

	return failureErr
}
