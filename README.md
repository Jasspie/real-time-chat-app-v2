#### What else is there besides REST / GraphQL?

+ **WebSockets** - Bidirectional communication over a single TCP connections. Messages are asynchronous and can be sent from client to server and server to client at the same time. This is natively supported by most browsers.
+ **gRPC** - High-performance bidirectional streaming. This not only supports simultaneous client/server and server/client messages like WebSockets, but it also offers multi-plexing, which means multiple messages can be sent in a single direction at the same time without blocking.
+ **MQTT** - Minimalistic publish/subscribe communication protocol. Designed for IoT use cases, MQTT offers multiple levels of delivery fidelity (Quality of Service) and defines what to do when a connection is dropped (Last Will and Testament).
+ **Server-Sent Events (SSE)** - Simple, one-way communication from a server to a client over HTTP. Mostly intended for browser-based interactions, this mechanism offers features like automatic reconnection and native browser support.

#### Use Cases
Now that we know what each of these are, we should talk about when to use them. As with any question in software, the answer of course is it depends, but as a general guideline, let’s take a look at certain situations where each one is the most appropriate.

+ **WebSockets** - Online multiplayer games, chat apps, collaborative tools, etc… Use cases are generally around syncing state between multiple users.
+ **gRPC** - Microservice communication, mobile apps, and high-performance systems. Typically you see use cases with requirements around low-latency and efficient bandwith usage.
+ **MQTT** - IoT devices. This is generally seen in environments where connections are intermittent and clients are devices that send frequent data.
+ **Server-Sent Events (SSE)** - Live news, sports scoring, monitoring dashboards. Use cases are around getting data from a server with no communication from the client.

Of course you can use these communication mechanisms for other use cases, but generally you will find them best suited for the tasks mentioned above.

### Decision Matrix

| API Decision making matrix | Client & server (same team) | client & server (different team) | client and server (different company) |
|---|---|---|---|
| client / server both use typescript | tRPC / GraphQL | GraphQL Federation | GraphQL / REST |
| Client / Server use different languages | GraphQL | GraphQL Federation | GraphQL / REST |
| Client is Backend Service or CLI | gRPC | gRPC | GraphQL / REST |
| Long-Running Operation Events | AsyncAPI (PubSub / stream) | AsyncAPI (PubSub / stream) | Webhooks |

+ Client & Server developed by the same team, using TypeScript only: tRPC or GraphQL
+ Client & Server developed by the same team, using different languages: GraphQL
+ Client is a Backend or CLI: gRPC
+ Long-Running Operations, within the same company: AsyncAPI (PubSub/Stream)
+ Long-Running Operations, across companies: WebHooks
+ Multiple Clients & Servers across teams: GraphQL Federation
+ Public APIs: REST / GraphQL
