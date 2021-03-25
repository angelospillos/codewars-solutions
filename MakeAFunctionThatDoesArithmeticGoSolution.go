package kata

func Arithmetic(a int, b int, operator string) int{
  
  result := 0
  
  switch operator {
  case "add":
		result = a + b
	case "subtract":
		result = a - b
  case "multiply":
    result = a * b
	default:
    result = a / b
	}

  return result
}
