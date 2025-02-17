package solid

import (
	g "maragu.dev/gomponents"

	h "maragu.dev/gomponents-heroicons/v3"
)

func ArrowDownLeft(children ...g.Node) g.Node {
	return h.Solid(g.Group(children),
		g.Raw(`<path fill-rule="evenodd" d="M20.03 3.97a.75.75 0 010 1.06L6.31 18.75h9.44a.75.75 0 010 1.5H4.5a.75.75 0 01-.75-.75V8.25a.75.75 0 011.5 0v9.44L18.97 3.97a.75.75 0 011.06 0z" clip-rule="evenodd"/>`))
}
