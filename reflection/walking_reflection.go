package reflection

func walk(x interface{}, fn func(input string)) {
	// walk the interface, find all strings, run the func on all the found strings
	fn("Roger the gopher")
}
