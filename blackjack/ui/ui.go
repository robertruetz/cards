package ui

import (
	"bufio"
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
	ui.writeLn(Banner)
}

func (ui *Ui) writeLn(msg string) {
	_, _ = ui.writer.WriteString(msg + "\n")
}

func (ui *Ui) write(msg string) {
	_, _ = ui.writer.WriteString(msg)
}

func (ui *Ui) readLn(prompt string) (string, error) {
	rdr := bufio.NewReader(os.Stdin)
	ui.write(prompt)
	s, err := rdr.ReadString('\n')
	if err != nil {
		return "", nil
	}
	s = strings.Replace(s, "\n", "", -1)
	return s, nil
}

func (ui *Ui) GetPlayerCount() (int, error) {
	var result int
	pCStr, err := ui.readLn("Enter number of players: ")
	if err != nil {
		return result, err
	}
	return strconv.Atoi(pCStr)
}

func (ui *Ui) HorizontalLine() {
	ui.writeLn(RuleHoriz)
}
