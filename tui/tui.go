// Package tui provides the terminal user interface
// this tui implement based on bubbletea, a terminal user interface library for Go
// check bubbletea at github.com/charmbracelet/bubbletea
// this tui is a chat room, you can type message and send it to the chat room
// the bot will respond to your message
package tui

import (
	"awesomeDSL/evaluator"
	"awesomeDSL/gpt"
	"awesomeDSL/lexer"
	"awesomeDSL/object"
	"awesomeDSL/parser"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var Env = object.NewEnvironment()
var History = []string{}

// GUI is the main function to start the GUI
func GUI() {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type (
	errMsg error
)

// Model is the Bubble Tea model, describe component of the UI
type model struct {
	viewport    viewport.Model
	messages    []string
	textarea    textarea.Model
	senderStyle lipgloss.Style
	err         error
}

// set initial layout of the UI
func initialModel() model {
	ta := textarea.New()
	ta.Placeholder = "Send a message..."
	ta.Focus()

	ta.Prompt = "┃ "
	ta.CharLimit = 450

	ta.SetWidth(90)
	ta.SetHeight(3)

	// Remove cursor line styling
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	ta.ShowLineNumbers = false

	vp := viewport.New(130, 5)
	vp.SetContent(`Welcome to the chat bot!
Type a message and press Enter to send.`)

	ta.KeyMap.InsertNewline.SetEnabled(false)

	return model{
		textarea:    ta,
		messages:    []string{},
		viewport:    vp,
		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		err:         nil,
	}
}

// Init is the initial command
func (m model) Init() tea.Cmd {
	return textarea.Blink
}

// Update define how the model update
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
	)

	m.textarea, tiCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)

	input := m.textarea.Value()
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		// Ctrl+C and ESC quits the program
		case tea.KeyCtrlC, tea.KeyEsc:
			//write history to file
			//file name is currently date and time
			t := time.Now()
			path := "./log/" + t.Format("2006-01-02 15:04:05.000") + ".txt"
			file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
				return m, nil
			}
			defer file.Close()
			for _, msg := range History {
				_, err := file.WriteString(msg + "\n")
				if err != nil {
					fmt.Println(err)
					return m, nil
				}
			}
			fmt.Println(m.textarea.Value())
			fmt.Println("Bye!")
			return m, tea.Quit
		case tea.KeyEnter:
			//echo input in terminal
			m.messages = append(m.messages, m.senderStyle.Render("You: ")+m.textarea.Value())
			m.viewport.SetContent(strings.Join(m.messages, "\n"))
			m.textarea.Reset()
			m.viewport.GotoBottom()

			//check if input is defined in script
			inputWithoutSpace := strings.Replace(input, " ", "", -1)
			_, ok := Env.Get(inputWithoutSpace)
			if !ok {
				//if not defined in script, use openai to generate response if OPENAI_FOR_DSL is set to true
				switch os.Getenv("OPENAI_FOR_DSL") {
				case "true":
					m.messages = append(m.messages, m.senderStyle.Render("Bot: ")+gpt.GPT(input))
				default:
					m.messages = append(m.messages, m.senderStyle.Render("Bot: ")+"I don't know what you mean :(")
				}
				History = append(History, strings.Join(m.messages, "\n"))
				m.viewport.SetContent(strings.Join(m.messages, "\n"))
				m.textarea.Reset()
				m.viewport.GotoBottom()
				break
			}
			//if defined in script, evaluate it and get response
			l := lexer.New(inputWithoutSpace)
			p := parser.New(l)
			program := p.ParseProgram()
			if len(p.Errors()) != 0 {
				m.err = fmt.Errorf("parser errors")
				for _, msg := range p.Errors() {
					m.err = fmt.Errorf("%s\t%s", m.err, msg)
				}
				return m, nil
			}
			evaluated := evaluator.Eval(program, Env)
			if evaluated != nil && evaluated.Type() != object.NULL_OBJ {
				//if evaluated is not null, echo it in terminal
				m.messages = append(m.messages, m.senderStyle.Render("Bot: ")+evaluated.Inspect())
				History = append(History, strings.Join(m.messages, "\n"))
				m.viewport.SetContent(strings.Join(m.messages, "\n"))
				m.viewport.GotoBottom()
				break
			}
			// here we handle puts buffer if we call puts in script
			if evaluator.PutsBuffer != "" {
				m.messages = append(m.messages, m.senderStyle.Render("Bot: ")+evaluator.PutsBuffer)
				History = append(History, strings.Join(m.messages, "\n"))
				m.viewport.SetContent(strings.Join(m.messages, "\n"))
				m.viewport.GotoBottom()
				evaluator.PutsBuffer = ""
				break
			}
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	if m.err != nil {
		return m, nil
	}

	return m, tea.Batch(tiCmd, vpCmd)
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s",
		m.viewport.View(),
		m.textarea.View(),
	) + "\n\n"
}

// LoadScript load script from file and evaluate it
func LoadScript(file *os.File) bool {
	//open file
	script := ""
	//read file in a string
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return false
	}
	scriptSize := fileInfo.Size()
	scriptBytes := make([]byte, scriptSize)
	_, err = file.Read(scriptBytes)
	if err != nil {
		fmt.Println(err)
		return false
	}
	script = string(scriptBytes)
	//parse script
	l := lexer.New(script)
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		fmt.Println("Woops! Something bad happens!\n parser errors:")
		for _, msg := range p.Errors() {
			fmt.Println(msg)
		}
		return false
	}
	evaluator.Eval(program, Env)
	return true
}
