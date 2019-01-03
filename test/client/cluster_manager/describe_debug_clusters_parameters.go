// Code generated by go-swagger; DO NOT EDIT.

package cluster_manager

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

// NewDescribeDebugClustersParams creates a new DescribeDebugClustersParams object
// with the default values initialized.
func NewDescribeDebugClustersParams() *DescribeDebugClustersParams {
	var ()
	return &DescribeDebugClustersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeDebugClustersParamsWithTimeout creates a new DescribeDebugClustersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDescribeDebugClustersParamsWithTimeout(timeout time.Duration) *DescribeDebugClustersParams {
	var ()
	return &DescribeDebugClustersParams{

		timeout: timeout,
	}
}

// NewDescribeDebugClustersParamsWithContext creates a new DescribeDebugClustersParams object
// with the default values initialized, and the ability to set a context for a request
func NewDescribeDebugClustersParamsWithContext(ctx context.Context) *DescribeDebugClustersParams {
	var ()
	return &DescribeDebugClustersParams{

		Context: ctx,
	}
}

// NewDescribeDebugClustersParamsWithHTTPClient creates a new DescribeDebugClustersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDescribeDebugClustersParamsWithHTTPClient(client *http.Client) *DescribeDebugClustersParams {
	var ()
	return &DescribeDebugClustersParams{
		HTTPClient: client,
	}
}

/*DescribeDebugClustersParams contains all the parameters to send to the API endpoint
for the describe debug clusters operation typically these are written to a http.Request
*/
type DescribeDebugClustersParams struct {

	/*AppID*/
	AppID []string
	/*ClusterID*/
	ClusterID []string
	/*ClusterType*/
	ClusterType *string
	/*CreatedDate*/
	CreatedDate *int64
	/*DisplayColumns*/
	DisplayColumns []string
	/*ExternalClusterID*/
	ExternalClusterID *string
	/*FrontgateID*/
	FrontgateID []string
	/*Limit*/
	Limit *int64
	/*Offset*/
	Offset *int64
	/*Owner*/
	Owner []string
	/*Reverse*/
	Reverse *bool
	/*RuntimeID*/
	RuntimeID []string
	/*SearchWord*/
	SearchWord *string
	/*SortKey*/
	SortKey *string
	/*Status*/
	Status []string
	/*VersionID*/
	VersionID []string
	/*WithDetail*/
	WithDetail *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the describe debug clusters params
func (o *DescribeDebugClustersParams) WithTimeout(timeout time.Duration) *DescribeDebugClustersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe debug clusters params
func (o *DescribeDebugClustersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe debug clusters params
func (o *DescribeDebugClustersParams) WithContext(ctx context.Context) *DescribeDebugClustersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe debug clusters params
func (o *DescribeDebugClustersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe debug clusters params
func (o *DescribeDebugClustersParams) WithHTTPClient(client *http.Client) *DescribeDebugClustersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe debug clusters params
func (o *DescribeDebugClustersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the describe debug clusters params
func (o *DescribeDebugClustersParams) WithAppID(appID []string) *DescribeDebugClustersParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the describe debug clusters params
func (o *DescribeDebugClustersParams) SetAppID(appID []string) {
	o.AppID = appID
}

// WithClusterID adds the clusterID to the describe debug clusters params
func (o *DescribeDebugClustersParams) WithClusterID(clusterID []string) *DescribeDebugClustersParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the describe debug clusters params
func (o *DescribeDebugClustersParams) SetClusterID(clusterID []string) {
	o.ClusterID = clusterID
}

// WithClusterType adds the clusterType to the describe debug clusters params
func (o *DescribeDebugClustersParams) WithClusterType(clusterType *string) *DescribeDebugClustersParams {
	o.SetClusterType(clusterType)
	return o
}

// SetClusterType adds the clusterType to the describe debug clusters params
func (o *DescribeDebugClustersParams) SetClusterType(clusterType *string) {
	o.ClusterType = clusterType
}

// WithCreatedDate adds the createdDate to the describe debug clusters params
func (o *DescribeDebugClustersParams) WithCreatedDate(createdDate *int64) *DescribeDebugClustersParams {
	o.SetCreatedDate(createdDate)
	return o
}

// SetCreatedDate adds the createdDate to the describe debug clusters params
func (o *DescribeDebugClustersParams) SetCreatedDate(createdDate *int64) {
	o.CreatedDate = createdDate
}

// WithDisplayColumns adds the displayColumns to the describe debug clusters params
func (o *DescribeDebugClustersParams) WithDisplayColumns(displayColumns []string) *DescribeDebugClustersParams {
	o.SetDisplayColumns(displayColumns)
	return o
}

// SetDisplayColumns adds the displayColumns to the describe debug clusters params
func (o *DescribeDebugClustersParams) SetDisplayColumns(displayColumns []string) {
	o.DisplayColumns = displayColumns
}

// WithExternalClusterID adds the externalClusterID to the describe debug clusters params
func (o *DescribeDebugClustersParams) WithExternalClusterID(externalClusterID *string) *DescribeDebugClustersParams {
	o.SetExternalClusterID(externalClusterID)
	return o
}

// SetExternalClusterID adds the externalClusterId to the describe debug clusters params
func (o *DescribeDebugClustersParams) SetExternalClusterID(externalClusterID *string) {
	o.ExternalClusterID = externalClusterID
}

// WithFrontgateID adds the frontgateID to the describe debug clusters params
func (o *DescribeDebugClustersParams) WithFrontgateID(frontgateID []string) *DescribeDebugClustersParams {
	o.SetFrontgateID(frontgateID)
	return o
}

// SetFrontgateID adds the frontgateId to the describe debug clusters params
func (o *DescribeDebugClustersParams) SetFrontgateID(frontgateID []string) {
	o.FrontgateID = frontgateID
}

// WithLimit adds the limit to the describe debug clusters params
func (o *DescribeDebugClustersParams) WithLimit(limit *int64) *DescribeDebugClustersParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the describe debug clusters params
func (o *DescribeDebugClustersParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the describe debug clusters params
func (o *DescribeDebugClustersParams) WithOffset(offset *int64) *DescribeDebugClustersParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the describe debug clusters params
func (o *DescribeDebugClustersParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOwner adds the owner to the describe debug clusters params
func (o *DescribeDebugClustersParams) WithOwner(owner []string) *DescribeDebugClustersParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the describe debug clusters params
func (o *DescribeDebugClustersParams) SetOwner(owner []string) {
	o.Owner = owner
}

// WithReverse adds the reverse to the describe debug clusters params
func (o *DescribeDebugClustersParams) WithReverse(reverse *bool) *DescribeDebugClustersParams {
	o.SetReverse(reverse)
	return o
}

// SetReverse adds the reverse to the describe debug clusters params
func (o *DescribeDebugClustersParams) SetReverse(reverse *bool) {
	o.Reverse = reverse
}

// WithRuntimeID adds the runtimeID to the describe debug clusters params
func (o *DescribeDebugClustersParams) WithRuntimeID(runtimeID []string) *DescribeDebugClustersParams {
	o.SetRuntimeID(runtimeID)
	return o
}

// SetRuntimeID adds the runtimeId to the describe debug clusters params
func (o *DescribeDebugClustersParams) SetRuntimeID(runtimeID []string) {
	o.RuntimeID = runtimeID
}

// WithSearchWord adds the searchWord to the describe debug clusters params
func (o *DescribeDebugClustersParams) WithSearchWord(searchWord *string) *DescribeDebugClustersParams {
	o.SetSearchWord(searchWord)
	return o
}

// SetSearchWord adds the searchWord to the describe debug clusters params
func (o *DescribeDebugClustersParams) SetSearchWord(searchWord *string) {
	o.SearchWord = searchWord
}

// WithSortKey adds the sortKey to the describe debug clusters params
func (o *DescribeDebugClustersParams) WithSortKey(sortKey *string) *DescribeDebugClustersParams {
	o.SetSortKey(sortKey)
	return o
}

// SetSortKey adds the sortKey to the describe debug clusters params
func (o *DescribeDebugClustersParams) SetSortKey(sortKey *string) {
	o.SortKey = sortKey
}

// WithStatus adds the status to the describe debug clusters params
func (o *DescribeDebugClustersParams) WithStatus(status []string) *DescribeDebugClustersParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the describe debug clusters params
func (o *DescribeDebugClustersParams) SetStatus(status []string) {
	o.Status = status
}

// WithVersionID adds the versionID to the describe debug clusters params
func (o *DescribeDebugClustersParams) WithVersionID(versionID []string) *DescribeDebugClustersParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the describe debug clusters params
func (o *DescribeDebugClustersParams) SetVersionID(versionID []string) {
	o.VersionID = versionID
}

// WithWithDetail adds the withDetail to the describe debug clusters params
func (o *DescribeDebugClustersParams) WithWithDetail(withDetail *bool) *DescribeDebugClustersParams {
	o.SetWithDetail(withDetail)
	return o
}

// SetWithDetail adds the withDetail to the describe debug clusters params
func (o *DescribeDebugClustersParams) SetWithDetail(withDetail *bool) {
	o.WithDetail = withDetail
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeDebugClustersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesAppID := o.AppID

	joinedAppID := swag.JoinByFormat(valuesAppID, "multi")
	// query array param app_id
	if err := r.SetQueryParam("app_id", joinedAppID...); err != nil {
		return err
	}

	valuesClusterID := o.ClusterID

	joinedClusterID := swag.JoinByFormat(valuesClusterID, "multi")
	// query array param cluster_id
	if err := r.SetQueryParam("cluster_id", joinedClusterID...); err != nil {
		return err
	}

	if o.ClusterType != nil {

		// query param cluster_type
		var qrClusterType string
		if o.ClusterType != nil {
			qrClusterType = *o.ClusterType
		}
		qClusterType := qrClusterType
		if qClusterType != "" {
			if err := r.SetQueryParam("cluster_type", qClusterType); err != nil {
				return err
			}
		}

	}

	if o.CreatedDate != nil {

		// query param created_date
		var qrCreatedDate int64
		if o.CreatedDate != nil {
			qrCreatedDate = *o.CreatedDate
		}
		qCreatedDate := swag.FormatInt64(qrCreatedDate)
		if qCreatedDate != "" {
			if err := r.SetQueryParam("created_date", qCreatedDate); err != nil {
				return err
			}
		}

	}

	valuesDisplayColumns := o.DisplayColumns

	joinedDisplayColumns := swag.JoinByFormat(valuesDisplayColumns, "multi")
	// query array param display_columns
	if err := r.SetQueryParam("display_columns", joinedDisplayColumns...); err != nil {
		return err
	}

	if o.ExternalClusterID != nil {

		// query param external_cluster_id
		var qrExternalClusterID string
		if o.ExternalClusterID != nil {
			qrExternalClusterID = *o.ExternalClusterID
		}
		qExternalClusterID := qrExternalClusterID
		if qExternalClusterID != "" {
			if err := r.SetQueryParam("external_cluster_id", qExternalClusterID); err != nil {
				return err
			}
		}

	}

	valuesFrontgateID := o.FrontgateID

	joinedFrontgateID := swag.JoinByFormat(valuesFrontgateID, "multi")
	// query array param frontgate_id
	if err := r.SetQueryParam("frontgate_id", joinedFrontgateID...); err != nil {
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

	if o.Reverse != nil {

		// query param reverse
		var qrReverse bool
		if o.Reverse != nil {
			qrReverse = *o.Reverse
		}
		qReverse := swag.FormatBool(qrReverse)
		if qReverse != "" {
			if err := r.SetQueryParam("reverse", qReverse); err != nil {
				return err
			}
		}

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

	if o.SortKey != nil {

		// query param sort_key
		var qrSortKey string
		if o.SortKey != nil {
			qrSortKey = *o.SortKey
		}
		qSortKey := qrSortKey
		if qSortKey != "" {
			if err := r.SetQueryParam("sort_key", qSortKey); err != nil {
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

	valuesVersionID := o.VersionID

	joinedVersionID := swag.JoinByFormat(valuesVersionID, "multi")
	// query array param version_id
	if err := r.SetQueryParam("version_id", joinedVersionID...); err != nil {
		return err
	}

	if o.WithDetail != nil {

		// query param with_detail
		var qrWithDetail bool
		if o.WithDetail != nil {
			qrWithDetail = *o.WithDetail
		}
		qWithDetail := swag.FormatBool(qrWithDetail)
		if qWithDetail != "" {
			if err := r.SetQueryParam("with_detail", qWithDetail); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
