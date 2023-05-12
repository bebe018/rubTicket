package WD

import (
	"github.com/tebeka/selenium"
)

func LoginWithGoogle(wd selenium.WebDriver) {
	wd.Get("https://tixcraft.com/login/google")
	wd.Get("https://tixcraft.com")
}
