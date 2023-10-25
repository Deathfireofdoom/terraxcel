package client

import "github.com/Deathfireofdoom/terraxcel/server/src/pkg/excel"

func (c *TerraxcelClient) GetExtensions() []string {
	return excel.GetExtensions()
}
