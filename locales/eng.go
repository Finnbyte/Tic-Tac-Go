package locales

import "tictacgo/funcs"

var (
	EngLang Texts = Texts{
		Lang:     "eng",
		Greeting: "\nGreetings, welcome to TicTacGo!",
		Side:     "\nYou start with " + string(funcs.GetSide()) + "!\n",
		Begin:    "\nLets begin! You can type commands to the prompt below. Get help with command \"help\"!",
		Prompt:   "Enter a command: ",
		Exit:     "Byebye, thanks for checking my project out! <3",
		Help:     "",
		Lost:     "You lost the game against the computer... :^(",
		Won:      "you won the game!!!",
		Draw:     "Your game has ended in a draw! Nice try!",
	}
)
