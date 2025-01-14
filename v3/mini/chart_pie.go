package mini

import (
	g "maragu.dev/gomponents"

	h "maragu.dev/gomponents-heroicons/v3"
)

func ChartPie(children ...g.Node) g.Node {
	return h.Mini(g.Group(children),
		g.Raw(`<path d="M12 9a1 1 0 01-1-1V3c0-.553.45-1.008.997-.93a7.004 7.004 0 015.933 5.933c.078.547-.378.997-.93.997h-5z"/>
  <path d="M8.003 4.07C8.55 3.992 9 4.447 9 5v5a1 1 0 001 1h5c.552 0 1.008.45.93.997A7.001 7.001 0 012 11a7.002 7.002 0 016.003-6.93z"/>`))
}
