# Lambda Layer with KubeCtl v1.36

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Stable](https://img.shields.io/badge/cdk--constructs-stable-success.svg?style=for-the-badge)

---
<!--END STABILITY BANNER-->

This module exports a single class called `KubectlV36Layer` which is a `lambda.LayerVersion` that
bundles the [`kubectl`](https://kubernetes.io/docs/reference/kubectl/kubectl/) and the
[`helm`](https://helm.sh/) command line.

> * Helm Version: 4.2.3
> * Kubectl Version: 1.36.3

Usage:

```go
// KubectlLayer bundles the 'kubectl' and 'helm' command lines
import "github.com/cdklabs/awscdk-kubectl-go/kubectlv36"
import lambda "github.com/aws/aws-cdk-go/awscdk"

var fn Function

kubectl := kubectlv36.NewKubectlV36Layer(this, jsii.String("KubectlLayer"))
fn.AddLayers(kubectl)
```

`kubectl` will be installed under `/opt/kubectl/kubectl`, and `helm` will be installed under `/opt/helm/helm`.
