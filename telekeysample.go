package main

// Sample use of telekey
// Change telekeysample to telekey and enter Telegram key in the " " below

type telekeysample struct {
	key string
}

func (tele telekeysample) setkey() {
	tele.key = "Enter key here"
}
