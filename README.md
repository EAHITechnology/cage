# cage
## Cage the raptor
Simple hierarchical architecture file directory example

```
.
├── README.md
├── conf                        # config
│   ├── config.toml                 # other config
│   └── server_config.toml          # metadata config
├── dao                         
│   ├── demo_task.go
│   ├── lock.go
│   └── model.go
├── go.mod
├── go.sum
├── handler                     # routing layer
│   └── demo_handler.go
├── logic                       # logic layer
│   └── manager                      # common logic layer (third-party interface, cache components, dao persistence strategy)
│       └── xxx.go
│   ├── backend_logic.go           
│   └── demo_logic.go
├── main.go
├── proto                       # protocol file
│   └── http_proto.go
├── server                     
│   ├── middleware.go
│   └── server.go
├── sql                         # sql file
│   └── demo_task.sql
└── utils                       # kits
    ├── config_model.go
    └── http_client.go

```

Compared with the traditional three-tier architecture, we have added the Manager general processing layer.
This lyer has two functions:
- Some general capabilities of the original logic layer can be lowered to this layer, such as interaction strategies with cache and storage, access to middleware;
- You can also encapsulate calls to third-party interfaces at this layer, such as calling payment services, calling audit services and other RPC interfaces.

The advantage of the layered architecture is that it is simple enough to meet the processing needs of the business application layer we are currently dealing with. Secondly, a clear layered structure can quickly locate problems and facilitate subsequent reading by students, and maintainability will be improved.