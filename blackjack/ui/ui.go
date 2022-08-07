package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	Banner = `*************
* BLACKJACK *
*************
-------------`
	RuleHoriz = `-------------`
)

type Ui struct {
	writer *os.File
}

func NewUi(w *os.File) Ui{
	ui := Ui{
		writer: w,
	}
	return ui
}

func (ui *Ui) PrintBanner() {
	ui.WriteLn(Banner)
}

func (ui *Ui) WriteLn(msg string) {
	_, _ = ui.writer.WriteString(msg + "\n")
}

func (ui *Ui) Write(msg string) {
	_, _ = ui.writer.WriteString(msg)
}

func (ui *Ui) Writef(msg string, a ...interface{}) {
	ui.Write(fmt.Sprintf(msg, a...))
}

func (ui *Ui) ReadLn(prompt string) (string, error) {
	rdr := bufio.NewReader(os.Stdin)
	ui.Write(prompt)
	s, err := rdr.ReadString('\n')
	if err != nil {
		return "", nil
	}
	s = strings.Replace(s, "\n", "", -1)
	return s, nil
}

func (ui *Ui) GetPlayerCount() (int, error) {
	var result int
	pCStr, err := ui.ReadLn("Enter number of players: ")
	if err != nil {
		return result, err
	}
	return strconv.Atoi(pCStr)
}

func (ui *Ui) HorizontalLine() {
	ui.WriteLn(RuleHoriz)
}

func (ui *Ui) GetMenuSelection(opts []string) (string, error){
	for ix, o := range opts {
		ui.Writef("%s %v\n", o, ix)
	}
	return ui.ReadLn("Enter a selection: ")
}
