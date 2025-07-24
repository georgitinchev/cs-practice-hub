package partyrobot

import (
	"fmt"
	"strconv"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

func HappyBirthday(name string, age int) string {
	return "Happy birthday " + name + "! You are now " + strconv.Itoa(age) + " years old!"
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	return "Welcome to my party, " + name + "!\n" +
		"You have been assigned to table " + fmt.Sprintf("%03d", table) + ". Your table is " + direction + ", exactly " +
		strconv.FormatFloat(distance, 'f', 1, 64) + " meters from here.\n" +
		"You will be sitting next to " + neighbor + "."
}
