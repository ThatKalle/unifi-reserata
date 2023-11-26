//go:generate goversioninfo

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/dixonwille/wmenu/v5"
)

type userInput struct {
	option wmenu.Opt
}

func (u *userInput) optFunc(option wmenu.Opt) error {
	u.option = option
	return nil
}

func createMenu(p string, m []string, u *userInput) {
	menu := wmenu.NewMenu(p)
	menu.ChangeReaderWriter(os.Stdin, os.Stdout, os.Stderr)
	for i, m := range m {
		menu.Option(m, i, false, u.optFunc)
	}

	err := menu.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func findFiles(p string, r string) ([]string, error) {
	var foundFiles []string
	dir, err := os.Open(p)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	files, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for _, f := range files {
		matched, _ := regexp.Match(r, []byte(f.Name()))
		if matched {
			foundFiles = append(foundFiles, f.Name())
		}
	}

	return foundFiles, nil
}

func AESDecrypt(ciphertext []byte) ([]byte, error) {
	key, _ := hex.DecodeString("626379616e676b6d6c756f686d617273")
	iv, _ := hex.DecodeString("75626e74656e74657270726973656170")

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("block size cant be zero")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	return ciphertext, nil
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	files, err := findFiles(path, "(.*?)\\.(supp|unf)$")
	if err != nil {
		fmt.Println(err)
	}

	if len(files) < 1 {
		fmt.Println("No UniFi files found in the current directory.\nLooking for .supp or .unf files.")
		return
	}

	u := &userInput{}
	createMenu("Select a File", files, u)

	selectedfile := u.option.Text
	file, err := os.ReadFile(selectedfile)
	if err != nil {
		fmt.Printf("Error reading file %s: %s\n", selectedfile, err)
		return
	}

	decrypted, err := AESDecrypt(file)
	if err != nil {
		fmt.Println("Error during decryption:", err)
		return
	}

	err = os.WriteFile((fmt.Sprintf("%s.zip", selectedfile)), decrypted, 0666)
	if err != nil {
		fmt.Println("Error during file write:", err)
		return
	}
}
