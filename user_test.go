package goapitest

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Name string `valid:"required~Name connot be blank"`
	Url  string `gorm:"uiqueIndex" valid:url"`
}

func TestVideoVali(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Check name not blank", func(t *testing.T) {
		u := Video{
			Name: "",
			Url:  "oo",
		}

		ok, err := govalidator.ValidateStruct(u)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("name not blank"))
	})
}
