package kinto

const (
	INFO_URI  = "/"
	BATCH_URI = "/batch"

	HEARTBEAT_URI    = "/__heartbeat__"
	LB_HEARTBEAT_URI = "/__lbheartbeat__"

	OPENAPI_URI = "/__api__"
	VERSION_URI = "/__version__"

	FLUSH_URI = "/__flush__"

	BUCKETS_URI = "/buckets"
	BUCKET_URI  = "/buckets/%s"

	COLLECTIONS_URI = "/buckets/%s/collections"
	COLLECTION_URI  = "/buckets/%s/collections/%s"

	RECORDS_URI = "/buckets/%s/collections/%s/records"
	RECORD_URI  = "/buckets/%s/collections/%s/records/%s"

	GROUPS_URI = "/buckets/%s/groups"
	GROUP_URI  = "/buckets/%s/groups/%s"

	ACCOUNT_URI = "/accounts/%s"
)
