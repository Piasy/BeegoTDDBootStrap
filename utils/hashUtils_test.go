package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Piasy/BeegoTDDBootStrap/utils"
)

func TestSha1(t *testing.T) {
	input := "123456"
	hash := "7c4a8d09ca3762af61e59520943dc26494f8941b"
	assert.Equal(t, hash, utils.Sha1(input))
	assert.Equal(t, "adc83b19e793491b1c6ea0fd8b46cd9f32e592fc", utils.Sha1(""))
}

func TestSecret2Password(t *testing.T) {
	username := "phone:18812345678"
	//origin := "123456abc"
	// sha1( 7c4a8d09ca3762af61e59520943dc26494f8941b + 18812345678 + a9993e364706816aba3e25717850c26c9cd0d89d )
	secret := "7b26ca1c7e6df85a7ede775c4f1e0e86abbe2425"
	assert.Equal(t, secret, utils.Sha1(utils.Sha1("123456") + "18812345678" + utils.Sha1("abc")))
	// sha1( sha1(7b26ca1c) + sha1(phone:18812345678) + sha1(7e6df85a7ede775c4f1e0e86abbe2425) )
	// sha1( dbc1e6c151d26dc283b1bf44babb1856763181eb + 512f34fdb790a6c710b032367b6d23a8a3795479 + f2ed1ea23aba8d4a074146d2b0bb671bb7354ad2 )
	password := "7da0c3ad2f3334a84bfa224e1b4431e14fe7f851"
	assert.Equal(t, password, utils.Secret2Password(username, secret))

	username2 := "phone:18812345678"
	//origin := "123456"
	// sha1( 7c4a8d09ca3762af61e59520943dc26494f8941b + 18812345678 + adc83b19e793491b1c6ea0fd8b46cd9f32e592fc )
	secret2 := "f7b3c18749f0b77b8993646018af9daab6dee948"
	assert.Equal(t, secret2, utils.Sha1(utils.Sha1("123456") + "18812345678" + utils.Sha1("")))
	// sha1( sha1(f7b3c187) + sha1(phone:18812345678) + sha1(49f0b77b8993646018af9daab6dee948) )
	// sha1( f54f913524b238cfc84ebd78d66f5fdc8d0579c1 + 512f34fdb790a6c710b032367b6d23a8a3795479 + 54938bf9b16ed14a14c61a6ddd49456b9119ab6f )
	password2 := "d346f02f85cbba9dc662a67f44d222c397712242"
	assert.Equal(t, password2, utils.Secret2Password(username2, secret2))
}