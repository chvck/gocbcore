package gocbcore

import (
	"fmt"
	"strconv"
	"strings"
)

type NodeVersion struct {
	Major   int
	Minor   int
	Patch   int
	Build   int
	Edition string
}

func (v NodeVersion) Equal(ov NodeVersion) bool {
	if v.Major == ov.Major && v.Minor == ov.Minor &&
		v.Patch == ov.Patch && v.Edition == ov.Edition {
		return true
	}
	return false
}

func (v NodeVersion) Higher(ov NodeVersion) bool {
	if v.Major > ov.Major {
		return true
	} else if v.Major < ov.Major {
		return false
	}

	if v.Minor > ov.Minor {
		return true
	} else if v.Minor < ov.Minor {
		return false
	}

	if v.Patch > ov.Patch {
		return true
	} else if v.Patch < ov.Patch {
		return false
	}

	if v.Build > ov.Build {
		return true
	} else if v.Build < ov.Build {
		return false
	}

	editionMap := map[string]int{
		"Community":  1,
		"Enterprise": 2,
	}

	vEditionVal, ok := editionMap[v.Edition]
	if !ok {
		vEditionVal = 0
	}

	ovEditionVal, ok := editionMap[ov.Edition]
	if !ok {
		ovEditionVal = 0
	}

	if vEditionVal > ovEditionVal {
		return true
	}

	return false
}

func (v NodeVersion) Lower(ov NodeVersion) bool {
	return !v.Higher(ov) && !v.Equal(ov)
}

func nodeVersionFromString(version string) (NodeVersion, error) {
	vSplit := strings.Split(version, ".")
	lenSplit := len(vSplit)
	if lenSplit == 0 {
		return NodeVersion{}, fmt.Errorf("Must prove at least a major version")
	}

	var err error
	nodeVersion := NodeVersion{}
	nodeVersion.Major, err = strconv.Atoi(vSplit[0])
	if err != nil {
		return NodeVersion{}, fmt.Errorf("Major version is not a valid integer")
	}
	if lenSplit == 1 {
		return nodeVersion, nil
	}

	nodeVersion.Minor, err = strconv.Atoi(vSplit[1])
	if err != nil {
		return NodeVersion{}, fmt.Errorf("Minor version is not a valid integer")
	}
	if lenSplit == 2 {
		return nodeVersion, nil
	}

	nodeVersion.Patch, err = strconv.Atoi(vSplit[2])
	if err != nil {
		return NodeVersion{}, fmt.Errorf("Patch version is not a valid integer")
	}
	if lenSplit == 3 {
		return nodeVersion, nil
	}

	buildEdition := strings.Split(vSplit[3], "_")
	nodeVersion.Build, err = strconv.Atoi(buildEdition[0])
	if err != nil {
		return NodeVersion{}, fmt.Errorf("Build version is not a valid integer")
	}
	if len(buildEdition) == 1 {
		return nodeVersion, nil
	}

	nodeVersion.Edition = buildEdition[1]

	return nodeVersion, nil
}
