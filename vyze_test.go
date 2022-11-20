package vyze

import (
	"testing"
)

type PersonalUser struct {
	ID           ID       `json:"id"`
	Name         string   `json:"name"`
	DisplayName  string   `json:"displayName"`
	Password     string   `json:"password"`
	PrimaryEmail string   `json:"primaryEmail"`
	Username     string   `json:"username"`
	Emails       []string `json:"emails"`
}

func TestGetObject1(t *testing.T) {
	v := NewClient(NewServiceClient("http://localhost:9150"), NewSystemClient("http://localhost:9131"))
	v.Service.SetToken("pjhpSnJpm-0CeqqZc-GDWJv1Y2MAAAAA____f_8BAAUAAAAFKi8qLyqqLIzePO8pQIQ7MK-RPeLD5x4iew")
	err := v.LoadUniverse("user")
	v.SelectedUniverse = "user"
	if err != nil {
		t.Fatal(err)
	}
	lp, err := v.Service.GetLayerProfile(MustParseID("d19af73a2b3e919f4915ca14d260a14b"))
	if err != nil {
		t.Fatal(err)
	}
	v.System.SetLayerProfile(lp)
	v.System.SetDefaultOptions(&AccessOptions{
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
	v := NewClient(NewServiceClient("http://localhost:9150"), NewSystemClient("http://localhost:9131"))
	v.Service.SetToken("pjhpSnJpm-0CeqqZc-GDWJv1Y2MAAAAA____f_8BAAUAAAAFKi8qLyqqLIzePO8pQIQ7MK-RPeLD5x4iew")
	err := v.LoadUniverse("user")
	v.SelectedUniverse = "user"
	if err != nil {
		t.Fatal(err)
	}
	lp, err := v.Service.GetLayerProfile(MustParseID("d19af73a2b3e919f4915ca14d260a14b"))
	if err != nil {
		t.Fatal(err)
	}
	v.System.SetLayerProfile(lp)
	v.System.SetDefaultOptions(&AccessOptions{
		Access:      "main_full",
		AccessNames: []string{"main_full"},
	})

	v.System.DeleteObject(MustParseID("4a2875b25f08b9bb43aae525f2108cc2"), nil)
	v.System.DeleteObject(MustParseID("bceecc839e532a24725db0c487a1bbbc"), nil)

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
