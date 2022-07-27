package hamming

import "errors"

func Distance(a, b string) (int, error) {
	count := 0
	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				count += 1
			}
		}
		return count, nil
	}
	return count, errors.New("Error")
}
