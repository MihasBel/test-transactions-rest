// Code generated by ogen, DO NOT EDIT.

package gen

import (
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func encodeGetTransactionByIdResponse(response GetTransactionByIdRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *Transaction:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *GetTransactionByIdBadRequest:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	case *GetTransactionByIdNotFound:
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodePlaceTransactionResponse(response PlaceTransactionRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PlaceTransactionAccepted:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(202)
		span.SetStatus(codes.Ok, http.StatusText(202))

		e := jx.GetEncoder()
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *PlaceTransactionBadRequest:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}
