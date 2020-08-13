// Copyright (c) 2020 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package vault

import (
	log "github.com/sirupsen/logrus"
)

func getOperStatusParams(vaultPath string) []string {
	args := []string{"/hostfs", "zfs", "get", "mounted,encryption,keystatus", vaultPath}
	return args
}

// CheckOperStatus returns ZFS status
func CheckOperStatus(vaultPath string) (string, error) {
	args := getOperStatusParams(vaultPath)
	if stdOut, stdErr, err := execCmd(ZfsPath, args...); err != nil {
		log.Errorf("oper status query for %s results in error=%v, %s, %s",
			vaultPath, err, stdOut, stdErr)
		return stdErr, err
	} else {
		return stdOut, nil
	}
}
