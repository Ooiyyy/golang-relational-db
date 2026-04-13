package main

import bootstrap "golang-sekolah/bootsrap"


func main() {
	// Titik masuk proses HTTP app: semua wiring dilakukan di bootstrap.Run().
	bootstrap.Run()
}
