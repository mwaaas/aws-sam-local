package cloudformation

// AWSElasticLoadBalancingV2LoadBalancer_LoadBalancerAttribute AWS CloudFormation Resource (AWS::ElasticLoadBalancingV2::LoadBalancer.LoadBalancerAttribute)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-loadbalancer-loadbalancerattributes.html
type AWSElasticLoadBalancingV2LoadBalancer_LoadBalancerAttribute struct {

	// Key AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-loadbalancer-loadbalancerattributes.html#cfn-elasticloadbalancingv2-loadbalancer-loadbalancerattributes-key
	Key string `json:"Key,omitempty"`

	// Value AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-loadbalancer-loadbalancerattributes.html#cfn-elasticloadbalancingv2-loadbalancer-loadbalancerattributes-value
	Value string `json:"Value,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingV2LoadBalancer_LoadBalancerAttribute) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancingV2::LoadBalancer.LoadBalancerAttribute"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticLoadBalancingV2LoadBalancer_LoadBalancerAttribute) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
