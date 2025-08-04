package encode

import (
	"math"
	"os"
	"strings"
	"time"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

var morseCode = map[rune]string{
	'A': ".-", 'B': "-...", 'C': "-.-.", 'D': "-..",
	'E': ".", 'F': "..-.", 'G': "--.", 'H': "....",
	'I': "..", 'J': ".---", 'K': "-.-", 'L': ".-..",
	'M': "--", 'N': "-.", 'O': "---", 'P': ".--.",
	'Q': "--.-", 'R': ".-.", 'S': "...", 'T': "-",
	'U': "..-", 'V': "...-", 'W': ".--", 'X': "-..-",
	'Y': "-.--", 'Z': "--..", ' ': "/",
}

const (
	sampleRate     = 44100
	freq           = 700.0
	dotLengthMs    = 100
	dashLengthMs   = dotLengthMs * 3
	intraCharSpace = dotLengthMs
	charSpace      = dotLengthMs * 3
	wordSpace      = dotLengthMs * 7
)

func silence(durationMs int) []int {
	samples := int(float64(sampleRate) * float64(durationMs) / 1000.0)
	return make([]int, samples)
}

func tone(durationMs int) []int {
	samples := int(float64(sampleRate) * float64(durationMs) / 1000.0)
	data := make([]int, samples)
	volume := 30000.0
	for i := 0; i < samples; i++ {
		angle := 2.0 * math.Pi * freq * float64(i) / float64(sampleRate)
		data[i] = int(volume * math.Sin(angle))
	}
	return data
}

func TextToAudio(text string) string {
	text = strings.ToUpper(text)
	fileName := time.Now().Format("2006-01-02 15:04:05") + ".wav"
	var result []int
	for _, ch := range text {
		code, ok := morseCode[ch]
		if !ok {
			continue
		}
		for i, c := range code {
			switch c {
			case '.':
				result = append(result, tone(dotLengthMs)...)
			case '-':
				result = append(result, tone(dashLengthMs)...)
			}
			// Intra-character spacing
			if i < len(code)-1 {
				result = append(result, silence(intraCharSpace)...)
			}
		}
		// After each letter
		if ch == ' ' {
			result = append(result, silence(wordSpace)...)
		} else {
			result = append(result, silence(charSpace)...)
		}
	}

	dir := "audios"

	// Katalog mavjud emasmi? Yaratsin
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			panic(err)
		}
	}
	fullPath := dir + "/" + fileName
	// Save to WAV file
	f, err := os.Create(fullPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	enc := wav.NewEncoder(f, sampleRate, 16, 1, 1)
	buf := &audio.IntBuffer{
		Data:           result,
		Format:         &audio.Format{NumChannels: 1, SampleRate: sampleRate},
		SourceBitDepth: 16,
	}
	if err := enc.Write(buf); err != nil {
		panic(err)
	}
	if err := enc.Close(); err != nil {
		panic(err)
	}

	return fileName
}
