# GitHub: SSH Authentication

The steps below are meant for more experienced GitHub users. They are necessary
when you want to use SSH instead HTTPS for GitHub authentication.

1. Register your public SSH key (typically `~/.ssh/id_rsa.pub`) at
   [GitHub](https://github.com/settings/ssh).

2. Run the following command: `git config --global
   url."git@github.com:".insteadOf https://github.com/`. This will make Git
   rewrite GitHub URLs to use SSH instead of HTTPS. This "hack" is necessary
   since it is not possible to specify the `go get` tool to use SSH
   authentication.

3. Run this `git remote add labs git@github.com:dat320-2019/username-labs.git`
   to set up the `labs` remote pointer to use ssh authentication. 
   (`username` should be replaced with your own GitHub username).
