package main

import encode "morse/decode"

func main() {
	//txtToAudio := decodetext.TxtToAudio{
	//	Text: "BUGUN KUN YAXSHI",
	//}
	//txtToAudio.Run()

	//stringToAudio := decodetext.StringToAudio{}
	//stringToAudio.Run()

	decode := encode.Decoder{}
	decode.Run()
}
