package encode

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Decoder struct{}

func (d *Decoder) Run() {
	dir := "audios"

	// Papkadagi barcha .wav fayllarni o'qib olish
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("❌ audios/ papkasini o'qib bo'lmadi:", err)
		return
	}

	// Faqat .wav fayllarni chiqaramiz
	var wavFiles []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".wav") {
			wavFiles = append(wavFiles, file.Name())
		}
	}

	// Hech qanday fayl topilmasa
	if len(wavFiles) == 0 {
		fmt.Println("⚠️ audios/ papkasida .wav fayllar topilmadi.")
		return
	}

	// Fayllar ro'yxatini ko'rsatamiz
	fmt.Println("🎧 Mavjud .wav fayllar:")
	for i, name := range wavFiles {
		fmt.Printf("%d. %s\n", i+1, name)
	}

	// Foydalanuvchi tanlovi
	var choice int
	fmt.Print("Tanlagan fayl raqamini kiriting: ")
	_, err = fmt.Scanf("%d", &choice)
	if err != nil || choice < 1 || choice > len(wavFiles) {
		fmt.Println("❌ Noto‘g‘ri tanlov.")
		return
	}

	selectedFile := filepath.Join(dir, wavFiles[choice-1])

	// Faylni o'qish
	data, err := os.ReadFile(selectedFile)
	if err != nil {
		fmt.Println("❌ Faylni o‘qishda xatolik:", err)
		return
	}

	// Fayl binarini oddiy stringga aylantirish
	str := string(data)
	fmt.Println("\n📝 Tanlangan fayl kontenti (binar -> string):")
	fmt.Println(str)
}
