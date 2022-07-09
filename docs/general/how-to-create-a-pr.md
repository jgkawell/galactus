# How to create a PR

## PRs into main

### Descriptions

When merging PRs we always squash the commits to retain a simple git history. This means that the commits to your branch will disappear and be replaced with a single commit consisting of the information in the PR itself.

The result of this is that your PR title, linked work items, comments, and (most importantly) **description** become the full description of the change you have made to the code base.

Be sure to include a thorough overview of **all** changes you are making in your PR within the PR description. The easiest thing is to create a list (remember markdown is supported). For example:

```md
<!-- EXAMPLE OF SIMPLE CHANGE -->

- added full cdn url (with container name) to handler layer
- removed container name from service layer in getBlobName


<!-- EXAMPLE OF COMPLICATED CHANGE -->

When the asset service delete function is called it does two things:

1. it deletes the db entry associated with that asset (specified by assetID, and version)
2. it deletes the blob associated with that db entry

We've discovered an issue with this approach. If you have two db entries pointing to the same blob (which can happen since blob ids (names) are simply set to the hash of the blob itself), then you may delete a blob associated with one db entry but another db entry is now pointing to empty space.

The solution to this is simply to stop deleting blobs when assets are deleted. Yes, this has the potential for extra cost in asset storage, but it is best practice not to delete anyway but instead "soft" delete by marking deleted. This approach will be done fully when we have versioning and updating fully implemented in the asset service.
```

### Versions

When creating PRs into `main` branch **ALWAYS** update the `VERSION` file within the directory where you are making changes.

The version follows the major.minor.patch syntax. This version must be changed on every PR so we don't overwrite artifact versions built off of `main`.
