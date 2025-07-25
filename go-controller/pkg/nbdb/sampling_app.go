// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package nbdb

import "github.com/ovn-kubernetes/libovsdb/model"

const SamplingAppTable = "Sampling_App"

type (
	SamplingAppType = string
)

var (
	SamplingAppTypeDrop   SamplingAppType = "drop"
	SamplingAppTypeACLNew SamplingAppType = "acl-new"
	SamplingAppTypeACLEst SamplingAppType = "acl-est"
)

// SamplingApp defines an object in Sampling_App table
type SamplingApp struct {
	UUID        string            `ovsdb:"_uuid"`
	ExternalIDs map[string]string `ovsdb:"external_ids"`
	ID          int               `ovsdb:"id"`
	Type        SamplingAppType   `ovsdb:"type"`
}

func (a *SamplingApp) GetUUID() string {
	return a.UUID
}

func (a *SamplingApp) GetExternalIDs() map[string]string {
	return a.ExternalIDs
}

func copySamplingAppExternalIDs(a map[string]string) map[string]string {
	if a == nil {
		return nil
	}
	b := make(map[string]string, len(a))
	for k, v := range a {
		b[k] = v
	}
	return b
}

func equalSamplingAppExternalIDs(a, b map[string]string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if w, ok := b[k]; !ok || v != w {
			return false
		}
	}
	return true
}

func (a *SamplingApp) GetID() int {
	return a.ID
}

func (a *SamplingApp) GetType() SamplingAppType {
	return a.Type
}

func (a *SamplingApp) DeepCopyInto(b *SamplingApp) {
	*b = *a
	b.ExternalIDs = copySamplingAppExternalIDs(a.ExternalIDs)
}

func (a *SamplingApp) DeepCopy() *SamplingApp {
	b := new(SamplingApp)
	a.DeepCopyInto(b)
	return b
}

func (a *SamplingApp) CloneModelInto(b model.Model) {
	c := b.(*SamplingApp)
	a.DeepCopyInto(c)
}

func (a *SamplingApp) CloneModel() model.Model {
	return a.DeepCopy()
}

func (a *SamplingApp) Equals(b *SamplingApp) bool {
	return a.UUID == b.UUID &&
		equalSamplingAppExternalIDs(a.ExternalIDs, b.ExternalIDs) &&
		a.ID == b.ID &&
		a.Type == b.Type
}

func (a *SamplingApp) EqualsModel(b model.Model) bool {
	c := b.(*SamplingApp)
	return a.Equals(c)
}

var _ model.CloneableModel = &SamplingApp{}
var _ model.ComparableModel = &SamplingApp{}
