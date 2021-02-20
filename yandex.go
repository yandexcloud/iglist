package main

import (
	"context"
	"fmt"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1/instancegroup"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

type sdk ycsdk.SDK

func ycNew(ctx context.Context, token string) (*sdk, error) {
	if token == "" {
		api, err := ycsdk.Build(ctx, ycsdk.Config{
			Credentials: ycsdk.InstanceServiceAccount(),
		})
		return (*sdk)(api), err
	}
	api, err := ycsdk.Build(ctx, ycsdk.Config{
		Credentials: ycsdk.NewIAMTokenCredentials(token),
	})
	return (*sdk)(api), err
}

func (api *sdk) GetGroupID(ctx context.Context, folderID, name string) (string, error) {
	req := instancegroup.ListInstanceGroupsRequest{
		FolderId: folderID,
		Filter:   fmt.Sprintf("name=\"%s\"", name),
		PageSize: 1,
	}

	l, err := (*ycsdk.SDK)(api).InstanceGroup().InstanceGroup().List(ctx, &req)
	if err != nil {
		return "", err
	}

	return l.InstanceGroups[0].Id, nil
}

func (api *sdk) GetGroupMembers(ctx context.Context, groupID string) ([]*instancegroup.ManagedInstance, error) {
	req := instancegroup.ListInstanceGroupInstancesRequest{
		InstanceGroupId: groupID,
		// NOTE: we do not expect large instance groups
		PageSize: 1000,
	}

	res, err := (*ycsdk.SDK)(api).InstanceGroup().InstanceGroup().ListInstances(ctx, &req)
	if err != nil {
		return nil, err
	}

	return res.Instances, nil
}
