//go:build darwin
// +build darwin

package keychain

import (
	"fmt"
	"testing"
)

func setUp() {

	// Create keychain
	// security create-keychain -p "" "blah"

	// return keychain
}

func tearDown() {

	// Delete keychain
	// security delete-keychain MyNew.keychain

}

func TestValidOpenKeychain(t *testing.T) {

	testKeychainPath := "/Users/uttie/Library/Keychains/login.keychain"
	testKeychain := Keychain{path: testKeychainPath}
	kref, err := openKeychainRef(testKeychain.path)
	if err != nil {
		t.Fatal(err)
	}

	if kref == 0 {
		t.Fatal("keychain is null")
	}

	fmt.Println(testKeychain.Status())
	//releaseKey(kref)
}
