//go:build no_runtime_type_checking

// A Lambda Layer that contains kubectl v1.26
package kubectlv26

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KubectlV26Layer) validateAddPermissionParameters(id *string, permission *awslambda.LayerVersionPermission) error {
	return nil
}

func (k *jsiiProxy_KubectlV26Layer) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (k *jsiiProxy_KubectlV26Layer) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (k *jsiiProxy_KubectlV26Layer) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateKubectlV26Layer_FromLayerVersionArnParameters(scope constructs.Construct, id *string, layerVersionArn *string) error {
	return nil
}

func validateKubectlV26Layer_FromLayerVersionAttributesParameters(scope constructs.Construct, id *string, attrs *awslambda.LayerVersionAttributes) error {
	return nil
}

func validateKubectlV26Layer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubectlV26Layer_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewKubectlV26LayerParameters(scope constructs.Construct, id *string) error {
	return nil
}
