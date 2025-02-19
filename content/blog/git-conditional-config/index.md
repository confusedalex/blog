+++
title= "TIL: git conditional config with and without home-manager"
date = 2025-02-19
[taxonomies]
tags = ["til", "git", "nix"]
+++

Today I learnt how to use git conditional config to use multiple identities easily.  
I use git for my private project and the projects I need to work at my day job.  
To keep these two projects apart, I have a `~/projects` folder in my home directory with the two subfolder `~/projects/work` and `~/projects/personal`.
Now we can tell git to always use the one identity for all projects under `~/projects/personal` and another for `~/projects/work`.
We just add two includeIfs in our git config. With includeIf we can specify other files to source if the condition is met. The prepended "gitdir:" just means every git repository under this directory. [^1]

```
[includeIf "gitdir:~/projects/personal/"]
	path = "./personal"

[includeIf "gitdir:~/projects/work/"]
	path = "./work"
```

In the personal config, we specify our personal identity

```
[user]
	email = "alex@confusedalex.dev"
	name = "confusedalex"
	signingKey = "05AF71643F6E2ED3"
```

and do the same for the work config.

## Integrating with home-manager

Because I use [home-manager](https://nix-community.github.io/home-manager/index.xhtml) I want to specify these conditionals with nix. We just need to add a list we attributes sets for every include.

> Everything under contents must follow the structure as in git-config.

```nix
{
  programs.git = {
    enable = true;
    # ...
    includes = [
      {
        condition = "gitdir:~/projects/personal/";
        contents = {
          user = {
            name = "confusedalex";
            email = "alex@confusedalex.dev";
            signingKey = "05AF71643F6E2ED3";
          };
          gpg.format = "openpgp";
        };
      }
      {
        condition = "gitdir:~/projects/work/";
        contents = {
            #...
        };
      }
    ];
  };
}
```

[^1]: [ https://git-scm.com/docs/git-config#\_includes ](https://git-scm.com/docs/git-config#_includes)
