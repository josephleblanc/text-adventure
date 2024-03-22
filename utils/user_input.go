package utils

import "text-adventure/mytypes"

func InitUserInput() []mytypes.UserInput {
	name_input := mytypes.UserInput{
		Selection:  "",
		PreVerify:  "Your name is ",
		PostVerify: ", is that right?",
		ToSelect:   "name",
		FieldName:  "Name",
	}

	all_inputs := []mytypes.UserInput{name_input}
	return all_inputs
}

func MapUserInput(all_inputs []mytypes.UserInput) map[string]*mytypes.UserInput {
	input_map := make(map[string]*mytypes.UserInput)

	for _, value := range all_inputs {
		input_map[value.ToSelect] = &value
	}
	return input_map
}

// References for using maps in go
//https://gobyexample.com/maps
// On using structs in maps
//https://golangspot.com/golang-map-of-struct-with-examples/
//
