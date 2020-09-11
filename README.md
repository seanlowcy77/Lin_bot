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

## 4. Requirements

-   IOS 13.0+
-   Running with XCode 11.5

## 5. Things to Note

### a. Possible Future Extensions

1. As mentioned earlier, for the `check` function, a rather simple algorithm was used to check if the contents of the text match the name / description of each medication in the database. Perhaps this algorithm could be improved in future especially since medication names and descriptions are difficult to remember.

2. While only basic text recognition is used in this protoype, it is interesting to note that for Apple Vision, there are already APIs available from the [NSDataDetector](https://developer.apple.com/documentation/foundation/nsdatadetector) class to detect properties such as

```
1. phone numbers
2. dates
3. addresses
4. links
5. transit information
```

More explanation of these APIs may be found in this [link](https://nshipster.com/nsdatadetector/).

Thus a possible future extension would be to code an algorithm that can effectively detect / seperate the name and description of medications which would greatly enhance the flow of the app.

3. Another possible extension would be to add more variables such as `frequency of medication` or the `prescription start / end date` of the medication which will be helpful in reminding patients when to take their medications as well as to identify them.

### b. General Comments

1. The deletion function was not implemented as I felt that patients should not be able to delete their medication after it is recorded.

## Further Questions?

This basic prototype was made by Sean (Intern) in July 2020. I will be happy to answer any questions you have - on Slack, [Github](https://github.com/seanlowcy77) or otherwise. Thank you!
