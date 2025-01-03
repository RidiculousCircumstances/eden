package source

const (
	VkProfileID uint = 1
)

const (
	VkProfileAlias = "vk-profile"
)

var sources = map[uint]string{
	VkProfileID: VkProfileAlias,
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
