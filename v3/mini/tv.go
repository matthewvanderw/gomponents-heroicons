package mini

import (
	g "maragu.dev/gomponents"

	h "maragu.dev/gomponents-heroicons/v3"
)

func Tv(children ...g.Node) g.Node {
	return h.Mini(g.Group(children),
		g.Raw(`<path d="M4 5h12v7H4V5z"/>
  <path fill-rule="evenodd" d="M1 3.5A1.5 1.5 0 012.5 2h15A1.5 1.5 0 0119 3.5v10a1.5 1.5 0 01-1.5 1.5H12v1.5h3.25a.75.75 0 010 1.5H4.75a.75.75 0 010-1.5H8V15H2.5A1.5 1.5 0 011 13.5v-10zm16.5 0h-15v10h15v-10z" clip-rule="evenodd"/>`))
}
