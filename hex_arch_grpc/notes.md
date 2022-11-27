# How to struct your Go App - Hexagonal Architecture
How to apply hex architecture (ports & adapters) to a `grpc api` app in Go - video [link](https://www.youtube.com/watch?v=MpFog2kZsHk).

Other hex-arch resources:
- DDD, Hexagonal, Onion, Clean, CQRS, How I put it all together [blog post](https://herbertograca.com/2017/11/16/explicit-architecture-01-ddd-hexagonal-onion-clean-cqrs-how-i-put-it-all-together/)
- Hexagon - a Kotlin microservices toolkit [page](https://hexagonkt.com/)


# What is hexagonal architecture
[source](https://www.youtube.com/watch?v=MpFog2kZsHk)  
<img src="./resources/hex_arch_0.png" width="500"/>

[source](https://hexagonkt.com/)  
<img src="./resources/hex_arch_1.png" width="500"/>

## Layers
- `Domain` - core of application, contains domain/business logic
- `Application` - orchestrates the use of our domain code and adapts requests from the framework layer to the domain layer by sitting between the two
- `Framework` - provides the logic for outside components such as db or grpc adapters


**Dependency flow**
- The outside layers depends on the inside layers!  
- The Domain layer does not depend on the Application layer on anything, and the Application layer cannot depend on the Framework layer.

**Dependency injection**  
To communicate with the database from the Application layer we need to use dependency injection, ie:  
instead of calling framework to create an instance of db from the application layer, we invert the control of db instantiation to the program at startup, and have the program inject db instance into the Application layer.

## Project structure
<!-- TODO: add tree from terminal after all is coded -->
<img src="./resources/tree.png" alt="drawing" width="300"/>

- `adapters` directory will contain the code for all our layers, each layer has its own subdirectory
    - `adapters/left` driving adapters
    - `adapters/right` driven adapters
- `ports` contain all ports organized by layer
- `main.go`
    - orchestrates the startup of the application,
    - contains code that connect all of the ports and adapters
    - contains code that injects dependencies into the layers that need them


# Stage 1 - Core layer
<img src="./resources/tree_s1.png" alt="drawing" width="250"/>

We defined:
- **arithmetic port**: `ArithmeticPort` in the core layer  
an interface for structs that implement basic arithmetic operations
- **arithmetic adapter**: `internal/adapters/core/arithmetic/Adapter` in the core layer  
that implements the interface

which completes core layer of our application.

