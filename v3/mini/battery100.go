package mini

import (
	g "maragu.dev/gomponents"

	h "maragu.dev/gomponents-heroicons/v3"
)

func Battery100(children ...g.Node) g.Node {
	return h.Mini(g.Group(children),
		g.Raw(`<path d="M4.75 8a.75.75 0 00-.75.75v3.5c0 .414.336.75.75.75h9.5a.75.75 0 00.75-.75v-2.5a.75.75 0 00-.75-.75h-9.5z"/>
  <path fill-rule="evenodd" d="M1 7.25A2.25 2.25 0 013.25 5h12.5A2.25 2.25 0 0118 7.25v1.085a1.5 1.5 0 011 1.415v.5a1.5 1.5 0 01-1 1.415v1.085A2.25 2.25 0 0115.75 15H3.25A2.25 2.25 0 011 12.75v-5.5zm2.25-.75a.75.75 0 00-.75.75v5.5c0 .414.336.75.75.75h12.5a.75.75 0 00.75-.75v-5.5a.75.75 0 00-.75-.75H3.25z" clip-rule="evenodd"/>`))
}
