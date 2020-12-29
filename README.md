# Astounding

This is the start of a project to create the cloud infrastructure around protocol buffers.  

Specifically, [Terraform](https://www.terraform.io/) is generated via .proto files and extending protoc.  

The first iteration will focus on [GCP](https://cloud.google.com/) and specifically data pipelines within.  

The services initially targeted will be:
* [Pub/Sub](https://cloud.google.com/pubsub)
* [Cloud Run](https://cloud.google.com/run)
* [Cloud Functions](https://cloud.google.com/functions)
* [BigQuery](https://cloud.google.com/bigquery)
* [Dataflow](https://cloud.google.com/dataflow)

For BigQuery, we will also rely on [protoc-gen-bq-schema](https://github.com/GoogleCloudPlatform/protoc-gen-bq-schema).

After initial bits of the easier stuff, extensions will folllow to support loadbalancers and perhaps everything to support the entire end-to-end pipelines.  

More difficult, but ideally longer-term, we'll get to a place where we can annotate the proto with the sort of desired grouping/stateful transformations, and some additional cool things that can be done with real-time data.  

This project extends protoc via [options](https://developers.google.com/protocol-buffers/docs/proto3#options).  

To avoid any future conflicts, we have a reserved an id (# 1090) for the protobuf extension.  We'll use a map to keep things flexible for the many desired options, rather than allocating more and more individual extenion id options.  
  

**Note on version of Terraform Generated:**
For now, we will generate Terraform v0.13.  TBD intent on whether/when to update what gets generated [seems more ideal]  and/or if will facilitate multiple versions based on additional options.  
  

**Note on Naming:**
The name is a reference to [Astounding Stories](https://en.wikipedia.org/wiki/Analog_Science_Fiction_and_Fact).   Specifically, [*Collision Orbit*](https://en.wikipedia.org/wiki/Collision_Orbit), written by [Jack Williamson](https://en.wikipedia.org/wiki/Jack_Williamson) -- *Collision Orbit* didn't seem a great name given the negative connotations with collisions :-) That work is relevant since it is the first claimed use of [Terraforming](https://en.wikipedia.org/wiki/Terraforming) (which is what much of this code creates!).
