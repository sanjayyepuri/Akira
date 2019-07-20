package router

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

// we will register a command with a handler
// for the time being let us say we simply route the first word to a handler
// ~<command> <arg1> <arg2> ...

// Router stores the map of register commands
type Router struct {
	commandPrefix string
	commands      map[string]*Command
}

// RegisterCommand will register a command with a particular handler
func (r *Router) RegisterCommand(name string, handler CommandHandler) {
	if _, ok := r.commands[name]; ok {
		log.Errorf("cannot create command with name %s", name)
		return
	}
	command := Command{name, handler}
	r.commands[name] = &command
}

// Handler will handle discordgo events
func (r *Router) Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// parse the message
	// TODO move out to a different function and make more robust
	var message string = m.Content

	// middleware
	if m.Author.Bot {
		log.Debugf("router received command from bot %s\n", m.Author)
		return
	}

	// middleware job
	if !strings.HasPrefix(message, r.commandPrefix) {
		return
	}

	// TODO build request object
	// TODO build message parser
	messageParts := strings.Fields(message)
	commandName := messageParts[0][1:]
	// TODO build response object

	command, ok := r.commands[commandName]
	if ok {
		log.Debugf("executing handler for command %s", commandName)
		command.handler(m, s)
	} else {
		log.Errorf("No handler registered for %s", commandName)
	}
}

// NewRouter returns a new Router
func NewRouter() *Router {
	var router Router
	router.commands = make(map[string]*Command)

	return &router
}

// WithPrefix builder to add prefix
func (r *Router) WithPrefix(prefix string) *Router {
	r.commandPrefix = prefix
	return r
}

// CommandHandler the function handle the command
type CommandHandler func(*discordgo.MessageCreate, *discordgo.Session)

// Command stores the information about a particular command to match and handle it
type Command struct {
	name    string
	handler CommandHandler
}

// Handler will set the handler for the command
func (c *Command) Handler(h CommandHandler) {
	c.handler = h
}
