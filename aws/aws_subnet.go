// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListSubnet(client *Client) ([]Resource, error) {
	req := client.ec2conn.DescribeSubnetsRequest(&ec2.DescribeSubnetsInput{})

	var result []Resource

	p := ec2.NewDescribeSubnetsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Subnets {
			if *r.OwnerId != client.accountid {
				continue
			}
			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}

			result = append(result, Resource{
				Type:   "aws_subnet",
				ID:     *r.SubnetId,
				Region: client.ec2conn.Config.Region,
				Tags:   tags,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
