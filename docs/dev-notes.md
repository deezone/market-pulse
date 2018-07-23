# forex-clock Development Notes

- [Git Basics - Tagging](https://git-scm.com/book/en/v2/Git-Basics-Tagging)

**List Tags**
```
$ git tag
v0.1
v1.3
```

**List Commit Checksums**
```
$ git log --pretty=oneline
59505f98b4f850e0d2223e684529182317df3dbe (HEAD -> master, origin/master) Update documentation to include test instructions
e158d0f69e3bbcdb4ef30d5a7f239b4c28a61aff Adds test coverage to server modules
71f1d6ca5c1741b0830d51fa7671287fa49f8984 (tag: v0.1.0) adds middleware items
6a1fcc9c81d9f120db430ac37ac730ce966ff64e first pass at adding middleware to server
```

**Tag Release**
```
git tag -a v0.1.0 -m "Version 0.1.0 - First functional version with basic status endpoints" 71f1d6ca5c1741b0830d51fa7671287fa49f8984
```
