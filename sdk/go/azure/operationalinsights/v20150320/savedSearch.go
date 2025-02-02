// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150320

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Value object for saved search results.
type SavedSearch struct {
	pulumi.CustomResourceState

	// The category of the saved search. This helps the user to find a saved search faster.
	Category pulumi.StringOutput `pulumi:"category"`
	// Saved search display name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The ETag of the saved search.
	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	// The name of the saved search.
	Name pulumi.StringOutput `pulumi:"name"`
	// The query expression for the saved search. Please see https://docs.microsoft.com/en-us/azure/log-analytics/log-analytics-search-reference for reference.
	Query pulumi.StringOutput `pulumi:"query"`
	// The tags attached to the saved search.
	Tags TagResponseArrayOutput `pulumi:"tags"`
	// The type of the saved search.
	Type pulumi.StringOutput `pulumi:"type"`
	// The version number of the query language. The current version is 2 and is the default.
	Version pulumi.Float64PtrOutput `pulumi:"version"`
}

// NewSavedSearch registers a new resource with the given unique name, arguments, and options.
func NewSavedSearch(ctx *pulumi.Context,
	name string, args *SavedSearchArgs, opts ...pulumi.ResourceOption) (*SavedSearch, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Category == nil {
		return nil, errors.New("invalid value for required argument 'Category'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Query == nil {
		return nil, errors.New("invalid value for required argument 'Query'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:operationalinsights/v20150320:SavedSearch"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights:SavedSearch"),
		},
		{
			Type: pulumi.String("azure-nextgen:operationalinsights:SavedSearch"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200301preview:SavedSearch"),
		},
		{
			Type: pulumi.String("azure-nextgen:operationalinsights/v20200301preview:SavedSearch"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200801:SavedSearch"),
		},
		{
			Type: pulumi.String("azure-nextgen:operationalinsights/v20200801:SavedSearch"),
		},
	})
	opts = append(opts, aliases)
	var resource SavedSearch
	err := ctx.RegisterResource("azure-native:operationalinsights/v20150320:SavedSearch", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSavedSearch gets an existing SavedSearch resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSavedSearch(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SavedSearchState, opts ...pulumi.ResourceOption) (*SavedSearch, error) {
	var resource SavedSearch
	err := ctx.ReadResource("azure-native:operationalinsights/v20150320:SavedSearch", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SavedSearch resources.
type savedSearchState struct {
	// The category of the saved search. This helps the user to find a saved search faster.
	Category *string `pulumi:"category"`
	// Saved search display name.
	DisplayName *string `pulumi:"displayName"`
	// The ETag of the saved search.
	ETag *string `pulumi:"eTag"`
	// The name of the saved search.
	Name *string `pulumi:"name"`
	// The query expression for the saved search. Please see https://docs.microsoft.com/en-us/azure/log-analytics/log-analytics-search-reference for reference.
	Query *string `pulumi:"query"`
	// The tags attached to the saved search.
	Tags []TagResponse `pulumi:"tags"`
	// The type of the saved search.
	Type *string `pulumi:"type"`
	// The version number of the query language. The current version is 2 and is the default.
	Version *float64 `pulumi:"version"`
}

type SavedSearchState struct {
	// The category of the saved search. This helps the user to find a saved search faster.
	Category pulumi.StringPtrInput
	// Saved search display name.
	DisplayName pulumi.StringPtrInput
	// The ETag of the saved search.
	ETag pulumi.StringPtrInput
	// The name of the saved search.
	Name pulumi.StringPtrInput
	// The query expression for the saved search. Please see https://docs.microsoft.com/en-us/azure/log-analytics/log-analytics-search-reference for reference.
	Query pulumi.StringPtrInput
	// The tags attached to the saved search.
	Tags TagResponseArrayInput
	// The type of the saved search.
	Type pulumi.StringPtrInput
	// The version number of the query language. The current version is 2 and is the default.
	Version pulumi.Float64PtrInput
}

func (SavedSearchState) ElementType() reflect.Type {
	return reflect.TypeOf((*savedSearchState)(nil)).Elem()
}

type savedSearchArgs struct {
	// The category of the saved search. This helps the user to find a saved search faster.
	Category string `pulumi:"category"`
	// Saved search display name.
	DisplayName string `pulumi:"displayName"`
	// The ETag of the saved search.
	ETag *string `pulumi:"eTag"`
	// The query expression for the saved search. Please see https://docs.microsoft.com/en-us/azure/log-analytics/log-analytics-search-reference for reference.
	Query string `pulumi:"query"`
	// The Resource Group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The id of the saved search.
	SavedSearchId *string `pulumi:"savedSearchId"`
	// The tags attached to the saved search.
	Tags []Tag `pulumi:"tags"`
	// The version number of the query language. The current version is 2 and is the default.
	Version *float64 `pulumi:"version"`
	// The Log Analytics Workspace name.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a SavedSearch resource.
type SavedSearchArgs struct {
	// The category of the saved search. This helps the user to find a saved search faster.
	Category pulumi.StringInput
	// Saved search display name.
	DisplayName pulumi.StringInput
	// The ETag of the saved search.
	ETag pulumi.StringPtrInput
	// The query expression for the saved search. Please see https://docs.microsoft.com/en-us/azure/log-analytics/log-analytics-search-reference for reference.
	Query pulumi.StringInput
	// The Resource Group name.
	ResourceGroupName pulumi.StringInput
	// The id of the saved search.
	SavedSearchId pulumi.StringPtrInput
	// The tags attached to the saved search.
	Tags TagArrayInput
	// The version number of the query language. The current version is 2 and is the default.
	Version pulumi.Float64PtrInput
	// The Log Analytics Workspace name.
	WorkspaceName pulumi.StringInput
}

func (SavedSearchArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*savedSearchArgs)(nil)).Elem()
}

type SavedSearchInput interface {
	pulumi.Input

	ToSavedSearchOutput() SavedSearchOutput
	ToSavedSearchOutputWithContext(ctx context.Context) SavedSearchOutput
}

func (*SavedSearch) ElementType() reflect.Type {
	return reflect.TypeOf((*SavedSearch)(nil))
}

func (i *SavedSearch) ToSavedSearchOutput() SavedSearchOutput {
	return i.ToSavedSearchOutputWithContext(context.Background())
}

func (i *SavedSearch) ToSavedSearchOutputWithContext(ctx context.Context) SavedSearchOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SavedSearchOutput)
}

type SavedSearchOutput struct {
	*pulumi.OutputState
}

func (SavedSearchOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SavedSearch)(nil))
}

func (o SavedSearchOutput) ToSavedSearchOutput() SavedSearchOutput {
	return o
}

func (o SavedSearchOutput) ToSavedSearchOutputWithContext(ctx context.Context) SavedSearchOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SavedSearchOutput{})
}
