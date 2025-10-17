//go:build no_runtime_type_checking

package kubectlv34

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KubectlV34Layer) validateAddPermissionParameters(id *string, permission *awslambda.LayerVersionPermission) error {
	return nil
}

func (k *jsiiProxy_KubectlV34Layer) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (k *jsiiProxy_KubectlV34Layer) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (k *jsiiProxy_KubectlV34Layer) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateKubectlV34Layer_FromLayerVersionArnParameters(scope constructs.Construct, id *string, layerVersionArn *string) error {
	return nil
}

func validateKubectlV34Layer_FromLayerVersionAttributesParameters(scope constructs.Construct, id *string, attrs *awslambda.LayerVersionAttributes) error {
	return nil
}

func validateKubectlV34Layer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubectlV34Layer_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateKubectlV34Layer_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewKubectlV34LayerParameters(scope constructs.Construct, id *string) error {
	return nil
}

