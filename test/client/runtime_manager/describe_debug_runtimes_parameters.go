// Code generated by go-swagger; DO NOT EDIT.

package runtime_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDescribeDebugRuntimesParams creates a new DescribeDebugRuntimesParams object
// with the default values initialized.
func NewDescribeDebugRuntimesParams() *DescribeDebugRuntimesParams {
	var ()
	return &DescribeDebugRuntimesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeDebugRuntimesParamsWithTimeout creates a new DescribeDebugRuntimesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDescribeDebugRuntimesParamsWithTimeout(timeout time.Duration) *DescribeDebugRuntimesParams {
	var ()
	return &DescribeDebugRuntimesParams{

		timeout: timeout,
	}
}

// NewDescribeDebugRuntimesParamsWithContext creates a new DescribeDebugRuntimesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDescribeDebugRuntimesParamsWithContext(ctx context.Context) *DescribeDebugRuntimesParams {
	var ()
	return &DescribeDebugRuntimesParams{

		Context: ctx,
	}
}

// NewDescribeDebugRuntimesParamsWithHTTPClient creates a new DescribeDebugRuntimesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDescribeDebugRuntimesParamsWithHTTPClient(client *http.Client) *DescribeDebugRuntimesParams {
	var ()
	return &DescribeDebugRuntimesParams{
		HTTPClient: client,
	}
}

/*DescribeDebugRuntimesParams contains all the parameters to send to the API endpoint
for the describe debug runtimes operation typically these are written to a http.Request
*/
type DescribeDebugRuntimesParams struct {

	/*DisplayColumns*/
	DisplayColumns []string
	/*Limit*/
	Limit *int64
	/*Offset*/
	Offset *int64
	/*Owner*/
	Owner []string
	/*Provider*/
	Provider []string
	/*RuntimeID*/
	RuntimeID []string
	/*SearchWord*/
	SearchWord *string
	/*Status*/
	Status []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the describe debug runtimes params
func (o *DescribeDebugRuntimesParams) WithTimeout(timeout time.Duration) *DescribeDebugRuntimesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe debug runtimes params
func (o *DescribeDebugRuntimesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe debug runtimes params
func (o *DescribeDebugRuntimesParams) WithContext(ctx context.Context) *DescribeDebugRuntimesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe debug runtimes params
func (o *DescribeDebugRuntimesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe debug runtimes params
func (o *DescribeDebugRuntimesParams) WithHTTPClient(client *http.Client) *DescribeDebugRuntimesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe debug runtimes params
func (o *DescribeDebugRuntimesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisplayColumns adds the displayColumns to the describe debug runtimes params
func (o *DescribeDebugRuntimesParams) WithDisplayColumns(displayColumns []string) *DescribeDebugRuntimesParams {
	o.SetDisplayColumns(displayColumns)
	return o
}

// SetDisplayColumns adds the displayColumns to the describe debug runtimes params
func (o *DescribeDebugRuntimesParams) SetDisplayColumns(displayColumns []string) {
	o.DisplayColumns = displayColumns
}

// WithLimit adds the limit to the describe debug runtimes params
func (o *DescribeDebugRuntimesParams) WithLimit(limit *int64) *DescribeDebugRuntimesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the describe debug runtimes params
func (o *DescribeDebugRuntimesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the describe debug runtimes params
func (o *DescribeDebugRuntimesParams) WithOffset(offset *int64) *DescribeDebugRuntimesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the describe debug runtimes params
func (o *DescribeDebugRuntimesParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOwner adds the owner to the describe debug runtimes params
func (o *DescribeDebugRuntimesParams) WithOwner(owner []string) *DescribeDebugRuntimesParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the describe debug runtimes params
func (o *DescribeDebugRuntimesParams) SetOwner(owner []string) {
	o.Owner = owner
}

// WithProvider adds the provider to the describe debug runtimes params
func (o *DescribeDebugRuntimesParams) WithProvider(provider []string) *DescribeDebugRuntimesParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the describe debug runtimes params
func (o *DescribeDebugRuntimesParams) SetProvider(provider []string) {
	o.Provider = provider
}

// WithRuntimeID adds the runtimeID to the describe debug runtimes params
func (o *DescribeDebugRuntimesParams) WithRuntimeID(runtimeID []string) *DescribeDebugRuntimesParams {
	o.SetRuntimeID(runtimeID)
	return o
}

// SetRuntimeID adds the runtimeId to the describe debug runtimes params
func (o *DescribeDebugRuntimesParams) SetRuntimeID(runtimeID []string) {
	o.RuntimeID = runtimeID
}

// WithSearchWord adds the searchWord to the describe debug runtimes params
func (o *DescribeDebugRuntimesParams) WithSearchWord(searchWord *string) *DescribeDebugRuntimesParams {
	o.SetSearchWord(searchWord)
	return o
}

// SetSearchWord adds the searchWord to the describe debug runtimes params
func (o *DescribeDebugRuntimesParams) SetSearchWord(searchWord *string) {
	o.SearchWord = searchWord
}

// WithStatus adds the status to the describe debug runtimes params
func (o *DescribeDebugRuntimesParams) WithStatus(status []string) *DescribeDebugRuntimesParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the describe debug runtimes params
func (o *DescribeDebugRuntimesParams) SetStatus(status []string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeDebugRuntimesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesDisplayColumns := o.DisplayColumns

	joinedDisplayColumns := swag.JoinByFormat(valuesDisplayColumns, "multi")
	// query array param display_columns
	if err := r.SetQueryParam("display_columns", joinedDisplayColumns...); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	valuesOwner := o.Owner

	joinedOwner := swag.JoinByFormat(valuesOwner, "multi")
	// query array param owner
	if err := r.SetQueryParam("owner", joinedOwner...); err != nil {
		return err
	}

	valuesProvider := o.Provider

	joinedProvider := swag.JoinByFormat(valuesProvider, "multi")
	// query array param provider
	if err := r.SetQueryParam("provider", joinedProvider...); err != nil {
		return err
	}

	valuesRuntimeID := o.RuntimeID

	joinedRuntimeID := swag.JoinByFormat(valuesRuntimeID, "multi")
	// query array param runtime_id
	if err := r.SetQueryParam("runtime_id", joinedRuntimeID...); err != nil {
		return err
	}

	if o.SearchWord != nil {

		// query param search_word
		var qrSearchWord string
		if o.SearchWord != nil {
			qrSearchWord = *o.SearchWord
		}
		qSearchWord := qrSearchWord
		if qSearchWord != "" {
			if err := r.SetQueryParam("search_word", qSearchWord); err != nil {
				return err
			}
		}

	}

	valuesStatus := o.Status

	joinedStatus := swag.JoinByFormat(valuesStatus, "multi")
	// query array param status
	if err := r.SetQueryParam("status", joinedStatus...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
