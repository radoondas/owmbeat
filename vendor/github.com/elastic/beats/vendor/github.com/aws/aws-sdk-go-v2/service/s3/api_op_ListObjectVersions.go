// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/ListObjectVersionsRequest
type ListObjectVersionsInput struct {
	_ struct{} `type:"structure"`

	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// A delimiter is a character you use to group keys.
	Delimiter *string `location:"querystring" locationName:"delimiter" type:"string"`

	// Requests Amazon S3 to encode the object keys in the response and specifies
	// the encoding method to use. An object key may contain any Unicode character;
	// however, XML 1.0 parser cannot parse some characters, such as characters
	// with an ASCII value from 0 to 10. For characters that are not supported in
	// XML 1.0, you can add this parameter to request that Amazon S3 encode the
	// keys in the response.
	EncodingType EncodingType `location:"querystring" locationName:"encoding-type" type:"string" enum:"true"`

	// Specifies the key to start with when listing objects in a bucket.
	KeyMarker *string `location:"querystring" locationName:"key-marker" type:"string"`

	// Sets the maximum number of keys returned in the response. The response might
	// contain fewer keys but will never contain more.
	MaxKeys *int64 `location:"querystring" locationName:"max-keys" type:"integer"`

	// Limits the response to keys that begin with the specified prefix.
	Prefix *string `location:"querystring" locationName:"prefix" type:"string"`

	// Specifies the object version you want to start listing from.
	VersionIdMarker *string `location:"querystring" locationName:"version-id-marker" type:"string"`
}

// String returns the string representation
func (s ListObjectVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListObjectVersionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListObjectVersionsInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *ListObjectVersionsInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListObjectVersionsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.Delimiter != nil {
		v := *s.Delimiter

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "delimiter", protocol.StringValue(v), metadata)
	}
	if len(s.EncodingType) > 0 {
		v := s.EncodingType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "encoding-type", v, metadata)
	}
	if s.KeyMarker != nil {
		v := *s.KeyMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "key-marker", protocol.StringValue(v), metadata)
	}
	if s.MaxKeys != nil {
		v := *s.MaxKeys

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "max-keys", protocol.Int64Value(v), metadata)
	}
	if s.Prefix != nil {
		v := *s.Prefix

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "prefix", protocol.StringValue(v), metadata)
	}
	if s.VersionIdMarker != nil {
		v := *s.VersionIdMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "version-id-marker", protocol.StringValue(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/ListObjectVersionsOutput
type ListObjectVersionsOutput struct {
	_ struct{} `type:"structure"`

	CommonPrefixes []CommonPrefix `type:"list" flattened:"true"`

	DeleteMarkers []DeleteMarkerEntry `locationName:"DeleteMarker" type:"list" flattened:"true"`

	Delimiter *string `type:"string"`

	// Encoding type used by Amazon S3 to encode object keys in the response.
	EncodingType EncodingType `type:"string" enum:"true"`

	// A flag that indicates whether or not Amazon S3 returned all of the results
	// that satisfied the search criteria. If your results were truncated, you can
	// make a follow-up paginated request using the NextKeyMarker and NextVersionIdMarker
	// response parameters as a starting place in another request to return the
	// rest of the results.
	IsTruncated *bool `type:"boolean"`

	// Marks the last Key returned in a truncated response.
	KeyMarker *string `type:"string"`

	MaxKeys *int64 `type:"integer"`

	Name *string `type:"string"`

	// Use this value for the key marker request parameter in a subsequent request.
	NextKeyMarker *string `type:"string"`

	// Use this value for the next version id marker parameter in a subsequent request.
	NextVersionIdMarker *string `type:"string"`

	Prefix *string `type:"string"`

	VersionIdMarker *string `type:"string"`

	Versions []ObjectVersion `locationName:"Version" type:"list" flattened:"true"`
}

// String returns the string representation
func (s ListObjectVersionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListObjectVersionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.CommonPrefixes) > 0 {
		v := s.CommonPrefixes

		metadata := protocol.Metadata{Flatten: true}
		ls0 := e.List(protocol.BodyTarget, "CommonPrefixes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if len(s.DeleteMarkers) > 0 {
		v := s.DeleteMarkers

		metadata := protocol.Metadata{Flatten: true}
		ls0 := e.List(protocol.BodyTarget, "DeleteMarker", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Delimiter != nil {
		v := *s.Delimiter

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Delimiter", protocol.StringValue(v), metadata)
	}
	if len(s.EncodingType) > 0 {
		v := s.EncodingType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "EncodingType", v, metadata)
	}
	if s.IsTruncated != nil {
		v := *s.IsTruncated

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IsTruncated", protocol.BoolValue(v), metadata)
	}
	if s.KeyMarker != nil {
		v := *s.KeyMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "KeyMarker", protocol.StringValue(v), metadata)
	}
	if s.MaxKeys != nil {
		v := *s.MaxKeys

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxKeys", protocol.Int64Value(v), metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.StringValue(v), metadata)
	}
	if s.NextKeyMarker != nil {
		v := *s.NextKeyMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextKeyMarker", protocol.StringValue(v), metadata)
	}
	if s.NextVersionIdMarker != nil {
		v := *s.NextVersionIdMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextVersionIdMarker", protocol.StringValue(v), metadata)
	}
	if s.Prefix != nil {
		v := *s.Prefix

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Prefix", protocol.StringValue(v), metadata)
	}
	if s.VersionIdMarker != nil {
		v := *s.VersionIdMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "VersionIdMarker", protocol.StringValue(v), metadata)
	}
	if len(s.Versions) > 0 {
		v := s.Versions

		metadata := protocol.Metadata{Flatten: true}
		ls0 := e.List(protocol.BodyTarget, "Version", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListObjectVersions = "ListObjectVersions"

// ListObjectVersionsRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Returns metadata about all of the versions of objects in a bucket.
//
//    // Example sending a request using ListObjectVersionsRequest.
//    req := client.ListObjectVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/ListObjectVersions
func (c *Client) ListObjectVersionsRequest(input *ListObjectVersionsInput) ListObjectVersionsRequest {
	op := &aws.Operation{
		Name:       opListObjectVersions,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?versions",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"KeyMarker", "VersionIdMarker"},
			OutputTokens:    []string{"NextKeyMarker", "NextVersionIdMarker"},
			LimitToken:      "MaxKeys",
			TruncationToken: "IsTruncated",
		},
	}

	if input == nil {
		input = &ListObjectVersionsInput{}
	}

	req := c.newRequest(op, input, &ListObjectVersionsOutput{})
	return ListObjectVersionsRequest{Request: req, Input: input, Copy: c.ListObjectVersionsRequest}
}

// ListObjectVersionsRequest is the request type for the
// ListObjectVersions API operation.
type ListObjectVersionsRequest struct {
	*aws.Request
	Input *ListObjectVersionsInput
	Copy  func(*ListObjectVersionsInput) ListObjectVersionsRequest
}

// Send marshals and sends the ListObjectVersions API request.
func (r ListObjectVersionsRequest) Send(ctx context.Context) (*ListObjectVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListObjectVersionsResponse{
		ListObjectVersionsOutput: r.Request.Data.(*ListObjectVersionsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListObjectVersionsRequestPaginator returns a paginator for ListObjectVersions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListObjectVersionsRequest(input)
//   p := s3.NewListObjectVersionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListObjectVersionsPaginator(req ListObjectVersionsRequest) ListObjectVersionsPaginator {
	return ListObjectVersionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListObjectVersionsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListObjectVersionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListObjectVersionsPaginator struct {
	aws.Pager
}

func (p *ListObjectVersionsPaginator) CurrentPage() *ListObjectVersionsOutput {
	return p.Pager.CurrentPage().(*ListObjectVersionsOutput)
}

// ListObjectVersionsResponse is the response type for the
// ListObjectVersions API operation.
type ListObjectVersionsResponse struct {
	*ListObjectVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListObjectVersions request.
func (r *ListObjectVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
