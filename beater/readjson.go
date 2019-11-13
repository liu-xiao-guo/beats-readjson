package beater

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"time"
	"os/signal"
    "syscall"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/liu-xiao-guo/readjson/config"
)

type Users struct {
    Users []User `json:"users"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
    Name   string `json:"name"`
    Type   string `json:"type"`
    Age    int    `json:"Age"`
    Social Social `json:"social"`
}

// Social struct which contains a
// list of links
type Social struct {
    Facebook string `json:"facebook"`
    Twitter  string `json:"twitter"`
}

// readjson configuration.
type readjson struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

// New creates an instance of readjson.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &readjson{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts readjson.
func (bt *readjson) Run(b *beat.Beat) error {
	logp.Info("readjson is running! Hit CTRL-C to stop it.")
	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	fmt.Println("Path: ", bt.config.Path)
	fmt.Println("Period: ", bt.config.Period)

	
	// Open our jsonFile
	jsonFile, err := os.Open(bt.config.Path)
	// if we os.Open returns an error then handle it
	if err != nil {
    	fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
    var users Users

    json.Unmarshal(byteValue, &users)

    // we iterate through every user within our users array and
    // print out the user Type, their name, and their facebook url
    // as just an example
    for i := 0; i < len(users.Users); i++ {
        fmt.Println("User Type: " + users.Users[i].Type)
        fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
        fmt.Println("User Name: " + users.Users[i].Name)
        fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)

        event := beat.Event{
			Timestamp: time.Now(),
			Fields: common.MapStr {
				"ostype":    	b.Info.Name,
				"name":		users.Users[i].Name,
				"type":		users.Users[i].Type,
				"age":		users.Users[i].Age,
				"social":	users.Users[i].Social,
			},
		}

		bt.client.Publish(event)
    }

    c := make(chan os.Signal)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        os.Exit(1)
    }()

    for {
        fmt.Println("sleeping...")
        time.Sleep(10 * time.Second)
    }	
}

// Stop stops readjson.
func (bt *readjson) Stop() {
	bt.client.Close()
	close(bt.done)
}
