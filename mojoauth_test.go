package mojoauth

import (
	"fmt"
	"testing"
	"time"
	"unicode/utf8"
)

func TestCreateSecret(t *testing.T) {
	const length = 93
	if secret := CreateSecret(); utf8.RuneCountInString(secret) != length {
		t.Errorf("CreateSecret created a string that wasn't %v characters", length)
	}
}

func TestCorrectCredentialsWithId(t *testing.T) {
	const user_id = "mojo"
	secret := CreateSecret()
	username, password := CreateCredentials(user_id, secret)
	result, id := TestCredentials(username, password, secret)
	if result == false {
		t.Errorf("Credentials should have been valid")
	}
	if id != user_id {
		t.Errorf("ID should have been %s, instead it was %s", user_id, id)
	}
}

func TestCorrectCredentialsWithIdAndTTL(t *testing.T) {
	const user_id = "mojo"
	secret := CreateSecret()
	username, password := CreateCredentials(user_id, secret, 10000)
	result, id := TestCredentials(username, password, secret)
	if result == false {
		t.Errorf("Credentials should have been valid")
	}
	if id != user_id {
		t.Errorf("ID should have been %s, instead it was %s", user_id, id)
	}
}

func TestCorrectCredentialsWithoutId(t *testing.T) {
	const user_id = ""
	secret := CreateSecret()
	username, password := CreateCredentials(user_id, secret)
	result, id := TestCredentials(username, password, secret)
	if result == false {
		t.Errorf("Credentials should have been valid")
	}
	if id != user_id {
		t.Errorf("ID should have been %s, instead it was %s", user_id, id)
	}
}

func TestIncorrectCredentialsWithId(t *testing.T) {
	const user_id = "mojo"
	secret := CreateSecret()
	username, password := CreateCredentials(user_id, secret)

	result, _ := TestCredentials(username, "password", secret)
	if result == true {
		t.Errorf("Credentials should have been invalid with a wrong password")
	}

	result, _ = TestCredentials("username", password, secret)
	if result == true {
		t.Errorf("Credentials should have been invalid with a wrong username")
	}

	result, _ = TestCredentials(username, password, "this is wrong")
	if result == true {
		t.Errorf("Credentials should have been invalid with a different secret")
	}

	new_timestamp := time.Date(2020, time.October, 10, 23, 0, 0, 0, time.UTC)
	forged_username := fmt.Sprintf("%d:%s", new_timestamp.Unix(), user_id)
	result, _ = TestCredentials(forged_username, password, secret)
	if result == true {
		t.Errorf("Credentials should have been invalid with an username with a forged TTL")
	}

	old_timestamp := int(time.Date(2010, time.October, 10, 23, 0, 0, 0, time.UTC).Unix())
	result, _ = TestCredentials(forged_username, password, secret, old_timestamp)
	if result == true {
		t.Errorf("Credentials should have been invalid with a timestamp that is expired")
	}
}
