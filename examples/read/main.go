package main

import (
	"github.com/Bitspark/go-vyze"
	"log"
)

func main() {
	serviceClient := go_vyze.NewServiceClient("https://api.vyze.io/service")
	systemClient := go_vyze.NewSystemClient("https://api.vyze.io/system")

	vyzeClient := go_vyze.NewClient(serviceClient, systemClient)
	// vyzeClient.Service.SetToken("...") // Set a token if you are using a private universe
	err := vyzeClient.LoadUniverse("example")
	if err != nil {
		log.Fatal(err)
	}

	profile, err := vyzeClient.Service.GetLayerProfile(go_vyze.MustParseID("df9b7b765c12082786df03ce98d3e6f4")) // Create a layer profile on the vyze website and put its ID here
	if err != nil {
		log.Fatal(err)
	}

	vyzeClient.System.SetLayerProfile(profile)
	vyzeClient.System.SetDefaultOptions(&go_vyze.AccessOptions{
		AccessNames: []string{"main_read"},
	})

	documents, err := go_vyze.GetSpecials[Document](vyzeClient, "getDocument")
	if err != nil {
		log.Fatal(err)
	}

	// Should print "Document 1", "Document 2"
	for _, doc := range documents {
		log.Println(doc.Name)
	}
}
