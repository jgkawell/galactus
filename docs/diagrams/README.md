# Diagrams

This directory holds subdirectories for each diagram group. Each diagram group should be a logical grouping of diagrams by their topic. These are things like system domains, components, or subsystems. The diagrams should be written in [PlantUML](https://plantuml.com/) (`.puml`) and should be rendered as PNG images within the `out/` subdirectory relative to the `.puml` files. This can be done with the `make diagrams` command in the top-level `Makefile`.

**PLEASE NOTE**: No diagrams files (`.puml`) should ever place placed in this directory. This will break the recursive image generation. Instead, all diagrams should be placed in the subdirectories of this directory.

## How to

Whenever you make edits or create new diagrams, you should run `make diagrams` from the root of `galactus` to generate their corresponding `.png` images so that PR reviewers can easily see the changes.

In order to run this command, you must have Docker installed and running.
