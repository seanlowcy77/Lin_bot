package linbot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	api "github.com/go-telegram-bot-api/telegram-bot-api"
)

type linbot struct {
	Name       string
	Bot        bot
	photo      photoconfig
	PokemonAns string
	stickers   stickers
	pictures   pictures
}

type pictures struct {
	sentpictures  []string
	arrofpictures []string
}
type photoconfig interface {
	NewPhotoShare(chatID int64, fileID string) api.PhotoConfig
	NewPhotoUpload(chatID int64, file interface{}) api.PhotoConfig
}

type bot interface {
	Send(c api.Chattable) (api.Message, error)
	GetUpdatesChan(config api.UpdateConfig) (api.UpdatesChannel, error)
	UploadFile(endpoint string, params map[string]string, fieldname string, file interface{}) (api.APIResponse, error)
	GetUserProfilePhotos(config api.UserProfilePhotosConfig) (api.UserProfilePhotos, error)
	KickChatMember(config api.KickChatMemberConfig) (api.APIResponse, error)
	SetWebhook(config api.WebhookConfig) (api.APIResponse, error)
	RemoveWebhook() (api.APIResponse, error)
}

// SendTextMessage sends a basic text message back to the specified user.
func (lin *linbot) SendTextMessage(recipient int, text string) error {
	msg := api.NewMessage(int64(recipient), text)
	msg.ReplyMarkup = api.NewRemoveKeyboard(true)
	msg.ParseMode = "Markdown"
	_, err := lin.Bot.Send(msg)
	return err
}

// Wrapper struct for a message
type message struct {
	Cmd  string
	Args []string
	*api.Message
}

type stickers struct {
	sentstickers []string
}

// InitLinBot initialises the bot
func InitLinBot() *linbot {

	bot, err := api.NewBotAPI("946928172:AAGst6bzVn4wb6_5j79DFYq-0yBPHedsMQk")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	lin := &linbot{
		Name: "Lin",
		Bot:  bot,
	}
	lin.PokemonAns = " "
	lin.Setpictures()
	return lin
}

// Listen exposes the telebot Listen API.
func (lin *linbot) Listen(timeout int) api.UpdatesChannel {
	u := api.NewUpdate(0)
	u.Timeout = timeout
	updates, err := lin.Bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalf("error creating updates channel: %s", err)
	}
	return updates
}

//Sends a Start message
func (lin *linbot) Start(msg *api.Message) {

	text := "Hello there " + msg.From.FirstName + "!\n\n" +
		"Im Linbot! YEEEEEEET!\n" +
		"Im always here to /help if you need it!"

	lin.SendTextMessage(int(msg.Chat.ID), text)
}

//Help releases list of commands for bot
func (lin *linbot) Help(msg *api.Message) {

	text := "UwU I'm really useful if you need my help " + msg.From.FirstName + " \n" +
		"Here's some things I can do :)" + "\n" + "\n" +
		"add a number to the the /time function and I'll do a timer for you!" + "\n" +
		"Feel free to play around with my other functions!" + "\n" +
		"/aboutme" + "\n" +
		"/lintime" + "\n" +
		"/wufan" + "\n" +
		"/shutup" + "\n" +
		"/echo" + "\n" +
		"/slurp" + "\n" +
		"/uwu" + "\n" + "\n" +
		"These are some of my SpeCiaL functions :D Try them out!" + "\n" +
		"Wanna find out your gender? Try /male or /female !!!" + "\n" +
		"Who I find /cute ???" + "\n" +
		"Lets play /WhosethatPokemon !!!"

	lin.SendTextMessage(int(msg.Chat.ID), text)
}

//Pokemons struct of an array of Pokemon
type Pokemons struct {
	Pokemons []Pokemon `json:"pokemonsters"`
}

//Pokemon struct
type Pokemon struct {
	Abilities string `json:"abilities"`
	Name      string `json:"Name"`
	Pic       string `json:"ThumbnailImage"`
}

//randomInt function to generate random ints
func randomInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

// pokemon function to generate a random picture of a pokemon
func (lin *linbot) Pokemon(msg *api.Message) {
	jsonFile, err := os.Open("pokemon.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var file Pokemons
	json.Unmarshal([]byte(byteValue), &file)

	rand.Seed(time.Now().UnixNano())
	i := randomInt(0, 386)

	newmsg := api.NewPhotoShare(int64(msg.Chat.ID), file.Pokemons[i].Pic)
	lin.PokemonAns = file.Pokemons[i].Name

	log.Println(lin.PokemonAns)
	_, errr := lin.Bot.Send(newmsg)
	lin.SendTextMessage(int(msg.Chat.ID), "Whose that Pokemon?!"+"\n"+
		"Type the command /pokemon along with your answer!"+"\n"+
		"If you're lousy click here for the right /answer :/")
	if errr != nil {
		fmt.Println(errr)
	}
}

// Returns a picture of user
func (lin *linbot) Handsome(msg *api.Message) {
	photo := api.NewUserProfilePhotos(msg.From.ID)
	photos, err := lin.Bot.GetUserProfilePhotos(photo)
	if err != nil {
		fmt.Println(err)
	}
	newmsg := api.NewPhotoShare(int64(msg.Chat.ID), photos.Photos[0][0].FileID)
	_, errr := lin.Bot.Send(newmsg)
	lin.SendTextMessage(int(msg.Chat.ID), "Its you, "+msg.From.FirstName+"!!!")
	if errr != nil {
		fmt.Println(errr)
	}
}

func (lin *linbot) Kick(msg *api.Message) {
	wuf := api.ChatMemberConfig{
		int64(-1001383326579),
		strconv.Itoa(int(-1001383326579)),
		"",
		int(249291763),
	}
	kicking := api.KickChatMemberConfig{
		wuf,
		int64(1),
	}
	resp, err := lin.Bot.KickChatMember(kicking)
	if err != nil {
		fmt.Println(err)

	}
	lin.SendTextMessage(int(msg.Chat.ID), resp.Description)

}

// Send text message when command male is entered
func (lin *linbot) Male(msg *api.Message) {
	text := "Hello? Gender is a spectrum okay" + "\n" + "*Gender fluid D a B*" + "\n" +
		"Down with the PaTriAchY >:( Not everyone is solely 1 gender." + "\n" + "\n" +
		"click /here to find out yours"
	lin.SendTextMessage(int(msg.Chat.ID), text)
}

// Returns a random number% when male
func (lin *linbot) Malegenderchecker(msg *api.Message) {
	rand.Seed(time.Now().UnixNano())
	i := randomInt(0, 100)
	text := "According to my CaLcuLatioNSSSSS" + "\n" +
		msg.From.FirstName + " ,you are " + strconv.Itoa(i) + "% male!!"
	lin.SendTextMessage(int(msg.Chat.ID), text)

}

// Send text message when command female is entered
func (lin *linbot) Female(msg *api.Message) {
	text := "You go girl!"
	lin.SendTextMessage(int(msg.Chat.ID), text)
}

// Setting the pictures that will be sent by the /lintime function
func (lin *linbot) Setpictures() {
	lin.pictures.arrofpictures = []string{"../pics/pic1.jpeg",
		"../pics/pic2.jpeg", "../pics/pic3.jpeg", "../pics/pic4.jpeg"}
}

// Returns a different picture of Lin each time function is called
func (lin *linbot) Lintime(msg *api.Message) {
	arr := lin.pictures.arrofpictures
	n := len(arr)
	i := randomInt(0, n)
	Sentpictures := lin.pictures.sentpictures
	for Contains(Sentpictures, arr[i]) {
		i = randomInt(0, n)
	}
	Sentpictures = append(Sentpictures, arr[i])
	newmsg := api.NewPhotoUpload(int64(msg.Chat.ID), arr[i])
	_, errr := lin.Bot.Send(newmsg)
	if errr != nil {
		fmt.Println(errr)
	}
}

// Function when the answer is correctly given
func (lin *linbot) Rightpokemon(msg *api.Message) {
	str := strings.Replace(msg.Text, "/pokemon ", "", 1)
	if strings.ToLower(str) == strings.ToLower(lin.PokemonAns) {
		lin.SendTextMessage(int(msg.Chat.ID), "That's right its "+lin.PokemonAns+"!!! WOOOOOOO"+"\n"+"/playagain ?")
	} else {
		lin.SendTextMessage(int(msg.Chat.ID), "Wrong :( U dont know your pokemon?!"+"\n"+
			"This is a helpful link to train your skills! "+"https://play.pokemonshowdown.com/"+"\n"+
			"If you really need help I'll let u know the /answer")
	}
}

func (lin *linbot) Aboutme(msg *api.Message) {
	str := "I was create by Sean YEEEEET" + "\n" + "https://www.facebook.com/sean.low.54"
	lin.SendTextMessage(int(msg.Chat.ID), str)

}
func (lin *linbot) Status(msg *api.Message) {
	str := "I love yeeting with you " + msg.Chat.FirstName
	lin.SendTextMessage(int(msg.Chat.ID), str)
}

func (lin *linbot) Uwu(msg *api.Message) {
	str := ""
	if msg.From.FirstName == "Wu Fan" {
		str = "This is especially for u <3 Hey UWU Fan"
	} else {
		str = "Hey " + msg.From.FirstName + " UWU"
	}
	lin.SendTextMessage(int(msg.Chat.ID), str)
}

// Funtion to send a different sticker each time from an array of stickers
func (lin *linbot) Sendsticker(msg *api.Message) {
	stickerarr := []string{"CAADAgADzWoBAAFji0YMJh7SqwnpNXQWBA", "CAADAgADk10BAAFji0YMrp5MBok7V1oWBA",
		"CAADAgADlV0BAAFji0YM4jBzLzwi3FYWBA", "CAADAgAD3nABAAFji0YMLpR9QayvR8oWBA"}
	n := len(stickerarr)
	i := randomInt(0, n)
	if len(lin.stickers.sentstickers) == n {
		lin.stickers.sentstickers = []string{}
	}
	for Contains(lin.stickers.sentstickers, stickerarr[i]) {
		i = randomInt(0, n)
	}
	lin.stickers.sentstickers = append(lin.stickers.sentstickers, stickerarr[i])
	newmsg := api.NewStickerShare(int64(msg.Chat.ID), stickerarr[i])
	_, errr := lin.Bot.Send(newmsg)
	if errr != nil {
		fmt.Println(errr)
	}
}

// Function to echo something
func (lin *linbot) Echo(msg *api.Message) {
	str1 := strings.Replace(msg.Text, "/echo", "", 1)
	lin.SendTextMessage(int(msg.Chat.ID), str1)
}

// Function to provide a countdown
func (lin *linbot) Time(msg *api.Message) {
	str1 := strings.Replace(msg.Text, "/time ", "", 1)
	seconds, err := strconv.Atoi(str1)
	if err != nil {
		lin.SendTextMessage(int(msg.Chat.ID), "Invalid number of seconds")
	} else if seconds < 100 {
		reply := " Timer for " + str1 + " seconds has started"
		lin.SendTextMessage(int(msg.Chat.ID), reply)
		time.Sleep(time.Duration(seconds) * time.Second)
		lin.SendTextMessage(int(msg.Chat.ID), "Time out!")
	} else {
		lin.SendTextMessage(int(msg.Chat.ID), "Its too long :'(")
	}
}

func (lin *linbot) Surprise(msg *api.Message) {
	text := "YEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEET"
	lin.SendTextMessage(int(msg.Chat.ID), text)

}

// Contains Checking if a string array contains a string
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
