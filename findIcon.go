package icons

func findIcon(iconName string) string {
	if icon, found := bootstrapMap[iconName]; found {
		if found {
			return icon
		}
	}

	if icon, found := boxiconsMap[iconName]; found {
		if found {
			return icon
		}
	}

	return ""
}
