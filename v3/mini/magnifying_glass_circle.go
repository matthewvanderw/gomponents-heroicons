package mini

import (
	g "maragu.dev/gomponents"

	h "maragu.dev/gomponents-heroicons/v3"
)

func MagnifyingGlassCircle(children ...g.Node) g.Node {
	return h.Mini(g.Group(children),
		g.Raw(`<path d="M6.5 9a2.5 2.5 0 115 0 2.5 2.5 0 01-5 0z"/>
  <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM9 5a4 4 0 102.248 7.309l1.472 1.471a.75.75 0 101.06-1.06l-1.471-1.472A4 4 0 009 5z" clip-rule="evenodd"/>`))
}
