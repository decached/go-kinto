package kinto

const (
	INFO_URI        = "/"
	BATCH_URI       = "/batch"
	HEARTBEAT_URI   = "/__heartbeat__"
	OPENAPI_URI     = "/__api__"
	BUCKETS_URI     = "/buckets"
	BUCKET_URI      = "/buckets/%s"
	GROUPS_URI      = "/buckets/%s/groups"
	GROUP_URI       = "/buckets/%s/groups/%s"
	COLLECTIONS_URI = "/buckets/%s/collections"
	COLLECTION_URI  = "/buckets/%s/collections/%s"
	RECORDS_URI     = "/buckets/%s/collections/%s/records"
	RECORD_URI      = "/buckets/%s/collections/%s/records/%s"
)
