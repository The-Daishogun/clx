package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var HOME_DIR, _ = os.UserHomeDir()
var CONFIG_FILES_DIR = filepath.Join(HOME_DIR, ".config", "clx")
var CONFIG_FILE_PATH = filepath.Join(CONFIG_FILES_DIR, "config.json")

type Config struct {
	Token   string `json:"token"`
	ModelID string `json:"model_id"`
}

func GetConfig() *Config {
	data, err := os.ReadFile(CONFIG_FILE_PATH)
	if err != nil {
		return &Config{}
	}
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		println(err, "Failed to read config file. is it valid json?")
		os.Exit(1)
	}
	return &config
}

func (c *Config) WriteConfig() error {
	err := os.MkdirAll(CONFIG_FILES_DIR, 0700)
	if err != nil {
		return err
	}
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	err = os.WriteFile(CONFIG_FILE_PATH, data, 0700)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) getModelID() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("No model selected. which model do you want to use?")
	for idx, model := range GROQ_MODELS {
		fmt.Fprintf(os.Stdout, "%d. %s by %s. context window: %d. id:%s\n", idx+1, model.Name, model.Developer, model.ContextWindow, model.ID)
	}
	fmt.Print("model index:")
	modelIdxString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	modelIdx, err := strconv.Atoi(strings.ReplaceAll(modelIdxString, "\n", ""))
	if err != nil {
		fmt.Println(err, "pls enter valid index.")
		os.Exit(1)
	}
	c.ModelID = GROQ_MODELS[modelIdx-1].ID
}

func (c *Config) getToken() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("no token set. visit https://console.groq.com/keys and generate a token. Enter it below:")
	fmt.Print("token: ")
	token, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	c.Token = strings.ReplaceAll(token, "\n", "")
}

func (c *Config) CheckConfig() error {
	if c.ModelID == "" || c.Token == "" {
		if c.ModelID == "" {
			c.getModelID()
		}
		if c.Token == "" {
			c.getToken()
		}
		fmt.Fprintf(os.Stdout, "config save at %s\n", CONFIG_FILE_PATH)
		return c.WriteConfig()
	}
	return nil
}
