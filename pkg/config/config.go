package config

import (
	"log"
	"os"
)

type vars struct {
	TwitterKey    string
	TwitterSecret string
}

func Load() vars {
	var v vars
	load("TWITTER_KEY", &v.TwitterKey, "", true)
	load("TWITTER_SECRET", &v.TwitterSecret, "", true)
	return v
}

func load(key string, v *string, def string, req bool) {
	val := os.Getenv(key)
	if val == "" && def == "" && req {
		log.Fatalf("required variable %q is not set", key)
	}
	if val != "" {
		*v = val
		log.Printf("Loaded envar %s with value %s", key, *v)
		return
	}
	*v = def
	log.Printf("Loaded envar %s with value %s", key, *v)
}
