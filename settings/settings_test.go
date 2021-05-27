package settings_test

import (
	"github.com/annakallo/parmtracker/settings"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSettingsGetUpdateDeleteVersion(t *testing.T) {
	tt := []struct {
		Name  string
		Key   string
		Value string
	}{
		{
			Name:  "New and Delete",
			Key:   "x \"-., hrg rew qwe bgf bfe  ",
			Value: "this is the value",
		},
		{
			Name:  "New and Delete 2",
			Key:   "y_fre_qwe_qwe_654_654_432_123",
			Value: "this is the other value!_.;,()?'¿¡=&%$",
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			result := settings.GetCurrentVersion(tc.Key)
			assert.Equal(t, "", result)
			settings.UpdateVersion(tc.Key, tc.Value)
			result = settings.GetCurrentVersion(tc.Key)
			assert.Equal(t, tc.Value, result)
			settings.Delete(tc.Key)
			result = settings.GetCurrentVersion(tc.Key)
			assert.Equal(t, "", result)
		})
	}
}
