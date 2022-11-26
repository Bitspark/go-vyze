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
	err := v.LoadUniverse("user")
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

	objs, err := Query[PersonalUser](v, "getPersonalUser").Equals("username", "juliann").GetObjects()
	if err != nil {
		t.Fatal(err)
	}
	for _, obj := range objs {
		t.Log(obj)
	}

	t.Log("SORTED")

	objs, err = Query[PersonalUser](v, "getPersonalUser").Sort("username", false).GetObjects()
	if err != nil {
		t.Fatal(err)
	}
	for _, obj := range objs {
		t.Log(obj)
	}

	t.Log("SLICED")

	objs, err = Query[PersonalUser](v, "getPersonalUser").Sort("username", true).Slice(0, 1).GetObjects()
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
	err := v.LoadUniverse("user")
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

	//objs, err := Query[PersonalUser](v, "putPersonalUser").PutObjects([]PersonalUser{
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
	v := NewClient(service.NewServiceClient("https://api.vyze.io/service"), system.NewSystemClient("https://api.vyze.io/system"))
	v.Service.SetToken("mt82e3R0cSsMOiRS9TLGpDsmgGMAAAAA____fwEABQAAAAUqLyovKnG9TMZbXxCdQ2kNIxKG98dMpLnM")
	err := v.LoadUniverse("vergleichsportal")
	v.SelectedUniverse = "vergleichsportal"
	if err != nil {
		t.Fatal(err)
	}
	lp, err := system.ReadLayerProfile("main_read:4924ea:9adf367b7474712b0c3a2452f532c6a48ce700d22de0697b5f22dacd84842941004924ea0000000000000000000000006380263b7ffffffe2e0b41f0769deb14644244924f394b0e58b48564,9adf367b7474712b0c3a2452f532c6a4962031f02a7ec76482a9a3b1ba4c8909004924ea0000000000000000000000006380263b7ffffffeab623f31975e2a8343291d8b9584e9b432bbe515;model_extend:492cea:9adf367b7474712b0c3a2452f532c6a457780ea25d5c139b1132dd66ecaa910a00492cea0000000000000000000000006380263b7ffffffecf185d2b6a3880e287a99703bc8349802fb6ba2d;main_full:1ffffff:9adf367b7474712b0c3a2452f532c6a4962031f02a7ec76482a9a3b1ba4c890901ffffff0000000000000000000000006380263b7ffffffe4d0c6ca98842bcc16835c2496d16990f67696eaa")
	if err != nil {
		t.Fatal(err)
	}
	v.System.SetLayerProfile(lp)
	v.System.SetDefaultOptions(&system.AccessOptions{
		Access:      "main_full",
		AccessNames: []string{"main_full", "main_read", "model_extend"},
	})

	objs, err := Query[any](v, "getCrawljob").Filter("distributor.id", "eq", "29aaed37553699858a7182994e318f71").GetObjects()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(objs)
}
