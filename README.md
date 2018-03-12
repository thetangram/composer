[![Build Status](https://travis-ci.org/thetangram/composer.svg)](https://travis-ci.org/thetangram/composer)
[![GoDoc](https://godoc.org/github.com/thetangram/composer?status.svg)](https://godoc.org/github.com/thetangram/composer)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/thetangram/composer)](https://goreportcard.com/report/thetangram/composer)

The Tangram Composer 
====================

The Tangram Composer is a server/edge side HTML composition service. Its main goal is solving the microservices frontend issue faced when you build a distributed microservices system, each microservice with it's own lifecycle and rollout. In this scenario it's usual to build the user interface as a monolith (aka. BFF), aggregator of microservices data. This model creates user interface-microservices one to one dependencies, so in a lot of cases the rollout of each piece is bound, then the service autonomy is harder and going to continous deployment can be compromissed.

### Features

  - Compose HTML based on HTML components/microservices
  - 100% stateless service.
  - Can be deplloyed as stand alone service or in a serverless architecture.
  - Request headers and cookies passthrough to components.
  - Request headers and cookies filtering to components.
  - Component response headers merging.
  - Component meta information, scripts and stylesheets merging. 
  - Concurrent composition.
  - Stand alone and container based artifacts.
  - ... TBD

### Documentation

You can find all the project documentation in [the documentation folder](./docs). We recommend you to start reading the [project description](./docs/description.md) and the [glossary](./docs/glossary.md). 

### About the name

[From Wikipedia](https://en.wikipedia.org/wiki/Tangram) *The tangram (Chinese: 七巧板; pinyin: qīqiǎobǎn; literally: "seven boards of skill") is a dissection puzzle consisting of seven flat shapes, called tans, which are put together to form shapes. The objective of the puzzle is to form a specific shape (given only an outline or silhouette) using all seven pieces, which may not overlap.*


Getting Started
---------------

### Requisites

In order to build The Tangram Composer you need [Go](https://golang.org), [Make](https://www.gnu.org/software/make/), [Dep](https://github.com/golang/dep), [Docker](https://www.docker.com/) and [zip](). The complete list of tools:

  - **[Go](https://golang.org) 1.9+**, as programming language.
  - **[Make](https://www.gnu.org/software/make/)**, as build automation tool.
  - **[Dep](https://github.com/golang/dep)**, as dependency management tool.
  - **[Docker](https://www.docker.com/) 17.09+**, to build container. 
  - **Zip**, to package *AWS Lambda deployment*.


### Project structure

This project follows the [Standard Go Project Layout](https://github.com/golang-standards/project-layout).


### Building 

This project uses `make` as build automation tool. The `Makefile` **rules** are:

  - **`clean`**, to remove all binaries from `dist/` directory.
  - **`compile`** (default rule), to complile the code.
  - **`run`**, to compile and run the standalone version.
  - **`test`**, to run the test and code coverage.
  - **`build`**, to compile and genrate the final binary artifacts. THis artifacts are full-independent binary files, and also ave some optimizations.
  - **`package`**, to generate the **docker image** (installed in the local repository) and the **AWS Lamdba** package.


Other projects like this
------------------------

This is not the one and only composition solution. We have a lot of friends out there, giving us a lot of inspiration. Take a look on them.

  - [Compoxure](https://github.com/tes/compoxure) is a composition middleware that acts as a proxy replacement for ESI os SSI for backend services and compose fragments from microservices into the response
  - [Skipper](https://github.com/zalando/skipper) is the Zalando's HTTP router and reverse proxy for service composition.
  - [Convergent UI](https://github.com/acesinc/convergent-ui) is a special Zuul Filter that aims to provide a solution to the Distributed Composition problem faced when building a GUI within a Micro Services Architecture. 
