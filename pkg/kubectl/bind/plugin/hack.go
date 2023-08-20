package plugin

import (
	"fmt"

	"github.com/headzoo/surf"
)

func SimulateBrowser(authURLCh chan string, resource string) {
	browser := surf.NewBrowser()
	authURL := <-authURLCh

	// t.Logf("Browsing to auth URL: %s", authURL)
	err := browser.Open(authURL)
	// require.NoError(t, err)
	if err != nil {
		fmt.Println("error occurred")
	}

	// t.Logf("Waiting for browser to be at /resources")
	// framework.BrowerEventuallyAtPath(t, browser, "/resources")

	// t.Logf("Clicking %s", resource)
	err = browser.Click("a." + resource)
	// require.NoError(t, err)

	// t.Logf("Waiting for browser to be forwarded to client")
	// framework.BrowerEventuallyAtPath(t, browser, "/callback")

	if err != nil {
		fmt.Println("error occurred")
	}
}
