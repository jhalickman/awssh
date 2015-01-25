package main

import (
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/ec2"
	"sort"
)

type Instance ec2.Instance
type InstanceList []Instance

func (a InstanceList) Len() int           { return len(a) }
func (a InstanceList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a InstanceList) Less(i, j int) bool { return a[i].getName() < a[j].getName() }

func listInstancesFromResponse(resp *ec2.InstancesResp) InstanceList {
	var list InstanceList

	for _, reservation := range resp.Reservations {
		for _, instance := range reservation.Instances {
			list = append(list, Instance(instance))
		}
	}

	sort.Sort(list)
	return list
}

func (instance *Instance) getName() string {
	for _, tag := range instance.Tags {
		if tag.Key == "Name" {
			return tag.Value
		}
	}
	return "No Name set"
}

func (instanceList InstanceList) getInstance(name string) *Instance {
	for _, instance := range instanceList {
		if instance.getName() == name {
			return &instance
		}
	}
	return nil
}

func getEC2Client(token string, secret string, region aws.Region) (*ec2.EC2, error) {
	auth, err := aws.GetAuth(token, secret)
	if err != nil {
		return nil, err
	}
	client := ec2.New(auth, region)
	return client, nil
}

func getRunningInstances(token string, secret string, region aws.Region) (InstanceList, error) {
	client, err := getEC2Client(token, secret, region)
	if err != nil {
		return nil, err
	}

	filter := ec2.NewFilter()
	filter.Add("instance-state-code", "16")
	resp, err := client.Instances(nil, filter)
	if err != nil {
		return nil, err
	}

	instances := listInstancesFromResponse(resp)
	return instances, nil
}
