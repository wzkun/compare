package compare

import (
	"errors"
	"strings"
)

// CompareStrVersion 比较字符串型版本号
func CompareStrVersion(verA, verB string) (int, error) {
	var err error
	verStrArrA := strings.Split(verA, ".")
	verStrArrB := strings.Split(verB, ".")

	lenStrA := len(verStrArrA)
	lenStrB := len(verStrArrB)

	if lenStrA != lenStrB {
		err = errors.New("版本号格式不一致")
		return 0, err
	}

	return compareArrStrVers(verStrArrA, verStrArrB), nil
}

// 比较版本号字符串数组
func compareArrStrVers(verA, verB []string) int {
	for index, _ := range verA {
		littleResult := compareLittleVer(verA[index], verB[index])
		if littleResult != VersionEqual {
			return littleResult
		}
	}

	return VersionEqual
}

// compareLittleVer 比较小版本号字符串
func compareLittleVer(verA, verB string) int {
	bytesA := []byte(verA)
	bytesB := []byte(verB)

	lenA := len(bytesA)
	lenB := len(bytesB)
	if lenA > lenB {
		return VersionBig
	}

	if lenA < lenB {
		return VersionSmall
	}

	//如果长度相等则按byte位进行比较
	return compareByBytes(bytesA, bytesB)
}

// compareByBytes 按byte位进行比较小版本号
func compareByBytes(verA, verB []byte) int {
	for index, _ := range verA {
		if verA[index] > verB[index] {
			return VersionBig
		}
		if verA[index] < verB[index] {
			return VersionSmall
		}
	}

	return VersionEqual
}
