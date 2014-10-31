package sort

func BucketSort(input []int) []int {
	// assume input ranges from 0 to 1000
	var buckets [1000]int // array are initialized with zero
	for _, val := range input {
		buckets[val]++
	}
	// use counter and pre-allocated array to reduce memory allocation overhead
	// but code gets ugly
	result := make([]int, len(input))
	count := 0
	for i, val := range buckets {
		for j := 0; j < val; func(j, count *int) {
			*j++
			*count++
		}(&j, &count) {
			result[count] = i
		}
	}
	return result
}
