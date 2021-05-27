## Bot

The bot package provides a set of functions that control a basic Twitch.tv chat bot. The package also exposes an interface which can be used to create a custom chat bot. See the following series for a step-by-step tutorial on [Building a Twitch.tv Chat Bot](https://dev.to/foresthoffman/building-a-twitchtv-chat-bot-with-go---part-1-i3k) with this package.

### Installation

Run `go get github.com/foresthoffman/bot`

### Importing

Import this package by including `github.com/foresthoffman/bot` in your import block.

e.g.

```go
package main

import(
    ...
    "github.com/foresthoffman/bot"
)
```

### Usage

Basic usage:

```go
package main

import (
	"github.com/foresthoffman/bot"
	"time"
)

func main() {

	// Replace the channel name, bot name, and the path to the private directory with your respective
	// values.
	myBot := bot.BasicBot{
		Channel:     "twitch",
		MsgRate:     time.Duration(20/30) * time.Millisecond,
		Name:        "TwitchBot",
		Port:        "6667",
		PrivatePath: "../private/oauth.json",
		Server:      "irc.chat.twitch.tv",
	}
	myBot.Start()
}
```

_That's all, enjoy!_
