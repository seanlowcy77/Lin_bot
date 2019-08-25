package linbot123

import (
	"log"

	api "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	bot, err := api.NewBotAPI("946928172:AAGst6bzVn4wb6_5j79DFYq-0yBPHedsMQk")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := api.NewUpdate(0)
	u.Timeout = 100

	updates, err := bot.GetUpdatesChan(u)
	lin := &linbot{
		Name: "Lin",
		bott: bot,
	}
	lin.PokemonAns = " "

	for update := range updates {

		/*if strings.Contains(strings.ToLower(update.Message.Text), "slurp") {
			lin.SendTextMessage(int(update.Message.Chat.ID), "YEAAAA ")
		}*/
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

		if err != nil {
			log.Panic(err)
		}
		bot.Send(msg)
	}
}

/*package main

import (
	"log"
	"strings"

	api "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	//tele := new(telekey)
	//tele.setkey()
	bot, err := api.NewBotAPI("946928172:AAGst6bzVn4wb6_5j79DFYq-0yBPHedsMQk")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := api.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	lin := &linbot{Name: "Lin", bott: bot}
	lin.PokemonAns = " "
	lin.setpictures()

	for update := range updates {

		if strings.Contains(strings.ToLower(update.Message.Text), "slurp") {
			lin.SendTextMessage(int(update.Message.Chat.ID), "YEAAAA ")
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
		}

		if err != nil {
			log.Panic(err)
		}
		bot.Send(msg)
	}
}
*/
