// Time Complexity : O(n) Space Complexity : O(1)

func leastInterval(tasks []byte, n int) int {

	freqs := make([]int, 26)

	for _, t := range tasks {
		freqs[t - 'A'] += 1
	}

	sort.Ints(freqs)

	mostFreqCount := 0
	mostFreq := freqs[25]

	for i := 25; i >= 0; i-- {
		if freqs[i] == mostFreq {
			mostFreqCount++
			
			continue
		}
		break
	}

	res := (mostFreq - 1) * (1 + n) + mostFreqCount
	if res < len(tasks) { res = len(tasks) }

	return res
}
