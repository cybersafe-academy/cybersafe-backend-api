package components

import (
	"log"
	"net/http"
	"reflect"

	"github.com/go-chi/render"
)

var (
	renderer = reflect.TypeOf(new(render.Renderer)).Elem()
)

type (
	Response struct {
		Error Error `json:"error"`
	}

	Error struct {
		Err            error         `json:"-"`
		HTTPStatusCode int           `json:"code,omitempty" example:"400"`
		ErrorText      string        `json:"description,omitempty" example:"Bad Request"`
		ErrorDetails   []ErrorDetail `json:"error_details,omitempty"`
	}

	ErrorDetail struct {
		Attribute string   `json:"attribute,omitempty" example:"field name with error or key for help messages"`
		Messages  []string `json:"messages,omitempty" example:"explanatory messages about the attribute error"`
	}
)

func (e *Response) Render(_ http.ResponseWriter, request *http.Request) error {
	render.Status(request, e.Error.HTTPStatusCode)
	return nil
}

func errorResponse(httpCode int, err error) render.Renderer {

	// Custom service responses
	vof := reflect.ValueOf(err)
	if vof.Type().Implements(renderer) {
		return vof.Interface().(render.Renderer)
	}

	// Default response
	response := &Response{
		Error: Error{
			Err:            err,
			HTTPStatusCode: httpCode,
			ErrorText:      err.Error(),
		},
	}

	return response
}

func errorResponseMessage(httpCode int, errText string) render.Renderer {

	// Default response
	response := &Response{
		Error: Error{
			HTTPStatusCode: httpCode,
			ErrorText:      errText,
		},
	}

	return response
}

func HttpErrorResponse(components *HTTPComponents, httpCode int, err error) {
	log.Printf("Error %s", err.Error())
	_ = render.Render(components.HttpResponse, components.HttpRequest, errorResponse(httpCode, err))
}

func HttpErrorLocalizedResponse(components *HTTPComponents, httpCode int, err string) {
	log.Printf("Error %s", err)
	_ = render.Render(components.HttpResponse, components.HttpRequest, errorResponseMessage(httpCode, err))
}

func HttpErrorMiddlewareResponse(w http.ResponseWriter, r *http.Request, httpCode int, err error) {
	log.Printf("Error %s", err.Error())
	_ = render.Render(w, r, errorResponse(httpCode, err))
}

func HttpResponse(components *HTTPComponents, httpCode int) {
	render.Status(components.HttpRequest, httpCode)
	render.Respond(components.HttpResponse, components.HttpRequest, nil)
}

func HttpResponseWithPayload(components *HTTPComponents, payload any, httpCode int) {
	render.Status(components.HttpRequest, httpCode)
	render.Respond(components.HttpResponse, components.HttpRequest, payload)
}

func ValidateRequest(components *HTTPComponents, value render.Binder) error {
	if err := render.Bind(components.HttpRequest, value); err != nil {
		return err
	}
	return nil
}
