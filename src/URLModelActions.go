package tnyuri

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"
	"strconv"

	base58 "github.com/itchyny/base58-go"
)

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

func (U *URL) isValid() bool {

	if len(U.Url) <= 5 || len(U.User) <= 3 {
		return false
	} else {
		return true
	}
}

func (U *URL) hash() {
	generatedNumber := new(big.Int).SetBytes(sha256Of(U.Url + U.User)).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	U.Short = finalString[:5]
}

func (U *URL) Save() int {
	if U.isValid() {
		U.hash()
		var id int = Insert("insert into url (url, short, user) VALUES ('" + U.Url + "','" + U.Short + "','" + U.User + "')")
		Insert("insert into stats (url_id, counter) VALUES ('" + strconv.Itoa(id) + "','0')")
		return id
	} else {
		fmt.Println("URL Object:")
		fmt.Println(U)
		fmt.Println("does not meet the requirements")
	}

	return -1
}

func (U *URL) Stats() Stats {
	var stats []Stats = By[Stats]("url_id", strconv.Itoa(U.Id))

	if len(stats) > 0 {
		return stats[0]
	}

	var stat Stats = *new(Stats)
	stat.Counter = 0
	stat.Uid = U.Id

	return stat
}
