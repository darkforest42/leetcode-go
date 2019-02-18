package leetcode451

func FrequencySort(s string) string {
	bytesForString := []byte(s)
	frequencyForString := make(map[byte]int, len(s))
	var frequency int
	var ok bool
	for _, b := range bytesForString {
		if frequency, ok = frequencyForString[b]; !ok {
			frequency = 0
		}
		frequencyForString[b] = frequency + 1
	}
	bucketsForByte := make([][]byte, len(bytesForString)+1)
	for b, frequency := range frequencyForString {
		bucketsForByte[frequency] = append(bucketsForByte[frequency], b)
	}

	tmp := []byte{}
	for i := len(bucketsForByte)-1; i >=0 ; i-- {
		if len(bucketsForByte[i]) == 0 {
			continue
		}
		for _, b := range bucketsForByte[i] {
			for j := i; j > 0; j-- {
				tmp = append(tmp, b)
			}
		}
	}
	return string(tmp)
}