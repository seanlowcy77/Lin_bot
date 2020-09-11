# Lin Bot

## 1. Introduction

A telegram bot coded during my own free time to experiment with Golang as well as some of the Golang Telegram Bot APIs.
The App has the following notable fun features amongst others:

```
1. /wufan which sends out a random sticker from a set of several stickers

2. /WhosethatPokemon which allows the user to play a pokemon guessing game!

3. An automatic 'kick' function which kicks a particular user out of the group should another user leave / get kicked

```

## 2. Quick Start

1. Ensure that you have installed golang. More info may be found [here](https://golang.org/doc/install)
2. Change the `Telegram_Key` variable in `linbot.go`
3. enter `go run main.go` in the terminal

### Customizable Features

1. `Telegram_Key` variable in `linbot.go`
2. Number of pokemon to be included in `linbot.go` under the `Pokemon` function
3. `chatId` variable in the `Kick` function in `linbot.go` whcih is the chat where the user is to be kicked from as well as the `userToBeKickedId` which is the id of the user to be kicked
4. Sticker Ids in the `Sendsticker` function in `linbot.go`
5. `Uwu` function in `linbot.go` which can be changed to send a customizable message to other users

### Note

-   Sticker Ids may be found through [this bot](https://t.me/idstickerbot?start=botostore)
-   ChatIds and UserIds may be found via the console when you run this bot

## 3. Demo of Features

### WhoseThatPokemon

  <img src="https://github.com/seanlowcy77/Lin_bot/blob/master/DemoPics/DemoPokemonPass.png" width="200" />
A successful guess of the pokemon shown

  <img src="https://github.com/seanlowcy77/Lin_bot/blob/master/DemoPics/DemoPokemonFail.png" width="200" />
An unsuccessful guess of the pokemon shown

### Kicking of another User

  <img src="https://github.com/seanlowcy77/Lin_bot/blob/master/DemoPics/DemoKick.png" width="150" />
The bot kicking someone when a person leaves / gets kicked from the group

## Further Questions?

This project was made by Sean just for fun / playing around with the Golang Telegram bot API. I will be happy to answer any questions if any. Thank you!
