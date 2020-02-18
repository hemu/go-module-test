package calc

func Add(nums ...int) int {
  sum := 0
  for _, num := range nums {
    sum += num
  }
  return sum
}

func Sub(a int, b int) int {
  return a - b
}
