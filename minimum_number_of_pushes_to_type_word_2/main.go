package main

import (
	"runtime/debug"
	"sort"
)

func main() {
	minimumPushes("abcde")
	minimumPushes("xyzxyzxyzxyz")
	minimumPushes("aabbccddeeffgghhiiiiii")
}

func minimumPushes(word string) int {
	cnt := make([]int, 26)
	for _, c := range word {
		cnt[c-'a']++
	}
	sort.Slice(cnt, func(a, b int) bool {
		return cnt[a] > cnt[b]
	})
	ans := 0
	for i, v := range cnt {
		ans += (i/8 + 1) * v
	}
	return ans
}

func init() {
	debug.SetGCPercent(-1)
	debug.FreeOSMemory()
}
