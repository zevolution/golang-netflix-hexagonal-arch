<p align="center">
<img src="https://user-images.githubusercontent.com/36534847/234449196-ff9cf890-bb63-40c8-9337-7a453829e52f.png" alt="FlixHexaArch-Logo" width="700">
<h1 align="center">Golang Netflix Hexagonal Arch</h1>

## Table of contents
- [About the project](#about-the-project)
- [Description](#description)
- [Built with](#built-with)
- [Installation](#installation)
- [Requirements to run](#requirements-to-run)
- [Run](#run)
- [License](#license)

## About the project
This project was created as a simple example to show how we can create a golang application. This project use Hexagonal Architecture as described in [zevolution/netflix-hexagonal-architecture](https://github.com/zevolution/netflix-hexagonal-architecture) and the same business rule applied in [zevolution/quarkus-job-java17-hexagonal-arch](https://github.com/zevolution/quarkus-job-java17-hexagonal-arch)

## Description
This project has as a business rule to get all software from one author in the main cloud DVCS such as Github, Gitlab, Bitbucket and persist this softwares in a mongodb grouped by author
<br>

### Describe folder structure
```markdown
📦 app-folder
┣ 📂 adapters: Adapters responsible for driving or being driven from external events or business-logic
┃ ┣ 📂 datasources: Also know as driven adapter, secondary adapter or outbound adapter. Responsible for dealing with our application's data sources such as databases, message brokers(producers), cache, http providers and etc. Normally driven by business-rule
┃ ┃ ┣ 📂 services: Services used by datasource
┃ ┃ ┃ ┣ 📂 data: DTO's used for any data source other than a database, such as http apis
┃ ┃ ┃ ┃ ┣ 📂 request:
┃ ┃ ┃ ┃ ┣ 📂 response:
┃ ┃ ┃ ┣ 📂 mappers: Responsible for mapping the DTO's, to business-logic(internal) entities
┃ ┃ ┃ ┣ 📂 models: Responsible for mapping entities that deal directly with any database
┃ ┣ 📂 transportlayers: Also known as a driving adapter, primary adapter or input adapter. Responsible for being the gateway to the external world of our application, whether via rest, graphql, message-broker(consumers), socket and application events as startup. Normally driving one or more business rule
┃ ┃ ┣ 📂 mappers: Responsible for mapping the DTO's, to business-logic(internal) entities
┃ ┃ ┣ 📂 http(example): Example of where we can add our http controllers/handlers such as REST and GraphQL
┃ ┃ ┣ 📂 consumers(example): Example of where we can add our Listener and Consumers from message brokes such as RabbitMQ, Kafka, SQS, Redis and etc
┃ ┃ ┣ 📂 cli(example): Example where we can have our entry point for CLI
┣ 📂 bootstrap: Application startup/configuration classes
┃ ┣ 📂 config: Application config such as database, migration, env variables
┃ ┣ 📜 entry.go: Main file for application startup
┣ 📂 internal: Contains everything related to Application Business-Logic
┃ ┣ 📂 entities: Contains entities that are normally used to represent our business domains/model
┃ ┣ 📂 interactors: Contains UseCases(or services) which is a list of actions/events that typically define the interactions required for the system to achieve a goal. Normally triggered by a driving adapter such as rest/graphql, and responsible for trigger a driven adapter, such as database or message-broker(as producer)
┃ ┣ 📂 repositories: Also know as Ports from Ports&Adapters is interfaces responsible for specifying the behavior of a datasource(driven/secondary/outbound adapter), such as which input and return parameters must follow. "How" datasource will do this, it is own responsibility, since it converts the result in the imposed specification
┣ 📜 .gitignore
┣ 📜 Dockerfile
┣ 📜 docker-compose.yml
┣ 📜 go.mod
┣ 📜 go.sum
┗ 📜 README.MD
```

## Built with
* [Golang](https://go.dev)
* [MongoDB](https://www.mongodb.com)

## Installation

To clone and run this application, you'll need Git installed on your computer(or no, if you want to download **.zip**). From your command line:
```bash
# Git CLI
git clone https://github.com/zevolution/golang-netflix-hexagonal-arch.git

# Github CLI
gh repo clone zevolution/golang-netflix-hexagonal-arch
```

## Requirements to run
* If you use Windows OS, is strongly recommended that you use WSL2 to perform all operations.
* [Docker Engine](https://www.docker.com/get-started)
* [Docker Compose](https://docs.docker.com/compose/install/)

## Run
1. Just open the application folder in a terminal and run `docker-compose up`

## License
[MIT](https://choosealicense.com/licenses/mit/)