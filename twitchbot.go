/**
 * twitchbot.go
 *
 * Copyright (c) 2017 Forest Hoffman. All Rights Reserved.
 * License: MIT License (see the included LICENSE file)
 */

package twitchbot

import (
	"time"
)

// TODO:
// 1. Connect to a Twitch.tv Chat channel.
// 	a. Pass along necessary information for the connection.
// 	 i.   The IRC (chat) server.
// 	 ii.  The port on the server.
// 	 iii. The channel we want the bot to join.
// 	 iv.  The bot's name.
// 	 v.   A secure key to allow the bot to connect indirectly (not through the website).
// 	 vi.  A maximum speed at which the bot can respond.
// 2. Listen for messages in the chat.
// 3. Do things based on what is happening in the chat.

type TwitchBot interface {
}

type BasicBot struct {

	// The channel that the bot is supposed to join. Note: The name MUST be lowercase, regardless
	// of how the username is displayed on Twitch.tv.
	Channel string

	// A forced delay between bot responses. This prevents the bot from breaking the message limit
	// rules. A 20/30 millisecond delay is enough for a non-modded bot. If you decrease the delay
	// make sure you're still within the limit!
	//
	// Message Rate Guidelines: https://dev.twitch.tv/docs/irc#irc-command-and-message-limits
	MsgRate time.Duration

	// The name that the bot will use in the chat that it's attempting to join.
	Name string

	// The port of the IRC server.
	Port string

	// A path to a limited-access directory containing the bot's OAuth credentials.
	PrivatePath string

	// The domain of the IRC server.
	Server string
}
