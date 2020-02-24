package main

import (

	api "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	
	lin := InitLinBot();
	updates := lin.Listen(60);
	

	for update := range updates {

		name := update.Message.LeftChatMember
		if name != nil {
			lin.kick(update.Message)
		}
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		msg := api.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "aboutme":
			lin.aboutme(update.Message)
		case "start":
			lin.Start(update.Message)
		case "help":
			lin.Help(update.Message)
		case "surprise":
			lin.Surprise(update.Message)
		case "talk":
			msg.Text = "You want me to talk? Nahhh"
		case "status":
			lin.status(update.Message)
		case "shutup":
			msg.Text = ">:( Lin bot smash"
		case "echo":
			lin.echo(update.Message)
		case "time":
			lin.time(update.Message)
		case "uwu":
			lin.uwu(update.Message)
		case "lintime":
			lin.lintime(update.Message)
		case "slurp":
			msg.Text = "YUMZ"
		case "WhosethatPokemon":
			lin.pokemon(update.Message)
		case "playagain":
			lin.pokemon(update.Message)
		case "pokemon":
			lin.rightpokemon(update.Message)
		case "answer":
			msg.Text = "Answer is " + lin.PokemonAns + "!!!" + "\n" + "Better luck next time " +
				update.Message.From.FirstName + " :(" +
				"\n" + "/playagain"
		case "male":
			lin.male(update.Message)
		case "female":
			lin.female(update.Message)
		case "here":
			lin.malegenderchecker(update.Message)
		case "cute":
			lin.handsome(update.Message)
		case "icebear":
			lin.sendsticker(update.Message)
		case "wufan":
			lin.sendsticker(update.Message)
		case "zijun":
			lin.kick(update.Message)
		}
		lin.bot.Send(msg)
	}
}
