package components

import (
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

// var (
// 	renderer = reflect.TypeOf(new(render.Renderer)).Elem()
// )

type (
	Response struct {
		Error Error `json:"error"`
	}

	Error struct {
		ID             uuid.UUID     `json:"id,omitempty" example:"c77fa521-99b1-4c54-9a8d-4b6902912eb0"`
		Err            error         `json:"-"`                                           // low-level runtime error
		HTTPStatusCode int           `json:"code,omitempty" example:"400"`                // http response status code
		ErrorText      string        `json:"description,omitempty" example:"Bad Request"` // application-level error message, for debugging
		ErrorDetails   []ErrorDetail `json:"error_details,omitempty"`
	}

	ErrorDetail struct {
		Attribute string   `json:"attribute,omitempty" example:"field name with error or key for help messages"`
		Messages  []string `json:"messages,omitempty" example:"explanatory messages about the attribute error"`
	}
)

// func (e *Response) Render(_ http.ResponseWriter, request *http.Request) error {
// 	render.Status(request, e.Error.HTTPStatusCode)
// 	return nil
// }

// func errorResponse(module string, httpCode int, err error) render.Renderer {

// 	// Custom service responses
// 	vof := reflect.ValueOf(err)
// 	if vof.Type().Implements(renderer) {
// 		return vof.Interface().(render.Renderer)
// 	}

// 	// Default response
// 	response := &Response{
// 		Error: Error{
// 			ID:             uuid.New(),
// 			Err:            err,
// 			HTTPStatusCode: httpCode,
// 			ErrorText:      err.Error(),
// 			ErrorDetails: []ErrorDetail{
// 				{
// 					Attribute: "module",
// 					Messages:  []string{module},
// 				},
// 			},
// 		},
// 	}

// 	return response
// }

// func errorBeautyResponse(module string, httpCode int, err error) render.Renderer {
// 	response := &Response{
// 		Error: Error{
// 			ID:             uuid.New(),
// 			Err:            err,
// 			HTTPStatusCode: httpCode,
// 			ErrorText:      err.Error(),
// 			ErrorDetails: []ErrorDetail{
// 				{
// 					Attribute: "module",
// 					Messages:  []string{module},
// 				},
// 			},
// 		},
// 	}

// 	return response
// }

// func HttpErrorResponse(components *HTTPComponents, httpCode int, err error) {
// 	LogError(components.HttpRequest, err, logger.EventLogLevelError)

// 	// Validation created to default the status code to 401 when the token is invalid.
// 	// As this status comes from Caradhras, this validation was performed in this function.
// 	if httpCode == http.StatusForbidden {
// 		httpCode = http.StatusUnauthorized
// 		err = errors.New(message.InvalidToken.String())
// 	}
// 	_ = render.Render(components.HttpResponse, components.HttpRequest, errorResponse(components.Components.Settings.String("application.name"), httpCode, err))
// }

// func HttpErrorBeautyResponse(components *HTTPComponents, httpCode int, err error, errorLevel ...logger.EventLogLevel) {
// 	LogError(components.HttpRequest, err, errorLevel...)
// 	_ = render.Render(components.HttpResponse, components.HttpRequest, errorBeautyResponse(components.Components.Settings.String("application.name"), httpCode, err))
// }

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
