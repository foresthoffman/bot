/**
 * doc.go
 *
 * Copyright (c) 2017 Forest Hoffman. All Rights Reserved.
 * License: MIT License (see the included LICENSE file)
 */

/*
The bot package provides a set of functions that control a basic Twitch.tv chat bot. The
package also exposes an interface which can be used to create a custom chat bot.

Basic usage:

```
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
*/
package bot
