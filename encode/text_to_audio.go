package encode

import (
	"fmt"
)

type TxtToAudio struct {
	Text string
}

func (t *TxtToAudio) Run() {
	fileName := TextToAudio(t.Text)
	fmt.Println("✅ Audio file created:", fileName)
}
