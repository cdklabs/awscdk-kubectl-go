# Lambda Layer with KubeCtl v1.33

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Stable](https://img.shields.io/badge/cdk--constructs-stable-success.svg?style=for-the-badge)

---
<!--END STABILITY BANNER-->

This module exports a single class called `KubectlV33Layer` which is a `lambda.LayerVersion` that
bundles the [`kubectl`](https://kubernetes.io/docs/reference/kubectl/kubectl/) and the
[`helm`](https://helm.sh/) command line.

> * Helm Version: 3.18.0
> * Kubectl Version: 1.33.0

Usage:

```go
// KubectlLayer bundles the 'kubectl' and 'helm' command lines
import "github.com/cdklabs/awscdk-kubectl-go/kubectlv33"
import lambda "github.com/aws/aws-cdk-go/awscdk"

var fn Function

kubectl := kubectlv33.NewKubectlV33Layer(this, jsii.String("KubectlLayer"))
fn.AddLayers(kubectl)
```

`kubectl` will be installed under `/opt/kubectl/kubectl`, and `helm` will be installed under `/opt/helm/helm`.
