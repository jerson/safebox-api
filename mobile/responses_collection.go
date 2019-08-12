package safebox

import (
	"encoding/json"
	"fmt"
)

// AccountSingleCollection ...
type AccountSingleCollection struct {
	s []*AccountSingle
}

// NewAccountSingleCollection ...
func NewAccountSingleCollection(s []*AccountSingle) *AccountSingleCollection {
	return &AccountSingleCollection{s: s}
}

// Clear ...
func (v *AccountSingleCollection) Clear() {
	v.s = v.s[:0]
}

// Equal ...
func (v *AccountSingleCollection) Equal(rhs *AccountSingleCollection) bool {
	if rhs == nil {
		return false
	}
	if len(v.s) != len(rhs.s) {
		return false
	}
	for i := range v.s {
		if !v.s[i].Equal(rhs.s[i]) {
			return false
		}
	}
	return true
}

// MarshalJSON ...
func (v *AccountSingleCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.s)
}

// UnmarshalJSON ...
func (v *AccountSingleCollection) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &v.s)
}

// Copy ...
func (v *AccountSingleCollection) Copy(rhs *AccountSingleCollection) {
	v.s = make([]*AccountSingle, len(rhs.s))
	copy(v.s, rhs.s)
}

// Clone ...
func (v *AccountSingleCollection) Clone() *AccountSingleCollection {
	return &AccountSingleCollection{
		s: v.s[:],
	}
}

// Index ...
func (v *AccountSingleCollection) Index(rhs *AccountSingle) int {
	for i, lhs := range v.s {
		if lhs == rhs {
			return i
		}
	}
	return -1
}

// Insert ...
func (v *AccountSingleCollection) Insert(i int, n *AccountSingle) {
	if i < 0 || i > len(v.s) {
		fmt.Printf("Vapi::AccountSingleCollection field_values.go error trying to insert at index %d\n", i)
		return
	}
	v.s = append(v.s, nil)
	copy(v.s[i+1:], v.s[i:])
	v.s[i] = n
}

// Remove ...
func (v *AccountSingleCollection) Remove(i int) {
	if i < 0 || i >= len(v.s) {
		fmt.Printf("Vapi::AccountSingleCollection field_values.go error trying to remove bad index %d\n", i)
		return
	}
	copy(v.s[i:], v.s[i+1:])
	v.s[len(v.s)-1] = nil
	v.s = v.s[:len(v.s)-1]
}

// Count ...
func (v *AccountSingleCollection) Count() int {
	return len(v.s)
}

// At ...
func (v *AccountSingleCollection) At(i int) *AccountSingle {
	if i < 0 || i >= len(v.s) {
		fmt.Printf("Vapi::AccountSingleCollection field_values.go invalid index %d\n", i)
	}
	return v.s[i]
}
