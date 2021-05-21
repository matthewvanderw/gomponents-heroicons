package outline

import (
	g "github.com/maragudk/gomponents"

	h "github.com/maragudk/gomponents-heroicons"
)

func ArrowCircleRightIcon(children ...g.Node) g.Node {
	return h.Outline(g.Group(children),
		g.Raw(`<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 9l3 3m0 0l-3 3m3-3H8m13 0a9 9 0 11-18 0 9 9 0 0118 0z"/>`))
}
