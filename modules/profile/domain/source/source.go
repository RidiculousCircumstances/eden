package source

const (
	VkProfileID        uint = 1
	InstagramProfileID uint = 2
)

const (
	VkProfileAlias        = "vk-profile"
	InstagramProfileAlias = "instagram-profile"
)

var sources = map[uint]string{
	VkProfileID:        VkProfileAlias,
	InstagramProfileID: InstagramProfileAlias,
}

func GetSourceAliasByID(id uint) (string, bool) {
	alias, exists := sources[id]
	return alias, exists
}

func GetIDBySourceAlias(alias string) (uint, bool) {
	for id, a := range sources {
		if a == alias {
			return id, true
		}
	}
	return 0, false
}
