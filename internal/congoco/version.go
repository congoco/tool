package congoco

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-git/go-git/v6/plumbing/object"
)

type Version struct {
	Tag    *object.Tag
	Major  int
	Minor  int
	Patch  int
	Prefix string
}

func ParseVersion(tag *object.Tag, str, tagPrefix string) (*Version, error) {
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
		Tag:    tag,
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
