package main

type Config struct {
	AuthData struct {
		ComsumerKey    string `toml:"ComsumerKey"`
		ComsumerSecret string `toml:"ComsumerSecret"`
		AccessToken    string `toml:"AccessToken"`
		AccessSecret   string `toml:"AccessSecret"`
		UserID         int    `toml:"UserID"`
	} `toml:"AuthData"`
}
