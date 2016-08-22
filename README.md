# cmm (Crimson MindMap) v0.1

## Specification

CMM files are archives (ZIP format) which store information about a node graph, with nodes and connections stored as TOML files, along with any attachments.

### Dependency version

* ZIP32
* TOML 0.4

### Archive folder structure

The CMM archive has the following  folder structure:

```
ROOT
 |
 +-- nodes
 |     |
 |     +-- 1.toml
 |     |
 |     \-- 2.toml
 |
 +-- connections
 |     |
 |     +-- 1.toml
 |     |
 |     \-- 2.toml
 |
  +-- attachments
 |     |
 |     +-- uniquefilename1.png
 |     |
 |     +-- uniquefilename2.png
 |     |
 |     \-- uniquefilename1.gif
 |
 \-- index.toml
```

### Node TOML format

Any non-reserved field can be used. The supported types are text, numbers, boolean and array of these. Nesting is currently not supported. The following fields are reserved for meta information:

```
_name        = "Example Node"                # The name of the node which can be used to represent the node in a graph view
_description = "A longer description"        # A helpful description
_connections = [4, 6, 11]                    # The outgoing connections
_attachments = ["firefox.png", "otters.gif"] # Attached files
```

### Connection TOML format

Any non-reserved field can be used. The supported types are text, numbers, boolean and array of these. Nesting is currently not supported. The following fields are reserved for meta information:

```
_name               = "Example Node"                # The name of the connnection which can be used to label the connection in a graph view
_description        = "A longer description"        # A helpful description
_source_node        = 11                            # The source node
_destination_node   = 12                            # The destination node
_attachments        = ["firefox.png", "otters.gif"] # Attached files
```