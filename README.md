# bazel-go-showcase

[CI pipeline](https://github.com/slamdev/bazel-go-showcase/actions/workflows/build.yaml) that runs the code with different OSes.

Targeted dev experience:

- Developer on MacOS can run go binary and container image via:

```shell
bazel run //:binary
bazel run //:image
```

- Developer on Linux can run go binary and container image via:

```shell
bazel run //:binary
bazel run //:image
```
