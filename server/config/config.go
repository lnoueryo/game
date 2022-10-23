package config

import (
	"log"
	"os"
)

var Infolog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
var Errorlog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
