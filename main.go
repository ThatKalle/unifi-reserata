//go:generate goversioninfo

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"os"
	"regexp"
)

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
	if len(os.Args) < 2 {
		fmt.Printf("Missing parameter, provide .supp or .unf file name.\nExample: %s network_support_20-11-2023.supp\n", os.Args[0])
		return
	}
	selectedfile := fmt.Sprintf("%s", os.Args[1])
	matched, _ := regexp.Match("(.*?)\\.(supp|unf)$", []byte(selectedfile))
	if matched == false {
		fmt.Printf("Provided parameter %s is not a .supp or .unf file name.\nExample: %s network_support_20-11-2023.supp\n", selectedfile, os.Args[0])
		return
	}

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
