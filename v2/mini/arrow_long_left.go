package mini

import (
	g "github.com/maragudk/gomponents"

	h "github.com/maragudk/gomponents-heroicons/v2"
)

func ArrowLongLeft(children ...g.Node) g.Node {
	return h.Mini(g.Group(children),
		g.Raw(`<path fill-rule="evenodd" d="M18 10a.75.75 0 01-.75.75H4.66l2.1 1.95a.75.75 0 11-1.02 1.1l-3.5-3.25a.75.75 0 010-1.1l3.5-3.25a.75.75 0 111.02 1.1l-2.1 1.95h12.59A.75.75 0 0118 10z" clip-rule="evenodd"/>`))
}