package mojoauth

import "testing"
import "unicode/utf8"

func TestCreateSecret(t *testing.T) {
  const length = 93
  if secret := CreateSecret(); utf8.RuneCountInString(secret) != length {
    t.Errorf("CreateSecret created a string that wasn't %v characters", length)
  }
}
