package members

import (
	"encoding/json"
	"os"

	"github.com/go-git/go-git/v5/plumbing/storer"
	"github.com/spf13/viper"
)

var members []Member
var env *Env

type Env struct {
	appEnv	string	`mapstructure:"APP_ENV"`
	store	string	`mapstructure:"STORE"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")
   
	err := viper.ReadInConfig()
	if err != nil {
	 panic(err)
	}
   
	err = viper.Unmarshal(&env)
	if err != nil {
	 panic(err)
	}
   
	return &env
   }

func init() {
	env = NewEnv()
	buf, err := os.ReadFile(env.store)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(buf, &members)
	if err != nil {
		panic(err)
	}
}