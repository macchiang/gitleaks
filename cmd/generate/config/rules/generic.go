package rules

import (
	"github.com/zricethezav/gitleaks/v8/config"
)

func GenericCredential() *config.Rule {
	// define rule
	r := config.Rule{
		RuleID:      "generic-api-key",
		Description: "Generic API Key",
		Regex: generateSemiGenericRegex([]string{
			"key",
			"api[^Version]",
			"token",
			"pat",
			"secret",
			"client",
			"password",
			"auth",
		}, `[0-9a-z\-_.=]{10,150}`),
		SecretGroup: 1,
		Keywords: []string{
			"key",
			"api",
			"token",
			"secret",
			"client",
			"pat",
			"password",
			"auth",
		},
		Entropy: 3.7,
	}

	// validate
	tps := []string{
		generateSampleSecret("generic", "CLOJARS_34bf0e88955ff5a1c328d6a7491acc4f48e865a7b8dd4d70a70749037443"),
		generateSampleSecret("generic", "Zf3D0LXCM3EIMbgJpUNnkRtOfOueHznB"),
		`"client_id" : "0afae57f3ccfd9d7f5767067bc48b30f719e271ba470488056e37ab35d4b6506"`,
		`"client_secret" : "6da89121079f83b2eb6acccf8219ea982c3d79bccc3e9c6a85856480661f8fde",`,
		// TODO add more
	}
	return validate(r, tps)
}
