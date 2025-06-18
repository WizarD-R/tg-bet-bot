package user

import "slices"

var Admins []string

func isAdmin(user string) bool {
	if slices.Contains(Admins, user) {
		return true
	}
	return false
}
