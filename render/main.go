package render

import (
    "github.com/nsf/termbox-go"
)

var title = "WarGame"

const bgColor = termbox.ColorGreen
const cardColor = termbox.ColorWhite
const cardBackColor = termbox.ColorRed

//Render functioon
func main(game *Game) {
    termbox.Clear(bgColor, bgColor)


    termbox.Flush()
}
