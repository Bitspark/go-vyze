package main

import (
	"github.com/Bitspark/go-vyze"
	"log"
)

func main() {
	serviceClient := vyze.NewServiceClient("https://api.vyze.io/service")
	systemClient := vyze.NewSystemClient("https://api.vyze.io/system")

	vyzeClient := vyze.NewClient(serviceClient, systemClient)
	// vyzeClient.Service.SetToken("...") // Set a token if you are using a private universe
	err := vyzeClient.LoadUniverse("example")
	if err != nil {
		log.Fatal(err)
	}

	profile, err := vyzeClient.Service.GetLayerProfile(vyze.MustParseID("df9b7b765c12082786df03ce98d3e6f4")) // Create a layer profile on the vyze website and put its ID here
	if err != nil {
		log.Fatal(err)
	}

	vyzeClient.System.SetLayerProfile(profile)
	vyzeClient.System.SetDefaultOptions(&vyze.AccessOptions{
		AccessNames: []string{"main_read"},
	})

	documents, err := vyze.GetSpecials[Document](vyzeClient, "getDocument")
	if err != nil {
		log.Fatal(err)
	}

	// Should print "Document 1", "Document 2"
	for _, doc := range documents {
		log.Println(doc.Name)
	}
}
