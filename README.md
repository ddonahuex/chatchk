# chatchk
Chatchk provides example source code demonstrating how to interact with
Open WebUI and Ollama API endpoints within a Private AI instance.

By default, it sends a "Why is the sky blue?" prompt to both Open WebUI
and Ollama APIs, using the gemma2:9b model, which is pre-installed in a
Private AI instance.

This project is written in Go (version 1.24).

## License
This project is licensed under the MIT License. See the LICENSE file for details.

Copyright © 2025. All rights reserved.

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
and pushing the Docker image to ddonahuex's namespace on Docker Hub.

In addition to building the Docker image, the build process generates an SBOM
and Vulnerability Report using syft and grype respectively.

For a complete list of build targets and their usage, refer to the Make Targets
section below.

There are two build options for chatchk: Standard and Chainguard. These build
types are detailed in follwoing sub-sections.

## make
The table below lists and describes the build targets and variables.

To list build targets:
> $ make help

To build a target:
> $ make \<target\>

### Build Targets
| Target | Description |
| --------- | ------- |
| build | Executes a Go build for the chatchk executable |
| clean | Executes a Go clean for all modules |
| docker-build | Docker build, SBOM generation, & Vulnerability report for ddonahuex/chatchk Docker image |
| docker-prod | Executes docker-build and docker-push make targets |
| docker-push | Docker push for of ddonahuex/chatchk to the ddonahuex Docker Hub namespace |
| docker-run | Executes docker-build then issues a docker run of the ddonahuex/chatchk image |
| help | Print this help menu |
| prod | Executes build and docker-prod make targets |
| test | Executes a Go test for all modules (none currently) |

### Build types - Standard and Chainguard
The chatchk docker image can either be built using a standard Dockerfile or 
Chainguard Dockerfile. The make variable that controls the build type is **TYPE**.

Valid values for *TYPE* are *standard* and *chainguard*. When not specified the
build defaults to standard, so no need to TYPE for that build.

The standard image uses **Dockerfile-standard** for the build. It 
intentionally includes Golang version 1.20.5 because that version contains
multiple CVEs, which is evident in grype's output during the build.

The Chainguard image uses **Dockerfile-chainguard** for the build. It uses a 
zero CVE Chainguard Go image, which is also evident by grype's output during
the build.

The two build types are for demo purposes.

#### Chainguard Demo
The straightforward demo contains 2 steps:
1. Standard build, see multiple CVEs
2. Chainguard build, see 0 CVEs

Execute the followming commands to run a Chainguard demo.
> $ make clean && make docker-build
See multiple CVEs

> $ make clean && make TYPE=chainguard docker-build
See 0 CVEs

# Run
Chatchk can be deployed on bare metal, in a Docker container, or within a
Kubernetes cluster. Detailed instructions for each deployment method are
provided in the subsections below.

# Software Bill of Materials (SBOM) & Vulnerability Report
The file chatchk-1.0.0-bom.json is a CycloneDX-formatted SBOM generated using Syft 
for the chatchk project.

The file chatchk:1.0.0-vuln-report.json is a vulnerability report generated
using grype.


