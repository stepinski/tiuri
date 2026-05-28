package main

import (
	"github.com/gdamore/tcell/v2"
)

var gameMap = [][]rune{
	[]rune("┌────────┐"),
	[]rune("│.....⚔..│"),
	[]rune("│..♟.....│"),
	[]rune("│....⚔...│"),
	[]rune("│..⚔.....│"),
	[]rune("│....⚔...│"),
	[]rune("│........D"),
	[]rune("│........│"),
	[]rune("└────────┘"),
}

func render(screen tcell.Screen, playerX, playerY int) {
	rosePineBg := tcell.NewRGBColor(25, 23, 36)       // #191724
	rosePineFg := tcell.NewRGBColor(224, 222, 244)    // #e0def4
	rosePinePine := tcell.NewRGBColor(49, 116, 143)   // #31748f
	rosePineMuted := tcell.NewRGBColor(110, 106, 134) // #6e6a86
	rosePineGold := tcell.NewRGBColor(246, 193, 119)  // #f6c177
	rosePineLove := tcell.NewRGBColor(235, 111, 146)  // #eb6f92

	wallStyle := tcell.StyleDefault.Foreground(rosePinePine).Background(rosePineBg)
	floorStyle := tcell.StyleDefault.Foreground(rosePineMuted).Background(rosePineBg)
	playerStyle := tcell.StyleDefault.Foreground(rosePineGold).Background(rosePineBg)
	guardStyle := tcell.StyleDefault.Foreground(rosePineLove).Background(rosePineBg)
	bgStyle := tcell.StyleDefault.Background(rosePineBg)

	_ = rosePineFg // reserved for text/UI later

	w, h := screen.Size()
	for y := range h {
		for x := range w {
			screen.SetContent(x, y, ' ', nil, bgStyle)
		}
	}

	for row, line := range gameMap {
		for col, ch := range line {
			switch ch {
			case '┌', '┐', '└', '┘', '─', '│':
				screen.SetContent(col, row, ch, nil, wallStyle)
			case '♟':
				screen.SetContent(col, row, ch, nil, guardStyle)
			default:
				screen.SetContent(col, row, ch, nil, floorStyle)
			}
		}
	}

	screen.SetContent(playerX, playerY, '@', nil, playerStyle)
	screen.Sync()
}

func isWalkable(ch rune) bool {
	return ch == '.'
}

func isDoor(ch rune) bool {
	return ch == 'D'
}

func isKnights(ch rune) bool {
	return ch == '⚔'
}

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := screen.Init(); err != nil {
		panic(err)
	}

	defer screen.Fini()
	defer screen.ShowCursor(0, 0)

	w, h := screen.Size()
	if w == 0 || h == 0 {
		panic("terminal size is zero")
	}

	playerX, playerY := 1, 1
	render(screen, playerX, playerY)

	for {
		switch ev := screen.PollEvent().(type) {
		case *tcell.EventKey:
			newX, newY := playerX, playerY
			switch ev.Key() {
			case tcell.KeyUp:
				newY--
			case tcell.KeyDown:
				newY++
			case tcell.KeyLeft:
				newX--
			case tcell.KeyRight:
				newX++
			case tcell.KeyEscape, tcell.KeyCtrlC:
				return
			}

			if newY >= 0 && newY < len(gameMap) &&
				newX >= 0 && newX < len(gameMap[newY]) &&
				isWalkable(gameMap[newY][newX]) {
				playerX, playerY = newX, newY
				render(screen, playerX, playerY)
			}

			if isKnights(gameMap[newY][newX]) {
				playerX, playerY = newX, newY
				render(screen, playerX, playerY)
				showMessage(screen, "The other knights see you. You return to your vigil in shame.")
				screen.PollEvent()
				return
			}

			if isDoor(gameMap[newY][newX]) {
				playerX, playerY = newX, newY
				render(screen, playerX, playerY)
				showMessage(screen, "A voice whispers: For the love of God, open this door...")
				screen.PollEvent()
				return
			}
		case *tcell.EventResize:
			screen.Sync()
			render(screen, playerX, playerY)
		}
	}
}

func showMessage(screen tcell.Screen, s string) {
	msgStyle := tcell.StyleDefault.Foreground(tcell.NewRGBColor(224, 222, 244)).Background(tcell.NewRGBColor(25, 23, 36))
	for i, ch := range s {
		screen.SetContent(i+1, 10, ch, nil, msgStyle)
	}
	screen.Sync()
}
