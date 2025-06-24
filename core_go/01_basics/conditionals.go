package basics

func conditionals(){
	var age int = 30
	var isAdult bool

	if age >= 18 {
		isAdult = true
	} else {
		isAdult = false
	}

	if isAdult {
		println("You are an adult.")
	} else {
		println("You are not an adult.")
	}

	switch age {
	case 0, 1, 2:
		println("You are a toddler.")
	case 3, 4, 5:
		println("You are a preschooler.")
	case 6, 7, 8, 9, 10:
		println("You are a child.")
	default:
		println("You are an adult or older.")
	}
}