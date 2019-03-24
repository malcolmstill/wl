// package zxdgv6 acts as a client for the xdg_shell_unstable_v6 wayland protocol.

// generated by wl-scanner
// https://github.com/dkolbly/wl-scanner
// from: /home/malcolm/git/ulubis/cl-wayland/xdg-shell-unstable-v6.xml
// on 2019-03-24 14:06:59 +0000
package zxdgv6

import (
	"sync"

	"github.com/malcolmstill/wl"
)

type ZxdgShellV6DestroyEvent struct {
}

type ZxdgShellV6DestroyHandler interface {
	HandleZxdgShellV6Destroy(ZxdgShellV6DestroyEvent)
}

func (p *ZxdgShellV6) AddDestroyHandler(h ZxdgShellV6DestroyHandler) {
	if h != nil {
		p.mu.Lock()
		p.destroyHandlers = append(p.destroyHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZxdgShellV6) RemoveDestroyHandler(h ZxdgShellV6DestroyHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.destroyHandlers {
		if e == h {
			p.destroyHandlers = append(p.destroyHandlers[:i], p.destroyHandlers[i+1:]...)
			break
		}
	}
}

type ZxdgShellV6CreatePositionerEvent struct {
	Id *ZxdgPositionerV6
}

type ZxdgShellV6CreatePositionerHandler interface {
	HandleZxdgShellV6CreatePositioner(ZxdgShellV6CreatePositionerEvent)
}

func (p *ZxdgShellV6) AddCreatePositionerHandler(h ZxdgShellV6CreatePositionerHandler) {
	if h != nil {
		p.mu.Lock()
		p.createPositionerHandlers = append(p.createPositionerHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZxdgShellV6) RemoveCreatePositionerHandler(h ZxdgShellV6CreatePositionerHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.createPositionerHandlers {
		if e == h {
			p.createPositionerHandlers = append(p.createPositionerHandlers[:i], p.createPositionerHandlers[i+1:]...)
			break
		}
	}
}

type ZxdgShellV6GetXdgSurfaceEvent struct {
	Id      *ZxdgSurfaceV6
	Surface *wl.Surface
}

type ZxdgShellV6GetXdgSurfaceHandler interface {
	HandleZxdgShellV6GetXdgSurface(ZxdgShellV6GetXdgSurfaceEvent)
}

func (p *ZxdgShellV6) AddGetXdgSurfaceHandler(h ZxdgShellV6GetXdgSurfaceHandler) {
	if h != nil {
		p.mu.Lock()
		p.getXdgSurfaceHandlers = append(p.getXdgSurfaceHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZxdgShellV6) RemoveGetXdgSurfaceHandler(h ZxdgShellV6GetXdgSurfaceHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.getXdgSurfaceHandlers {
		if e == h {
			p.getXdgSurfaceHandlers = append(p.getXdgSurfaceHandlers[:i], p.getXdgSurfaceHandlers[i+1:]...)
			break
		}
	}
}

type ZxdgShellV6PongEvent struct {
	Serial uint32
}

type ZxdgShellV6PongHandler interface {
	HandleZxdgShellV6Pong(ZxdgShellV6PongEvent)
}

func (p *ZxdgShellV6) AddPongHandler(h ZxdgShellV6PongHandler) {
	if h != nil {
		p.mu.Lock()
		p.pongHandlers = append(p.pongHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZxdgShellV6) RemovePongHandler(h ZxdgShellV6PongHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.pongHandlers {
		if e == h {
			p.pongHandlers = append(p.pongHandlers[:i], p.pongHandlers[i+1:]...)
			break
		}
	}
}

func (p *ZxdgShellV6) Dispatch(event *wl.Event) {
	switch event.Opcode {
	case 0:
		if len(p.destroyHandlers) > 0 {
			ev := ZxdgShellV6DestroyEvent{}
			p.mu.RLock()
			for _, h := range p.destroyHandlers {
				h.HandleZxdgShellV6Destroy(ev)
			}
			p.mu.RUnlock()
		}
	case 1:
		if len(p.createPositionerHandlers) > 0 {
			ev := ZxdgShellV6CreatePositionerEvent{}
			ev.Id = NewZxdgPositionerV6(p.Context(), int(event.Uint32()))
			p.mu.RLock()
			for _, h := range p.createPositionerHandlers {
				h.HandleZxdgShellV6CreatePositioner(ev)
			}
			p.mu.RUnlock()
		}
	case 2:
		if len(p.getXdgSurfaceHandlers) > 0 {
			ev := ZxdgShellV6GetXdgSurfaceEvent{}
			ev.Id = NewZxdgSurfaceV6(p.Context(), int(event.Uint32()))
			ev.Surface = event.Proxy(p.Context()).(*wl.Surface)
			p.mu.RLock()
			for _, h := range p.getXdgSurfaceHandlers {
				h.HandleZxdgShellV6GetXdgSurface(ev)
			}
			p.mu.RUnlock()
		}
	case 3:
		if len(p.pongHandlers) > 0 {
			ev := ZxdgShellV6PongEvent{}
			ev.Serial = event.Uint32()
			p.mu.RLock()
			for _, h := range p.pongHandlers {
				h.HandleZxdgShellV6Pong(ev)
			}
			p.mu.RUnlock()
		}
	}
}

type ZxdgShellV6 struct {
	wl.BaseProxy
	mu                       sync.RWMutex
	destroyHandlers          []ZxdgShellV6DestroyHandler
	createPositionerHandlers []ZxdgShellV6CreatePositionerHandler
	getXdgSurfaceHandlers    []ZxdgShellV6GetXdgSurfaceHandler
	pongHandlers             []ZxdgShellV6PongHandler
}

func NewZxdgShellV6(ctx *wl.Context, id int) *ZxdgShellV6 {
	ret := new(ZxdgShellV6)
	ctx.RegisterId(ret, id)
	return ret
}

// Ping will check if the client is alive.
//
//
// The ping event asks the client if it's still alive. Pass the
// serial specified in the event back to the compositor by sending
// a "pong" request back with the specified serial. See xdg_shell.ping.
//
// Compositors can use this to determine if the client is still
// alive. It's unspecified what will happen if the client doesn't
// respond to the ping request, or in what timeframe. Clients should
// try to respond in a reasonable amount of time.
//
// A compositor is free to ping in any way it wants, but a client must
// always respond to any xdg_shell object it created.
//
func (p *ZxdgShellV6) Ping(serial uint32) error {
	return p.Context().SendRequest(p, 0, serial)
}

const (
	ZxdgShellV6ErrorRole                = 0
	ZxdgShellV6ErrorDefunctSurfaces     = 1
	ZxdgShellV6ErrorNotTheTopmostPopup  = 2
	ZxdgShellV6ErrorInvalidPopupParent  = 3
	ZxdgShellV6ErrorInvalidSurfaceState = 4
	ZxdgShellV6ErrorInvalidPositioner   = 5
)

type ZxdgPositionerV6DestroyEvent struct {
}

type ZxdgPositionerV6DestroyHandler interface {
	HandleZxdgPositionerV6Destroy(ZxdgPositionerV6DestroyEvent)
}

func (p *ZxdgPositionerV6) AddDestroyHandler(h ZxdgPositionerV6DestroyHandler) {
	if h != nil {
		p.mu.Lock()
		p.destroyHandlers = append(p.destroyHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZxdgPositionerV6) RemoveDestroyHandler(h ZxdgPositionerV6DestroyHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.destroyHandlers {
		if e == h {
			p.destroyHandlers = append(p.destroyHandlers[:i], p.destroyHandlers[i+1:]...)
			break
		}
	}
}

type ZxdgPositionerV6SetSizeEvent struct {
	Width  int32
	Height int32
}

type ZxdgPositionerV6SetSizeHandler interface {
	HandleZxdgPositionerV6SetSize(ZxdgPositionerV6SetSizeEvent)
}

func (p *ZxdgPositionerV6) AddSetSizeHandler(h ZxdgPositionerV6SetSizeHandler) {
	if h != nil {
		p.mu.Lock()
		p.setSizeHandlers = append(p.setSizeHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZxdgPositionerV6) RemoveSetSizeHandler(h ZxdgPositionerV6SetSizeHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.setSizeHandlers {
		if e == h {
			p.setSizeHandlers = append(p.setSizeHandlers[:i], p.setSizeHandlers[i+1:]...)
			break
		}
	}
}

type ZxdgPositionerV6SetAnchorRectEvent struct {
	X      int32
	Y      int32
	Width  int32
	Height int32
}

type ZxdgPositionerV6SetAnchorRectHandler interface {
	HandleZxdgPositionerV6SetAnchorRect(ZxdgPositionerV6SetAnchorRectEvent)
}

func (p *ZxdgPositionerV6) AddSetAnchorRectHandler(h ZxdgPositionerV6SetAnchorRectHandler) {
	if h != nil {
		p.mu.Lock()
		p.setAnchorRectHandlers = append(p.setAnchorRectHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZxdgPositionerV6) RemoveSetAnchorRectHandler(h ZxdgPositionerV6SetAnchorRectHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.setAnchorRectHandlers {
		if e == h {
			p.setAnchorRectHandlers = append(p.setAnchorRectHandlers[:i], p.setAnchorRectHandlers[i+1:]...)
			break
		}
	}
}

type ZxdgPositionerV6SetAnchorEvent struct {
	Anchor uint32
}

type ZxdgPositionerV6SetAnchorHandler interface {
	HandleZxdgPositionerV6SetAnchor(ZxdgPositionerV6SetAnchorEvent)
}

func (p *ZxdgPositionerV6) AddSetAnchorHandler(h ZxdgPositionerV6SetAnchorHandler) {
	if h != nil {
		p.mu.Lock()
		p.setAnchorHandlers = append(p.setAnchorHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZxdgPositionerV6) RemoveSetAnchorHandler(h ZxdgPositionerV6SetAnchorHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.setAnchorHandlers {
		if e == h {
			p.setAnchorHandlers = append(p.setAnchorHandlers[:i], p.setAnchorHandlers[i+1:]...)
			break
		}
	}
}

type ZxdgPositionerV6SetGravityEvent struct {
	Gravity uint32
}

type ZxdgPositionerV6SetGravityHandler interface {
	HandleZxdgPositionerV6SetGravity(ZxdgPositionerV6SetGravityEvent)
}

func (p *ZxdgPositionerV6) AddSetGravityHandler(h ZxdgPositionerV6SetGravityHandler) {
	if h != nil {
		p.mu.Lock()
		p.setGravityHandlers = append(p.setGravityHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZxdgPositionerV6) RemoveSetGravityHandler(h ZxdgPositionerV6SetGravityHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.setGravityHandlers {
		if e == h {
			p.setGravityHandlers = append(p.setGravityHandlers[:i], p.setGravityHandlers[i+1:]...)
			break
		}
	}
}

type ZxdgPositionerV6SetConstraintAdjustmentEvent struct {
	ConstraintAdjustment uint32
}

type ZxdgPositionerV6SetConstraintAdjustmentHandler interface {
	HandleZxdgPositionerV6SetConstraintAdjustment(ZxdgPositionerV6SetConstraintAdjustmentEvent)
}

func (p *ZxdgPositionerV6) AddSetConstraintAdjustmentHandler(h ZxdgPositionerV6SetConstraintAdjustmentHandler) {
	if h != nil {
		p.mu.Lock()
		p.setConstraintAdjustmentHandlers = append(p.setConstraintAdjustmentHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZxdgPositionerV6) RemoveSetConstraintAdjustmentHandler(h ZxdgPositionerV6SetConstraintAdjustmentHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.setConstraintAdjustmentHandlers {
		if e == h {
			p.setConstraintAdjustmentHandlers = append(p.setConstraintAdjustmentHandlers[:i], p.setConstraintAdjustmentHandlers[i+1:]...)
			break
		}
	}
}

type ZxdgPositionerV6SetOffsetEvent struct {
	X int32
	Y int32
}

type ZxdgPositionerV6SetOffsetHandler interface {
	HandleZxdgPositionerV6SetOffset(ZxdgPositionerV6SetOffsetEvent)
}

func (p *ZxdgPositionerV6) AddSetOffsetHandler(h ZxdgPositionerV6SetOffsetHandler) {
	if h != nil {
		p.mu.Lock()
		p.setOffsetHandlers = append(p.setOffsetHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZxdgPositionerV6) RemoveSetOffsetHandler(h ZxdgPositionerV6SetOffsetHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.setOffsetHandlers {
		if e == h {
			p.setOffsetHandlers = append(p.setOffsetHandlers[:i], p.setOffsetHandlers[i+1:]...)
			break
		}
	}
}

func (p *ZxdgPositionerV6) Dispatch(event *wl.Event) {
	switch event.Opcode {
	case 0:
		if len(p.destroyHandlers) > 0 {
			ev := ZxdgPositionerV6DestroyEvent{}
			p.mu.RLock()
			for _, h := range p.destroyHandlers {
				h.HandleZxdgPositionerV6Destroy(ev)
			}
			p.mu.RUnlock()
		}
	case 1:
		if len(p.setSizeHandlers) > 0 {
			ev := ZxdgPositionerV6SetSizeEvent{}
			ev.Width = event.Int32()
			ev.Height = event.Int32()
			p.mu.RLock()
			for _, h := range p.setSizeHandlers {
				h.HandleZxdgPositionerV6SetSize(ev)
			}
			p.mu.RUnlock()
		}
	case 2:
		if len(p.setAnchorRectHandlers) > 0 {
			ev := ZxdgPositionerV6SetAnchorRectEvent{}
			ev.X = event.Int32()
			ev.Y = event.Int32()
			ev.Width = event.Int32()
			ev.Height = event.Int32()
			p.mu.RLock()
			for _, h := range p.setAnchorRectHandlers {
				h.HandleZxdgPositionerV6SetAnchorRect(ev)
			}
			p.mu.RUnlock()
		}
	case 3:
		if len(p.setAnchorHandlers) > 0 {
			ev := ZxdgPositionerV6SetAnchorEvent{}
			ev.Anchor = event.Uint32()
			p.mu.RLock()
			for _, h := range p.setAnchorHandlers {
				h.HandleZxdgPositionerV6SetAnchor(ev)
			}
			p.mu.RUnlock()
		}
	case 4:
		if len(p.setGravityHandlers) > 0 {
			ev := ZxdgPositionerV6SetGravityEvent{}
			ev.Gravity = event.Uint32()
			p.mu.RLock()
			for _, h := range p.setGravityHandlers {
				h.HandleZxdgPositionerV6SetGravity(ev)
			}
			p.mu.RUnlock()
		}
	case 5:
		if len(p.setConstraintAdjustmentHandlers) > 0 {
			ev := ZxdgPositionerV6SetConstraintAdjustmentEvent{}
			ev.ConstraintAdjustment = event.Uint32()
			p.mu.RLock()
			for _, h := range p.setConstraintAdjustmentHandlers {
				h.HandleZxdgPositionerV6SetConstraintAdjustment(ev)
			}
			p.mu.RUnlock()
		}
	case 6:
		if len(p.setOffsetHandlers) > 0 {
			ev := ZxdgPositionerV6SetOffsetEvent{}
			ev.X = event.Int32()
			ev.Y = event.Int32()
			p.mu.RLock()
			for _, h := range p.setOffsetHandlers {
				h.HandleZxdgPositionerV6SetOffset(ev)
			}
			p.mu.RUnlock()
		}
	}
}

type ZxdgPositionerV6 struct {
	wl.BaseProxy
	mu                              sync.RWMutex
	destroyHandlers                 []ZxdgPositionerV6DestroyHandler
	setSizeHandlers                 []ZxdgPositionerV6SetSizeHandler
	setAnchorRectHandlers           []ZxdgPositionerV6SetAnchorRectHandler
	setAnchorHandlers               []ZxdgPositionerV6SetAnchorHandler
	setGravityHandlers              []ZxdgPositionerV6SetGravityHandler
	setConstraintAdjustmentHandlers []ZxdgPositionerV6SetConstraintAdjustmentHandler
	setOffsetHandlers               []ZxdgPositionerV6SetOffsetHandler
}

func NewZxdgPositionerV6(ctx *wl.Context, id int) *ZxdgPositionerV6 {
	ret := new(ZxdgPositionerV6)
	ctx.RegisterId(ret, id)
	return ret
}

const (
	ZxdgPositionerV6ErrorInvalidInput = 0
)

const (
	ZxdgPositionerV6AnchorNone   = 0
	ZxdgPositionerV6AnchorTop    = 1
	ZxdgPositionerV6AnchorBottom = 2
	ZxdgPositionerV6AnchorLeft   = 4
	ZxdgPositionerV6AnchorRight  = 8
)

const (
	ZxdgPositionerV6GravityNone   = 0
	ZxdgPositionerV6GravityTop    = 1
	ZxdgPositionerV6GravityBottom = 2
	ZxdgPositionerV6GravityLeft   = 4
	ZxdgPositionerV6GravityRight  = 8
)

const (
	ZxdgPositionerV6ConstraintAdjustmentNone    = 0
	ZxdgPositionerV6ConstraintAdjustmentSlideX  = 1
	ZxdgPositionerV6ConstraintAdjustmentSlideY  = 2
	ZxdgPositionerV6ConstraintAdjustmentFlipX   = 4
	ZxdgPositionerV6ConstraintAdjustmentFlipY   = 8
	ZxdgPositionerV6ConstraintAdjustmentResizeX = 16
	ZxdgPositionerV6ConstraintAdjustmentResizeY = 32
)

type ZxdgSurfaceV6DestroyEvent struct {
}

type ZxdgSurfaceV6DestroyHandler interface {
	HandleZxdgSurfaceV6Destroy(ZxdgSurfaceV6DestroyEvent)
}

func (p *ZxdgSurfaceV6) AddDestroyHandler(h ZxdgSurfaceV6DestroyHandler) {
	if h != nil {
		p.mu.Lock()
		p.destroyHandlers = append(p.destroyHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZxdgSurfaceV6) RemoveDestroyHandler(h ZxdgSurfaceV6DestroyHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.destroyHandlers {
		if e == h {
			p.destroyHandlers = append(p.destroyHandlers[:i], p.destroyHandlers[i+1:]...)
			break
		}
	}
}

type ZxdgSurfaceV6GetToplevelEvent struct {
	Id *ZxdgToplevelV6
}

type ZxdgSurfaceV6GetToplevelHandler interface {
	HandleZxdgSurfaceV6GetToplevel(ZxdgSurfaceV6GetToplevelEvent)
}

func (p *ZxdgSurfaceV6) AddGetToplevelHandler(h ZxdgSurfaceV6GetToplevelHandler) {
	if h != nil {
		p.mu.Lock()
		p.getToplevelHandlers = append(p.getToplevelHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZxdgSurfaceV6) RemoveGetToplevelHandler(h ZxdgSurfaceV6GetToplevelHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.getToplevelHandlers {
		if e == h {
			p.getToplevelHandlers = append(p.getToplevelHandlers[:i], p.getToplevelHandlers[i+1:]...)
			break
		}
	}
}

type ZxdgSurfaceV6GetPopupEvent struct {
	Id         *ZxdgPopupV6
	Parent     *ZxdgSurfaceV6
	Positioner *ZxdgPositionerV6
}

type ZxdgSurfaceV6GetPopupHandler interface {
	HandleZxdgSurfaceV6GetPopup(ZxdgSurfaceV6GetPopupEvent)
}

func (p *ZxdgSurfaceV6) AddGetPopupHandler(h ZxdgSurfaceV6GetPopupHandler) {
	if h != nil {
		p.mu.Lock()
		p.getPopupHandlers = append(p.getPopupHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZxdgSurfaceV6) RemoveGetPopupHandler(h ZxdgSurfaceV6GetPopupHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.getPopupHandlers {
		if e == h {
			p.getPopupHandlers = append(p.getPopupHandlers[:i], p.getPopupHandlers[i+1:]...)
			break
		}
	}
}

type ZxdgSurfaceV6SetWindowGeometryEvent struct {
	X      int32
	Y      int32
	Width  int32
	Height int32
}

type ZxdgSurfaceV6SetWindowGeometryHandler interface {
	HandleZxdgSurfaceV6SetWindowGeometry(ZxdgSurfaceV6SetWindowGeometryEvent)
}

func (p *ZxdgSurfaceV6) AddSetWindowGeometryHandler(h ZxdgSurfaceV6SetWindowGeometryHandler) {
	if h != nil {
		p.mu.Lock()
		p.setWindowGeometryHandlers = append(p.setWindowGeometryHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZxdgSurfaceV6) RemoveSetWindowGeometryHandler(h ZxdgSurfaceV6SetWindowGeometryHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.setWindowGeometryHandlers {
		if e == h {
			p.setWindowGeometryHandlers = append(p.setWindowGeometryHandlers[:i], p.setWindowGeometryHandlers[i+1:]...)
			break
		}
	}
}

type ZxdgSurfaceV6AckConfigureEvent struct {
	Serial uint32
}

type ZxdgSurfaceV6AckConfigureHandler interface {
	HandleZxdgSurfaceV6AckConfigure(ZxdgSurfaceV6AckConfigureEvent)
}

func (p *ZxdgSurfaceV6) AddAckConfigureHandler(h ZxdgSurfaceV6AckConfigureHandler) {
	if h != nil {
		p.mu.Lock()
		p.ackConfigureHandlers = append(p.ackConfigureHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZxdgSurfaceV6) RemoveAckConfigureHandler(h ZxdgSurfaceV6AckConfigureHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.ackConfigureHandlers {
		if e == h {
			p.ackConfigureHandlers = append(p.ackConfigureHandlers[:i], p.ackConfigureHandlers[i+1:]...)
			break
		}
	}
}

func (p *ZxdgSurfaceV6) Dispatch(event *wl.Event) {
	switch event.Opcode {
	case 0:
		if len(p.destroyHandlers) > 0 {
			ev := ZxdgSurfaceV6DestroyEvent{}
			p.mu.RLock()
			for _, h := range p.destroyHandlers {
				h.HandleZxdgSurfaceV6Destroy(ev)
			}
			p.mu.RUnlock()
		}
	case 1:
		if len(p.getToplevelHandlers) > 0 {
			ev := ZxdgSurfaceV6GetToplevelEvent{}
			ev.Id = NewZxdgToplevelV6(p.Context(), int(event.Uint32()))
			p.mu.RLock()
			for _, h := range p.getToplevelHandlers {
				h.HandleZxdgSurfaceV6GetToplevel(ev)
			}
			p.mu.RUnlock()
		}
	case 2:
		if len(p.getPopupHandlers) > 0 {
			ev := ZxdgSurfaceV6GetPopupEvent{}
			ev.Id = NewZxdgPopupV6(p.Context(), int(event.Uint32()))
			ev.Parent = event.Proxy(p.Context()).(*ZxdgSurfaceV6)
			ev.Positioner = event.Proxy(p.Context()).(*ZxdgPositionerV6)
			p.mu.RLock()
			for _, h := range p.getPopupHandlers {
				h.HandleZxdgSurfaceV6GetPopup(ev)
			}
			p.mu.RUnlock()
		}
	case 3:
		if len(p.setWindowGeometryHandlers) > 0 {
			ev := ZxdgSurfaceV6SetWindowGeometryEvent{}
			ev.X = event.Int32()
			ev.Y = event.Int32()
			ev.Width = event.Int32()
			ev.Height = event.Int32()
			p.mu.RLock()
			for _, h := range p.setWindowGeometryHandlers {
				h.HandleZxdgSurfaceV6SetWindowGeometry(ev)
			}
			p.mu.RUnlock()
		}
	case 4:
		if len(p.ackConfigureHandlers) > 0 {
			ev := ZxdgSurfaceV6AckConfigureEvent{}
			ev.Serial = event.Uint32()
			p.mu.RLock()
			for _, h := range p.ackConfigureHandlers {
				h.HandleZxdgSurfaceV6AckConfigure(ev)
			}
			p.mu.RUnlock()
		}
	}
}

type ZxdgSurfaceV6 struct {
	wl.BaseProxy
	mu                        sync.RWMutex
	destroyHandlers           []ZxdgSurfaceV6DestroyHandler
	getToplevelHandlers       []ZxdgSurfaceV6GetToplevelHandler
	getPopupHandlers          []ZxdgSurfaceV6GetPopupHandler
	setWindowGeometryHandlers []ZxdgSurfaceV6SetWindowGeometryHandler
	ackConfigureHandlers      []ZxdgSurfaceV6AckConfigureHandler
}

func NewZxdgSurfaceV6(ctx *wl.Context, id int) *ZxdgSurfaceV6 {
	ret := new(ZxdgSurfaceV6)
	ctx.RegisterId(ret, id)
	return ret
}

// Configure will suggest a surface change.
//
//
// The configure event marks the end of a configure sequence. A configure
// sequence is a set of one or more events configuring the state of the
// xdg_surface, including the final xdg_surface.configure event.
//
// Where applicable, xdg_surface surface roles will during a configure
// sequence extend this event as a latched state sent as events before the
// xdg_surface.configure event. Such events should be considered to make up
// a set of atomically applied configuration states, where the
// xdg_surface.configure commits the accumulated state.
//
// Clients should arrange their surface for the new states, and then send
// an ack_configure request with the serial sent in this configure event at
// some point before committing the new surface.
//
// If the client receives multiple configure events before it can respond
// to one, it is free to discard all but the last event it received.
//
func (p *ZxdgSurfaceV6) Configure(serial uint32) error {
	return p.Context().SendRequest(p, 0, serial)
}

const (
	ZxdgSurfaceV6ErrorNotConstructed     = 1
	ZxdgSurfaceV6ErrorAlreadyConstructed = 2
	ZxdgSurfaceV6ErrorUnconfiguredBuffer = 3
)

type ZxdgToplevelV6DestroyEvent struct {
}

type ZxdgToplevelV6DestroyHandler interface {
	HandleZxdgToplevelV6Destroy(ZxdgToplevelV6DestroyEvent)
}

func (p *ZxdgToplevelV6) AddDestroyHandler(h ZxdgToplevelV6DestroyHandler) {
	if h != nil {
		p.mu.Lock()
		p.destroyHandlers = append(p.destroyHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZxdgToplevelV6) RemoveDestroyHandler(h ZxdgToplevelV6DestroyHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.destroyHandlers {
		if e == h {
			p.destroyHandlers = append(p.destroyHandlers[:i], p.destroyHandlers[i+1:]...)
			break
		}
	}
}

type ZxdgToplevelV6SetParentEvent struct {
	Parent *ZxdgToplevelV6
}

type ZxdgToplevelV6SetParentHandler interface {
	HandleZxdgToplevelV6SetParent(ZxdgToplevelV6SetParentEvent)
}

func (p *ZxdgToplevelV6) AddSetParentHandler(h ZxdgToplevelV6SetParentHandler) {
	if h != nil {
		p.mu.Lock()
		p.setParentHandlers = append(p.setParentHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZxdgToplevelV6) RemoveSetParentHandler(h ZxdgToplevelV6SetParentHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.setParentHandlers {
		if e == h {
			p.setParentHandlers = append(p.setParentHandlers[:i], p.setParentHandlers[i+1:]...)
			break
		}
	}
}

type ZxdgToplevelV6SetTitleEvent struct {
	Title string
}

type ZxdgToplevelV6SetTitleHandler interface {
	HandleZxdgToplevelV6SetTitle(ZxdgToplevelV6SetTitleEvent)
}

func (p *ZxdgToplevelV6) AddSetTitleHandler(h ZxdgToplevelV6SetTitleHandler) {
	if h != nil {
		p.mu.Lock()
		p.setTitleHandlers = append(p.setTitleHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZxdgToplevelV6) RemoveSetTitleHandler(h ZxdgToplevelV6SetTitleHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.setTitleHandlers {
		if e == h {
			p.setTitleHandlers = append(p.setTitleHandlers[:i], p.setTitleHandlers[i+1:]...)
			break
		}
	}
}

type ZxdgToplevelV6SetAppIdEvent struct {
	AppId string
}

type ZxdgToplevelV6SetAppIdHandler interface {
	HandleZxdgToplevelV6SetAppId(ZxdgToplevelV6SetAppIdEvent)
}

func (p *ZxdgToplevelV6) AddSetAppIdHandler(h ZxdgToplevelV6SetAppIdHandler) {
	if h != nil {
		p.mu.Lock()
		p.setAppIdHandlers = append(p.setAppIdHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZxdgToplevelV6) RemoveSetAppIdHandler(h ZxdgToplevelV6SetAppIdHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.setAppIdHandlers {
		if e == h {
			p.setAppIdHandlers = append(p.setAppIdHandlers[:i], p.setAppIdHandlers[i+1:]...)
			break
		}
	}
}

type ZxdgToplevelV6ShowWindowMenuEvent struct {
	Seat   *wl.Seat
	Serial uint32
	X      int32
	Y      int32
}

type ZxdgToplevelV6ShowWindowMenuHandler interface {
	HandleZxdgToplevelV6ShowWindowMenu(ZxdgToplevelV6ShowWindowMenuEvent)
}

func (p *ZxdgToplevelV6) AddShowWindowMenuHandler(h ZxdgToplevelV6ShowWindowMenuHandler) {
	if h != nil {
		p.mu.Lock()
		p.showWindowMenuHandlers = append(p.showWindowMenuHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZxdgToplevelV6) RemoveShowWindowMenuHandler(h ZxdgToplevelV6ShowWindowMenuHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.showWindowMenuHandlers {
		if e == h {
			p.showWindowMenuHandlers = append(p.showWindowMenuHandlers[:i], p.showWindowMenuHandlers[i+1:]...)
			break
		}
	}
}

type ZxdgToplevelV6MoveEvent struct {
	Seat   *wl.Seat
	Serial uint32
}

type ZxdgToplevelV6MoveHandler interface {
	HandleZxdgToplevelV6Move(ZxdgToplevelV6MoveEvent)
}

func (p *ZxdgToplevelV6) AddMoveHandler(h ZxdgToplevelV6MoveHandler) {
	if h != nil {
		p.mu.Lock()
		p.moveHandlers = append(p.moveHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZxdgToplevelV6) RemoveMoveHandler(h ZxdgToplevelV6MoveHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.moveHandlers {
		if e == h {
			p.moveHandlers = append(p.moveHandlers[:i], p.moveHandlers[i+1:]...)
			break
		}
	}
}

type ZxdgToplevelV6ResizeEvent struct {
	Seat   *wl.Seat
	Serial uint32
	Edges  uint32
}

type ZxdgToplevelV6ResizeHandler interface {
	HandleZxdgToplevelV6Resize(ZxdgToplevelV6ResizeEvent)
}

func (p *ZxdgToplevelV6) AddResizeHandler(h ZxdgToplevelV6ResizeHandler) {
	if h != nil {
		p.mu.Lock()
		p.resizeHandlers = append(p.resizeHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZxdgToplevelV6) RemoveResizeHandler(h ZxdgToplevelV6ResizeHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.resizeHandlers {
		if e == h {
			p.resizeHandlers = append(p.resizeHandlers[:i], p.resizeHandlers[i+1:]...)
			break
		}
	}
}

type ZxdgToplevelV6SetMaxSizeEvent struct {
	Width  int32
	Height int32
}

type ZxdgToplevelV6SetMaxSizeHandler interface {
	HandleZxdgToplevelV6SetMaxSize(ZxdgToplevelV6SetMaxSizeEvent)
}

func (p *ZxdgToplevelV6) AddSetMaxSizeHandler(h ZxdgToplevelV6SetMaxSizeHandler) {
	if h != nil {
		p.mu.Lock()
		p.setMaxSizeHandlers = append(p.setMaxSizeHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZxdgToplevelV6) RemoveSetMaxSizeHandler(h ZxdgToplevelV6SetMaxSizeHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.setMaxSizeHandlers {
		if e == h {
			p.setMaxSizeHandlers = append(p.setMaxSizeHandlers[:i], p.setMaxSizeHandlers[i+1:]...)
			break
		}
	}
}

type ZxdgToplevelV6SetMinSizeEvent struct {
	Width  int32
	Height int32
}

type ZxdgToplevelV6SetMinSizeHandler interface {
	HandleZxdgToplevelV6SetMinSize(ZxdgToplevelV6SetMinSizeEvent)
}

func (p *ZxdgToplevelV6) AddSetMinSizeHandler(h ZxdgToplevelV6SetMinSizeHandler) {
	if h != nil {
		p.mu.Lock()
		p.setMinSizeHandlers = append(p.setMinSizeHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZxdgToplevelV6) RemoveSetMinSizeHandler(h ZxdgToplevelV6SetMinSizeHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.setMinSizeHandlers {
		if e == h {
			p.setMinSizeHandlers = append(p.setMinSizeHandlers[:i], p.setMinSizeHandlers[i+1:]...)
			break
		}
	}
}

type ZxdgToplevelV6SetMaximizedEvent struct {
}

type ZxdgToplevelV6SetMaximizedHandler interface {
	HandleZxdgToplevelV6SetMaximized(ZxdgToplevelV6SetMaximizedEvent)
}

func (p *ZxdgToplevelV6) AddSetMaximizedHandler(h ZxdgToplevelV6SetMaximizedHandler) {
	if h != nil {
		p.mu.Lock()
		p.setMaximizedHandlers = append(p.setMaximizedHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZxdgToplevelV6) RemoveSetMaximizedHandler(h ZxdgToplevelV6SetMaximizedHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.setMaximizedHandlers {
		if e == h {
			p.setMaximizedHandlers = append(p.setMaximizedHandlers[:i], p.setMaximizedHandlers[i+1:]...)
			break
		}
	}
}

type ZxdgToplevelV6UnsetMaximizedEvent struct {
}

type ZxdgToplevelV6UnsetMaximizedHandler interface {
	HandleZxdgToplevelV6UnsetMaximized(ZxdgToplevelV6UnsetMaximizedEvent)
}

func (p *ZxdgToplevelV6) AddUnsetMaximizedHandler(h ZxdgToplevelV6UnsetMaximizedHandler) {
	if h != nil {
		p.mu.Lock()
		p.unsetMaximizedHandlers = append(p.unsetMaximizedHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZxdgToplevelV6) RemoveUnsetMaximizedHandler(h ZxdgToplevelV6UnsetMaximizedHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.unsetMaximizedHandlers {
		if e == h {
			p.unsetMaximizedHandlers = append(p.unsetMaximizedHandlers[:i], p.unsetMaximizedHandlers[i+1:]...)
			break
		}
	}
}

type ZxdgToplevelV6SetFullscreenEvent struct {
	Output *wl.Output
}

type ZxdgToplevelV6SetFullscreenHandler interface {
	HandleZxdgToplevelV6SetFullscreen(ZxdgToplevelV6SetFullscreenEvent)
}

func (p *ZxdgToplevelV6) AddSetFullscreenHandler(h ZxdgToplevelV6SetFullscreenHandler) {
	if h != nil {
		p.mu.Lock()
		p.setFullscreenHandlers = append(p.setFullscreenHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZxdgToplevelV6) RemoveSetFullscreenHandler(h ZxdgToplevelV6SetFullscreenHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.setFullscreenHandlers {
		if e == h {
			p.setFullscreenHandlers = append(p.setFullscreenHandlers[:i], p.setFullscreenHandlers[i+1:]...)
			break
		}
	}
}

type ZxdgToplevelV6UnsetFullscreenEvent struct {
}

type ZxdgToplevelV6UnsetFullscreenHandler interface {
	HandleZxdgToplevelV6UnsetFullscreen(ZxdgToplevelV6UnsetFullscreenEvent)
}

func (p *ZxdgToplevelV6) AddUnsetFullscreenHandler(h ZxdgToplevelV6UnsetFullscreenHandler) {
	if h != nil {
		p.mu.Lock()
		p.unsetFullscreenHandlers = append(p.unsetFullscreenHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZxdgToplevelV6) RemoveUnsetFullscreenHandler(h ZxdgToplevelV6UnsetFullscreenHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.unsetFullscreenHandlers {
		if e == h {
			p.unsetFullscreenHandlers = append(p.unsetFullscreenHandlers[:i], p.unsetFullscreenHandlers[i+1:]...)
			break
		}
	}
}

type ZxdgToplevelV6SetMinimizedEvent struct {
}

type ZxdgToplevelV6SetMinimizedHandler interface {
	HandleZxdgToplevelV6SetMinimized(ZxdgToplevelV6SetMinimizedEvent)
}

func (p *ZxdgToplevelV6) AddSetMinimizedHandler(h ZxdgToplevelV6SetMinimizedHandler) {
	if h != nil {
		p.mu.Lock()
		p.setMinimizedHandlers = append(p.setMinimizedHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZxdgToplevelV6) RemoveSetMinimizedHandler(h ZxdgToplevelV6SetMinimizedHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.setMinimizedHandlers {
		if e == h {
			p.setMinimizedHandlers = append(p.setMinimizedHandlers[:i], p.setMinimizedHandlers[i+1:]...)
			break
		}
	}
}

func (p *ZxdgToplevelV6) Dispatch(event *wl.Event) {
	switch event.Opcode {
	case 0:
		if len(p.destroyHandlers) > 0 {
			ev := ZxdgToplevelV6DestroyEvent{}
			p.mu.RLock()
			for _, h := range p.destroyHandlers {
				h.HandleZxdgToplevelV6Destroy(ev)
			}
			p.mu.RUnlock()
		}
	case 1:
		if len(p.setParentHandlers) > 0 {
			ev := ZxdgToplevelV6SetParentEvent{}
			ev.Parent = event.Proxy(p.Context()).(*ZxdgToplevelV6)
			p.mu.RLock()
			for _, h := range p.setParentHandlers {
				h.HandleZxdgToplevelV6SetParent(ev)
			}
			p.mu.RUnlock()
		}
	case 2:
		if len(p.setTitleHandlers) > 0 {
			ev := ZxdgToplevelV6SetTitleEvent{}
			ev.Title = event.String()
			p.mu.RLock()
			for _, h := range p.setTitleHandlers {
				h.HandleZxdgToplevelV6SetTitle(ev)
			}
			p.mu.RUnlock()
		}
	case 3:
		if len(p.setAppIdHandlers) > 0 {
			ev := ZxdgToplevelV6SetAppIdEvent{}
			ev.AppId = event.String()
			p.mu.RLock()
			for _, h := range p.setAppIdHandlers {
				h.HandleZxdgToplevelV6SetAppId(ev)
			}
			p.mu.RUnlock()
		}
	case 4:
		if len(p.showWindowMenuHandlers) > 0 {
			ev := ZxdgToplevelV6ShowWindowMenuEvent{}
			ev.Seat = event.Proxy(p.Context()).(*wl.Seat)
			ev.Serial = event.Uint32()
			ev.X = event.Int32()
			ev.Y = event.Int32()
			p.mu.RLock()
			for _, h := range p.showWindowMenuHandlers {
				h.HandleZxdgToplevelV6ShowWindowMenu(ev)
			}
			p.mu.RUnlock()
		}
	case 5:
		if len(p.moveHandlers) > 0 {
			ev := ZxdgToplevelV6MoveEvent{}
			ev.Seat = event.Proxy(p.Context()).(*wl.Seat)
			ev.Serial = event.Uint32()
			p.mu.RLock()
			for _, h := range p.moveHandlers {
				h.HandleZxdgToplevelV6Move(ev)
			}
			p.mu.RUnlock()
		}
	case 6:
		if len(p.resizeHandlers) > 0 {
			ev := ZxdgToplevelV6ResizeEvent{}
			ev.Seat = event.Proxy(p.Context()).(*wl.Seat)
			ev.Serial = event.Uint32()
			ev.Edges = event.Uint32()
			p.mu.RLock()
			for _, h := range p.resizeHandlers {
				h.HandleZxdgToplevelV6Resize(ev)
			}
			p.mu.RUnlock()
		}
	case 7:
		if len(p.setMaxSizeHandlers) > 0 {
			ev := ZxdgToplevelV6SetMaxSizeEvent{}
			ev.Width = event.Int32()
			ev.Height = event.Int32()
			p.mu.RLock()
			for _, h := range p.setMaxSizeHandlers {
				h.HandleZxdgToplevelV6SetMaxSize(ev)
			}
			p.mu.RUnlock()
		}
	case 8:
		if len(p.setMinSizeHandlers) > 0 {
			ev := ZxdgToplevelV6SetMinSizeEvent{}
			ev.Width = event.Int32()
			ev.Height = event.Int32()
			p.mu.RLock()
			for _, h := range p.setMinSizeHandlers {
				h.HandleZxdgToplevelV6SetMinSize(ev)
			}
			p.mu.RUnlock()
		}
	case 9:
		if len(p.setMaximizedHandlers) > 0 {
			ev := ZxdgToplevelV6SetMaximizedEvent{}
			p.mu.RLock()
			for _, h := range p.setMaximizedHandlers {
				h.HandleZxdgToplevelV6SetMaximized(ev)
			}
			p.mu.RUnlock()
		}
	case 10:
		if len(p.unsetMaximizedHandlers) > 0 {
			ev := ZxdgToplevelV6UnsetMaximizedEvent{}
			p.mu.RLock()
			for _, h := range p.unsetMaximizedHandlers {
				h.HandleZxdgToplevelV6UnsetMaximized(ev)
			}
			p.mu.RUnlock()
		}
	case 11:
		if len(p.setFullscreenHandlers) > 0 {
			ev := ZxdgToplevelV6SetFullscreenEvent{}
			ev.Output = event.Proxy(p.Context()).(*wl.Output)
			p.mu.RLock()
			for _, h := range p.setFullscreenHandlers {
				h.HandleZxdgToplevelV6SetFullscreen(ev)
			}
			p.mu.RUnlock()
		}
	case 12:
		if len(p.unsetFullscreenHandlers) > 0 {
			ev := ZxdgToplevelV6UnsetFullscreenEvent{}
			p.mu.RLock()
			for _, h := range p.unsetFullscreenHandlers {
				h.HandleZxdgToplevelV6UnsetFullscreen(ev)
			}
			p.mu.RUnlock()
		}
	case 13:
		if len(p.setMinimizedHandlers) > 0 {
			ev := ZxdgToplevelV6SetMinimizedEvent{}
			p.mu.RLock()
			for _, h := range p.setMinimizedHandlers {
				h.HandleZxdgToplevelV6SetMinimized(ev)
			}
			p.mu.RUnlock()
		}
	}
}

type ZxdgToplevelV6 struct {
	wl.BaseProxy
	mu                      sync.RWMutex
	destroyHandlers         []ZxdgToplevelV6DestroyHandler
	setParentHandlers       []ZxdgToplevelV6SetParentHandler
	setTitleHandlers        []ZxdgToplevelV6SetTitleHandler
	setAppIdHandlers        []ZxdgToplevelV6SetAppIdHandler
	showWindowMenuHandlers  []ZxdgToplevelV6ShowWindowMenuHandler
	moveHandlers            []ZxdgToplevelV6MoveHandler
	resizeHandlers          []ZxdgToplevelV6ResizeHandler
	setMaxSizeHandlers      []ZxdgToplevelV6SetMaxSizeHandler
	setMinSizeHandlers      []ZxdgToplevelV6SetMinSizeHandler
	setMaximizedHandlers    []ZxdgToplevelV6SetMaximizedHandler
	unsetMaximizedHandlers  []ZxdgToplevelV6UnsetMaximizedHandler
	setFullscreenHandlers   []ZxdgToplevelV6SetFullscreenHandler
	unsetFullscreenHandlers []ZxdgToplevelV6UnsetFullscreenHandler
	setMinimizedHandlers    []ZxdgToplevelV6SetMinimizedHandler
}

func NewZxdgToplevelV6(ctx *wl.Context, id int) *ZxdgToplevelV6 {
	ret := new(ZxdgToplevelV6)
	ctx.RegisterId(ret, id)
	return ret
}

// Configure will suggest a surface change.
//
//
// This configure event asks the client to resize its toplevel surface or
// to change its state. The configured state should not be applied
// immediately. See xdg_surface.configure for details.
//
// The width and height arguments specify a hint to the window
// about how its surface should be resized in window geometry
// coordinates. See set_window_geometry.
//
// If the width or height arguments are zero, it means the client
// should decide its own window dimension. This may happen when the
// compositor need to configure the state of the surface but doesn't
// have any information about any previous or expected dimension.
//
// The states listed in the event specify how the width/height
// arguments should be interpreted, and possibly how it should be
// drawn.
//
// Clients must send an ack_configure in response to this event. See
// xdg_surface.configure and xdg_surface.ack_configure for details.
//
func (p *ZxdgToplevelV6) Configure(width int32, height int32, states []int32) error {
	return p.Context().SendRequest(p, 0, width, height, states)
}

// Close will surface wants to be closed.
//
//
// The close event is sent by the compositor when the user
// wants the surface to be closed. This should be equivalent to
// the user clicking the close button in client-side decorations,
// if your application has any...
//
// This is only a request that the user intends to close your
// window. The client may choose to ignore this request, or show
// a dialog to ask the user to save their data...
//
func (p *ZxdgToplevelV6) Close() error {
	return p.Context().SendRequest(p, 1)
}

const (
	ZxdgToplevelV6ResizeEdgeNone        = 0
	ZxdgToplevelV6ResizeEdgeTop         = 1
	ZxdgToplevelV6ResizeEdgeBottom      = 2
	ZxdgToplevelV6ResizeEdgeLeft        = 4
	ZxdgToplevelV6ResizeEdgeTopLeft     = 5
	ZxdgToplevelV6ResizeEdgeBottomLeft  = 6
	ZxdgToplevelV6ResizeEdgeRight       = 8
	ZxdgToplevelV6ResizeEdgeTopRight    = 9
	ZxdgToplevelV6ResizeEdgeBottomRight = 10
)

const (
	ZxdgToplevelV6StateMaximized  = 1
	ZxdgToplevelV6StateFullscreen = 2
	ZxdgToplevelV6StateResizing   = 3
	ZxdgToplevelV6StateActivated  = 4
)

type ZxdgPopupV6DestroyEvent struct {
}

type ZxdgPopupV6DestroyHandler interface {
	HandleZxdgPopupV6Destroy(ZxdgPopupV6DestroyEvent)
}

func (p *ZxdgPopupV6) AddDestroyHandler(h ZxdgPopupV6DestroyHandler) {
	if h != nil {
		p.mu.Lock()
		p.destroyHandlers = append(p.destroyHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZxdgPopupV6) RemoveDestroyHandler(h ZxdgPopupV6DestroyHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.destroyHandlers {
		if e == h {
			p.destroyHandlers = append(p.destroyHandlers[:i], p.destroyHandlers[i+1:]...)
			break
		}
	}
}

type ZxdgPopupV6GrabEvent struct {
	Seat   *wl.Seat
	Serial uint32
}

type ZxdgPopupV6GrabHandler interface {
	HandleZxdgPopupV6Grab(ZxdgPopupV6GrabEvent)
}

func (p *ZxdgPopupV6) AddGrabHandler(h ZxdgPopupV6GrabHandler) {
	if h != nil {
		p.mu.Lock()
		p.grabHandlers = append(p.grabHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZxdgPopupV6) RemoveGrabHandler(h ZxdgPopupV6GrabHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.grabHandlers {
		if e == h {
			p.grabHandlers = append(p.grabHandlers[:i], p.grabHandlers[i+1:]...)
			break
		}
	}
}

func (p *ZxdgPopupV6) Dispatch(event *wl.Event) {
	switch event.Opcode {
	case 0:
		if len(p.destroyHandlers) > 0 {
			ev := ZxdgPopupV6DestroyEvent{}
			p.mu.RLock()
			for _, h := range p.destroyHandlers {
				h.HandleZxdgPopupV6Destroy(ev)
			}
			p.mu.RUnlock()
		}
	case 1:
		if len(p.grabHandlers) > 0 {
			ev := ZxdgPopupV6GrabEvent{}
			ev.Seat = event.Proxy(p.Context()).(*wl.Seat)
			ev.Serial = event.Uint32()
			p.mu.RLock()
			for _, h := range p.grabHandlers {
				h.HandleZxdgPopupV6Grab(ev)
			}
			p.mu.RUnlock()
		}
	}
}

type ZxdgPopupV6 struct {
	wl.BaseProxy
	mu              sync.RWMutex
	destroyHandlers []ZxdgPopupV6DestroyHandler
	grabHandlers    []ZxdgPopupV6GrabHandler
}

func NewZxdgPopupV6(ctx *wl.Context, id int) *ZxdgPopupV6 {
	ret := new(ZxdgPopupV6)
	ctx.RegisterId(ret, id)
	return ret
}

// Configure will configure the popup surface.
//
//
// This event asks the popup surface to configure itself given the
// configuration. The configured state should not be applied immediately.
// See xdg_surface.configure for details.
//
// The x and y arguments represent the position the popup was placed at
// given the xdg_positioner rule, relative to the upper left corner of the
// window geometry of the parent surface.
//
func (p *ZxdgPopupV6) Configure(x int32, y int32, width int32, height int32) error {
	return p.Context().SendRequest(p, 0, x, y, width, height)
}

// PopupDone will popup interaction is done.
//
//
// The popup_done event is sent out when a popup is dismissed by the
// compositor. The client should destroy the xdg_popup object at this
// point.
//
func (p *ZxdgPopupV6) PopupDone() error {
	return p.Context().SendRequest(p, 1)
}

const (
	ZxdgPopupV6ErrorInvalidGrab = 0
)