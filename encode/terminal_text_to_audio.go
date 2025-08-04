package encode

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type StringToAudio struct {
	Text string
}

func (t *StringToAudio) Run() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("💬 Enter your message: ")
	text, _ := reader.ReadString('\n')
	t.Text = strings.TrimSpace(text)
	t.Text = strings.ToUpper(t.Text)
	fileName := TextToAudio(t.Text)
	fmt.Println("✅ Audio file created:", fileName)

}
