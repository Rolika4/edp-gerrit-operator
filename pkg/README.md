# Autogenerated files

Directory `apis/v2/v1alpha` contains autogenerated files. To generate `deepcopy*` use external tool [deepcopy-gen](https://pkg.go.dev/k8s.io/gengo).

## tool installation

```bash
$ GO111MODULE="off" go get -u -v k8s.io/gengo/...
```

Make sure `deepcopy-gen` is installed into `$GOPATH/bin/deepcopy-gen`, example:
```bash
$ ls -la $(go env GOPATH)/bin
-rwxr-xr-x  1 user  user  7834896 Feb  4 10:37 deepcopy-gen
```

## tool usage

```bash
$ cd <PROJECT_ROOT>
$ deepcopy-gen -i ./pkg/apis/v2/v1alpha1 -o ./ 
```

## autogenerated files fixes

* The result filename is `deepcopy_generated.go` and it conflicts with `zz_generated.deepcopy.go`. The last one should be removed.

* Don't know why, but result-files are not compile-able, because they contain `DeepCopyInto` methods for simple go-types, like [time.Time](https://pkg.go.dev/time#Time). So you need edit autogenerated files: remove these method-calls.