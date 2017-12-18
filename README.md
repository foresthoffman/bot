## TwitchBot

The twitchbot package provides a set of functions that control a basic Twitch.tv chat bot. The package also exposes an interface which can be used to create a custom chat bot.

### Installation

Run `go get github.com/foresthoffman/twitchbot`

### Importing

Import this package by including `github.com/foresthoffman/twitchbot` in your import block.

e.g.

```go
package main

import(
    ...
    "github.com/foresthoffman/twitchbot"
)
```

### Usage

Basic usage:

```go
package main

import "github.com/foresthoffman/twitchbot"

func main() {

	// Replace the channel name, bot name, and the path to the private directory with your respective
	// values.
	myBot := twitchbot.TwitchBot{
		Channel:      "twitch",
		MsgRate:      time.Duration(20/30) * time.Millisecond,
		Name:         "TwitchBot",
		Port:         "6667",
		PrivatePath:  "./private",
		Server:       "irc.chat.twitch.tv",
	}
	myBot.Start()
}
```

_That's all, enjoy!_
