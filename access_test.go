package vyze

import (
	"encoding/hex"
	"testing"
)

func TestReadLayerToken(t *testing.T) {
	lt1, err := ReadLayerToken("971f8594101187a92e866d2a3c8211ec21658b2d047d281cc515c7ff90670b5501ffffff00000000000000000000000062b786c17ffffffff3caa6adc7a83eea1780e9b68f9dac8ef22f0ce7")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(lt1.UserID)
	t.Log(lt1.LayerID)
	t.Log(lt1.Granted)
	t.Log(lt1.Created.Unix())
	t.Log(lt1.Expiry.Seconds())
	t.Log(lt1.Expires().Unix())
	t.Log(hex.EncodeToString(lt1.Signature))
}

func TestReadLayerToken__1(t *testing.T) {
	td, err := ReadLayerToken("971f8594101187a92e866d2a3c8211ec21658b2d047d281cc515c7ff90670b5501ffffff00000000000000000000000062b786c17ffffffff3caa6adc7a83eea1780e9b68f9dac8ef22f0ce7")
	if err != nil {
		t.Fatal(err)
	}
	if td.Granted != uint32(PermAll) {
		t.Fatal()
	}
}

func TestReadLayerToken__2(t *testing.T) {
	userID, _ := ParseID("971f8594101187a92e866d2a3c8211ec")
	layerID, _ := ParseID("0fc053e11df777aa9f0e9b2480e14020")
	td, err := ReadLayerToken("971f8594101187a92e866d2a3c8211ec0fc053e11df777aa9f0e9b2480e1402001ffffff00000000000000000000000062b7874f7fffffff9af7d3f0d91ad83d31129a12c5909c186bb53851")
	if err != nil {
		t.Fatal(err)
	}
	if td.UserID != userID {
		t.Fatal()
	}
	if td.LayerID != layerID {
		t.Fatal()
	}
	if td.Granted != uint32(PermAll) {
		t.Fatal()
	}
}

func TestReadLayerToken__Expired(t *testing.T) {
	_, err := ReadLayerToken("971f8594101187a92e866d2a3c8211ec21658b2d047d281cc515c7ff90670b5501ffffff00000000000000000000000062b786cb0000003b2dd92c2276bae1a023a67c35f3890ad3fee0d060")
	if err == nil {
		t.Fatal()
	}
}

func TestReadAccessGroup(t *testing.T) {
	ag := newAccessGroup("test", uint32(PermAll))

	lt1, _ := ReadLayerToken("971f8594101187a92e866d2a3c8211ec0fc053e11df777aa9f0e9b2480e1402001ffffff00000000000000000000000062bc55007fffffff2848959e1c1b7fbaf5ca2ac0dbeeaa81f3bbaea2")
	err := ag.RegisterLayerToken(lt1)
	if err != nil {
		t.Fatal(err)
	}

	lt2, _ := ReadLayerToken("971f8594101187a92e866d2a3c8211ec21658b2d047d281cc515c7ff90670b5501ffffff00000000000000000000000062bc550f7fffffff8f8f653c5aa188701ee1a2ea369e206c39adf94d")
	err = ag.RegisterLayerToken(lt2)
	if err != nil {
		t.Fatal(err)
	}

	ag2, err := ReadAccessGroup(ag.String())
	if err != nil {
		t.Fatal(err)
	}
	if ag2.Name != "test" {
		t.Fatal()
	}
	if ag2.Permissions != uint32(PermAll) {
		t.Fatal()
	}
	if len(ag2.Tokens) != 2 {
		t.Fatal()
	}
}

func TestReadLayerProfile(t *testing.T) {
	lp := NewLayerProfile()
	ag1, _ := lp.AddAccessGroup("a", uint32(PermAll))
	lt1, _ := ReadLayerToken("971f8594101187a92e866d2a3c8211ec21658b2d047d281cc515c7ff90670b5501ffffff00000000000000000000000062b786d87fffffffe6679ef20c46f438ce07b2f84e2d032f6922d210")
	_ = ag1.RegisterLayerToken(lt1)
	ag2, _ := lp.AddAccessGroup("b", uint32(PermView))
	lt2, _ := ReadLayerToken("971f8594101187a92e866d2a3c8211ec0fc053e11df777aa9f0e9b2480e14020004924ea00000000000000000000000062b786ea7fffffffabff5481b09e8c08a3d8aeb663194db2bf9baed9")
	_ = ag2.RegisterLayerToken(lt2)

	lp2, err := ReadLayerProfile(lp.String())
	if err != nil {
		t.Fatal()
	}
	ag3 := lp2.GetAccessGroup("b")
	if ag3 == nil {
		t.Fatal()
	}
	if ag3.Permissions != uint32(PermView) {
		t.Fatal()
	}
	if len(ag3.Tokens) != 1 {
		t.Fatal()
	}
}
