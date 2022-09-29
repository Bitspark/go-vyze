package vyze

import (
	"crypto/rand"
	"encoding/json"
	"gopkg.in/yaml.v3"
	"strings"
	"testing"
	"time"
)

func TestBinary_Int(t *testing.T) {
	d := Binary{}

	// Positive

	number := int64(0)

	err := d.FromInt64(number)
	if err != nil {
		t.Fatal(err)
	}
	if len(d) != 0 {
		t.Fatal("should have length 0")
	}

	number = int64(543864736554)

	err = d.FromInt64(number)
	if err != nil {
		t.Fatal(err)
	}

	i, err := d.ToInt64()
	if err != nil {
		t.Fatal(err)
	}

	if i != number {
		t.Fatal()
	}

	// Negative

	number = int64(-553)

	err = d.FromInt64(number)
	if err != nil {
		t.Fatal(err)
	}

	i, err = d.ToInt64()
	if err != nil {
		t.Fatal(err)
	}

	if i != number {
		t.Fatal()
	}
}

func TestBinary_Uint(t *testing.T) {
	d := Binary{}

	number := uint64(0)

	err := d.FromUint(number)
	if err != nil {
		t.Fatal(err)
	}
	if len(d) != 0 {
		t.Fatal("should have length 0")
	}

	number = uint64(543864736554)

	err = d.FromUint(number)
	if err != nil {
		t.Fatal(err)
	}

	i, err := d.ToUint()
	if err != nil {
		t.Fatal(err)
	}

	if i != number {
		t.Fatal()
	}

	number = uint64(344)

	err = d.FromUint(number)
	if err != nil {
		t.Fatal(err)
	}

	i, err = d.ToUint()
	if err != nil {
		t.Fatal(err)
	}

	if i != number {
		t.Fatal()
	}
}

func TestBinary_Float64(t *testing.T) {
	d := Binary{}

	number := float64(0)

	err := d.FromFloat64(number)
	if err != nil {
		t.Fatal(err)
	}
	if len(d) != 0 {
		t.Fatal("should have length 0")
	}

	number = 543864.736554

	err = d.FromFloat64(number)
	if err != nil {
		t.Fatal(err)
	}

	i, err := d.ToFloat64()
	if err != nil {
		t.Fatal(err)
	}

	if i != number {
		t.Fatal()
	}

	number = 45.4

	err = d.FromFloat64(number)
	if err != nil {
		t.Fatal(err)
	}

	i, err = d.ToFloat64()
	if err != nil {
		t.Fatal(err)
	}

	if i != number {
		t.Fatal()
	}
}

func TestBinary_Time(t *testing.T) {
	d := Binary{}

	now := time.Now()
	err := d.FromTime(now)
	if err != nil {
		t.Fatal(err)
	}

	err = d.FromTime(time.Date(2000, 0, 0, 0, 0, 0, 0, time.Local))
	if err != nil {
		t.Fatal(err)
	}

	i, err := d.ToTime()
	if err != nil {
		t.Fatal(err)
	}

	if i.Equal(now) {
		t.Fatal()
	}

	err = d.FromTime(time.Date(2020, 0, 0, 0, 0, 0, 0, time.Local))
	if err != nil {
		t.Fatal(err)
	}

	i, err = d.ToTime()
	if err != nil {
		t.Fatal(err)
	}

	if i.Equal(now) {
		t.Fatal()
	}
}

func TestBinary_Boolean(t *testing.T) {
	d := Binary{}

	err := d.FromBoolean(true)
	if err != nil {
		t.Fatal(err)
	}

	err = d.FromBoolean(false)
	if err != nil {
		t.Fatal(err)
	}

	i, err := d.ToBoolean()
	if err != nil {
		t.Fatal(err)
	}

	if i {
		t.Fatal()
	}

	err = d.FromBoolean(true)
	if err != nil {
		t.Fatal(err)
	}

	i, err = d.ToBoolean()
	if err != nil {
		t.Fatal(err)
	}

	if !i {
		t.Fatal()
	}
}

func TestBinary_String(t *testing.T) {
	d := Binary{}

	err := d.FromString("test123")
	if err != nil {
		t.Fatal(err)
	}

	if len(d) != 7 {
		t.Fatal()
	}

	err = d.FromString("false")
	if err != nil {
		t.Fatal(err)
	}

	i, err := d.ToString()
	if err != nil {
		t.Fatal(err)
	}

	if i != "false" {
		t.Fatal()
	}

	err = d.FromString("gfgäö")
	if err != nil {
		t.Fatal(err)
	}

	i, err = d.ToString()
	if err != nil {
		t.Fatal(err)
	}

	if i != "gfgäö" {
		t.Fatal()
	}
}

func TestBinary_Hex(t *testing.T) {
	d := Binary{0, 1, 255}

	if len(d.Hex()) != 6 {
		t.Fatal()
	}
	if d.Hex()[0:2] != "00" {
		t.Fatal()
	}
	if d.Hex()[2:4] != "01" {
		t.Fatal()
	}
	if d.Hex()[4:6] != "ff" {
		t.Fatal()
	}
}

func TestBinary_Base64(t *testing.T) {
	d := Binary{}
	_ = d.FromString("TEST")

	if d.Base64() != "VEVTVA" {
		t.Fatal()
	}
}

func TestBinary_JSON(t *testing.T) {
	a := Binary{}
	b := Binary{}
	c := Binary{}

	a.FromBoolean(true)
	b.FromUint(12)
	c.FromString("hello world")

	aB, _ := json.Marshal(a)
	bB, _ := json.Marshal(b)
	cB, _ := json.Marshal(c)

	a = Binary{}
	b = Binary{}
	c = Binary{}

	json.Unmarshal(aB, &a)
	json.Unmarshal(bB, &b)
	json.Unmarshal(cB, &c)

	if v, _ := a.ToBoolean(); v != true {
		t.Fatal()
	}
	if v, _ := b.ToUint(); v != 12 {
		t.Fatal()
	}
	if v, _ := c.ToString(); v != "hello world" {
		t.Fatal()
	}
}

func TestBinary_YAML(t *testing.T) {
	a := Binary{}
	b := Binary{}
	c := Binary{}

	a.FromBoolean(true)
	b.FromUint(12)
	c.FromString("hello world")

	aB, _ := yaml.Marshal(a)
	bB, _ := yaml.Marshal(b)
	cB, _ := yaml.Marshal(c)

	a = Binary{}
	b = Binary{}
	c = Binary{}

	yaml.Unmarshal(aB, &a)
	yaml.Unmarshal(bB, &b)
	yaml.Unmarshal(cB, &c)

	if v, _ := a.ToBoolean(); v != true {
		t.Fatal()
	}
	if v, _ := b.ToUint(); v != 12 {
		t.Fatal()
	}
	if v, _ := c.ToString(); v != "hello world" {
		t.Fatal()
	}
}

func Test_EncodeLength(t *testing.T) {
	b := encodeLength(0)
	if len(b) != 1 {
		t.Fatal()
	}
	b = encodeLength(128)
	if len(b) != 2 {
		t.Fatal()
	}
	b = encodeLength(128 * 128)
	if len(b) != 3 {
		t.Fatal()
	}
	b = encodeLength(128 * 128 * 128)
	if len(b) != 4 {
		t.Fatal()
	}
}

func Test_DecodeLength(t *testing.T) {
	b, l := decodeLength(encodeLength(0))
	if b != 0 || l != 1 {
		t.Fatal()
	}
	b, l = decodeLength(encodeLength(1))
	if b != 1 || l != 1 {
		t.Fatal()
	}
	b, l = decodeLength(encodeLength(127))
	if b != 127 || l != 1 {
		t.Fatal()
	}
	b, l = decodeLength(encodeLength(128))
	if b != 128 || l != 2 {
		t.Fatal()
	}
	b, l = decodeLength(encodeLength(128 * 128))
	if b != 128*128 || l != 3 {
		t.Fatal()
	}
	b, l = decodeLength(encodeLength(128 * 128 * 128 * 128))
	if b != 128*128*128*128 || l != 5 {
		t.Fatal()
	}

	for i := int64(0); i < 1024; i++ {
		b, _ = decodeLength(encodeLength(i))
		if b != i {
			t.Fatal()
		}
	}

	for i := int64(0); i < 1024*1024; i += 1023 {
		b, _ = decodeLength(encodeLength(i))
		if b != i {
			t.Fatal()
		}
	}

	for i := int64(0); i < 1024*1024*1024; i += 1023 * 1023 {
		b, _ = decodeLength(encodeLength(i))
		if b != i {
			t.Fatal()
		}
	}
}

func TestBinary_Encrypt(t *testing.T) {
	k := Key{}
	_, _ = rand.Read(k[:])

	d := Binary{}
	_ = d.FromString("Hallo Welt!")
	if s, _ := d.ToString(); s != "Hallo Welt!" {
		t.Fatal()
	}
	d.Encrypt(k)
	if s, _ := d.ToString(); s == "Hallo Welt!" {
		t.Fatal()
	}

	d = Binary{}
	_ = d.FromString(strings.Repeat("Hallo Welt!", 1000))
	if s, _ := d.ToString(); s != strings.Repeat("Hallo Welt!", 1000) {
		t.Fatal()
	}
	d.Encrypt(k)
	if s, _ := d.ToString(); s == strings.Repeat("Hallo Welt!", 1000) {
		t.Fatal()
	}
}

func TestBinary_Decrypt(t *testing.T) {
	k := Key{}
	_, _ = rand.Read(k[:])

	d := Binary{}
	_ = d.FromString("Hallo Welt!")
	d.Encrypt(k)
	d.Decrypt(k)
	if s, _ := d.ToString(); s != "Hallo Welt!" {
		t.Fatal()
	}

	d = Binary{}
	_ = d.FromString(strings.Repeat("Hallo Welt!", 1000))
	d.Encrypt(k)
	d.Decrypt(k)
	if s, _ := d.ToString(); s != strings.Repeat("Hallo Welt!", 1000) {
		t.Fatal()
	}
}
