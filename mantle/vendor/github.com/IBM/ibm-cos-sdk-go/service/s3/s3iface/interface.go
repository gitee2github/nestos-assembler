// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package s3iface provides an interface to enable mocking the Amazon Simple Storage Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package s3iface

import (
	"github.com/IBM/ibm-cos-sdk-go/aws"
	"github.com/IBM/ibm-cos-sdk-go/aws/request"
	"github.com/IBM/ibm-cos-sdk-go/service/s3"
)

// S3API provides an interface to enable mocking the
// s3.S3 service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Simple Storage Service.
//    func myFunc(svc s3iface.S3API) bool {
//        // Make svc.AbortMultipartUpload request
//    }
//
// IBM COS SDK Code -- START
//    func main() {
//        sess := session.Must(session.NewSession())
//        svc := s3.New(sess)
//
//        myFunc(svc)
//    }
// IBM COS SDK Code -- END
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockS3Client struct {
//        s3iface.S3API
//    }
//    func (m *mockS3Client) AbortMultipartUpload(input *s3.AbortMultipartUploadInput) (*s3.AbortMultipartUploadOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockS3Client{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type S3API interface {
	AbortMultipartUpload(*s3.AbortMultipartUploadInput) (*s3.AbortMultipartUploadOutput, error)
	AbortMultipartUploadWithContext(aws.Context, *s3.AbortMultipartUploadInput, ...request.Option) (*s3.AbortMultipartUploadOutput, error)
	AbortMultipartUploadRequest(*s3.AbortMultipartUploadInput) (*request.Request, *s3.AbortMultipartUploadOutput)

	AddLegalHold(*s3.AddLegalHoldInput) (*s3.AddLegalHoldOutput, error)
	AddLegalHoldWithContext(aws.Context, *s3.AddLegalHoldInput, ...request.Option) (*s3.AddLegalHoldOutput, error)
	AddLegalHoldRequest(*s3.AddLegalHoldInput) (*request.Request, *s3.AddLegalHoldOutput)

	CompleteMultipartUpload(*s3.CompleteMultipartUploadInput) (*s3.CompleteMultipartUploadOutput, error)
	CompleteMultipartUploadWithContext(aws.Context, *s3.CompleteMultipartUploadInput, ...request.Option) (*s3.CompleteMultipartUploadOutput, error)
	CompleteMultipartUploadRequest(*s3.CompleteMultipartUploadInput) (*request.Request, *s3.CompleteMultipartUploadOutput)

	CopyObject(*s3.CopyObjectInput) (*s3.CopyObjectOutput, error)
	CopyObjectWithContext(aws.Context, *s3.CopyObjectInput, ...request.Option) (*s3.CopyObjectOutput, error)
	CopyObjectRequest(*s3.CopyObjectInput) (*request.Request, *s3.CopyObjectOutput)

	CreateBucket(*s3.CreateBucketInput) (*s3.CreateBucketOutput, error)
	CreateBucketWithContext(aws.Context, *s3.CreateBucketInput, ...request.Option) (*s3.CreateBucketOutput, error)
	CreateBucketRequest(*s3.CreateBucketInput) (*request.Request, *s3.CreateBucketOutput)

	CreateMultipartUpload(*s3.CreateMultipartUploadInput) (*s3.CreateMultipartUploadOutput, error)
	CreateMultipartUploadWithContext(aws.Context, *s3.CreateMultipartUploadInput, ...request.Option) (*s3.CreateMultipartUploadOutput, error)
	CreateMultipartUploadRequest(*s3.CreateMultipartUploadInput) (*request.Request, *s3.CreateMultipartUploadOutput)

	DeleteBucket(*s3.DeleteBucketInput) (*s3.DeleteBucketOutput, error)
	DeleteBucketWithContext(aws.Context, *s3.DeleteBucketInput, ...request.Option) (*s3.DeleteBucketOutput, error)
	DeleteBucketRequest(*s3.DeleteBucketInput) (*request.Request, *s3.DeleteBucketOutput)

	DeleteBucketCors(*s3.DeleteBucketCorsInput) (*s3.DeleteBucketCorsOutput, error)
	DeleteBucketCorsWithContext(aws.Context, *s3.DeleteBucketCorsInput, ...request.Option) (*s3.DeleteBucketCorsOutput, error)
	DeleteBucketCorsRequest(*s3.DeleteBucketCorsInput) (*request.Request, *s3.DeleteBucketCorsOutput)

	DeleteBucketLifecycle(*s3.DeleteBucketLifecycleInput) (*s3.DeleteBucketLifecycleOutput, error)
	DeleteBucketLifecycleWithContext(aws.Context, *s3.DeleteBucketLifecycleInput, ...request.Option) (*s3.DeleteBucketLifecycleOutput, error)
	DeleteBucketLifecycleRequest(*s3.DeleteBucketLifecycleInput) (*request.Request, *s3.DeleteBucketLifecycleOutput)

	DeleteBucketWebsite(*s3.DeleteBucketWebsiteInput) (*s3.DeleteBucketWebsiteOutput, error)
	DeleteBucketWebsiteWithContext(aws.Context, *s3.DeleteBucketWebsiteInput, ...request.Option) (*s3.DeleteBucketWebsiteOutput, error)
	DeleteBucketWebsiteRequest(*s3.DeleteBucketWebsiteInput) (*request.Request, *s3.DeleteBucketWebsiteOutput)

	DeleteLegalHold(*s3.DeleteLegalHoldInput) (*s3.DeleteLegalHoldOutput, error)
	DeleteLegalHoldWithContext(aws.Context, *s3.DeleteLegalHoldInput, ...request.Option) (*s3.DeleteLegalHoldOutput, error)
	DeleteLegalHoldRequest(*s3.DeleteLegalHoldInput) (*request.Request, *s3.DeleteLegalHoldOutput)

	DeleteObject(*s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error)
	DeleteObjectWithContext(aws.Context, *s3.DeleteObjectInput, ...request.Option) (*s3.DeleteObjectOutput, error)
	DeleteObjectRequest(*s3.DeleteObjectInput) (*request.Request, *s3.DeleteObjectOutput)

	DeleteObjectTagging(*s3.DeleteObjectTaggingInput) (*s3.DeleteObjectTaggingOutput, error)
	DeleteObjectTaggingWithContext(aws.Context, *s3.DeleteObjectTaggingInput, ...request.Option) (*s3.DeleteObjectTaggingOutput, error)
	DeleteObjectTaggingRequest(*s3.DeleteObjectTaggingInput) (*request.Request, *s3.DeleteObjectTaggingOutput)

	DeleteObjects(*s3.DeleteObjectsInput) (*s3.DeleteObjectsOutput, error)
	DeleteObjectsWithContext(aws.Context, *s3.DeleteObjectsInput, ...request.Option) (*s3.DeleteObjectsOutput, error)
	DeleteObjectsRequest(*s3.DeleteObjectsInput) (*request.Request, *s3.DeleteObjectsOutput)

	ExtendObjectRetention(*s3.ExtendObjectRetentionInput) (*s3.ExtendObjectRetentionOutput, error)
	ExtendObjectRetentionWithContext(aws.Context, *s3.ExtendObjectRetentionInput, ...request.Option) (*s3.ExtendObjectRetentionOutput, error)
	ExtendObjectRetentionRequest(*s3.ExtendObjectRetentionInput) (*request.Request, *s3.ExtendObjectRetentionOutput)

	GetBucketAcl(*s3.GetBucketAclInput) (*s3.GetBucketAclOutput, error)
	GetBucketAclWithContext(aws.Context, *s3.GetBucketAclInput, ...request.Option) (*s3.GetBucketAclOutput, error)
	GetBucketAclRequest(*s3.GetBucketAclInput) (*request.Request, *s3.GetBucketAclOutput)

	GetBucketCors(*s3.GetBucketCorsInput) (*s3.GetBucketCorsOutput, error)
	GetBucketCorsWithContext(aws.Context, *s3.GetBucketCorsInput, ...request.Option) (*s3.GetBucketCorsOutput, error)
	GetBucketCorsRequest(*s3.GetBucketCorsInput) (*request.Request, *s3.GetBucketCorsOutput)

	GetBucketLifecycleConfiguration(*s3.GetBucketLifecycleConfigurationInput) (*s3.GetBucketLifecycleConfigurationOutput, error)
	GetBucketLifecycleConfigurationWithContext(aws.Context, *s3.GetBucketLifecycleConfigurationInput, ...request.Option) (*s3.GetBucketLifecycleConfigurationOutput, error)
	GetBucketLifecycleConfigurationRequest(*s3.GetBucketLifecycleConfigurationInput) (*request.Request, *s3.GetBucketLifecycleConfigurationOutput)

	GetBucketLocation(*s3.GetBucketLocationInput) (*s3.GetBucketLocationOutput, error)
	GetBucketLocationWithContext(aws.Context, *s3.GetBucketLocationInput, ...request.Option) (*s3.GetBucketLocationOutput, error)
	GetBucketLocationRequest(*s3.GetBucketLocationInput) (*request.Request, *s3.GetBucketLocationOutput)

	GetBucketLogging(*s3.GetBucketLoggingInput) (*s3.GetBucketLoggingOutput, error)
	GetBucketLoggingWithContext(aws.Context, *s3.GetBucketLoggingInput, ...request.Option) (*s3.GetBucketLoggingOutput, error)
	GetBucketLoggingRequest(*s3.GetBucketLoggingInput) (*request.Request, *s3.GetBucketLoggingOutput)

	GetBucketProtectionConfiguration(*s3.GetBucketProtectionConfigurationInput) (*s3.GetBucketProtectionConfigurationOutput, error)
	GetBucketProtectionConfigurationWithContext(aws.Context, *s3.GetBucketProtectionConfigurationInput, ...request.Option) (*s3.GetBucketProtectionConfigurationOutput, error)
	GetBucketProtectionConfigurationRequest(*s3.GetBucketProtectionConfigurationInput) (*request.Request, *s3.GetBucketProtectionConfigurationOutput)

	GetBucketWebsite(*s3.GetBucketWebsiteInput) (*s3.GetBucketWebsiteOutput, error)
	GetBucketWebsiteWithContext(aws.Context, *s3.GetBucketWebsiteInput, ...request.Option) (*s3.GetBucketWebsiteOutput, error)
	GetBucketWebsiteRequest(*s3.GetBucketWebsiteInput) (*request.Request, *s3.GetBucketWebsiteOutput)

	GetObject(*s3.GetObjectInput) (*s3.GetObjectOutput, error)
	GetObjectWithContext(aws.Context, *s3.GetObjectInput, ...request.Option) (*s3.GetObjectOutput, error)
	GetObjectRequest(*s3.GetObjectInput) (*request.Request, *s3.GetObjectOutput)

	GetObjectAcl(*s3.GetObjectAclInput) (*s3.GetObjectAclOutput, error)
	GetObjectAclWithContext(aws.Context, *s3.GetObjectAclInput, ...request.Option) (*s3.GetObjectAclOutput, error)
	GetObjectAclRequest(*s3.GetObjectAclInput) (*request.Request, *s3.GetObjectAclOutput)

	GetObjectTagging(*s3.GetObjectTaggingInput) (*s3.GetObjectTaggingOutput, error)
	GetObjectTaggingWithContext(aws.Context, *s3.GetObjectTaggingInput, ...request.Option) (*s3.GetObjectTaggingOutput, error)
	GetObjectTaggingRequest(*s3.GetObjectTaggingInput) (*request.Request, *s3.GetObjectTaggingOutput)

	HeadBucket(*s3.HeadBucketInput) (*s3.HeadBucketOutput, error)
	HeadBucketWithContext(aws.Context, *s3.HeadBucketInput, ...request.Option) (*s3.HeadBucketOutput, error)
	HeadBucketRequest(*s3.HeadBucketInput) (*request.Request, *s3.HeadBucketOutput)

	HeadObject(*s3.HeadObjectInput) (*s3.HeadObjectOutput, error)
	HeadObjectWithContext(aws.Context, *s3.HeadObjectInput, ...request.Option) (*s3.HeadObjectOutput, error)
	HeadObjectRequest(*s3.HeadObjectInput) (*request.Request, *s3.HeadObjectOutput)

	ListBuckets(*s3.ListBucketsInput) (*s3.ListBucketsOutput, error)
	ListBucketsWithContext(aws.Context, *s3.ListBucketsInput, ...request.Option) (*s3.ListBucketsOutput, error)
	ListBucketsRequest(*s3.ListBucketsInput) (*request.Request, *s3.ListBucketsOutput)

	ListBucketsExtended(*s3.ListBucketsExtendedInput) (*s3.ListBucketsExtendedOutput, error)
	ListBucketsExtendedWithContext(aws.Context, *s3.ListBucketsExtendedInput, ...request.Option) (*s3.ListBucketsExtendedOutput, error)
	ListBucketsExtendedRequest(*s3.ListBucketsExtendedInput) (*request.Request, *s3.ListBucketsExtendedOutput)

	ListBucketsExtendedPages(*s3.ListBucketsExtendedInput, func(*s3.ListBucketsExtendedOutput, bool) bool) error
	ListBucketsExtendedPagesWithContext(aws.Context, *s3.ListBucketsExtendedInput, func(*s3.ListBucketsExtendedOutput, bool) bool, ...request.Option) error

	ListLegalHolds(*s3.ListLegalHoldsInput) (*s3.ListLegalHoldsOutput, error)
	ListLegalHoldsWithContext(aws.Context, *s3.ListLegalHoldsInput, ...request.Option) (*s3.ListLegalHoldsOutput, error)
	ListLegalHoldsRequest(*s3.ListLegalHoldsInput) (*request.Request, *s3.ListLegalHoldsOutput)

	ListMultipartUploads(*s3.ListMultipartUploadsInput) (*s3.ListMultipartUploadsOutput, error)
	ListMultipartUploadsWithContext(aws.Context, *s3.ListMultipartUploadsInput, ...request.Option) (*s3.ListMultipartUploadsOutput, error)
	ListMultipartUploadsRequest(*s3.ListMultipartUploadsInput) (*request.Request, *s3.ListMultipartUploadsOutput)

	ListMultipartUploadsPages(*s3.ListMultipartUploadsInput, func(*s3.ListMultipartUploadsOutput, bool) bool) error
	ListMultipartUploadsPagesWithContext(aws.Context, *s3.ListMultipartUploadsInput, func(*s3.ListMultipartUploadsOutput, bool) bool, ...request.Option) error

	ListObjects(*s3.ListObjectsInput) (*s3.ListObjectsOutput, error)
	ListObjectsWithContext(aws.Context, *s3.ListObjectsInput, ...request.Option) (*s3.ListObjectsOutput, error)
	ListObjectsRequest(*s3.ListObjectsInput) (*request.Request, *s3.ListObjectsOutput)

	ListObjectsPages(*s3.ListObjectsInput, func(*s3.ListObjectsOutput, bool) bool) error
	ListObjectsPagesWithContext(aws.Context, *s3.ListObjectsInput, func(*s3.ListObjectsOutput, bool) bool, ...request.Option) error

	ListObjectsV2(*s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error)
	ListObjectsV2WithContext(aws.Context, *s3.ListObjectsV2Input, ...request.Option) (*s3.ListObjectsV2Output, error)
	ListObjectsV2Request(*s3.ListObjectsV2Input) (*request.Request, *s3.ListObjectsV2Output)

	ListObjectsV2Pages(*s3.ListObjectsV2Input, func(*s3.ListObjectsV2Output, bool) bool) error
	ListObjectsV2PagesWithContext(aws.Context, *s3.ListObjectsV2Input, func(*s3.ListObjectsV2Output, bool) bool, ...request.Option) error

	ListParts(*s3.ListPartsInput) (*s3.ListPartsOutput, error)
	ListPartsWithContext(aws.Context, *s3.ListPartsInput, ...request.Option) (*s3.ListPartsOutput, error)
	ListPartsRequest(*s3.ListPartsInput) (*request.Request, *s3.ListPartsOutput)

	ListPartsPages(*s3.ListPartsInput, func(*s3.ListPartsOutput, bool) bool) error
	ListPartsPagesWithContext(aws.Context, *s3.ListPartsInput, func(*s3.ListPartsOutput, bool) bool, ...request.Option) error

	PutBucketAcl(*s3.PutBucketAclInput) (*s3.PutBucketAclOutput, error)
	PutBucketAclWithContext(aws.Context, *s3.PutBucketAclInput, ...request.Option) (*s3.PutBucketAclOutput, error)
	PutBucketAclRequest(*s3.PutBucketAclInput) (*request.Request, *s3.PutBucketAclOutput)

	PutBucketCors(*s3.PutBucketCorsInput) (*s3.PutBucketCorsOutput, error)
	PutBucketCorsWithContext(aws.Context, *s3.PutBucketCorsInput, ...request.Option) (*s3.PutBucketCorsOutput, error)
	PutBucketCorsRequest(*s3.PutBucketCorsInput) (*request.Request, *s3.PutBucketCorsOutput)

	PutBucketLifecycleConfiguration(*s3.PutBucketLifecycleConfigurationInput) (*s3.PutBucketLifecycleConfigurationOutput, error)
	PutBucketLifecycleConfigurationWithContext(aws.Context, *s3.PutBucketLifecycleConfigurationInput, ...request.Option) (*s3.PutBucketLifecycleConfigurationOutput, error)
	PutBucketLifecycleConfigurationRequest(*s3.PutBucketLifecycleConfigurationInput) (*request.Request, *s3.PutBucketLifecycleConfigurationOutput)

	PutBucketLogging(*s3.PutBucketLoggingInput) (*s3.PutBucketLoggingOutput, error)
	PutBucketLoggingWithContext(aws.Context, *s3.PutBucketLoggingInput, ...request.Option) (*s3.PutBucketLoggingOutput, error)
	PutBucketLoggingRequest(*s3.PutBucketLoggingInput) (*request.Request, *s3.PutBucketLoggingOutput)

	PutBucketProtectionConfiguration(*s3.PutBucketProtectionConfigurationInput) (*s3.PutBucketProtectionConfigurationOutput, error)
	PutBucketProtectionConfigurationWithContext(aws.Context, *s3.PutBucketProtectionConfigurationInput, ...request.Option) (*s3.PutBucketProtectionConfigurationOutput, error)
	PutBucketProtectionConfigurationRequest(*s3.PutBucketProtectionConfigurationInput) (*request.Request, *s3.PutBucketProtectionConfigurationOutput)

	PutBucketWebsite(*s3.PutBucketWebsiteInput) (*s3.PutBucketWebsiteOutput, error)
	PutBucketWebsiteWithContext(aws.Context, *s3.PutBucketWebsiteInput, ...request.Option) (*s3.PutBucketWebsiteOutput, error)
	PutBucketWebsiteRequest(*s3.PutBucketWebsiteInput) (*request.Request, *s3.PutBucketWebsiteOutput)

	PutObject(*s3.PutObjectInput) (*s3.PutObjectOutput, error)
	PutObjectWithContext(aws.Context, *s3.PutObjectInput, ...request.Option) (*s3.PutObjectOutput, error)
	PutObjectRequest(*s3.PutObjectInput) (*request.Request, *s3.PutObjectOutput)

	PutObjectAcl(*s3.PutObjectAclInput) (*s3.PutObjectAclOutput, error)
	PutObjectAclWithContext(aws.Context, *s3.PutObjectAclInput, ...request.Option) (*s3.PutObjectAclOutput, error)
	PutObjectAclRequest(*s3.PutObjectAclInput) (*request.Request, *s3.PutObjectAclOutput)

	PutObjectTagging(*s3.PutObjectTaggingInput) (*s3.PutObjectTaggingOutput, error)
	PutObjectTaggingWithContext(aws.Context, *s3.PutObjectTaggingInput, ...request.Option) (*s3.PutObjectTaggingOutput, error)
	PutObjectTaggingRequest(*s3.PutObjectTaggingInput) (*request.Request, *s3.PutObjectTaggingOutput)

	RestoreObject(*s3.RestoreObjectInput) (*s3.RestoreObjectOutput, error)
	RestoreObjectWithContext(aws.Context, *s3.RestoreObjectInput, ...request.Option) (*s3.RestoreObjectOutput, error)
	RestoreObjectRequest(*s3.RestoreObjectInput) (*request.Request, *s3.RestoreObjectOutput)

	UploadPart(*s3.UploadPartInput) (*s3.UploadPartOutput, error)
	UploadPartWithContext(aws.Context, *s3.UploadPartInput, ...request.Option) (*s3.UploadPartOutput, error)
	UploadPartRequest(*s3.UploadPartInput) (*request.Request, *s3.UploadPartOutput)

	UploadPartCopy(*s3.UploadPartCopyInput) (*s3.UploadPartCopyOutput, error)
	UploadPartCopyWithContext(aws.Context, *s3.UploadPartCopyInput, ...request.Option) (*s3.UploadPartCopyOutput, error)
	UploadPartCopyRequest(*s3.UploadPartCopyInput) (*request.Request, *s3.UploadPartCopyOutput)

	WaitUntilBucketExists(*s3.HeadBucketInput) error
	WaitUntilBucketExistsWithContext(aws.Context, *s3.HeadBucketInput, ...request.WaiterOption) error

	WaitUntilBucketNotExists(*s3.HeadBucketInput) error
	WaitUntilBucketNotExistsWithContext(aws.Context, *s3.HeadBucketInput, ...request.WaiterOption) error

	WaitUntilObjectExists(*s3.HeadObjectInput) error
	WaitUntilObjectExistsWithContext(aws.Context, *s3.HeadObjectInput, ...request.WaiterOption) error

	WaitUntilObjectNotExists(*s3.HeadObjectInput) error
	WaitUntilObjectNotExistsWithContext(aws.Context, *s3.HeadObjectInput, ...request.WaiterOption) error
}

var _ S3API = (*s3.S3)(nil)
