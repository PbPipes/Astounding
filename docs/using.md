# HOW TO USE

need to create the binary:

from the root, run -->
`go build -o protoc-gen-pipeline`

then run:
`protoc  --plugin=./protoc-gen-pipeline --pipeline_out=. example/example.proto`

This will generate individual tf files for each pubsub topic, in a a pubsub folder.

As README says, this will be extended to include bigquery tables, and needed insert logic (functions, run, or dataflow)
