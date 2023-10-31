package client

import "github.com/Deathfireofdoom/terraxcel/server/pkg/excel"

func (c *TerraxcelClient) GetExtensions() []string {
	return excel.GetExtensions()
}
