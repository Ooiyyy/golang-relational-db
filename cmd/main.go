package main

import bootstrap "golang-relational-db/bootsrap"

func main() {
	// Titik masuk proses HTTP app: semua wiring dilakukan di bootstrap.Run().
	bootstrap.Run()
}
