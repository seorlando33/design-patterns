package factory

import (
	"fmt"
	"testing"
)

func TestCreateWeapon(t *testing.T) {
	wf := &WeaponFactory{}

	tests := []struct {
		name     string
		t        int
		wantType string
		wantErr  error
	}{
		{"Sword", SwordType, "*factory.Sword", nil},
		{"Bow", BowType, "*factory.Bow", nil},
		{"Gun", GunType, "*factory.Gun", nil},
		{"Invalid", 128, "nil", fmt.Errorf("weapon %d not implemented", 128)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := wf.CreateWeapon(test.t)
			if err != nil && err.Error() != test.wantErr.Error() {
				t.Errorf("CreateWeapon() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if got != nil {
				if gotType := fmt.Sprintf("%T", got); gotType != test.wantType {
					t.Errorf("CreateWeapon() = %v, want %v", gotType, test.wantType)
				}
			} else {
				if test.wantErr == nil {
					t.Errorf("CreateWeapon() = %v, want %v", got, test.wantType)
				}
			}
		})
	}
}
