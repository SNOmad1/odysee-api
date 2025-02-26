// Code generated by goa v3.5.2, DO NOT EDIT.
//
// reporter HTTP server
//
// Command:
// $ goa gen github.com/lbryio/lbrytv/apps/watchman/design -o apps/watchman

package server

import (
	"context"
	"net/http"
	"regexp"

	reporter "github.com/lbryio/lbrytv/apps/watchman/gen/reporter"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the reporter service endpoint HTTP handlers.
type Server struct {
	Mounts  []*MountPoint
	Add     http.Handler
	Healthz http.Handler
	CORS    http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the reporter service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *reporter.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Add", "POST", "/reports/playback"},
			{"Healthz", "GET", "/healthz"},
			{"CORS", "OPTIONS", "/reports/playback"},
			{"CORS", "OPTIONS", "/healthz"},
		},
		Add:     NewAddHandler(e.Add, mux, decoder, encoder, errhandler, formatter),
		Healthz: NewHealthzHandler(e.Healthz, mux, decoder, encoder, errhandler, formatter),
		CORS:    NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "reporter" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Add = m(s.Add)
	s.Healthz = m(s.Healthz)
	s.CORS = m(s.CORS)
}

// Mount configures the mux to serve the reporter endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountAddHandler(mux, h.Add)
	MountHealthzHandler(mux, h.Healthz)
	MountCORSHandler(mux, h.CORS)
}

// MountAddHandler configures the mux to serve the "reporter" service "add"
// endpoint.
func MountAddHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleReporterOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/reports/playback", f)
}

// NewAddHandler creates a HTTP handler which loads the HTTP request and calls
// the "reporter" service "add" endpoint.
func NewAddHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeAddRequest(mux, decoder)
		encodeResponse = EncodeAddResponse(encoder)
		encodeError    = EncodeAddError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "add")
		ctx = context.WithValue(ctx, goa.ServiceKey, "reporter")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountHealthzHandler configures the mux to serve the "reporter" service
// "healthz" endpoint.
func MountHealthzHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleReporterOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/healthz", f)
}

// NewHealthzHandler creates a HTTP handler which loads the HTTP request and
// calls the "reporter" service "healthz" endpoint.
func NewHealthzHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeHealthzResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "healthz")
		ctx = context.WithValue(ctx, goa.ServiceKey, "reporter")
		var err error
		res, err := endpoint(ctx, nil)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service reporter.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = HandleReporterOrigin(h)
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("OPTIONS", "/reports/playback", f)
	mux.Handle("OPTIONS", "/healthz", f)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// HandleReporterOrigin applies the CORS response headers corresponding to the
// origin for the service reporter.
func HandleReporterOrigin(h http.Handler) http.Handler {
	spec0 := regexp.MustCompile("(http:\\/\\/localhost:\\d+)|(https:\\/\\/odysee.com)|(https:\\/\\/.+\\.odysee.com)|(https:\\/\\/.+\\.lbry.tv)")
	origHndlr := h.(http.HandlerFunc)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			origHndlr(w, r)
			return
		}
		if cors.MatchOriginRegexp(origin, spec0) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Max-Age", "600")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
				w.Header().Set("Access-Control-Allow-Headers", "content-type")
			}
			origHndlr(w, r)
			return
		}
		origHndlr(w, r)
		return
	})
}
