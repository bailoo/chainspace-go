package main

import (
	"github.com/tav/golly/optparse"
)

const logo = `          [0;1;34;94m█[0m               [0;1;31;91m▀[0m                                             
    [0;1;32;92m▄▄[0;1;36;96m▄[0m   [0;1;35;95m█[0m [0;1;31;91m▄▄[0m    [0;1;36;96m▄▄[0;1;34;94m▄[0m   [0;1;31;91m▄▄[0;1;33;93m▄[0m    [0;1;36;96m▄[0m [0;1;34;94m▄[0;1;35;95m▄[0m    [0;1;33;93m▄[0;1;32;92m▄▄[0m   [0;1;34;94m▄[0;1;35;95m▄▄[0;1;31;91m▄[0m    [0;1;32;92m▄[0;1;36;96m▄▄[0m    [0;1;31;91m▄▄[0;1;33;93m▄[0m    [0;1;36;96m▄[0;1;34;94m▄▄[0m  
   [0;1;32;92m█[0;1;36;96m▀[0m  [0;1;34;94m▀[0m  [0;1;31;91m█▀[0m  [0;1;32;92m█[0m  [0;1;36;96m▀[0m   [0;1;35;95m█[0m    [0;1;32;92m█[0m    [0;1;34;94m█[0;1;35;95m▀[0m  [0;1;31;91m█[0m  [0;1;32;92m█[0m   [0;1;34;94m▀[0m  [0;1;35;95m█[0;1;31;91m▀[0m [0;1;33;93m▀█[0m  [0;1;36;96m▀[0m   [0;1;35;95m█[0m  [0;1;31;91m█[0;1;33;93m▀[0m  [0;1;32;92m▀[0m  [0;1;34;94m█▀[0m  [0;1;31;91m█[0m 
   [0;1;36;96m█[0m      [0;1;33;93m█[0m   [0;1;36;96m█[0m  [0;1;34;94m▄[0;1;35;95m▀▀[0;1;31;91m▀█[0m    [0;1;36;96m█[0m    [0;1;35;95m█[0m   [0;1;33;93m█[0m   [0;1;36;96m▀[0;1;34;94m▀▀[0;1;35;95m▄[0m  [0;1;31;91m█[0m   [0;1;32;92m█[0m  [0;1;34;94m▄▀[0;1;35;95m▀▀[0;1;31;91m█[0m  [0;1;33;93m█[0m      [0;1;35;95m█▀[0;1;31;91m▀▀[0;1;33;93m▀[0m 
   [0;1;34;94m▀[0;1;35;95m█▄[0;1;31;91m▄▀[0m  [0;1;32;92m█[0m   [0;1;34;94m█[0m  [0;1;35;95m▀[0;1;31;91m▄▄[0;1;33;93m▀█[0m  [0;1;36;96m▄▄[0;1;34;94m█▄[0;1;35;95m▄[0m  [0;1;31;91m█[0m   [0;1;32;92m█[0m  [0;1;34;94m▀▄[0;1;35;95m▄▄[0;1;31;91m▀[0m  [0;1;33;93m█[0;1;32;92m█▄[0;1;36;96m█▀[0m  [0;1;35;95m▀▄[0;1;31;91m▄▀[0;1;33;93m█[0m  [0;1;32;92m▀[0;1;36;96m█▄[0;1;34;94m▄▀[0m  [0;1;31;91m▀█[0;1;33;93m▄▄[0;1;32;92m▀[0m 
                                             [0;1;32;92m█[0m                          
                                             [0;1;36;96m▀[0m`

func main() {
	cmds := map[string]func([]string, string){
		"genkeys": cmdGenKeys,
		"init":    cmdInit,
		"run":     cmdRun,
	}
	info := map[string]string{
		"genkeys": "Generate new keys for a node",
		"init":    "Initialise a new Chainspace network",
		"run":     "Run a node in a Chainspace network",
	}
	optparse.Commands("chainspace", "0.0.1", cmds, info, logo)
}
