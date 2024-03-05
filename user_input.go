package main

func initUserInput() []UserInput {
	name_input := UserInput{"", "Your name is ", ", is that right?", "name"}

	all_inputs := []UserInput{name_input}
	return all_inputs
}

func mapUserInput(all_inputs []UserInput) map[string]*UserInput {
	input_map := make(map[string]*UserInput)

	for _, value := range all_inputs {
		input_map[value.to_select] = &value
	}
	return input_map
}
