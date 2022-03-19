#!/bin/bash

protoc api_post_operations_service.proto --go_out=plugins=grpc:.