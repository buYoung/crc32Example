package main

import (
"fmt"
"io"
"encoding/hex"
"hash/crc32"
)

func main() {
	fmt.Println(hash_password_crc32("test"))
  fmt.Println(hash_file_crc32("/home/pi/test.txt"))
}

func hash_file_crc32(filePath string) (string, error) {
	var returnCRC32String string
	file, err := os.Open(filePath)
	if err != nil {
		return returnCRC32String, err
	}
	defer file.Close()
	tablePolynomial := crc32.MakeTable(crc32.IEEE)
	hash := crc32.New(tablePolynomial)
	if _, err := io.Copy(hash, file); err != nil {
		return returnCRC32String, err
	}
	hashInBytes := hash.Sum(nil)[:]
	returnCRC32String = hex.EncodeToString(hashInBytes)
	return returnCRC32String, nil
}
func hash_password_crc32(password string) (string, error) {
	var returnCRC32String string

	tablePolynomial := crc32.MakeTable(crc32.IEEE)
	hash := crc32.New(tablePolynomial)
	if _, err := hash.Write([]byte(password)); err != nil {
		return returnCRC32String, err
	}
	hashInBytes := hash.Sum(nil)[:]
	returnCRC32String = hex.EncodeToString(hashInBytes)
	return returnCRC32String, nil
}
