### Basic usage

```
you[at]ci-runner$ operator release gitlab workloads/develop -a git.5kbps.io -t [your_pat_token]
## Changelog(time="2019-09-23 17:37", ancestor="1.0.0")

* Changed:
  * 53afd60f #changed refactor cmd release structure

* Added:
  * c4034d56  #added cmd release github
  * 72c57fd4 #added github repository support

* Fixed:
  * 5982680a #fixed gitlab 1st release failed
  * af2237ad  #fixed does not need compare diff for first release

* Missing Defination:
  * 1931d76a fix first release

you[at]ci-runner$ operator release gitlab workloads/develop -a https://git.5kbps.io -t [your_pat_token] --submit
A release was submmited!!!
```
