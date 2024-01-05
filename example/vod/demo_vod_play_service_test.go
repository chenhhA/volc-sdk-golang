// Code generated by protoc-gen-volcengine-sdk
// source: VodPlayService
// DO NOT EDIT!

package vod

import (
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/vod"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"
)

func Test_GetAllPlayInfo(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodGetAllPlayInfoRequest{
		Vids:                     "your Vids",
		Formats:                  "your Formats",
		Codecs:                   "your Codecs",
		Definitions:              "your Definitions",
		FileTypes:                "your FileTypes",
		LogoTypes:                "your LogoTypes",
		NeedEncryptStream:        "your NeedEncryptStream",
		Ssl:                      "your Ssl",
		NeedThumbs:               "your NeedThumbs",
		NeedBarrageMask:          "your NeedBarrageMask",
		CdnType:                  "your CdnType",
		UnionInfo:                "your UnionInfo",
		PlayScene:                "your PlayScene",
		DrmExpireTimestamp:       "your DrmExpireTimestamp",
		HDRType:                  "your HDRType",
		KeyFrameAlignmentVersion: "your KeyFrameAlignmentVersion",
		UserAction:               "your UserAction",
		Quality:                  "your Quality",
	}

	resp, status, err := instance.GetAllPlayInfo(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_GetPlayInfo(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodGetPlayInfoRequest{
		Vid:                "your Vid",
		Format:             "your Format",
		Codec:              "your Codec",
		Definition:         "your Definition",
		FileType:           "your FileType",
		LogoType:           "your LogoType",
		Base64:             "your Base64",
		Ssl:                "your Ssl",
		NeedThumbs:         "your NeedThumbs",
		NeedBarrageMask:    "your NeedBarrageMask",
		CdnType:            "your CdnType",
		UnionInfo:          "your UnionInfo",
		HDRDefinition:      "your HDRDefinition",
		PlayScene:          "your PlayScene",
		DrmExpireTimestamp: "your DrmExpireTimestamp",
		Quality:            "your Quality",
		PlayConfig:         "your PlayConfig",
		NeedOriginal:       "your NeedOriginal",
	}

	resp, status, err := instance.GetPlayInfo(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_GetPrivateDrmPlayAuth(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodGetPrivateDrmPlayAuthRequest{
		DrmType:     "your DrmType",
		Vid:         "your Vid",
		PlayAuthIds: "your PlayAuthIds",
		UnionInfo:   "your UnionInfo",
	}

	resp, status, err := instance.GetPrivateDrmPlayAuth(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_GetHlsDecryptionKey(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodGetHlsDecryptionKeyRequest{
		DrmAuthToken: "your DrmAuthToken",
		Ak:           "your Ak",
		Source:       "your Source",
	}

	resp, status, err := instance.GetHlsDecryptionKey(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_GetPlayInfoWithLiveTimeShiftScene(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodGetPlayInfoWithLiveTimeShiftSceneRequest{
		StoreUris:             "your StoreUris",
		SpaceName:             "your SpaceName",
		Ssl:                   "your Ssl",
		ExpireTimestamp:       "your ExpireTimestamp",
		NeedComposeBucketName: "your NeedComposeBucketName",
	}

	resp, status, err := instance.GetPlayInfoWithLiveTimeShiftScene(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_DescribeDrmDataKey(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodDescribeDrmDataKeyRequest{
		Ak: "your Ak",
	}

	resp, status, err := instance.DescribeDrmDataKey(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}
