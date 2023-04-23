package zsys

import "github.com/gonutz/w32/v2"

func ShowConsole(cmdShow int) {
	console := w32.GetConsoleWindow()
	if 0 != console {
		_, consoleProcID := w32.GetWindowThreadProcessId(console)
		if w32.GetCurrentProcessId() == consoleProcID {
			w32.ShowWindowAsync(console, cmdShow)
		}
	}
}
