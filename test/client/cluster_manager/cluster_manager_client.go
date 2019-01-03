// Code generated by go-swagger; DO NOT EDIT.

package cluster_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new cluster manager API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cluster manager API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddClusterNodes adds cluster nodes
*/
func (a *Client) AddClusterNodes(params *AddClusterNodesParams, authInfo runtime.ClientAuthInfoWriter) (*AddClusterNodesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddClusterNodesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddClusterNodes",
		Method:             "POST",
		PathPattern:        "/v1/clusters/add_nodes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddClusterNodesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddClusterNodesOK), nil

}

/*
AttachKeyPairs attaches key pairs
*/
func (a *Client) AttachKeyPairs(params *AttachKeyPairsParams, authInfo runtime.ClientAuthInfoWriter) (*AttachKeyPairsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAttachKeyPairsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AttachKeyPairs",
		Method:             "POST",
		PathPattern:        "/v1/clusters/key_pair/attach",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AttachKeyPairsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AttachKeyPairsOK), nil

}

/*
CeaseClusters ceases clusters
*/
func (a *Client) CeaseClusters(params *CeaseClustersParams, authInfo runtime.ClientAuthInfoWriter) (*CeaseClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCeaseClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CeaseClusters",
		Method:             "POST",
		PathPattern:        "/v1/clusters/cease",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CeaseClustersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CeaseClustersOK), nil

}

/*
CreateCluster creates cluster
*/
func (a *Client) CreateCluster(params *CreateClusterParams, authInfo runtime.ClientAuthInfoWriter) (*CreateClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateCluster",
		Method:             "POST",
		PathPattern:        "/v1/clusters/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateClusterOK), nil

}

/*
CreateDebugCluster creates debug cluster
*/
func (a *Client) CreateDebugCluster(params *CreateDebugClusterParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDebugClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDebugClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateDebugCluster",
		Method:             "POST",
		PathPattern:        "/v1/debug_clusters/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateDebugClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateDebugClusterOK), nil

}

/*
CreateKeyPair creates key pair
*/
func (a *Client) CreateKeyPair(params *CreateKeyPairParams, authInfo runtime.ClientAuthInfoWriter) (*CreateKeyPairOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateKeyPairParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateKeyPair",
		Method:             "POST",
		PathPattern:        "/v1/clusters/key_pairs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateKeyPairReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateKeyPairOK), nil

}

/*
DeleteClusterNodes deletes cluster nodes
*/
func (a *Client) DeleteClusterNodes(params *DeleteClusterNodesParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteClusterNodesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteClusterNodesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteClusterNodes",
		Method:             "POST",
		PathPattern:        "/v1/clusters/delete_nodes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteClusterNodesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteClusterNodesOK), nil

}

/*
DeleteClusters deletes clusters
*/
func (a *Client) DeleteClusters(params *DeleteClustersParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteClusters",
		Method:             "POST",
		PathPattern:        "/v1/clusters/delete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteClustersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteClustersOK), nil

}

/*
DeleteKeyPairs deletes key pairs
*/
func (a *Client) DeleteKeyPairs(params *DeleteKeyPairsParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteKeyPairsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteKeyPairsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteKeyPairs",
		Method:             "DELETE",
		PathPattern:        "/v1/clusters/key_pairs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteKeyPairsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteKeyPairsOK), nil

}

/*
DescribeAppClusters describes app clusters
*/
func (a *Client) DescribeAppClusters(params *DescribeAppClustersParams, authInfo runtime.ClientAuthInfoWriter) (*DescribeAppClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeAppClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DescribeAppClusters",
		Method:             "GET",
		PathPattern:        "/v1/clusters/apps",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DescribeAppClustersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DescribeAppClustersOK), nil

}

/*
DescribeClusterNodes describes cluster nodes
*/
func (a *Client) DescribeClusterNodes(params *DescribeClusterNodesParams, authInfo runtime.ClientAuthInfoWriter) (*DescribeClusterNodesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeClusterNodesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DescribeClusterNodes",
		Method:             "GET",
		PathPattern:        "/v1/clusters/nodes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DescribeClusterNodesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DescribeClusterNodesOK), nil

}

/*
DescribeClusters describes clusters
*/
func (a *Client) DescribeClusters(params *DescribeClustersParams, authInfo runtime.ClientAuthInfoWriter) (*DescribeClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DescribeClusters",
		Method:             "GET",
		PathPattern:        "/v1/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DescribeClustersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DescribeClustersOK), nil

}

/*
DescribeDebugClusters describes debug clusters
*/
func (a *Client) DescribeDebugClusters(params *DescribeDebugClustersParams, authInfo runtime.ClientAuthInfoWriter) (*DescribeDebugClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeDebugClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DescribeDebugClusters",
		Method:             "GET",
		PathPattern:        "/v1/debug_clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DescribeDebugClustersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DescribeDebugClustersOK), nil

}

/*
DescribeKeyPairs describes key pairs
*/
func (a *Client) DescribeKeyPairs(params *DescribeKeyPairsParams, authInfo runtime.ClientAuthInfoWriter) (*DescribeKeyPairsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeKeyPairsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DescribeKeyPairs",
		Method:             "GET",
		PathPattern:        "/v1/clusters/key_pairs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DescribeKeyPairsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DescribeKeyPairsOK), nil

}

/*
DescribeSubnets describes subnets
*/
func (a *Client) DescribeSubnets(params *DescribeSubnetsParams, authInfo runtime.ClientAuthInfoWriter) (*DescribeSubnetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeSubnetsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DescribeSubnets",
		Method:             "GET",
		PathPattern:        "/v1/clusters/subnets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DescribeSubnetsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DescribeSubnetsOK), nil

}

/*
DetachKeyPairs detaches key pairs
*/
func (a *Client) DetachKeyPairs(params *DetachKeyPairsParams, authInfo runtime.ClientAuthInfoWriter) (*DetachKeyPairsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDetachKeyPairsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DetachKeyPairs",
		Method:             "POST",
		PathPattern:        "/v1/clusters/key_pair/detach",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DetachKeyPairsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DetachKeyPairsOK), nil

}

/*
GetClusterStatistics gets cluster statistics
*/
func (a *Client) GetClusterStatistics(params *GetClusterStatisticsParams, authInfo runtime.ClientAuthInfoWriter) (*GetClusterStatisticsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterStatisticsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClusterStatistics",
		Method:             "GET",
		PathPattern:        "/v1/clusters/statistics",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetClusterStatisticsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterStatisticsOK), nil

}

/*
ModifyClusterAttributes modifies cluster attributes
*/
func (a *Client) ModifyClusterAttributes(params *ModifyClusterAttributesParams, authInfo runtime.ClientAuthInfoWriter) (*ModifyClusterAttributesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyClusterAttributesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ModifyClusterAttributes",
		Method:             "POST",
		PathPattern:        "/v1/clusters/modify",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ModifyClusterAttributesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifyClusterAttributesOK), nil

}

/*
ModifyClusterNodeAttributes modifies cluster node attributes
*/
func (a *Client) ModifyClusterNodeAttributes(params *ModifyClusterNodeAttributesParams, authInfo runtime.ClientAuthInfoWriter) (*ModifyClusterNodeAttributesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyClusterNodeAttributesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ModifyClusterNodeAttributes",
		Method:             "POST",
		PathPattern:        "/v1/clusters/modify_nodes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ModifyClusterNodeAttributesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifyClusterNodeAttributesOK), nil

}

/*
RecoverClusters recovers clusters
*/
func (a *Client) RecoverClusters(params *RecoverClustersParams, authInfo runtime.ClientAuthInfoWriter) (*RecoverClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRecoverClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RecoverClusters",
		Method:             "POST",
		PathPattern:        "/v1/clusters/recover",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RecoverClustersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RecoverClustersOK), nil

}

/*
ResizeCluster resizes cluster
*/
func (a *Client) ResizeCluster(params *ResizeClusterParams, authInfo runtime.ClientAuthInfoWriter) (*ResizeClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResizeClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ResizeCluster",
		Method:             "POST",
		PathPattern:        "/v1/clusters/resize",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ResizeClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ResizeClusterOK), nil

}

/*
RollbackCluster rollbacks cluster
*/
func (a *Client) RollbackCluster(params *RollbackClusterParams, authInfo runtime.ClientAuthInfoWriter) (*RollbackClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRollbackClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RollbackCluster",
		Method:             "POST",
		PathPattern:        "/v1/clusters/rollback",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RollbackClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RollbackClusterOK), nil

}

/*
StartClusters starts clusters
*/
func (a *Client) StartClusters(params *StartClustersParams, authInfo runtime.ClientAuthInfoWriter) (*StartClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StartClusters",
		Method:             "POST",
		PathPattern:        "/v1/clusters/start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartClustersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StartClustersOK), nil

}

/*
StopClusters stops clusters
*/
func (a *Client) StopClusters(params *StopClustersParams, authInfo runtime.ClientAuthInfoWriter) (*StopClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StopClusters",
		Method:             "POST",
		PathPattern:        "/v1/clusters/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopClustersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StopClustersOK), nil

}

/*
UpdateClusterEnv updates cluster env
*/
func (a *Client) UpdateClusterEnv(params *UpdateClusterEnvParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateClusterEnvOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateClusterEnvParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateClusterEnv",
		Method:             "PATCH",
		PathPattern:        "/v1/clusters/update_env",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateClusterEnvReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateClusterEnvOK), nil

}

/*
UpgradeCluster upgrades cluster
*/
func (a *Client) UpgradeCluster(params *UpgradeClusterParams, authInfo runtime.ClientAuthInfoWriter) (*UpgradeClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpgradeClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpgradeCluster",
		Method:             "POST",
		PathPattern:        "/v1/clusters/upgrade",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpgradeClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpgradeClusterOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
