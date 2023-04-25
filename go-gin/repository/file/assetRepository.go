package fileRepository

import (
	"com.xpwk/go-gin/model"
	"com.xpwk/go-gin/repository"
)

var AssetRepository = new(FileAssetRepository)

type FileAssetRepository struct {
}

func (*FileAssetRepository) Insert(asset *model.FileAsset) error {
	if err := repository.GetDB().Create(asset).Error; err != nil {
		return err
	}
	return nil
}

func (*FileAssetRepository) MultiInsert(assets *[]model.FileAsset) error {
	if err := repository.GetDB().Create(assets).Error; err != nil {
		return err
	}
	return nil
}
