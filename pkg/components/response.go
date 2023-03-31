package components

import (
	"net/http"
	"reflect"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

var (
	renderer = reflect.TypeOf(new(render.Renderer)).Elem()
)

type (
	Response struct {
		Error Error `json:"error"`
	}

	Error struct {
		ID             uuid.UUID     `json:"id,omitempty" example:"c77fa521-99b1-4c54-9a8d-4b6902912eb0"`
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

func errorResponse(module string, httpCode int, err error) render.Renderer {

	// Custom service responses
	vof := reflect.ValueOf(err)
	if vof.Type().Implements(renderer) {
		return vof.Interface().(render.Renderer)
	}

	// Default response
	response := &Response{
		Error: Error{
			ID:             uuid.New(),
			Err:            err,
			HTTPStatusCode: httpCode,
			ErrorText:      err.Error(),
			ErrorDetails: []ErrorDetail{
				{
					Attribute: "module",
					Messages:  []string{module},
				},
			},
		},
	}

	return response
}

func HttpErrorResponse(components *HTTPComponents, httpCode int, err error) {
	_ = render.Render(components.HttpResponse, components.HttpRequest, errorResponse(components.Components.Settings.String("application.name"), httpCode, err))
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
