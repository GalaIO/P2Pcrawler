package test

import (
	"encoding/hex"
	"fmt"
	"github.com/GalaIO/P2Pcrawler/dht"
	"log"
	"testing"
)

func TestDecode(t *testing.T) {
	raw := "64313a7264323a696432303a3c00727348b3b8ed70baa1e1411b3869d8481321353a6e6f6465733230383ae362cc8bc798e9b616348dd482139ac5f5a24c838e04d163c8d5e362cc8bc798e9b616348dd482139ac5f5a24c838e04d163c8d5e362cc8bc798e9b616348dd482139ac5f5a24c838e04d163c8d5e362cc8bc798e9b616348dd482139ac5f5a24c838e04d163c8d5e362cc8bc798e9b616348dd482139ac5f5a24c838e04d163c8d5e362cc8bc798e9b616348dd482139ac5f5a24c838e04d163c8d5e362cc8bc798e9b616348dd482139ac5f5a24c838e04d163c8d5e362cc8bc798e9b616348dd482139ac5f5a24c838e04d163c8d565313a74323a6161313a76343a4a420000313a79313a7265"

	bytes, err := hex.DecodeString(raw)
	if err != nil {
		log.Fatal(err)
	}

	dicts, err := dht.DecodeDict(string(bytes))
	if err != nil {
		log.Fatal(err)
	}

	body := dicts.GetDict("r")
	fmt.Println(hex.EncodeToString([]byte(body.GetString("id"))))
	fmt.Println(hex.EncodeToString([]byte(body.GetString("nodes"))))
}
