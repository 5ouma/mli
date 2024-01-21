# Contributing Guide<!-- omit in toc -->

> [!NOTE]
> You must follow the [Code of Conduct](./CODE_OF_CONDUCT.md).

I happily welcome your contributions!
Before you contribute,
I would recommend reading this guideline for a better development experience.

<br />

- [🔠 Branch Prefix](#-branch-prefix)
- [🧪 Testing Advice](#-testing-advice)

<br />

## 🔠 Branch Prefix

You should follow these branch name rules
because [Pull Request Labeler] automatically adds labels to your Pull Requests.
<br />
For the details of each label, please see [Labels](https://github.com/5ouma/mli/labels).

[Pull Request Labeler]: https://github.com/actions/labeler

|        Label        |         Branch Prefix RegEx          |
| :-----------------: | :----------------------------------: |
|    Type: Feature    |            `^feat(ure)?-`            |
|      Type: Bug      |               `^fix-`                |
|   Type: Security    |           `^sec(urity)?-`            |
| Type: Documentation |          `^doc(ument)?s?-`           |
|  Type: Refactoring  |          `^refactor(ing)?-`          |
|    Type: Testing    |          `^test(ing\|s)?-`           |
|  Type: Maintenance  | `^maintenance-` , `^maintain(ing)?-` |
|      Type: CI       |                `^ci-`                |
| Type: Dependencies  |   `^dep(endency\|endencies\|s)?-`    |
|     Type: Meta      |               `^meta-`               |

> Labels were generated with [@azu / github-label-setup](https://github.com/azu/github-label-setup)

## 🧪 Testing Advice

There is a test function called [`Test_OAScript`](../lib/oascript_test.go)
for AppleScript commands.<br />
When you run this test, your Login Items all will be removed.
Therefore, I highly recommend backing up your current Login Items by `mli save`.
