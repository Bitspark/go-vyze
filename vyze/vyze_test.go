package vyze

import (
	"github.com/Bitspark/go-vyze/core"
	"github.com/Bitspark/go-vyze/service"
	"github.com/Bitspark/go-vyze/system"
	"testing"
)

type PersonalUser struct {
	ID           core.ID  `json:"id"`
	Name         string   `json:"name"`
	DisplayName  string   `json:"displayName"`
	Password     string   `json:"password"`
	PrimaryEmail string   `json:"primaryEmail"`
	Username     string   `json:"username"`
	Emails       []string `json:"emails"`
}

func TestGetObject1(t *testing.T) {
	v := NewClient(service.NewServiceClient("http://localhost:9150"), system.NewSystemClient("http://localhost:9131"))
	v.Service.SetToken("pjhpSnJpm-0CeqqZc-GDWJv1Y2MAAAAA____f_8BAAUAAAAFKi8qLyqqLIzePO8pQIQ7MK-RPeLD5x4iew")
	_, err := v.LoadUniverse("user")
	v.SelectedUniverse = "user"
	if err != nil {
		t.Fatal(err)
	}
	lp, err := v.Service.GetLayerProfile(core.MustParseID("d19af73a2b3e919f4915ca14d260a14b"))
	if err != nil {
		t.Fatal(err)
	}
	v.System.SetLayerProfile(lp)
	v.System.SetDefaultOptions(&system.AccessOptions{
		Access:      "main_full",
		AccessNames: []string{"main_full"},
	})

	t.Log("FILTERED")

	objs, err := QueryReference[PersonalUser](v, "getPersonalUser").Equals("username", "juliann").GetObjects()
	if err != nil {
		t.Fatal(err)
	}
	for _, obj := range objs {
		t.Log(obj)
	}

	t.Log("SORTED")

	objs, err = QueryReference[PersonalUser](v, "getPersonalUser").Sort("username", false).GetObjects()
	if err != nil {
		t.Fatal(err)
	}
	for _, obj := range objs {
		t.Log(obj)
	}

	t.Log("SLICED")

	objs, err = QueryReference[PersonalUser](v, "getPersonalUser").Sort("username", true).Slice(0, 1).GetObjects()
	if err != nil {
		t.Fatal(err)
	}
	for _, obj := range objs {
		t.Log(obj)
	}
}

func TestPutObjects(t *testing.T) {
	v := NewClient(service.NewServiceClient("http://localhost:9150"), system.NewSystemClient("http://localhost:9131"))
	v.Service.SetToken("pjhpSnJpm-0CeqqZc-GDWJv1Y2MAAAAA____f_8BAAUAAAAFKi8qLyqqLIzePO8pQIQ7MK-RPeLD5x4iew")
	_, err := v.LoadUniverse("user")
	v.SelectedUniverse = "user"
	if err != nil {
		t.Fatal(err)
	}
	lp, err := v.Service.GetLayerProfile(core.MustParseID("d19af73a2b3e919f4915ca14d260a14b"))
	if err != nil {
		t.Fatal(err)
	}
	v.System.SetLayerProfile(lp)
	v.System.SetDefaultOptions(&system.AccessOptions{
		Access:      "main_full",
		AccessNames: []string{"main_full"},
	})

	v.System.DeleteObject(core.MustParseID("4a2875b25f08b9bb43aae525f2108cc2"), nil)
	v.System.DeleteObject(core.MustParseID("bceecc839e532a24725db0c487a1bbbc"), nil)

	//objs, err := QueryReference[PersonalUser](v, "putPersonalUser").PutObjects([]PersonalUser{
	//	{
	//		Name:        "new user",
	//		DisplayName: "new user",
	//		Password:    "testabc",
	//		Username:    "new user",
	//	},
	//})
	//if err != nil {
	//	t.Fatal(err)
	//}
	//for _, obj := range objs {
	//	t.Log(obj)
	//}
}

func TestObjects1(t *testing.T) {
	client := NewClient(
		service.NewServiceClient("https://api.vyze.io/service"),
		system.NewSystemClient("https://api.vyze.io/system"),
	)

	client.Service.SetToken("...")
	lp, _ := system.ReadLayerProfile("...")
	client.System.SetLayerProfile(lp)
	client.System.SetDefaultOptions(&system.AccessOptions{
		Access:      "main_full",
		AccessNames: []string{"main_full", "main_read", "model_read", "model_extend"},
	})
	_, _ = client.LoadUniverse("vergleichsportal")

	objs, err := QueryReference[any](client, "getCrawljob").GetObjects()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(objs)
}
