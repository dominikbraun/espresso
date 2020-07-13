![Espresso Static Site Generator](https://sternentstehung.de/espresso-ssg-logo.png)
<br /><br />
[![CircleCI](https://circleci.com/gh/dominikbraun/espresso.svg?style=shield)](https://circleci.com/gh/dominikbraun/espresso)
[![Go Report Card](https://goreportcard.com/badge/github.com/dominikbraun/espresso)](https://goreportcard.com/report/github.com/dominikbraun/espresso)
[![GitHub release](https://img.shields.io/github/v/release/dominikbraun/espresso?include_prereleases&sort=semver)](https://github.com/dominikbraun/espresso/releases)
[![License](https://img.shields.io/github/license/dominikbraun/espresso)](https://github.com/dominikbraun/espresso/blob/master/LICENSE)

---

Espresso is a Static Site Generator designed for Markdown-based content with focus on simplicity and performance.

## <img src="https://sternentstehung.de/espresso-ssg-dot.png"> Features

* Flexibility: Create and provide your own Espresso templates
* Simplicity: Build your entire website with a single command
* Zero Configuration: Provide a configuration only if you want to
* Performance: Render thousands of articles in a few seconds

## <img src="https://sternentstehung.de/espresso-ssg-dot.png"> Real-World Example

A good example for an Espresso-built website is my own one: [dominikbraun.io](https://dominikbraun.io).

## <img src="https://sternentstehung.de/espresso-ssg-dot.png"> Installation

### Running Espresso natively

Download the [latest release](https://github.com/dominikbraun/espresso/releases) for your platform. **On Linux and macOS**,
move the binary into a directory like `/usr/local/bin`. Make sure the directory is in `PATH`. **On Windows**, create a
directory like `C:\Program Files\espresso` and copy the executable into it. [Add the directory to `Path`.](https://www.computerhope.com/issues/ch000549.htm)

### Running Espresso via Docker

Assuming your project directory is called `my-blog`, just run the following command:

```shell script
$ docker container run -v $(pwd)/my-blog:/project dominikbraun/espresso
```

This will run the `build` command internally and you'll find the built website in `my-blog/target`. In order to run
another Espresso command, append it to the image name like so:

```shell script
$ docker container run dominikbraun/espresso version
```
