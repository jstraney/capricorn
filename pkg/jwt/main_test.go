package jwt

import "testing"

// tests
func TestGenerateJWT(t *testing.T) {

  token, err := GenerateJWT("jstraney", true)

  if err != nil {
    t.Errorf("%v", err)
  }

  t.Logf("jwt good! %v", token)

}

