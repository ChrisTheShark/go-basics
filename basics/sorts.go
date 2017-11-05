package basics

import (
	"sort"
	"fmt"
)

type People []string
func (p People) Len() int { return len(p) }
func (p People) Less(i, j int) bool { return p[i] < p[j] }
func (p People) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func main() {
	studyGroup := People{"Zeno", "John", "Al", "Jenny"}
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
	sort.Sort(sort.Reverse(studyGroup))
	fmt.Println(studyGroup)

	names := []string{"Zeno", "John", "Al", "Jenny"}
	sort.Strings(names)
	fmt.Println(names)
	sort.Sort(sort.Reverse(sort.StringSlice(names)))
	fmt.Println(names)

	nums := []int{3, 5, 2, 4, 9}
	sort.Ints(nums)
	fmt.Println(nums)
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println(nums)
}

