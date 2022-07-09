# `/test`

This directory is for application testing code. Primarily it is where you can implement Argo tests which will be performed as Experiments during an Argo Rollout in Kubernetes, but other application testing can also be added here.

There are two main subfolders here:

- `/integration`: This is for integration tests to verify services are operating properly in the environment they are deployed too. These tests focus on things like inter-service calls and database connections.
- `/functional`: This is for application testing from the user's perspective. These tests focus on end-to-end testing like what happens when a user clicks "Place Order" from the client.

Feel free to add more types of tests as needed, but keep in mind that unit tests belong beside the code it is testing, not in this directory.
