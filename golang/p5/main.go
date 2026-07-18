package main
import "fmt"

type MediaPlayer interface{
	Play(title string) string
}

type AudioPlayer struct {
	AudioFormat string
}
type VideoPlayer struct {
	Resolution string
}

func (a AudioPlayer) Play(title string) string{
	return fmt.Sprintf("Playing audio %s in %s format",title, a.AudioFormat)
}

func (v VideoPlayer) Play(title string) string{
	return fmt.Sprintf("Playing video %s at %s resolution",title, v.Resolution)
}

func StartPlayback(m MediaPlayer, title string) {
	fmt.Println(m.Play(title))
}
func main(){
	a1 := AudioPlayer{AudioFormat: "mp3"}
	v1 := VideoPlayer{Resolution: "1080"}
	StartPlayback(a1,"music")
	StartPlayback(v1,"video")
}