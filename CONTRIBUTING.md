# Contributing Guidelines

The following is a set of guidelines for contributing to the project. These are 
mostly guidelines, not rules. Feel free to propose changes to this document in a
pull request. 

## General

I welcome all contributions to this project, no matter how small. But please
keep a few things in mind when you do contribute:

* Please leave the place nicer than you found it. If something can be made
  simpler or more readable, please make the change.
* While all contributions are welcome, not all will be accepted. This may be for
  any number of reasons, including work that is not aligned with the goal of the
  project, work that introduces unnecessary complexity, and a number of other 
  reasons
* If you are providing feedback on another's contribution, please remember that
  there is a human on the other side of your computer and strive to be:
    * Kind
    * Respectful
    * Constructive

## Branching Strategy

I use a single branch branching strategy where all branches
are made from `master` and and merge back into `master`. New feature should be
developed on a branch with the prefix `feature`, and bugfixes should be developed
on a branch prefixed with `bugfix`. If you've forked the repo and want to make
a PR, it doesn't matter what you've called your branch, it'll be judged on its
merits.

## Pull Requests

Please confirm that:

1. There isn't already a PR to fix the issue
2. If you are fixing a bug, that there is a issue related to it (feel free to 
raise one if it doesn't exist)
3. You provide some context around your change, why it was made and how to test
it
   
## Styleguides

### Coding Styleguides

There isn't one. This is my first Go project, so I'm not sure what the consensus
standards are.

### Git Commit Messages

I like [Conventional Commits][conv-commits] to ensure consistency of commit 
messages, although I don't always abide by them.

A compliant commit looks like this:

```
<type>(optional scope): <description>

<optional body>
```

* The first line of the commit becomes the subject line
* Don't capitalize the first letter of the description
* Use the imperative mood ("add feature" not "adds feature")
* Use present tense ("add feature" not "added feature")
* No period (.) at the end of the subject line
* Limit the subject line to 72 characters or less.


[conv-commits]: https://www.conventionalcommits.org/en/v1.0.0-beta.3/