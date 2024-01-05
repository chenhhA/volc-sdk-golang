// Code generated by protoc-gen-volcengine-sdk
// source: VodSpaceService
// DO NOT EDIT!

package vod

import (
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/vod"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"
)

func Test_DeleteSpace(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodDeleteSpaceRequest{
		SpaceName: "your SpaceName",
	}

	resp, status, err := instance.DeleteSpace(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_CreateSpace(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodCreateSpaceRequest{
		SpaceName:   "your SpaceName",
		ProjectName: "your ProjectName",
		Description: "your Description",
		Region:      "your Region",
		UserName:    "your UserName",
	}

	resp, status, err := instance.CreateSpace(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_ListSpace(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodListSpaceRequest{
		Offset: 0,
		Limit:  0,
	}

	resp, status, err := instance.ListSpace(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_GetSpaceDetail(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodGetSpaceDetailRequest{
		SpaceName: "your SpaceName",
	}

	resp, status, err := instance.GetSpaceDetail(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_UpdateSpace(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodUpdateSpaceRequest{
		SpaceName:         "your SpaceName",
		SourceProjectName: "your SourceProjectName",
		TargetProjectName: "your TargetProjectName",
		Description:       "your Description",
	}

	resp, status, err := instance.UpdateSpace(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_UpdateSpaceUploadConfig(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodUpdateSpaceUploadConfigRequest{
		SpaceName:   "your SpaceName",
		ConfigKey:   "your ConfigKey",
		ConfigValue: "your ConfigValue",
	}

	resp, status, err := instance.UpdateSpaceUploadConfig(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}
