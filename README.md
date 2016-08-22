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
