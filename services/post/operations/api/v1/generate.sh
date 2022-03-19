#!/bin/bash

protoc post_operations.proto --go_out=plugins=grpc:.