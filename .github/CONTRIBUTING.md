# Contributing Guide<!-- omit in toc -->

> [!NOTE]
> You must follow the [Code of Conduct](./CODE_OF_CONDUCT.md).

I happily welcome your contributions!
Before you contribute,
I would recommend reading this guideline for a better development experience.

<br />

- [🧪 Testing Advice](#-testing-advice)
- [💬 Commit Message](#-commit-message)
- [❓ Pull Requests Title](#-pull-requests-title)
- [🪵 Commit Log](#-commit-log)

<br /><br />

## 🧪 Testing Advice

There is a test function called [`Test_OAScript`](../lib/oascript_test.go)
for AppleScript commands.<br />
When you run this test, your Login Items all will be removed.
Therefore, I highly recommend backing up your current Login Items by `mli save`.

<br />

## 💬 Commit Message

I recommend you to follow [Conventional Commits] with this format.

```commit message
type(scope): Description

Body
```

[Conventional Commits]: https://www.conventionalcommits.org

<br />

## ❓ Pull Requests Title

You don't need to add any prefixes like `feature` or `bug fix`
to the Pull Requests title.
I can recognize what kind of PR it is from labels.
Please give a clear title.

<br />

## 🪵 Commit Log

I do squash merge to the main branch to keep the commit history clean.
When merging your Pull Request, I'll add the Conventional Commits type and scope.
