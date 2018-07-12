package server

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Gnidleif/discripter/invoker"
	"github.com/Gnidleif/discripter/logger"
	"github.com/bwmarrin/discordgo"
)

var (
	dg  *discordgo.Session
	inv *invoker.Invoker
)

type MsgType uint

const (
	user MsgType = iota + 1
	bot
	out
)

type DiscordMsg struct {
	Channel string
	Author  string
	Msg     string
	Error   error      `json:",omitempty"`
	Action  *ScriptMsg `json:",omitempty"`
}

type ScriptMsg struct {
	Script string
	Args   []string
	Result []byte
}

func Start(token, scriptdir string) error {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return err
	}

	dg.AddHandler(messageCreate)

	inv, err = invoker.New(scriptdir)
	if err != nil {
		return err
	}

	return dg.Open()
}

var rgx = regexp.MustCompile(`^\!(\w+)`)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	dm := &DiscordMsg{
		Channel: m.ChannelID,
		Author:  m.Author.ID,
		Msg:     m.Content,
	}

	defer logger.Write(dm)

	if ok := rgx.MatchString(dm.Msg); !ok {
		return
	}

	cmds := strings.Split(dm.Msg[1:], " ")
	if len(cmds) == 0 {
		return
	}

	res, err := inv.Run(cmds[0], cmds[1:]...)
	if err != nil {
		dm.Error = err
		return
	}

	dm.Action = &ScriptMsg{
		Script: cmds[0],
		Args:   cmds[1:],
		Result: res,
	}

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("`%s`", dm.Action.Result))
}

func Stop() error {
	return dg.Close()
}
