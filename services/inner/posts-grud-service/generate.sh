#!/bin/bash

protoc postpb/post.proto --go_out=plugins=grpc:.
protoc -I . post.proto --go_out=plugins=grpc:.