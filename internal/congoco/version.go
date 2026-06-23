package congoco

import (
	"fmt"
	"strconv"
	"strings"
)

type Version struct {
	Major  int
	Minor  int
	Patch  int
	Prefix string
}

func VersionFromString(str string, tagPrefix string) (*Version, error) {
	if !strings.HasPrefix(str, tagPrefix) {
		return nil, fmt.Errorf("Version tag without prefix")
	}

	vStr := strings.TrimPrefix(str, fmt.Sprintf("%s", tagPrefix))
	parts := strings.Split(vStr, ".")

	majorStr := parts[0]
	major, err := strconv.Atoi(majorStr)
	if err != nil {
		return nil, err
	}

	minorStr := parts[1]
	minor, err := strconv.Atoi(minorStr)
	if err != nil {
		return nil, err
	}

	patchStr := parts[2]
	patch, err := strconv.Atoi(patchStr)
	if err != nil {
		return nil, err
	}

	v := Version{
		Major:  major,
		Minor:  minor,
		Patch:  patch,
		Prefix: tagPrefix,
	}
	return &v, nil
}

func (v *Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func (v *Version) Tag() string {
	return fmt.Sprintf("%s%d.%d.%d", v.Prefix, v.Major, v.Minor, v.Patch)
}
