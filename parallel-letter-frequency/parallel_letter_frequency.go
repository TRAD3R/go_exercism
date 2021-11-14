package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	rchan := make(chan FreqMap)
	freqMap := FreqMap{}

	for _, line := range l {
		go func(line string, rchan chan<- FreqMap) {

			rchan <- Frequency(line)

		}(line, rchan)
	}

	for i := 0; i < len(l); i++ {
		for key, value := range <-rchan {
			freqMap[key] += value

		}
	}

	return freqMap
}
