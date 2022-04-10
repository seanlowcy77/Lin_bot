package linbot

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/seanlowcy77/Lin_bot/linbot"
)

func main() {

	lin := linbot.InitLinBot()
	updates := lin.Listen(60)

	for update := range updates {

		// If a user gets kicked / leaves, kicks another user of your choice
		name := update.Message.LeftChatMember
		if name != nil {
			lin.Kick(update.Message)
		}
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		msg := api.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "aboutme":
			lin.Aboutme(update.Message)
		case "start":
			lin.Start(update.Message)
		case "help":
			lin.Help(update.Message)
		case "surprise":
			lin.Surprise(update.Message)
		case "talk":
			msg.Text = "You want me to talk? Nahhh"
		case "status":
			lin.Status(update.Message)
		case "shutup":
			msg.Text = ">:( Lin bot smash"
		case "echo":
			lin.Echo(update.Message)
		case "time":
			lin.Time(update.Message)
		case "uwu":
			lin.Uwu(update.Message)
		case "lintime":
			lin.Lintime(update.Message)
		case "slurp":
			msg.Text = "YUMZ"
		case "WhosethatPokemon":
			lin.Pokemon(update.Message)
		case "playagain":
			lin.Pokemon(update.Message)
		case "pokemon":
			lin.Rightpokemon(update.Message)
		case "answer":
			msg.Text = "Answer is " + lin.PokemonAns + "!!!" + "\n" + "Better luck next time " +
				update.Message.From.FirstName + " :(" +
				"\n" + "/playagain"
		case "male":
			lin.Male(update.Message)
		case "female":
			lin.Female(update.Message)
		case "here":
			lin.Malegenderchecker(update.Message)
		case "cute":
			lin.Handsome(update.Message)
		case "icebear":
			lin.Sendsticker(update.Message)
		case "wufan":
			lin.Sendsticker(update.Message)
		case "zijun":
			lin.Kick(update.Message)
		}
		lin.Bot.Send(msg)
	}
}
