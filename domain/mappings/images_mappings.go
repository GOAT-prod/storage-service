package mappings

import (
	"github.com/samber/lo"
	"storage-service/database"
)

func ToImages(images []database.DbImage) []string {
	return lo.Map(images, func(item database.DbImage, _ int) string {
		return item.Url
	})
}
