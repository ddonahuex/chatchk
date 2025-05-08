# chatchk
Chatchk provides example source code demonstrating how to interact with
Open WebUI and Ollama API endpoints within a Nethopper Private AI
instance.

By default, it sends a "Why is the sky blue?" prompt to both Open WebUI
and Ollama APIs, using the gemma2:9b model, which is pre-installed in a
Nethopper Private AI instance.

This project is written in Go (version 1.24).

## License
This project is licensed under the MIT License. See the LICENSE file for details.

Copyright © 2025 [Nethopper, Inc.](nethopper.io). All rights reserved.

## Modules
Each package listed below resides in ./src/\<package name\>. They are listed in
order of relevance.

| Package | Description |
| --------- | ------- |
| chatchk | Main module. Entry point for the chatchk program |
| open_webui | Functions implementing Open WebUI API endpoints |
| ollama | Functions implementing Ollama API endpoints |
| utils | Support functions for chatchk packages |
| ingest | Implements file transfer and ingestion methods in Open WebUI |
| knowledge | Manages Workspace Knowledge bases in Open WebUI |
| prompts | Manages Workspace Prompts in Open WebUI |
| admin | Manages administrator functions in Open WebUI |

# Building the Project
The chatchk project is built from the software/chatchk directory.

This directory includes a Makefile that supports building a Go executable and
creating a Docker image. The Makefile provides targets for building, running,
and pushing the Docker image to Nethopper's namespace on Docker Hub.

For a complete list of build targets and their usage, refer to the Make Targets
section below.

## make
The table below lists and describes the build targets.

To list build targets:
> $ make help

To build a target:
> $ make \<target\>

### Build Targets
| Target | Description |
| --------- | ------- |
| build | Executes a Go build for the chatchk executable |
| clean | Executes a Go clean for all modules |
| docker-build | Docker build for nethopper/chatchk Docker image |
| docker-prod | Executes docker-build and docker-push make targets |
| docker-push | Docker push for of nethopper/chatchk to the Nethopper Docker Hub namespace |
| docker-run | Executes docker-build then issues a docker run of the nethopper/chatchk image |
| help | Print this help menu |
| prod | Executes build and docker-prod make targets |
| test | Executes a Go test for all modules (none currently) |
 
# Run
Chatchk can be deployed on bare metal, in a Docker container, or within a
Kubernetes cluster. Detailed instructions for each deployment method are
provided in the subsections below.



