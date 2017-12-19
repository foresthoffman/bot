/**
 * twitchbot.go
 *
 * Copyright (c) 2017 Forest Hoffman. All Rights Reserved.
 * License: MIT License (see the included LICENSE file)
 */

package twitchbot

import (
	"fmt"
	"net"
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

const PSTFormat = "Jan 2 15:04:05 PST"

type OAuthCred struct {

	// The bot account's OAuth password. Thanks to the JSON syntax after the data type, this field
	// will be filled with the value of the field with the specified key.
	Password string `json:"password",omitempty`
}

type TwitchBot interface {
	Connect()
	Disconnect()
	HandleChat() error
	JoinChannel()
	ReadCredentials() (*OAuthCred, error)
	Say(msg string) error
	Start()
}

type BasicBot struct {

	// The channel that the bot is supposed to join. Note: The name MUST be lowercase, regardless
	// of how the username is displayed on Twitch.tv.
	Channel string

	// A reference to the bot's connection to the server.
	conn net.Conn

	// The credentials necessary for authentication.
	Credentials *OAuthCred

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

	// The time at which the bot achieved a connection to the server.
	startTime time.Time
}

// Connects the bot to the Twitch IRC server. The bot will continue to try to connect until it
// succeeds or is manually shutdown.
func (bb *BasicBot) Connect() {
	var err error
	fmt.Printf("[%s] Connecting to %s...\n", timeStamp(), bb.Server)

	// makes connection to Twitch IRC server
	bb.conn, err = net.Dial("tcp", bb.Server+":"+bb.Port)
	if nil != err {
		fmt.Printf("[%s] Cannot connect to %s, retrying.\n", timeStamp(), bb.Server)
		bb.Connect()
		return
	}
	fmt.Printf("[%s] Connected to %s!\n", timeStamp(), bb.Server)
	bb.startTime = time.Now()
}

// Officially disconnects the bot from the Twitch IRC server.
func (bb *BasicBot) Disconnect() {
	bb.conn.Close()
	upTime := time.Now().Sub(bb.startTime).Seconds()
	fmt.Printf("[%s] Closed connection from %s! | Live for: %fs\n", timeStamp(), bb.Server, upTime)
}

func timeStamp() string {
	return TimeStamp(PSTFormat)
}

func TimeStamp(format string) string {
	return time.Now().Format(format)
}
