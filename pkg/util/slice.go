package util

import "k8s.io/apimachinery/pkg/util/sets"

func SliceIntersection(a, b []string) (c []string) {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}

	return
}

func SliceUnion(a, b []string) []string {
	s := sets.NewString()

	s.Insert(a...)
	s.Insert(b...)

	return s.List()
}
