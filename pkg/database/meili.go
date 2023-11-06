package database

import (
	"fmt"

	"github.com/go-microservice-boilerplate/pkg/config"
	"github.com/meilisearch/meilisearch-go"
)

func NewMeiliSearch(c *config.CommonConfig) *meilisearch.Client {
	meili := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   fmt.Sprintf("http://%s:%s", c.Meili.Host, c.Meili.Port),
		APIKey: c.Meili.Key,
	})
	// meili.CreateIndex(&meilisearch.IndexConfig{
	// 	Uid:        "issues",
	// 	PrimaryKey: "id",
	// })

	// mc := meili.Index("issues")

	// mc.UpdateSettings(&meilisearch.Settings{
	// 	FilterableAttributes: []string{
	// 		"id",
	// 	},
	// 	SortableAttributes: []string{
	// 		"created_at",
	// 		"updated_at",
	// 	},
	// })
	return meili
}
