package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/bsync-tech/util/xalgorithms/bf"
	"github.com/bsync-tech/util/xcompressions"
	"github.com/bsync-tech/util/xconcurrency"
	"github.com/bsync-tech/util/xconditions"
	"github.com/bsync-tech/util/xconversions"
	"github.com/bsync-tech/util/xencodings"
	"github.com/bsync-tech/util/xerrors"
	"github.com/bsync-tech/util/xhashes"
	"github.com/bsync-tech/util/xmanipulations"
	"github.com/bsync-tech/util/xmetrics/cpu"
	"github.com/bsync-tech/util/xmetrics/disk"
	"github.com/bsync-tech/util/xmetrics/memory"
	"github.com/bsync-tech/util/xos"
	"github.com/bsync-tech/util/xstrings"
)

func main() {
	Hashes()
	Errors()
	Strings()
	Encodings()
	Conditions()
	Conversions()
	Manipulations()
	Compressions()
	Concurrency()
	BF()
	Cpu()
	Memory()
	Disk()
	ListSubFiles()
}

func Conversions() {
	x := map[string]interface{}{"number": 1, "string": "cool", "bool": true, "float": 1.5}
	fmt.Println(xconversions.PrettyJson(x))
	fmt.Println(xconversions.Stringify(x))

	data := "{\"bool\":true,\"float\":1.5,\"number\":1,\"string\":\"cool\"}"
	var results map[string]interface{}
	fmt.Println(xconversions.Structify(data, &results))
	fmt.Println(results)
}

func Errors() {
	fmt.Println(xerrors.DefaultErrorIfNil(nil, "Cool"))
	fmt.Println(xerrors.DefaultErrorIfNil(xerrors.New("Oops"), "Cool"))
}

func Strings() {
	fmt.Println(xstrings.IsEmpty(""))
	fmt.Println(xstrings.IsEmpty("text"))
	fmt.Println(xstrings.IsEmpty("	"))

	fmt.Println(xstrings.IsNotEmpty(""))
	fmt.Println(xstrings.IsNotEmpty("text"))
	fmt.Println(xstrings.IsNotEmpty("	"))

	fmt.Println(xstrings.IsBlank(""))
	fmt.Println(xstrings.IsBlank("	"))
	fmt.Println(xstrings.IsBlank("text"))

	fmt.Println(xstrings.IsNotBlank(""))
	fmt.Println(xstrings.IsNotBlank("	"))
	fmt.Println(xstrings.IsNotBlank("text"))

	fmt.Println(xstrings.Left("", 5))
	fmt.Println(xstrings.Left("X", 5))
	fmt.Println(xstrings.Left("😎⚽", 4))
	fmt.Println(xstrings.Left("ab\u0301cde", 8))

	fmt.Println(xstrings.Right("", 5))
	fmt.Println(xstrings.Right("X", 5))
	fmt.Println(xstrings.Right("😎⚽", 4))
	fmt.Println(xstrings.Right("ab\u0301cde", 8))

	fmt.Println(xstrings.Center("", 5))
	fmt.Println(xstrings.Center("X", 5))
	fmt.Println(xstrings.Center("😎⚽", 4))
	fmt.Println(xstrings.Center("ab\u0301cde", 8))

	fmt.Println(xstrings.Length(""))
	fmt.Println(xstrings.Length("X"))
	fmt.Println(xstrings.Length("b\u0301"))
	fmt.Println(xstrings.Length("😎⚽"))
	fmt.Println(xstrings.Length("Les Mise\u0301rables"))
	fmt.Println(xstrings.Length("ab\u0301cde"))
	fmt.Println(xstrings.Length("This `\xc5` is an invalid UTF8 character"))
	fmt.Println(xstrings.Length("The quick bròwn 狐 jumped over the lazy 犬"))

	fmt.Println(xstrings.Reverse(""))
	fmt.Println(xstrings.Reverse("X"))
	fmt.Println(xstrings.Reverse("😎⚽"))
	fmt.Println(xstrings.Reverse("Les Mise\u0301rables"))
	fmt.Println(xstrings.Reverse("This `\xc5` is an invalid UTF8 character"))
	fmt.Println(xstrings.Reverse("The quick bròwn 狐 jumped over the lazy 犬"))
	fmt.Println(xstrings.Reverse("رائد شوملي"))
}

func Conditions() {
	fmt.Println(xconditions.IfThen(1 == 1, "Yes"))
	fmt.Println(xconditions.IfThen(1 != 1, "Woo"))
	fmt.Println(xconditions.IfThen(1 < 2, "Less"))

	fmt.Println(xconditions.IfThenElse(1 == 1, "Yes", false))
	fmt.Println(xconditions.IfThenElse(1 != 1, nil, 1))
	fmt.Println(xconditions.IfThenElse(1 < 2, nil, "No"))

	fmt.Println(xconditions.DefaultIfNil(nil, nil))
	fmt.Println(xconditions.DefaultIfNil(nil, ""))
	fmt.Println(xconditions.DefaultIfNil("A", "B"))
	fmt.Println(xconditions.DefaultIfNil(true, "B"))
	fmt.Println(xconditions.DefaultIfNil(1, false))

	fmt.Println(xconditions.FirstNonNil(nil, nil))
	fmt.Println(xconditions.FirstNonNil(nil, ""))
	fmt.Println(xconditions.FirstNonNil("A", "B"))
	fmt.Println(xconditions.FirstNonNil(true, "B"))
	fmt.Println(xconditions.FirstNonNil(1, false))
	fmt.Println(xconditions.FirstNonNil(nil, nil, nil, 10))
	fmt.Println(xconditions.FirstNonNil(nil, nil, nil, nil, nil))
	fmt.Println(xconditions.FirstNonNil())
}

func Compressions() {
	fmt.Println(xcompressions.Compress([]byte("Raed Shomali")))
}

func Encodings() {
	fmt.Println(xencodings.Base32Encode([]byte("Raed Shomali")))
	fmt.Println(xencodings.Base64Encode([]byte("Raed Shomali")))
}

func Hashes() {
	fmt.Println(xhashes.FNV32("Raed Shomali"))
	fmt.Println(xhashes.FNV32a("Raed Shomali"))
	fmt.Println(xhashes.FNV64("Raed Shomali"))
	fmt.Println(xhashes.FNV64a("Raed Shomali"))
	fmt.Println(xhashes.MD5("Raed Shomali"))
	fmt.Println(xhashes.SHA1("Raed Shomali"))
	fmt.Println(xhashes.SHA256("Raed Shomali"))
	fmt.Println(xhashes.SHA512("Raed Shomali"))
}

func Manipulations() {
	source := rand.NewSource(time.Now().UnixNano())

	array := []interface{}{"a", "b", "c"}
	xmanipulations.Shuffle(array, source)

	fmt.Println(array)
}

func Concurrency() {
	func1 := func() {
		for char := 'a'; char < 'a'+3; char++ {
			fmt.Printf("%c ", char)
		}
	}

	func2 := func() {
		for number := 1; number < 4; number++ {
			fmt.Printf("%d ", number)
		}
	}

	xconcurrency.Parallelize(func1, func2)
}

func Cpu() (string, error) {
	before, err := cpu.Get()
	if err != nil {
		return "", err
	}
	time.Sleep(time.Duration(1) * time.Second)
	after, err := cpu.Get()
	if err != nil {
		return "", err
	}
	total := float64(after.Total - before.Total)
	user := fmt.Sprintf("%.2f%%", float64(after.User-before.User)/total*100)
	system := fmt.Sprintf("%.2f%%", (float64(after.System-before.System)/total)*100)
	idle := fmt.Sprintf("%.2f%%", (float64(after.Idle-before.Idle)/total)*100)
	info := map[string]interface{}{
		"User":   user,
		"System": system,
		"Idle":   idle,
	}
	d, err := json.Marshal(&info)
	if err != nil {
		return "", err
	}
	fmt.Println(string(d))
	return string(d), nil
}

func Memory() (string, error) {
	memory, err := memory.Get()
	if err != nil {
		return "", err
	}
	total := fmt.Sprintf("%dMB", memory.Total/1024/1024)
	used := fmt.Sprintf("%dMB", memory.Used/1024/1024)
	cached := fmt.Sprintf("%dMB", memory.Cached/1024/1024)
	free := fmt.Sprintf("%dMB", memory.Free/1024/1024)
	info := map[string]interface{}{
		"Total":  total,
		"Used":   used,
		"Cached": cached,
		"Free":   free,
	}
	d, err := json.Marshal(&info)
	if err != nil {
		return "", err
	}
	fmt.Println(string(d))
	return string(d), nil
}

func Disk() (string, error){
	before, err := disk.Get()
	begin := time.Now()
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	time.Sleep(time.Duration(1000) * time.Millisecond)
	after, err := disk.Get()
	elapse := time.Since(begin)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	ds := make([]*map[string]interface{}, len(after))
	for i := 0; i < len(after); i++ {
		info := map[string]interface{}{
			"DeviceName": after[i].DeviceName,
			"ReadUtil":   fmt.Sprintf("%.2f%%", 100*float64(after[i].TimeSpentReading-before[i].TimeSpentReading)/float64(elapse.Milliseconds())),
			"WriteUtil":  fmt.Sprintf("%.2f%%", 100*float64(after[i].TimeSpentWriting-before[i].TimeSpentWriting)/float64(elapse.Milliseconds())),
		}
		ds = append(ds, &info)
	}

	d, err := json.Marshal(&ds)
	if err != nil {
		return "", err
	}
	fmt.Println(string(d))
	return string(d), nil
}

func ListSubFiles() {
	fmt.Println(xos.ListSubFiles(".", xos.MODE_DIR))
}

func BF() {
	key1 := []byte{0x00, 0x01, 0x02, 0x03}
	key2 := []byte{0x00, 0x01, 0x02, 0x04}
	s := "hello this is a simple world"
	bf.WithKey(key1)
	fmt.Println(string(bf.EncryptData([]byte(s))))
	bf.WithKey(key2)
	fmt.Println(string(bf.EncryptData([]byte(s))))
}
