package suites

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tebeka/selenium"
)

func (wds *WebDriverSession) verifyURLIs(ctx context.Context, t *testing.T, url string) {
	err := wds.Wait(ctx, func(driver selenium.WebDriver) (bool, error) {
		currentURL, err := driver.CurrentURL()

		if err != nil {
			return false, err
		}

		return currentURL == url, nil
	})

	require.NoError(t, err)
}
