# third_party/pipelines

This directory services as a place for pipelines and related files (e.g. `VERSION`) for git submodules. The reasoning for this is that we don't want to have to add pipelines or other files to submodules as they're often public forks of third party repositories. Instead we place pipeline files here so that we can track their changes without altering the upstream forks.