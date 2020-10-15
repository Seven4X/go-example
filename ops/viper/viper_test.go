package config

import (
	"github.com/spf13/viper"
	"testing"
)

/*
export acm_ak="91be99196ef54b"
export acmak="91be99196f4fef54b"
export link_acmak="91be99196fe8dd7d4fef54b"
export link_acm_ak="91be99196f8d4fef54b"
export acm_sk="Mfmv6iZ"


export ACM_AK="da91be99196f8d41b0a45e8dd7d4fef54b"
export ACMAK="da91be991ef54b"
export LINK_ACMAK="da91be991954b"
export LINK_ACM_AK="da91be99ef54b"
export ACM_SK="daMf"
*/
func TestViperUse(t *testing.T) {

	viper.AutomaticEnv()

	get := func(key string) {
		println(key)
		println(viper.Get(key))
		println(viper.GetString(key))
		println("-----------")
	}

	get("LINK_DSN")
	get("acm_ak")
	get("acmak")
	get("link_acmak")
	get("link_acm_ak")
	get("acm_sk")
	println("大写")

	get("ACM_AK")
	get("ACMAK")
	get("LINK_ACMAK")
	get("LINK_ACM_AK")
	get("ACM_SK")

}
func TestViperUseBind(t *testing.T) {

	viper.AutomaticEnv()
	viper.SetEnvPrefix("link")

	get := func(key string) {
		println(key)
		println(viper.Get(key))
		println(viper.GetString(key))
		println("-----------")
	}

	get("acm_ak")
	get("acmak")
	get("link_acmak")
	get("link_acm_ak")
	get("acm_sk")

	println("大写。")
	get("ACM_AK")
	get("ACMAK")
	get("ACMAK")
	get("LACM_AK")
	get("ACM_SK")
}
