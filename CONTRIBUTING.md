
# Contribute Code

You are welcome to contribute to project BFE. To contribute to BFE, you have to agree with the 
[Contributor License Agreement](https://cla-assistant.io/bfenetworks/conf-agent).

We sincerely appreciate your contribution. This document explains our workflow and work style.

## Workflow

BFE uses this [Git branching model](http://nvie.com/posts/a-successful-git-branching-model/). The following steps guide usual contributions.

1. Fork

   Our development community has been growing fastly; it doesn't make sense for everyone to write into the official repo.  So, please file Pull Requests from your fork.  To make a fork, just head over to the GitHub page and click the ["Fork" button](https://help.github.com/articles/fork-a-repo/).

1. Clone

   To make a copy of your fork to your local computers, please run

   ```bash
   git clone https://github.com/your-github-account/conf-agent
   cd bfe
   ```

1. Create the local feature branch

   For daily works like adding a new feature or fixing a bug, please open your feature branch before coding:

   ```bash
   git checkout -b my-cool-stuff
   ```


1. Build and test

   Users can build Conf Agent natively on Linux. 

   ```bash
   make
   ```

1. Keep pulling

   An experienced Git user pulls from the official repo often -- daily or even hourly, so they notice conflicts with others work early, and it's easier to resolve smaller conflicts.

   ```bash