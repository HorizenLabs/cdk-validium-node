# Generate Protobuf GRPC Source Files
This document outlines how to automatically generate the source files to process through GRPC the Protobuf messages employed for communication among different CDK components. The structure of such messages is specified in the `src` folder.

The generation of the source files can be performed through a Docker container, which is created with the `Dockerfile-Protoc` file.

To generate the source files in the current repository, go to the root folder of the repository and run the following commands:
```bash
docker build -t build-proto -f proto/Dockerfile-Protoc .
docker run -v $(pwd):/usr/src build-proto
```
