package main

import "log"

func main() {
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	animals := []string{"dog", "cat", "cow"}
	for idx, animal := range animals {
		log.Println(idx, animal)
	}

	toys := make(map[string]string)
	toys["marvel"] = "cap america"
	toys["dc"] = "superman"
	for comic, char := range toys {
		log.Println(comic, char)
	}
}