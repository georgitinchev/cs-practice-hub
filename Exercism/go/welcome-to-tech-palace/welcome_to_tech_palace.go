package techpalace

import "strings"

func WelcomeMessage(customer string) string {
	customer = strings.ToUpper(customer)
	return "Welcome to the Tech Palace, " + customer
}

func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := strings.Repeat("*", numStarsPerLine)
	return border + "\n" + welcomeMsg + "\n" + border
}

func CleanupMessage(oldMsg string) string {
oldMsg = strings.ReplaceAll(oldMsg, "*", "")
oldMsg = strings.TrimSpace(oldMsg)
return oldMsg
}
