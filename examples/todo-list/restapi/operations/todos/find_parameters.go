// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewFindParams creates a new FindParams object
// no default values defined in spec.
func NewFindParams() FindParams {

	return FindParams{}
}

// FindParams contains all the bound params for the find operation
// typically these are obtained from a http.Request
//
// swagger:parameters find
type FindParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: header
	*/
	XRateLimit int32
	/*
	  Required: true
	  In: formData
	*/
	Limit int32
	/*
	  Required: true
	  In: formData
	  Collection Format: multi
	*/
	Tags []int32
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewFindParams() beforehand.
func (o *FindParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		if err != http.ErrNotMultipart {
			return errors.New(400, "%v", err)
		} else if err := r.ParseForm(); err != nil {
			return errors.New(400, "%v", err)
		}
	}
	fds := runtime.Values(r.Form)

	if err := o.bindXRateLimit(r.Header[http.CanonicalHeaderKey("X-Rate-Limit")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	fdLimit, fdhkLimit, _ := fds.GetOK("limit")
	if err := o.bindLimit(fdLimit, fdhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	fdTags, fdhkTags, _ := fds.GetOK("tags")
	if err := o.bindTags(fdTags, fdhkTags, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FindParams) bindXRateLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("X-Rate-Limit", "header")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("X-Rate-Limit", "header", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("X-Rate-Limit", "header", "int32", raw)
	}
	o.XRateLimit = value

	return nil
}

func (o *FindParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("limit", "formData")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("limit", "formData", "int32", raw)
	}
	o.Limit = value

	return nil
}

func (o *FindParams) bindTags(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("tags", "formData")
	}

	// CollectionFormat: multi
	tagsIC := rawData

	if len(tagsIC) == 0 {
		return nil
	}

	var tagsIR []int32
	for i, tagsIV := range tagsIC {
		// items.Format: "int32"
		tagsI, err := swag.ConvertInt32(tagsIV)
		if err != nil {
			return errors.InvalidType(fmt.Sprintf("%s.%v", "tags", i), "formData", "int32", tagsI)
		}

		tagsIR = append(tagsIR, tagsI)
	}

	o.Tags = tagsIR

	return nil
}
