package util

import (
	"sort"
	"strconv"
)

func RemoveDuplicatesInt[T int | int8 | int16 | int32 | int64](slice []T) []T {
	unique := make(map[T]bool)
	result := []T{}

	for _, item := range slice {
		if !unique[item] {
			unique[item] = true
			result = append(result, item)
		}
	}
	return result
}

func StrToint[U int | int8 | int16 | int32 | int64](slice []string) ([]U, error) {
	result := make([]U, len(slice))
	for i, v := range slice {

		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		n := U(num)
		result[i] = n
	}
	return result, nil
}

func SortAndRemoveDuplicates(ids []int32) []int32 {
	ids = RemoveDuplicatesInt(ids)
	sort.Slice(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})
	return ids
}
