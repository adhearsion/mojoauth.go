package mojoauth

import "github.com/dchest/uniuri"

func CreateSecret() string {
  return uniuri.NewLen(93)
}
